package api

import (
	"net/http"
	"strings"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
)

//BatchSetBalanceAPI 批量设置接口负载
func BatchSetBalanceAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiIDList := httpRequest.PostFormValue("apiIDList")
	balanceName := httpRequest.PostFormValue("balance")

	result, err := api.BatchEditAPIBalance(strings.Split(apiIDList, ","), balanceName)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", result, err)
		return
	}
	controller.WriteResultInfo(httpResponse, "api", "", nil)
	return
}
