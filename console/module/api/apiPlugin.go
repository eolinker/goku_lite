package api

import (
	"strconv"
	"strings"

	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//BatchEditAPIPluginStatus 批量修改接口插件状态
func BatchEditAPIPluginStatus(connIDList, strategyID string, pluginStatus, userID int) (bool, string, error) {
	idList := []string{}
	plugins := []string{}
	flag, apiIDList, _ := console_mysql.CheckAPIPluginIsExistByConnIDList(connIDList, "goku-circuit_breaker")
	if flag {
		for _, id := range apiIDList {
			idList = append(idList, strconv.Itoa(id))
		}

		plugins = append(plugins, "goku-circuit_breaker")
	}
	flag, apiIDList, _ = console_mysql.CheckAPIPluginIsExistByConnIDList(connIDList, "goku-proxy_caching")
	if flag {
		for _, id := range apiIDList {
			idList = append(idList, strconv.Itoa(id))
		}
		plugins = append(plugins, "goku-proxy_caching")
	}
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.BatchEditAPIPluginStatus(connIDList, strategyID, pluginStatus, userID)
	if flag {
		dao.UpdateTable(name)
		p := strings.Join(plugins, ",")
		ids := strings.Join(idList, ",")
		console_mysql.UpdateAPITagByPluginName(strategyID, ids, p)
	}
	return flag, result, err
}

//BatchDeleteAPIPlugin 批量删除接口插件
func BatchDeleteAPIPlugin(connIDList, strategyID string) (bool, string, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.BatchDeleteAPIPlugin(connIDList, strategyID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//AddPluginToAPI 新增插件到接口
func AddPluginToAPI(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.AddPluginToAPI(pluginName, config, strategyID, apiID, userID)
	if flag {
		dao.UpdateTable(name)
		console_mysql.UpdateAPITagByPluginName(strategyID, strconv.Itoa(apiID), pluginName)
	}
	return flag, result, err
}

//EditAPIPluginConfig 修改接口插件配置
func EditAPIPluginConfig(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.EditAPIPluginConfig(pluginName, config, strategyID, apiID, userID)
	if flag {
		dao.UpdateTable(name)
		console_mysql.UpdateAPITagByPluginName(strategyID, strconv.Itoa(apiID), pluginName)
	}
	return flag, result, err
}

//GetAPIPluginList 获取接口插件列表
func GetAPIPluginList(apiID int, strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetAPIPluginList(apiID, strategyID)
}

//GetPluginIndex 获取插件优先级
func GetPluginIndex(pluginName string) (bool, int, error) {
	return console_mysql.GetPluginIndex(pluginName)
}

//GetAPIPluginConfig 通过APIID获取配置信息
func GetAPIPluginConfig(apiID int, strategyID, pluginName string) (bool, map[string]string, error) {
	return console_mysql.GetAPIPluginConfig(apiID, strategyID, pluginName)
}

//CheckPluginIsExistInAPI 检查策略组是否绑定插件
func CheckPluginIsExistInAPI(strategyID, pluginName string, apiID int) (bool, error) {
	return console_mysql.CheckPluginIsExistInAPI(strategyID, pluginName, apiID)
}

//GetAllAPIPluginInStrategy 获取策略组中所有接口插件列表
func GetAllAPIPluginInStrategy(strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetAllAPIPluginInStrategy(strategyID)
}

// GetAPIPluginInStrategyByAPIID 获取策略组中所有接口插件列表
func GetAPIPluginInStrategyByAPIID(strategyID string, apiID int, keyword string, condition int) (bool, []map[string]interface{}, map[string]interface{}, error) {
	return console_mysql.GetAPIPluginInStrategyByAPIID(strategyID, apiID, keyword, condition)
}

//GetAPIPluginListWithNotAssignAPIList 获取没有绑定插件的接口列表
func GetAPIPluginListWithNotAssignAPIList(strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetAPIPluginListWithNotAssignAPIList(strategyID)
}

//UpdateAllAPIPluginUpdateTag 更新所有接口插件更新标识
func UpdateAllAPIPluginUpdateTag() error {
	return console_mysql.UpdateAllAPIPluginUpdateTag()
}
