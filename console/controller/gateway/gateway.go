package gateway

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/gateway"
	monitor_read "github.com/eolinker/goku-api-gateway/server/monitor/monitor-read"
)

func GetGatewayConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationGatewayConfig, controller.OperationREAD)
	if e != nil {
		return
	}

	result, err := gateway.GetGatewayConfig()
	if err != nil {

		controller.WriteError(httpResponse,
			"320000",
			"gateway",
			"[ERROR]The gateway config does not exist",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "gateway", "gatewayConfig", result)

	return
}

func EditGatewayBaseConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationGatewayConfig, controller.OperationEDIT)
	if e != nil {
		return
	}

	successCode := httpRequest.PostFormValue("successCode")
	nodeUpdatePeriod := httpRequest.PostFormValue("nodeUpdatePeriod")
	monitorUpdatePeriod := httpRequest.PostFormValue("monitorUpdatePeriod")
	monitorTimeout := httpRequest.PostFormValue("monitorTimeout")

	nodePeriod, err := strconv.Atoi(nodeUpdatePeriod)
	if err != nil {

		controller.WriteError(httpResponse,
			"320001",
			"gateway",
			"[ERROR]Illegal nodeUpdatePeriod!",
			err)
		return
	}
	monitorPeriod, err := strconv.Atoi(monitorUpdatePeriod)
	if err != nil && monitorUpdatePeriod != "" {

		controller.WriteError(httpResponse,
			"320002",
			"gateway",
			"[ERROR]Illegal monitorUpdatePeriod!",
			err)
		return
	}
	if monitorUpdatePeriod == "" {
		monitorPeriod = 30
	}
	timeout, err := strconv.Atoi(monitorTimeout)
	if (err != nil && monitorTimeout != "") || (timeout < 1 && timeout > 30) {

		controller.WriteError(httpResponse,
			"320011",
			"gateway",
			"[ERROR]Illegal monitorTimeout!",
			err)
		return
	}
	if monitorTimeout == "" {
		timeout = 5
	}
	flag, result, err := gateway.EditGatewayBaseConfig(successCode, nodePeriod, monitorPeriod, timeout)
	if !flag {
		monitor_read.SetPeriod(monitorPeriod)

		controller.WriteError(httpResponse,
			"320000",
			"gateway",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "gateway", "", nil)

	return
}

func EditGatewayAlarmConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationGatewayConfig, controller.OperationEDIT)
	if e != nil {
		return
	}

	alertStatus := httpRequest.PostFormValue("alertStatus")
	apiAlertInfo := httpRequest.PostFormValue("apiAlertInfo")
	sender := httpRequest.PostFormValue("sender")
	senderPassword := httpRequest.PostFormValue("senderPassword")
	smtpAddress := httpRequest.PostFormValue("smtpAddress")
	smtpPort := httpRequest.PostFormValue("smtpPort")
	smtpProtocol := httpRequest.PostFormValue("smtpProtocol")

	aStatus, err := strconv.Atoi(alertStatus)
	if err != nil {

		controller.WriteError(httpResponse,
			"320003",
			"gateway",
			"[ERROR]Illegal alertStatus!",
			err)
		return
	}
	port, err := strconv.Atoi(smtpPort)
	if err != nil {

		controller.WriteError(httpResponse,
			"320005",
			"gateway",
			"[ERROR]Illegal smtpPort!",
			err)
		return
	}
	proto, err := strconv.Atoi(smtpProtocol)
	if err != nil {
		controller.WriteError(httpResponse,
			"320006",
			"gateway",
			"[ERROR]Illegal smtpProtocol!",
			err)
		return
	}
	flag, result, err := gateway.EditGatewayAlarmConfig(apiAlertInfo, sender, senderPassword, smtpAddress, aStatus, port, proto)
	if !flag {
		controller.WriteError(httpResponse,
			"320000",
			"gateway",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "gateway", "", nil)

	return
}
