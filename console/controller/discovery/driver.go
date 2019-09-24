package discovery

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	driver2 "github.com/eolinker/goku-api-gateway/server/driver"
)

func getDrivices(w http.ResponseWriter, r *http.Request) {

	ds := driver2.GetByType(driver2.Discovery)
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationREAD)
	if err != nil {
		return
	}

	controller.WriteResultInfo(w, "serviceDiscovery", "data", ds)

}
