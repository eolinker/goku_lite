package balance

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	"net/http"
	"strings"

	"github.com/eolinker/goku-api-gateway/console/module/balance"
)

//DeleteBalance 删除负载配置
func DeleteBalance(httpResponse http.ResponseWriter, httpRequest *http.Request) {


	if err := httpRequest.ParseForm(); err != nil {
		controller.WriteError(httpResponse, "260000", "data", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}
	balanceName := httpRequest.PostFormValue("balanceName")

	restlt, err := balance.Delete(balanceName)
	if err != nil {

		controller.WriteError(httpResponse, "260000", "balance", restlt, err)
		return

	}

	controller.WriteResultInfo(httpResponse, "balance", "", nil)

	return
}

//BatchDeleteBalance 批量删除负载
func BatchDeleteBalance(httpResponse http.ResponseWriter, httpRequest *http.Request) {


	balanceNames := httpRequest.PostFormValue("balanceNames")
	result, err := balance.BatchDeleteBalance(strings.Split(balanceNames, ","))
	if err != nil {
		controller.WriteError(httpResponse, "260000,", "balance", result, err)
		return
	}

	controller.WriteResultInfo(httpResponse, "balance", "", nil)

	return
}
