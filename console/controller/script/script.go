package script

import (
	"net/http"

	"github.com/eolinker/goku/console/controller"

	"github.com/eolinker/goku/console/module/script"
)

var initTables = []string{"goku_gateway", "goku_plugin", "goku_balance", "goku_gateway_api", "goku_gateway_strategy", "goku_conn_plugin_strategy", "goku_conn_plugin_api", "goku_conn_strategy_api"}

// 新建项目
func RefreshApiInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	hashCode := httpRequest.PostFormValue("hashCode")

	if hashCode != "cf70df9de35d556cb5eea88e422ec6cb" {

		controller.WriteError(httpResponse, "800000", "script", "[Error] Illegal hashCode", nil)
		return
	}
	if !script.RefreshApiInfo() {

		controller.WriteError(httpResponse, "800000", "script", "[Error] Fail to refresh!", nil)

	}

	controller.WriteResultInfo(httpResponse, "script", "", nil)
}

// 刷新网关告警信息
func RefreshGatewayAlertConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	hashCode := httpRequest.PostFormValue("hashCode")

	if hashCode != "cf70df9de35d556cb5eea88e422ec6cb" {

		controller.WriteError(httpResponse, "800000", "script", "[Error] Illegal hashCode", nil)

	} else {
		if !script.RefreshGatewayAlertConfig() {

			controller.WriteError(httpResponse, "800000", "script", "[Error] Fail to refresh!", nil)

		}
	}
	controller.WriteResultInfo(httpResponse, "script", "", nil)
}

func UpdateTables() {
	script.UpdateTables(initTables)
}
