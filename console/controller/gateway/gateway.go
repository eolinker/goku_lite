package gateway

import (
	"encoding/json"
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/gateway"
)

const operationGateway = "gateway"

//Handlers hendlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/getSummaryInfo": factory.NewAccountHandleFunction(operationGateway, false, GetGatewayBasicInfo),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

//GetGatewayBasicInfo 获取网关基本信息
func GetGatewayBasicInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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
