package balance

import (
	"net/http"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/balance"
)

func GetSimpleList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationLoadBalance, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := balance.GetBalancNames()

	if !flag {
		controller.WriteError(httpResponse, "260000,", "balance", "[ERROR]Empty balance list!", err)
		return
	}
	controller.WriteResultInfo(httpResponse, "balance", "balanceNames", result)

	return

}
