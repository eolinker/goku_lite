package plugin_manager

import (
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

// 获取默认插件列表
func GetDefaultPlugins() []*entity.PluginHandlerExce {
	return defaultPlugins
}

// 获取默认插件列表
func GetBeforPlugins() []*entity.PluginHandlerExce {
	return beforPlugins
}

func Check(name string) (int, error) {
	return globalPluginManager.check(name)
}
