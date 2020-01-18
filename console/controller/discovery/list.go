package discovery

import (
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/service"
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	keyword := r.FormValue("keyword")
	infos, err := service.List(keyword)
	if err != nil {
		controller.WriteError(w, "100000", "data", err.Error(), err)
		return
	}

	controller.WriteResultInfo(w, "serviceDiscovery", "data", infos)

}
