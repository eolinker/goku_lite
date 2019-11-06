package updater

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"

	"github.com/eolinker/goku-api-gateway/console/module/updater"
)

//IsTableExist 检查table是否存在
func IsTableExist(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	exist := updater.IsTableExist(name)
	controller.WriteResultInfo(httpResponse, "updater", "exist", exist)
	return
}

//IsColumnExist 检查列是否存在
func IsColumnExist(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	column := httpRequest.Form.Get("column")
	exist := updater.IsColumnExist(name, column)
	controller.WriteResultInfo(httpResponse, "updater", "exist", exist)
	return
}
