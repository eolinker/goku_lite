package discovery

import (
	"fmt"
	"net/http"

	"github.com/eolinker/goku/common/auto-form"
	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/service"
	driver2 "github.com/eolinker/goku/server/driver"
)

func edit(w http.ResponseWriter, r *http.Request) {
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationEDIT)
	if err != nil {
		return
	}
	if err != r.ParseForm() {
		controller.WriteError(w, "260000", "serviceDiscovery", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}

	param := new(service.AddParam)
	err = auto.SetValues(r.PostForm, param)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] %s", err.Error()), err)
		return
	}

	d, has := driver2.Get(param.Driver)
	if !has {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] invalid  [driver]"), nil)
		return
	}

	if d.Type == driver2.Discovery {

		if param.Config == "" {
			controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] invalid  [driver]"), nil)
			return
		}
	}

	err = service.Save(param)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[mysql]:%s", err.Error()), err)
		return
	}

	controller.WriteResultInfo(w, "serviceDiscovery", "data", nil)
}
