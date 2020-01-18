package discovery

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	driver2 "github.com/eolinker/goku-api-gateway/server/driver"
	"net/http"
)

func getDrivices(w http.ResponseWriter, r *http.Request) {

	ds := driver2.GetByType(driver2.Discovery)
	controller.WriteResultInfo(w, "serviceDiscovery", "data", ds)

}
