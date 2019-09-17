package config_log

import (
	"fmt"
	"net/http"

	"github.com/eolinker/goku/common/auto-form"
	"github.com/eolinker/goku/console/controller"
	module "github.com/eolinker/goku/console/module/config-log"
)

type AccessLogHandler struct {
}

func (h *AccessLogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := controller.CheckLogin(w, r, controller.OperationGatewayConfig, controller.OperationEDIT)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		{
			h.get(w, r)

		}

	case http.MethodPut:
		{
			h.set(w, r)

		}
	default:
		w.WriteHeader(404)
	}
}

func (h *AccessLogHandler) get(w http.ResponseWriter, r *http.Request) {
	config, e := module.GetAccess()
	if e = r.ParseForm(); e != nil {
		controller.WriteError(w, "270000", "data", "[Get]未知错误："+e.Error(), e)
		return
	}

	controller.WriteResultInfo(w,
		"data",
		"data",
		config)

}
func (h *AccessLogHandler) set(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err = r.ParseForm(); err != nil {
		controller.WriteError(w, "260000", "data", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}

	param := new(module.AccessParam)
	err = auto.SetValues(r.Form, param)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] %s", err.Error()), err)
		return
	}
	if param.Dir == "" {
		controller.WriteError(w, "260000", "data", "[param_check] inval dir", err)
		return
	}
	if param.File == "" {
		controller.WriteError(w, "260000", "data", "[param_check] inval file name", err)
		return
	}
	if param.Expire < module.ExpireDefault {
		controller.WriteError(w, "260000", "data", "[param_check] inval expire", nil)
		return
	}
	paramBase, err := param.Format()
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] %s", err.Error()), err)
		return
	}

	err = module.Set(module.AccessLog, paramBase)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[mysql_error] %s", err.Error()), err)
		return
	}
	controller.WriteResultInfo(w,
		"data",
		"",
		nil)
}
