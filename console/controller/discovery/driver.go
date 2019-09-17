package discovery

import (
	"net/http"

	"github.com/eolinker/goku/console/controller"
	driver2 "github.com/eolinker/goku/server/driver"
)

func getDrivices(w http.ResponseWriter, r *http.Request) {

	ds := driver2.GetByType(driver2.Discovery)
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationREAD)
	if err != nil {
		return
	}

	controller.WriteResultInfo(w, "serviceDiscovery", "data", ds)

}
