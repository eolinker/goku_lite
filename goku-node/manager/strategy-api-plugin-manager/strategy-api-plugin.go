package strategy_api_plugin_manager

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
	updater.Add(loadStategyApiPlugin, 8, "goku_conn_plugin_strategy", "goku_conn_plugin_api", "goku_conn_strategy_api", "goku_gateway_strategy", "goku_gateway_api", "goku_plugin")
}

var (
	pluginsOfApi = make(map[string][]*entity.PluginHandlerExce)
	locker       sync.RWMutex
)

func get(strategyId string, apiId int) ([]*entity.PluginHandlerExce, bool) {
	locker.RLock()
	defer locker.RUnlock()
	key := strategyId + ":" + strconv.Itoa(apiId)
	p, has := pluginsOfApi[key]
	return p, has

}
func GetPluginsOfApi(strategyId string, apiId int) ([]*entity.PluginHandlerExce, bool) {
	p, has := get(strategyId, apiId)
	if !has {
		return strategy_plugin_manager.GetPluginsOfStrategy(strategyId)
	}
	return p, true
}

func reset(plugins map[string][]*entity.PluginHandlerExce) {
	for name := range plugins {
		sort.Sort(sort.Reverse(entity.PluginSlice(plugins[name])))
	}

	locker.Lock()
	defer locker.Unlock()

	pluginsOfApi = plugins

}

func loadStategyApiPlugin() {
	plugins, err := dao_strategy.GetApiPlugin()
	if err != nil {
		return
	}

	phsa := make(map[string]map[int][]*entity.PluginHandlerExce)
	for _, p := range plugins {
		apiId, _ := strconv.Atoi(p.ApiId)

		phs, has := phsa[p.StrategyID]
		if !has {
			phs = make(map[int][]*entity.PluginHandlerExce)
		}
		phsa[p.StrategyID] = phs

		handle := plugin_manager.GetPluginHandle(p.PluginName)
		if handle == nil {
			continue
		}

		obj, err := handle.Factory.Create(p.PluginConfig, node_common.ClusterName(), p.UpdateTag, p.StrategyID, apiId)
		if err != nil {
			continue
		}
		handleExec := &entity.PluginHandlerExce{
			PluginObj: obj,
			Priority:  handle.Info.Priority,
			Name:      p.PluginName,
			IsStop:    handle.Info.IsStop,
		}
		list, has := phsa[p.StrategyID][apiId]
		if !has {
			list = make([]*entity.PluginHandlerExce, 0)
			phsa[p.StrategyID][apiId] = list
		}
		phsa[p.StrategyID][apiId] = append(list, handleExec)

	}

	phl := make(map[string][]*entity.PluginHandlerExce)
	for strategyId, pla := range phsa {
		pls, ok := strategy_plugin_manager.GetPluginsOfStrategy(strategyId)

		for apiId, list := range pla {
			key := fmt.Sprintf("%s:%d", strategyId, apiId)

			if ok {
				list = append(list, pls...)
			}
			phl[key] = list

		}
	}

	reset(phl)
}
