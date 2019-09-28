package pluginflow

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/plugin-manager"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"

	"reflect"
	"time"
)

//ProxyFunc 执行插件的Proxy函数
func ProxyFunc(ctx *common.Context, handleFunc []*entity.PluginHandlerExce) (bool, int) {
	requestID := ctx.RequestId()
	defer func(ctx *common.Context) {
		log.Debug(requestID, " Proxy plugin default: begin")
		for _, handler := range plugin_manager.GetDefaultPlugins() {
			if handler.PluginObj.Proxy == nil || reflect.ValueOf(handler.PluginObj.Proxy).IsNil() {
				continue
			}
			ctx.SetPlugin(handler.Name)
			log.Debug(requestID, " Proxy plugin :", handler.Name, " start")
			now := time.Now()
			_, err := handler.PluginObj.Proxy.Proxy(ctx)
			log.Debug(requestID, " Proxy plugin :", handler.Name, " Duration:", time.Since(now))
			log.Debug(requestID, " Proxy plugin :", handler.Name, " end")

			if err != nil {
				log.Warn(requestID, " Proxy plugin:", handler.Name, " error:", err.Error())
			}
		}
		log.Debug(requestID, " Proxy plugin default: begin")
	}(ctx)
	lastIndex := 0
	log.Debug(requestID, " Proxy plugin : begin")
	for index, handler := range handleFunc {
		lastIndex = index
		if handler.PluginObj.Proxy == nil || reflect.ValueOf(handler.PluginObj.Proxy).IsNil() {
			continue
		}

		ctx.SetPlugin(handler.Name)
		log.Debug(requestID, " Proxy plugin :", handler.Name, " start")
		now := time.Now()
		flag, err := handler.PluginObj.Proxy.Proxy(ctx)
		log.Debug(requestID, " Proxy plugin :", handler.Name, " Duration:", time.Since(now))
		log.Debug(requestID, " Proxy plugin :", handler.Name, " end")

		if err != nil {
			log.Warn(requestID, " Proxy plugin :", handler.Name, " error: ", err.Error())
		}
		if flag == false && handler.IsStop == true {

			return false, lastIndex
		}
	}
	log.Debug(requestID, " Proxy plugin : end")
	return true, lastIndex
}
