package goku_node

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	config_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/config-manager"
	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	monitor_write "github.com/eolinker/goku-api-gateway/server/monitor/monitor-write"

	goku_plugin "github.com/eolinker/goku-plugin"
	redis_plugin_proxy "github.com/eolinker/goku-api-gateway/goku-node/redis-plugin-proxy"

	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
	"strings"
)

func InitPluginUtils() {
	goku_plugin.SetRedisManager(redis_plugin_proxy.Create())
	goku_plugin.InitLog(log.GetLogger())
	//goku_plugin.SetLog(new(log_plugin_proxy.LoggerGeneral))
}
func InitDiscovery() {

	all := discovery.AllDrivers()
	log.Infof("install service discovery driver:[%s]\n", strings.Join(all, ","))
}
func InitLog() {
	config_manager.InitLog()
}
func InitServer() {
	log.Debug("init InitServer start")

	InitPluginUtils()

	log.Debug("init InitPluginUtils done")
	InitDiscovery()
	log.Debug("init InitDiscovery done")
	// 注册自动更新，并保证第一次全量拉取完数据
	updater.InitUpdate()

	//// 执行插件初始化函数
	// 插件初始化放在 plugin-manager中
	log.Debug("init updater.InitUpdate done")

	monitor_write.InitMonitorWrite(node_common.ClusterName())
	log.Debug("init server done")
}
