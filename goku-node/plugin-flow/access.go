package plugin_flow

import (
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/goku-node/common"
	plugin_manager "github.com/eolinker/goku/goku-node/manager/plugin-manager"
	entity "github.com/eolinker/goku/server/entity/node-entity"
	"reflect"
	"time"
)

var (
	authNames = map[string]string{
		"Oauth2": "goku-oauth2_auth",
		"Apikey": "goku-apikey_auth",
		"Basic":  "goku-basic_auth",
		"Jwt":    "goku-jwt_auth",
	}
	authPluginNames = map[string]string{
		"goku-oauth2_auth": "Oauth2",
		"goku-apikey_auth": "Apikey",
		"goku-basic_auth":  "Basic",
		"goku-jwt_auth":    "Jwt",
	}
)

func getPluginNameByType(authType string) (string, bool) {
	name, has := authNames[authType]
	return name, has
}

// 执行插件的Access函数
func AccessFunc(ctx *common.Context, handleFunc []*entity.PluginHandlerExce) (bool, int) {
	requestId := ctx.RequestId()
	authType := ctx.Request().GetHeader("Authorization-Type")
	authName, _ := getPluginNameByType(authType)
	defer func(ctx *common.Context) {
		log.Debug(requestId, " access plugin default: begin")
		for _, handler := range plugin_manager.GetDefaultPlugins() {
			if handler.PluginObj.Access == nil || reflect.ValueOf(handler.PluginObj.Access).IsNil() {
				continue
			}
			ctx.SetPlugin(handler.Name)
			log.Info(requestId, " access plugin:", handler.Name)
			now := time.Now()
			_, err := handler.PluginObj.Access.Access(ctx)
			log.Debug(requestId, " access plugin:", handler.Name, " Duration", time.Since(now))
			if err != nil {
				log.Warn(requestId, " access plugin:", handler.Name, " error:", err.Error())
			}
		}
		log.Debug(requestId, " access plugin default: end")
	}(ctx)
	isAuthSucess := false
	isNeedAuth := false

	log.Debug(requestId, " access plugin auth check: begin")
	for _, handler := range handleFunc {
		if _, has := authPluginNames[handler.Name]; has {
			isNeedAuth = true
			if handler.Name != authName {
				continue
			}
			if handler.PluginObj.Access == nil || reflect.ValueOf(handler.PluginObj.Access).IsNil() {
				continue
			}
			ctx.SetPlugin(handler.Name)
			log.Debug(requestId, " access plugin:", handler.Name, " begin")
			now := time.Now()
			flag, err := handler.PluginObj.Access.Access(ctx)
			log.Debug(requestId, " access plugin:", handler.Name, " Duration", time.Since(now))
			if flag == false {
				// 校验失败
				if err != nil {
					log.Warn(requestId, " access auth:[", handler.Name, "] error:", err.Error())
				}
				log.Info(requestId, " auth [", authName, "] refuse")

				return false, 0
			}
			log.Debug(requestId, " auth [", authName, "] pass")
			isAuthSucess = true
		}
	}
	log.Debug(requestId, " access plugin auth check: end")
	// 需要校验但是没有执行校验
	if isNeedAuth && !isAuthSucess {
		log.Warn(requestId, " Illegal authorization type:", authType)
		ctx.SetStatus(403, "403")
		ctx.SetBody([]byte("[ERROR]Illegal authorization type!"))
		return false, 0
	}
	lastIndex := 0
	log.Debug(requestId, " access plugin : begin")
	// 执行校验以外的插件
	for index, handler := range handleFunc {
		lastIndex = index
		if _, has := authPluginNames[handler.Name]; has {
			continue
		}

		if handler.PluginObj.Access == nil || reflect.ValueOf(handler.PluginObj.Access).IsNil() {
			continue
		}

		ctx.SetPlugin(handler.Name)
		log.Debug(requestId, " access plugin:", handler.Name)
		now := time.Now()
		flag, err := handler.PluginObj.Access.Access(ctx)
		log.Debug(requestId, " access plugin:", handler.Name, " Duration:", time.Since(now))
		if err != nil {
			log.Warn(requestId, " access plugin:", handler.Name, " error:", err.Error())
		}
		if flag == false && handler.IsStop {
			log.Info(requestId, " access plugin:", handler.Name, " stop")
			return false, index
		}
		log.Debug(requestId, " access plugin:", handler.Name, " continue")
	}
	log.Debug(requestId, " access plugin : end")
	return true, lastIndex
}
