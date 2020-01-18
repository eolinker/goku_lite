package discovery

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/module/service"
)

func simple(w http.ResponseWriter, r *http.Request) {


	vs, def, err := service.SimpleList()
	if err != nil {
		controller.WriteError(w, "100000", "data", err.Error(), err)
		return
	}
	data := make(map[string]interface{})
	data["default"] = def
	data["list"] = vs
	controller.WriteResultInfoWithPage(w, "serviceDiscovery", "data", data, controller.NewItemNum(len(vs)))
}
