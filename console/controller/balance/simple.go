package balance

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/balance"
	"net/http"
)

//GetSimpleList 获取简易列表
func GetSimpleList(httpResponse http.ResponseWriter, httpRequest *http.Request) {


	flag, result, err := balance.GetBalancNames()

	if !flag {
		controller.WriteError(httpResponse, "260000,", "balance", "[ERROR]Empty balance list!", err)
		return
	}
	controller.WriteResultInfo(httpResponse, "balance", "balanceNames", result)

	return

}
