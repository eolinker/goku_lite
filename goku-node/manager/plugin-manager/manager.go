package plugin_manager

import (
	"sort"
	"sync"

	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

// 写插件管理包对外接口
var (
	pluginHandles  = make(map[string]*entity.PluginFactoryHandler)
	defaultPlugins []*entity.PluginHandlerExce
	locker         = sync.RWMutex{}

	beforPlugins []*entity.PluginHandlerExce
)

// 获取单一插件handle
func GetPluginHandle(name string) *entity.PluginFactoryHandler {
	locker.RLock()
	handle := pluginHandles[name]
	locker.RUnlock()

	return handle
}
func reset(pis map[string]*entity.PluginInfo) {
	plugins, def, before := LoadPlugin(pis)

	sort.Sort(sort.Reverse(entity.PluginSlice(def)))
	sort.Sort(sort.Reverse(entity.PluginSlice(before)))
	locker.Lock()
	defer locker.Unlock()

	pluginHandles = plugins
	defaultPlugins = def
	beforPlugins = before
}
