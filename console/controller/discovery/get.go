package discovery

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/service"
	"net/http"
)

func getInfo(w http.ResponseWriter, r *http.Request) {


	name := r.URL.Query().Get("name")
	if !service.ValidateName(name) {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] invalid  [name]=%s", name), nil)
		return
	}

	info, err := service.Get(name)
	if err != nil {
		controller.WriteError(w, "71002", "data", fmt.Sprintf("[param_check] error:%s ", err.Error()), err)
		return
	}

	info.Decode()
	controller.WriteResultInfo(w, "serviceDiscovery", "data", info)

}
