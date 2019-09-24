package cluster

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	cluster2 "github.com/eolinker/goku-api-gateway/server/cluster"
)

func GetClusterList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	list := cluster2.GetList()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		list)

}

func GetClusterInfoList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	all := cluster2.GetAll()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		all)
}
