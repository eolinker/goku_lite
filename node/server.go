package node

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"

	redis_plugin_proxy "github.com/eolinker/goku-api-gateway/goku-node/redis-plugin-proxy"
	goku_plugin "github.com/eolinker/goku-plugin"
)

func InitPluginUtils() {
	goku_plugin.SetRedisManager(redis_plugin_proxy.Create())
	goku_plugin.InitLog(log.GetLogger())
	//goku_plugin.SetLog(new(log_plugin_proxy.LoggerGeneral))
}
