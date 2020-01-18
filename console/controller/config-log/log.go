package config_log

import (
	"fmt"
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/common/auto-form"
	"github.com/eolinker/goku-api-gateway/console/controller"
	module "github.com/eolinker/goku-api-gateway/console/module/config-log"
)

//LogHandlerGet 日志配置获取处理器
type LogHandlerGet struct {
	name string
}

func (h *LogHandlerGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	config, e := module.Get(h.name)
	if e = r.ParseForm(); e != nil {
		controller.WriteError(w, "270000", "data", "[Get]未知错误："+e.Error(), e)
		return
	}

	controller.WriteResultInfo(w,
		"data",
		"data",
		config)

}

//LogHandlerSet 设置日志处理器
type LogHandlerSet struct {
	name string
}

func (h *LogHandlerSet) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err = r.ParseForm(); err != nil {
		controller.WriteError(w, "260000", "data", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}

	param := new(module.PutParam)
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
	err = module.Set(h.name, paramBase)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[db_error] %s", err.Error()), err)
		return
	}
	controller.WriteResultInfo(w,
		"data",
		"",
		nil)
}

//NewLogHandler 日志handler
func NewLogHandler(name string, factory *goku_handler.AccountHandlerFactory) http.Handler {
	return &LogHandler{
		getHandler: factory.NewAccountHandler(operationLog, false, &LogHandlerGet{name: name}),
		setHandler: factory.NewAccountHandler(operationLog, true, &LogHandlerSet{name: name}),
	}
}
