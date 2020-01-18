package monitor

import (
	"net/http"
	"strconv"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
	"github.com/eolinker/goku-api-gateway/ksitigarbha"

	"github.com/pkg/errors"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/monitor"
)

const operationMonitorModule = "monitorModuleManagement"

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/get": factory.NewAccountHandleFunction(operationMonitorModule, false, GetMonitorModules),
		"/set": factory.NewAccountHandleFunction(operationMonitorModule, true, SetMonitorModule),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

//GetMonitorModules 获取监控模块列表
func GetMonitorModules(httpResponse http.ResponseWriter, httpRequest *http.Request) {
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

//SetMonitorModule 设置监控模块
func SetMonitorModule(httpResponse http.ResponseWriter, httpRequest *http.Request) {
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
	if status == 1 {
		ksitigarbha.Open(moduleName, config)
	} else {
		ksitigarbha.Close(moduleName)
	}
	controller.WriteResultInfo(httpResponse, "monitorModule", "", nil)
	return
}
