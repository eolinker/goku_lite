package discovery

import (
	"net/http"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/service"
)

func simple(w http.ResponseWriter, r *http.Request) {
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationREAD)
	if err != nil {
		return
	}

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
