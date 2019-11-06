package gateway

import (
	"encoding/json"
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/gateway"
)

//GetGatewayBasicInfo 获取网关基本信息
func GetGatewayBasicInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := gateway.GetGatewayMonitorSummaryByPeriod()
	if !flag {
		controller.WriteError(httpResponse,
			"340000",
			"monitor",
			"[ERROR]The gateway basic information does not exist!",
			err)
		return

	}
	monitorInfo := map[string]interface{}{
		"statusCode": "000000",
		"type":       "monitor",
		"baseInfo":   result.BaseInfo,
	}
	info, _ := json.Marshal(monitorInfo)

	httpResponse.Write(info)
	return

}
