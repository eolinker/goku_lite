package strategypluginmanager

import (
	"github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"sort"
	"sync"

	plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/plugin-manager"
	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_strategy "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func init() {
	updater.Add(loadStategyPlugin, 7, "goku_conn_plugin_strategy", "goku_gateway_strategy", "goku_plugin")
}

var (
	pluginsOfStrategy = make(map[string][]*entity.PluginHandlerExce)
	locker            sync.RWMutex
)

//GetPluginsOfStrategy 获取策略插件
func GetPluginsOfStrategy(strategyID string) ([]*entity.PluginHandlerExce, bool) {
	locker.RLock()
	defer locker.RUnlock()
	p, has := pluginsOfStrategy[strategyID]
	if !has {
		return nil, false
	}
	return p, true
}

func reset(plugins map[string][]*entity.PluginHandlerExce) {
	for name := range plugins {
		//plugins[name] = append(list,def...)
		sort.Sort(sort.Reverse(entity.PluginSlice(plugins[name])))
	}
	locker.Lock()
	defer locker.Unlock()

	pluginsOfStrategy = plugins

}

func loadStategyPlugin() {
	plugins, err := dao_strategy.GetAllStrategyPluginList()
	if err != nil {
		return
	}
	phl := make(map[string][]*entity.PluginHandlerExce)
	for _, p := range plugins {
		if _, ok := phl[p.StrategyID]; !ok {
			phl[p.StrategyID] = make([]*entity.PluginHandlerExce, 0)
		}
		handle := plugin_manager.GetPluginHandle(p.PluginName)
		if handle == nil {
			continue
		}
		excer, err := handle.Factory.Create(p.PluginConfig, nodecommon.ClusterName(), p.UpdateTag, p.StrategyID, 0)
		if err != nil {

			continue
		}
		handleExec := &entity.PluginHandlerExce{
			PluginObj: excer,
			Name:      p.PluginName,
			Priority:  handle.Info.Priority,
			IsStop:    handle.Info.IsStop,
		}
		phl[p.StrategyID] = append(phl[p.StrategyID], handleExec)
	}
	reset(phl)
}
