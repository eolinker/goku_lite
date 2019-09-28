package strategyapipluginmanager

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"sort"
	"strconv"
	"sync"

	plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/plugin-manager"
	strategy_plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/strategy-plugin-manager"
	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_strategy "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func init() {
	updater.Add(loadStategyAPIPlugin, 8, "goku_conn_plugin_strategy", "goku_conn_plugin_api", "goku_conn_strategy_api", "goku_gateway_strategy", "goku_gateway_api", "goku_plugin")
}

var (
	pluginsOfAPI = make(map[string][]*entity.PluginHandlerExce)
	locker       sync.RWMutex
)

func get(strategyID string, apiID int) ([]*entity.PluginHandlerExce, bool) {
	locker.RLock()
	defer locker.RUnlock()
	key := strategyID + ":" + strconv.Itoa(apiID)
	p, has := pluginsOfAPI[key]
	return p, has

}

//GetPluginsOfAPI 获取接口插件
func GetPluginsOfAPI(strategyID string, apiID int) ([]*entity.PluginHandlerExce, bool) {
	p, has := get(strategyID, apiID)
	if !has {
		return strategy_plugin_manager.GetPluginsOfStrategy(strategyID)
	}
	return p, true
}

func reset(plugins map[string][]*entity.PluginHandlerExce) {
	for name := range plugins {
		sort.Sort(sort.Reverse(entity.PluginSlice(plugins[name])))
	}

	locker.Lock()
	defer locker.Unlock()

	pluginsOfAPI = plugins

}

func loadStategyAPIPlugin() {
	plugins, err := dao_strategy.GetAPIPlugin()
	if err != nil {
		return
	}

	phsa := make(map[string]map[int][]*entity.PluginHandlerExce)
	for _, p := range plugins {
		apiID, _ := strconv.Atoi(p.APIId)

		phs, has := phsa[p.StrategyID]
		if !has {
			phs = make(map[int][]*entity.PluginHandlerExce)
		}
		phsa[p.StrategyID] = phs

		handle := plugin_manager.GetPluginHandle(p.PluginName)
		if handle == nil {
			continue
		}

		obj, err := handle.Factory.Create(p.PluginConfig, nodecommon.ClusterName(), p.UpdateTag, p.StrategyID, apiID)
		if err != nil {
			continue
		}
		handleExec := &entity.PluginHandlerExce{
			PluginObj: obj,
			Priority:  handle.Info.Priority,
			Name:      p.PluginName,
			IsStop:    handle.Info.IsStop,
		}
		list, has := phsa[p.StrategyID][apiID]
		if !has {
			list = make([]*entity.PluginHandlerExce, 0)
			phsa[p.StrategyID][apiID] = list
		}
		phsa[p.StrategyID][apiID] = append(list, handleExec)

	}

	phl := make(map[string][]*entity.PluginHandlerExce)
	for strategyID, pla := range phsa {
		pls, ok := strategy_plugin_manager.GetPluginsOfStrategy(strategyID)

		for apiID, list := range pla {
			key := fmt.Sprintf("%s:%d", strategyID, apiID)

			if ok {
				list = append(list, pls...)
			}
			phl[key] = list

		}
	}

	reset(phl)
}
