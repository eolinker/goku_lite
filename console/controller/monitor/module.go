package monitor

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/monitor"
)

//GetMonitorModules 获取监控模块列表
func GetMonitorModules(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}

	result, err := monitor.GetMonitorModules()
	if err != nil {
		controller.WriteError(httpResponse,
			"410000",
			"monitor",
			err.Error(),
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "monitorModule", "moduleList", result)
	return
}

func SetMonitorModule(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	moduleName := httpRequest.Form.Get("moduleName")
	moduleStatus := httpRequest.Form.Get("moduleStatus")
	config := httpRequest.Form.Get("config")

	status, err := strconv.Atoi(moduleStatus)
	if err != nil && moduleStatus != "" {
		errInfo := "[error]illegal moduleStatus"
		controller.WriteError(httpResponse,
			"410001",
			"monitor",
			errInfo,
			errors.New(errInfo))
		return
	}

	err = monitor.SetMonitorModule(moduleName, config, status)
	if err != nil {
		controller.WriteError(httpResponse,
			"410000",
			"monitor",
			err.Error(),
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "monitorModule", "", nil)
	return
}
