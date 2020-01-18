package config_log

import (
	"fmt"
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/common/auto-form"
	"github.com/eolinker/goku-api-gateway/console/controller"
	module "github.com/eolinker/goku-api-gateway/console/module/config-log"
)

//LogHandler access日志处理器
type LogHandler struct {
	getHandler http.Handler
	setHandler http.Handler
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		{
			h.getHandler.ServeHTTP(w, r)
		}

	case http.MethodPut:
		{
			h.setHandler.ServeHTTP(w, r)
		}
	default:
		w.WriteHeader(404)
	}
}

//AccessLogGet 获取access日志配置
func AccessLogGet(w http.ResponseWriter, r *http.Request) {
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

//AccessLogSet 设置access日志内容
func AccessLogSet(w http.ResponseWriter, r *http.Request) {
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
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[db_error] %s", err.Error()), err)
		return
	}
	controller.WriteResultInfo(w,
		"data",
		"",
		nil)
}

//NewAccessHandler accessHandler
func NewAccessHandler(factory *goku_handler.AccountHandlerFactory) http.Handler {
	return &LogHandler{
		getHandler: factory.NewAccountHandleFunction(operationLog, false, AccessLogGet),
		setHandler: factory.NewAccountHandleFunction(operationLog, true, AccessLogSet),
	}
}
