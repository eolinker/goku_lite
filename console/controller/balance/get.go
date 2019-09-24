package balance

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/balance"
)

// 获取负载信息
func GetBalanceInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationLoadBalance, controller.OperationREAD)
	if e != nil {
		return
	}
	if err := httpRequest.ParseForm(); err != nil {
		controller.WriteError(httpResponse, "501", "balance", "[ERROR]参数解析错误t!", err)
		return
	}
	balanceName := httpRequest.FormValue("balanceName")
	info, err := balance.Get(balanceName)

	if err != nil {
		controller.WriteError(httpResponse, "260000", "balance", "[ERROR]The balance does not exist!", err)
		return
	}

	controller.WriteResultInfo(httpResponse, "balance", "balanceInfo", info)

	return
}
