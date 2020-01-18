package dao_version_config

import (
	"strconv"

	"github.com/eolinker/goku-api-gateway/config"
)

//GetGlobalPlugin 获取全局插件
func (d *VersionConfigDao) GetGlobalPlugin() (*config.GatewayPluginConfig, error) {
	db := d.db
	sql := "SELECT pluginName,isStop,IFNULL(pluginConfig,''),pluginType FROM goku_plugin"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pluginConfigs := config.GatewayPluginConfig{
		BeforePlugins: make([]*config.PluginConfig, 0, 20),
		GlobalPlugins: make([]*config.PluginConfig, 0, 20),
	}
	for rows.Next() {
		var pluginName, pluginConfig string
		var isStop bool
		var pluginType int
		err = rows.Scan(&pluginName, &isStop, &pluginConfig, &pluginType)
		if err != nil {
			return nil, err
		}
		if pluginType == 0 {
			pluginConfigs.GlobalPlugins = append(pluginConfigs.GlobalPlugins, &config.PluginConfig{
				Name:   pluginName,
				IsStop: isStop,
				Config: pluginConfig,
			})
		} else {
			pluginConfigs.BeforePlugins = append(pluginConfigs.BeforePlugins, &config.PluginConfig{
				Name:   pluginName,
				IsStop: isStop,
				Config: pluginConfig,
			})
		}
	}
	return &pluginConfigs, nil
}

//GetAPIPlugins 获取接口插件
func (d *VersionConfigDao) GetAPIPlugins() (map[string][]*config.PluginConfig, error) {
	db := d.db
	sql := "SELECT goku_conn_plugin_api.apiID,goku_conn_plugin_api.strategyID,goku_conn_plugin_api.pluginName,goku_conn_plugin_api.pluginConfig,goku_plugin.isStop FROM goku_conn_plugin_api INNER JOIN goku_plugin ON goku_conn_plugin_api.pluginName = goku_plugin.pluginName"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pluginMaps := make(map[string][]*config.PluginConfig)
	for rows.Next() {
		var apiID int
		var isStop bool
		var pluginName, pluginConfig, strategyID string
		err = rows.Scan(&apiID, &strategyID, &pluginName, &pluginConfig, &isStop)
		if err != nil {
			return nil, err
		}
		key := strategyID + ":" + strconv.Itoa(apiID)
		if _, ok := pluginMaps[key]; !ok {
			pluginMaps[key] = make([]*config.PluginConfig, 0, 20)
		}
		pluginMaps[key] = append(pluginMaps[key], &config.PluginConfig{
			Name:   pluginName,
			IsStop: isStop,
			Config: pluginConfig,
		})
	}
	return pluginMaps, nil

}

//GetStrategyPlugins 获取策略插件
func (d *VersionConfigDao) GetStrategyPlugins() (map[string][]*config.PluginConfig, map[string]map[string]string, error) {
	db := d.db
	sql := "SELECT goku_conn_plugin_strategy.strategyID,goku_conn_plugin_strategy.pluginName,goku_conn_plugin_strategy.pluginConfig,goku_plugin.isStop FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_conn_plugin_strategy.pluginName = goku_plugin.pluginName WHERE goku_plugin.pluginStatus = 1 AND goku_conn_plugin_strategy.pluginStatus = 1"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	pluginMaps := make(map[string][]*config.PluginConfig)
	authMaps := make(map[string]map[string]string)
	for rows.Next() {
		var isStop bool
		var pluginName, pluginConfig, strategyID string
		err = rows.Scan(&strategyID, &pluginName, &pluginConfig, &isStop)
		if err != nil {
			return nil, nil, err
		}
		key := strategyID
		if _, ok := pluginMaps[key]; !ok {
			pluginMaps[key] = make([]*config.PluginConfig, 0, 20)
		}
		if v, ok := autoAuthNames[pluginName]; ok {
			if _, ok := authMaps[key]; !ok {
				authMaps[key] = make(map[string]string)
			}
			authMaps[key][v] = pluginConfig
		}

		pluginMaps[key] = append(pluginMaps[key], &config.PluginConfig{
			Name:   pluginName,
			IsStop: isStop,
			Config: pluginConfig,
		})
	}
	return pluginMaps, authMaps, nil

}
