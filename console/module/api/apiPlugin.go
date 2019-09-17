package api

import (
	"strconv"
	"strings"

	"github.com/eolinker/goku/server/dao"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
)

// 批量修改接口插件状态
func BatchEditApiPluginStatus(connIDList, strategyID string, pluginStatus, userID int) (bool, string, error) {
	idList := []string{}
	plugins := []string{}
	flag, apiIDList, _ := console_mysql.CheckApiPluginIsExistByConnIDList(connIDList, "goku-circuit_breaker")
	if flag {
		for _, id := range apiIDList {
			idList = append(idList, strconv.Itoa(id))
		}

		plugins = append(plugins, "goku-circuit_breaker")
	}
	flag, apiIDList, _ = console_mysql.CheckApiPluginIsExistByConnIDList(connIDList, "goku-proxy_caching")
	if flag {
		for _, id := range apiIDList {
			idList = append(idList, strconv.Itoa(id))
		}
		plugins = append(plugins, "goku-proxy_caching")
	}
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.BatchEditApiPluginStatus(connIDList, strategyID, pluginStatus, userID)
	if flag {
		dao.UpdateTable(name)
		p := strings.Join(plugins, ",")
		ids := strings.Join(idList, ",")
		console_mysql.UpdateApiTagByPluginName(strategyID, ids, p)
	}
	return flag, result, err
}

// 批量删除接口插件
func BatchDeleteApiPlugin(connIDList, strategyID string) (bool, string, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.BatchDeleteApiPlugin(connIDList, strategyID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 新增插件到接口
func AddPluginToApi(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.AddPluginToApi(pluginName, config, strategyID, apiID, userID)
	if flag {
		dao.UpdateTable(name)
		console_mysql.UpdateApiTagByPluginName(strategyID, strconv.Itoa(apiID), pluginName)
	}
	return flag, result, err
}

// 修改接口插件配置
func EditApiPluginConfig(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	name := "goku_conn_plugin_api"
	flag, result, err := console_mysql.EditApiPluginConfig(pluginName, config, strategyID, apiID, userID)
	if flag {
		dao.UpdateTable(name)
		console_mysql.UpdateApiTagByPluginName(strategyID, strconv.Itoa(apiID), pluginName)
	}
	return flag, result, err
}

func GetApiPluginList(apiID int, strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetApiPluginList(apiID, strategyID)
}

// 获取插件优先级
func GetPluginIndex(pluginName string) (bool, int, error) {
	return console_mysql.GetPluginIndex(pluginName)
}

// 通过ApiID获取配置信息
func GetApiPluginConfig(apiID int, strategyID, pluginName string) (bool, map[string]string, error) {
	return console_mysql.GetApiPluginConfig(apiID, strategyID, pluginName)
}

// 检查策略组是否绑定插件
func CheckPluginIsExistInApi(strategyID, pluginName string, apiID int) (bool, error) {
	return console_mysql.CheckPluginIsExistInApi(strategyID, pluginName, apiID)
}

// 获取策略组中所有接口插件列表
func GetAllApiPluginInStrategy(strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetAllApiPluginInStrategy(strategyID)
}

// GetAPIPluginInStrategyByAPIID 获取策略组中所有接口插件列表
func GetAPIPluginInStrategyByAPIID(strategyID string, apiID int, keyword string, condition int) (bool, []map[string]interface{}, map[string]interface{}, error) {
	return console_mysql.GetAPIPluginInStrategyByAPIID(strategyID, apiID, keyword, condition)
}

func GetApiPluginListWithNotAssignApiList(strategyID string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetApiPluginListWithNotAssignApiList(strategyID)
}

func UpdateAllApiPluginUpdateTag() error {
	return console_mysql.UpdateAllApiPluginUpdateTag()
}
