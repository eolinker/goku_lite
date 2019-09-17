package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/eolinker/goku/console/controller"
	alert_module "github.com/eolinker/goku/console/module/alert"
	"github.com/eolinker/goku/utils"
)

// 新增告警信息
func AddAlertMsg(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	nodeIP := httpRequest.RemoteAddr
	requestID := httpRequest.PostFormValue("requestID")
	apiName := httpRequest.PostFormValue("apiName")
	requestURL := httpRequest.PostFormValue("requestURL")
	targetServer := httpRequest.PostFormValue("targetServer")
	targetURL := httpRequest.PostFormValue("targetURL")
	requestMethod := httpRequest.PostFormValue("requestMethod")
	proxyMethod := httpRequest.PostFormValue("proxyMethod")
	alertPeriodType := httpRequest.PostFormValue("alertPeriodType")
	alertCount := httpRequest.PostFormValue("alertCount")
	headerList := httpRequest.PostFormValue("headerList")
	queryParamList := httpRequest.PostFormValue("queryParamList")
	formParamList := httpRequest.PostFormValue("formParamList")
	responseHeaderList := httpRequest.PostFormValue("responseHeaderList")
	responseStatus := httpRequest.PostFormValue("responseStatus")
	isAlert := httpRequest.PostFormValue("isAlert")
	clusterName := httpRequest.PostFormValue("clusterName")
	nodePort := httpRequest.PostFormValue("nodePort")
	apiID := httpRequest.PostFormValue("apiID")
	strategyID := httpRequest.PostFormValue("strategyID")
	strategyName := httpRequest.PostFormValue("strategyName")

	ip := utils.InterceptIP(nodeIP, ":") + ":" + nodePort
	if realIP := strings.TrimSpace(httpRequest.Header.Get("X-Real-Ip")); realIP != "" {
		ip = realIP + ":" + nodePort
	}
	period, err := strconv.Atoi(alertPeriodType)
	if err != nil {
		controller.WriteError(httpResponse,
			"330002",
			"alert",
			"[ERROR]Illegal alertPeriodType!",
			err)
		return
	}
	count, err := strconv.Atoi(alertCount)
	if err != nil {
		controller.WriteError(httpResponse,
			"330003",
			"alert",
			"[ERROR]Illegal alertCount!",
			err)
		return
	}
	status, err := strconv.Atoi(responseStatus)
	if err != nil {
		controller.WriteError(httpResponse,
			"330004",
			"alert",
			"[ERROR]Illegal responseStatus!",
			err)
		return
	}
	alert := false
	if isAlert == "true" {
		alert = true
	}

	flag, result, err := alert_module.AddAlertMsg(apiID, apiName, requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList, strategyID, strategyName, requestID, period, count, status, alert, ip, clusterName)
	if !flag {
		controller.WriteError(httpResponse,
			"330000",
			"alert",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "alert", "", nil)
}
