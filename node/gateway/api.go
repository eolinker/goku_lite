package gateway

import (
	"fmt"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/node/gateway/application"
	plugin_executor "github.com/eolinker/goku-api-gateway/node/gateway/plugin-executor"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
)

//API 接口
type API struct {
	strategyID string
	//api *config.APIContent
	app                application.Application
	pluginAccess       []plugin_executor.Executor
	pluginAccessGlobal []plugin_executor.Executor

	pluginProxies       []plugin_executor.Executor
	pluginProxiesGlobal []plugin_executor.Executor

	apiID   int
	apiName string
}

//Router router
func (h *API) Router(ctx *common.Context) {

	ctx.SetAPIID(h.apiID)
	ctx.LogFields[access_field.API] = fmt.Sprintf("\"%d %s\"", h.apiID, h.apiName)

	isAccess := h.accessFlow(ctx)
	h.accessGlobalFlow(ctx)
	if !isAccess {
		return
	}

	h.app.Execute(ctx)

	isproxy := h.proxyFlow(ctx)
	h.proxyGlobalFlow(ctx)

	if !isproxy {
		return
	}

	return
}

func (h *API) accessFlow(ctx *common.Context) bool {
	for _, handler := range h.pluginAccess {
		flag, err := handler.Execute(ctx)
		if err != nil {
			fmt.Println(err)
		}
		if flag == false && handler.IsStop() {

			return false
		}
	}
	return true
}

func (h *API) accessGlobalFlow(ctx *common.Context) {
	// 全局插件不中断
	for _, handler := range h.pluginAccessGlobal {
		_, _ = handler.Execute(ctx)
	}
}

func (h *API) proxyFlow(ctx *common.Context) bool {
	for _, handler := range h.pluginProxies {
		flag, _ := handler.Execute(ctx)

		if flag == false && handler.IsStop() {

			return false
		}
	}
	return true
}

func (h *API) proxyGlobalFlow(ctx *common.Context) {
	// 全局插件不中断
	for _, handler := range h.pluginProxiesGlobal {
		_, _ = handler.Execute(ctx)

	}
}
