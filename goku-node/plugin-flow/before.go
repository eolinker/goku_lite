package plugin_flow

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/plugin-manager"
	"reflect"
	"time"
)

// 执行插件的BeforeMatch函数
func BeforeMatch(ctx *common.Context) bool {
	requestId := ctx.RequestId()
	defer func(ctx *common.Context) {
		log.Debug(requestId, " before plugin default: begin")
		for _, handler := range plugin_manager.GetDefaultPlugins() {
			if handler.PluginObj.BeforeMatch == nil || reflect.ValueOf(handler.PluginObj.BeforeMatch).IsNil() {
				continue
			}
			ctx.SetPlugin(handler.Name)
			log.Debug(requestId, " before plugin :", handler.Name, " start")
			now := time.Now()
			_, err := handler.PluginObj.BeforeMatch.BeforeMatch(ctx)
			log.Debug(requestId, " before plugin :", handler.Name, " Duration:", time.Since(now))
			log.Debug(requestId, " before plugin :", handler.Name, " end")
			if err != nil {
				log.Warn(requestId, " before plugin:", handler.Name, " error:", err.Error())
			}
		}
		log.Debug(requestId, " before plugin default: end")
	}(ctx)
	log.Debug(requestId, " before plugin : begin")
	for _, handler := range plugin_manager.GetBeforPlugins() {

		if handler.PluginObj.BeforeMatch == nil || reflect.ValueOf(handler.PluginObj.BeforeMatch).IsNil() {
			continue
		}

		ctx.SetPlugin(handler.Name)
		log.Debug(requestId, " before plugin :", handler.Name, " start")
		now := time.Now()
		flag, err := handler.PluginObj.BeforeMatch.BeforeMatch(ctx)
		log.Debug(requestId, " before plugin :", handler.Name, " Duration:", time.Since(now))
		log.Debug(requestId, " before plugin :", handler.Name, " end")

		if err != nil {
			log.Warn(requestId, " before plugin:", handler.Name, " error:", err.Error())
		}
		if flag == false {
			if handler.IsStop == true {
				return false
			}
		}
	}
	log.Debug(requestId, " before plugin : end")
	return true
}
