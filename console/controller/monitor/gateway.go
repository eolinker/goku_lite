package monitor

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/monitor"
	"github.com/eolinker/goku-api-gateway/server/cluster"
)

func GetGatewayMonitorSummaryByPeriod(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}

	clusterName := httpRequest.PostFormValue("cluster")
	clusterId, has := cluster.GetId(clusterName)
	if !has && clusterName != "" {
		controller.WriteError(httpResponse, "340003", "", "[ERROR]Illegal cluster!", nil)
		return
	}

	beginTime := httpRequest.PostFormValue("beginTime")
	endTime := httpRequest.PostFormValue("endTime")
	period := httpRequest.PostFormValue("period")
	p, err := strconv.Atoi(period)
	if err != nil {
		p = 0
	}
	bt, err := time.ParseInLocation("2006-01-02", beginTime, time.Local)
	if err != nil && p == 3 {
		controller.WriteError(httpResponse,
			"340001",
			"monitor",
			"[ERROR]Illegal beginTime!",
			err)
		return

	}
	et, err := time.ParseInLocation("2006-01-02", endTime, time.Local)
	if err != nil && p == 3 {
		controller.WriteError(httpResponse,
			"340002",
			"monitor",
			"[ERROR]Illegal endTime!",
			err)
		return

	}
	if bt.After(et) && p == 3 {
		controller.WriteError(httpResponse,
			"340003",
			"monitor",
			"[ERROR]beginTime should be before the endTime!",
			nil)
		return

	}
	flag, result, err := monitor.GetGatewayMonitorSummaryByPeriod(clusterId, beginTime, endTime, p)
	if !flag {
		controller.WriteError(httpResponse,
			"340000",
			"monitor",
			"[ERROR]The gateway monitor information does not exist!",
			err)
		return

	}
	monitorInfo := map[string]interface{}{
		"statusCode":         "000000",
		"type":               "monitor",
		"baseInfo":           result.BaseInfo,
		"gatewayRequestInfo": result.GatewayRequestInfo,
		"proxyRequestInfo":   result.ProxyInfo,
	}
	info, _ := json.Marshal(monitorInfo)

	httpResponse.Write(info)
	return

}
