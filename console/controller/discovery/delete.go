package discovery

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/service"
	dao_service2 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-service"
	"net/http"
	"strings"
)

func delete(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm() ; err!= nil{
		controller.WriteError(w, "260000", "serviceDiscovery", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}
	nameStr := r.FormValue("names")

	names := strings.Split(nameStr, ",")

	err := service.Delete(names)

	if err != nil {
		if en, ok := err.(dao_service2.DeleteError); ok {
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
