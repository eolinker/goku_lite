package api

import (
	"net/http"
	"strings"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/api"
)

func BatchSetBalanceApi(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiIDList := httpRequest.PostFormValue("apiIDList")
	balanceName := httpRequest.PostFormValue("balance")

	result, err := api.BatchEditApiBalance(strings.Split(apiIDList, ","), balanceName)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", result, err)
		return
	}
	controller.WriteResultInfo(httpResponse, "api", "", nil)
	return
}
