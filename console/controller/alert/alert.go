package alert

import (
	"net/http"
	"strconv"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/console/controller"
	alert_module "github.com/eolinker/goku-api-gateway/console/module/alert"
)

// GetAlertMsgList 获取告警信息列表
func GetAlertMsgList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAlert, controller.OperationREAD)
	if e != nil {
		return
	}

	page := httpRequest.PostFormValue("page")
	pageSize := httpRequest.PostFormValue("pageSize")

	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	pSize, err := strconv.Atoi(pageSize)
	if err != nil {
		pSize = 15
	}
	_, result, count, err := alert_module.GetAlertMsgList(p, pSize)

	controller.WriteResultInfoWithPage(httpResponse, "alert", "alertMessageList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: count,
		Page:     p,
		PageSize: pSize,
	})
	return
}

//ClearAlertMsg 清空告警信息列表
func ClearAlertMsg(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAlert, controller.OperationEDIT)
	if e != nil {
		return
	}

	flag, result, err := alert_module.ClearAlertMsg()
	if !flag {

		controller.WriteError(httpResponse,
			"330000",
			"alert",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "alert", "", nil)
	return
}

//DeleteAlertMsg 删除告警信息
func DeleteAlertMsg(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAlert, controller.OperationEDIT)
	if e != nil {
		return
	}

	alertID := httpRequest.PostFormValue("alertID")

	id, err := strconv.Atoi(alertID)
	if err != nil {

		log.Debug(err)

		controller.WriteError(httpResponse,
			"330001",
			"alert",
			"[ERROR]Illegal alertID!",
			err)
		return
	}
	flag, result, err := alert_module.DeleteAlertMsg(id)
	if !flag {

		controller.WriteError(httpResponse,
			"330000",
			"alert",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "alert", "", nil)

	return
}

//GetAlertConfig 获取告警配置
func GetAlertConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAlert, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := alert_module.GetAlertConfig()
	if !flag {

		controller.WriteError(httpResponse,
			"320000",
			"gateway",
			"[ERROR]The gateway config does not exist",
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "alert", "gatewayConfig", result)
	return

}
