package discovery

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/service"
)

func list(w http.ResponseWriter, r *http.Request) {
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationREAD)
	if err != nil {
		return
	}

	_ = r.ParseForm()

	keyword := r.FormValue("keyword")
	infos, err := service.List(keyword)
	if err != nil {
		controller.WriteError(w, "100000", "data", err.Error(), err)
		return
	}

	controller.WriteResultInfo(w, "serviceDiscovery", "data", infos)

}
