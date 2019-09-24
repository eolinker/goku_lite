package strategy

import (
	"strings"

	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

// 新增策略组插件
func AddPluginToStrategy(pluginName, config, strategyID string) (bool, interface{}, error) {
	tableName := "goku_conn_plugin_strategy"
	flag, result, err := console_mysql.AddPluginToStrategy(pluginName, config, strategyID)
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.UpdateStrategyTagByPluginName(strategyID, pluginName)
	}
	return flag, result, err
}

// 新增策略组插件配置
func EditStrategyPluginConfig(pluginName, config, strategyID string) (bool, string, error) {
	flag, result, err := console_mysql.EditStrategyPluginConfig(pluginName, config, strategyID)
	tableName := "goku_conn_plugin_strategy"
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.UpdateStrategyTagByPluginName(strategyID, pluginName)
	}
	return flag, result, err
}

// 批量修改策略组插件状态
func BatchEditStrategyPluginStatus(connIDList, strategyID string, pluginStatus int) (bool, string, error) {
	plugins := []string{}
	flag, _ := console_mysql.CheckStrategyPluginIsExistByConnIDList(connIDList, "goku-rate_limiting")
	if flag {
		plugins = append(plugins, "goku-rate_limiting")
	}
	flag, _ = console_mysql.CheckStrategyPluginIsExistByConnIDList(connIDList, "goku-replay_attack_defender")
	if flag {
		plugins = append(plugins, "goku-replay_attack_defender")
	}

	tableName := "goku_conn_plugin_strategy"
	flag, result, err := console_mysql.BatchEditStrategyPluginStatus(connIDList, strategyID, pluginStatus)
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.UpdateStrategyTagByPluginName(strategyID, strings.Join(plugins, ","))
	}
	return flag, result, err
}

// 批量删除策略组插件
func BatchDeleteStrategyPlugin(connIDList, strategyID string) (bool, string, error) {
	tableName := "goku_conn_plugin_strategy"
	flag, result, err := console_mysql.BatchDeleteStrategyPlugin(connIDList, strategyID)
	if flag {
		dao.UpdateTable(tableName)
	}
	return flag, result, err
}

// GetStrategyPluginList 获取策略插件列表
func GetStrategyPluginList(strategyID, keyword string, condition int) (bool, []map[string]interface{}, error) {
	return console_mysql.GetStrategyPluginList(strategyID, keyword, condition)
}

// 通过策略组ID获取配置信息
func GetStrategyPluginConfig(strategyID, pluginName string) (bool, string, error) {
	return console_mysql.GetStrategyPluginConfig(strategyID, pluginName)
}

// 检查策略组是否绑定插件
func CheckPluginIsExistInStrategy(strategyID, pluginName string) (bool, error) {
	return console_mysql.CheckPluginIsExistInStrategy(strategyID, pluginName)
}

// 检查策略组插件是否开启
func GetStrategyPluginStatus(strategyID, pluginName string) (bool, error) {
	return console_mysql.GetStrategyPluginStatus(strategyID, pluginName)
}

// 获取Connid
func GetConnIDFromStrategyPlugin(pluginName, strategyID string) (bool, int, error) {
	return console_mysql.GetConnIDFromStrategyPlugin(pluginName, strategyID)
}
