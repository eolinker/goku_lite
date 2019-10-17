package strategy

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//AddPluginToStrategy 新增策略组插件
func AddPluginToStrategy(pluginName, config, strategyID string) (bool, interface{}, error) {
	flag, result, err := console_sqlite3.AddPluginToStrategy(pluginName, config, strategyID)

	return flag, result, err
}

//EditStrategyPluginConfig 新增策略组插件配置
func EditStrategyPluginConfig(pluginName, config, strategyID string) (bool, string, error) {
	flag, result, err := console_sqlite3.EditStrategyPluginConfig(pluginName, config, strategyID)

	return flag, result, err
}

//BatchEditStrategyPluginStatus 批量修改策略组插件状态
func BatchEditStrategyPluginStatus(connIDList, strategyID string, pluginStatus int) (bool, string, error) {

	flag, result, err := console_sqlite3.BatchEditStrategyPluginStatus(connIDList, strategyID, pluginStatus)

	return flag, result, err
}

//BatchDeleteStrategyPlugin 批量删除策略组插件
func BatchDeleteStrategyPlugin(connIDList, strategyID string) (bool, string, error) {
	flag, result, err := console_sqlite3.BatchDeleteStrategyPlugin(connIDList, strategyID)

	return flag, result, err
}

// GetStrategyPluginList 获取策略插件列表
func GetStrategyPluginList(strategyID, keyword string, condition int) (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetStrategyPluginList(strategyID, keyword, condition)
}

//GetStrategyPluginConfig 通过策略组ID获取配置信息
func GetStrategyPluginConfig(strategyID, pluginName string) (bool, string, error) {
	return console_sqlite3.GetStrategyPluginConfig(strategyID, pluginName)
}

//CheckPluginIsExistInStrategy 检查策略组是否绑定插件
func CheckPluginIsExistInStrategy(strategyID, pluginName string) (bool, error) {
	return console_sqlite3.CheckPluginIsExistInStrategy(strategyID, pluginName)
}

//GetStrategyPluginStatus 检查策略组插件是否开启
func GetStrategyPluginStatus(strategyID, pluginName string) (bool, error) {
	return console_sqlite3.GetStrategyPluginStatus(strategyID, pluginName)
}

//GetConnIDFromStrategyPlugin 获取Connid
func GetConnIDFromStrategyPlugin(pluginName, strategyID string) (bool, int, error) {
	return console_sqlite3.GetConnIDFromStrategyPlugin(pluginName, strategyID)
}
