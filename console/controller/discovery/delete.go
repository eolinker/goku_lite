package discovery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/service"
	dao_service "github.com/eolinker/goku/server/dao/console-mysql/dao-service"
)

func delete(w http.ResponseWriter, r *http.Request) {
	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationEDIT)
	if err != nil {
		return
	}
	if err != r.ParseForm() {
		controller.WriteError(w, "260000", "serviceDiscovery", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}
	nameStr := r.FormValue("names")

	names := strings.Split(nameStr, ",")

	err = service.Delete(names)

	if err != nil {
		if en, ok := err.(dao_service.DeleteError); ok {
			controller.WriteError(w, "260000", "", fmt.Sprintf("删除[%s]失败，请先从负载中移除对该服务的引用", string(en)), err)
		} else {
			controller.WriteError(w, "260000", "serviceDiscovery", fmt.Sprintf("[error] %s", err.Error()), err)

		}
		return
	}

	controller.WriteResultInfo(w,
		"serviceDiscovery",
		"data",
		nil)

}
