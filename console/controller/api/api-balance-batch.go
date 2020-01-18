package api

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
	"net/http"
	"strings"
)

//BatchSetBalanceAPI 批量设置接口负载
func BatchSetBalanceAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {


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
