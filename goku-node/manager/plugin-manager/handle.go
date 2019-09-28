package pluginmanager

import (
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

//GetDefaultPlugins 获取默认插件列表
func GetDefaultPlugins() []*entity.PluginHandlerExce {
	return defaultPlugins
}

//GetBeforPlugins 获取默认插件列表
func GetBeforPlugins() []*entity.PluginHandlerExce {
	return beforPlugins
}

//Check check
func Check(name string) (int, error) {
	return globalPluginManager.check(name)
}
