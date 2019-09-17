package plugin

import (
	"github.com/eolinker/goku/server/dao"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
	entity "github.com/eolinker/goku/server/entity/console-entity"
)

// 获取插件配置信息
func GetPluginInfo(pluginName string) (bool, *entity.Plugin, error) {
	return console_mysql.GetPluginInfo(pluginName)
}

// 获取插件列表
func GetPluginList(keyword string, condition int) (bool, []*entity.Plugin, error) {
	return console_mysql.GetPluginList(keyword, condition)
}

// 新增插件信息
func AddPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error) {
	return console_mysql.AddPlugin(pluginName, pluginConfig, pluginDesc, version, pluginPriority, isStop, pluginType)
}

// 修改插件信息
func EditPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error) {
	return console_mysql.EditPlugin(pluginName, pluginConfig, pluginDesc, version, pluginPriority, isStop, pluginType)
}

// 删除插件信息
func DeletePlugin(pluginName string) (bool, string, error) {
	return console_mysql.DeletePlugin(pluginName)
}

// 判断插件ID是否存在
func CheckIndexIsExist(pluginName string, pluginPriority int) (bool, error) {
	return console_mysql.CheckIndexIsExist(pluginName, pluginPriority)
}

// 获取插件配置及插件信息
func GetPluginConfig(pluginName string) (bool, string, error) {
	return console_mysql.GetPluginConfig(pluginName)
}

// 检查插件名称是否存在
func CheckNameIsExist(pluginName string) (bool, error) {
	return console_mysql.CheckNameIsExist(pluginName)
}

// 修改插件开启状态
func EditPluginStatus(pluginName string, pluginStatus int) (bool, error) {
	tableName := "goku_plugin"
	flag, err := console_mysql.EditPluginStatus(pluginName, pluginStatus)
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.UpdatePluginTagByPluginName(pluginName)
	}
	return flag, err
}

// 获取不同类型的插件列表
func GetPluginListByPluginType(pluginType int) (bool, []map[string]interface{}, error) {
	return console_mysql.GetPluginListByPluginType(pluginType)
}

// 批量关闭插件
func BatchStopPlugin(pluginNameList string) (bool, string, error) {
	//if strings.Contains(pluginNameList, "goku-rate_limiting") {
	//	updateFlag, errInfo := console_mysql.DeleteRateInfoInRedis("")
	//	if !updateFlag {
	//		utils.SystemLog(errInfo)
	//	}
	//} else if strings.Contains(pluginNameList, "goku-replay_attack_defender") {
	//	updateFlag, errInfo := console_mysql.DeleteReplayAttackTokenInRedis("")
	//	if !updateFlag {
	//		utils.SystemLog(errInfo)
	//	}
	//} else if strings.Contains(pluginNameList, "goku-oauth2_auth") {
	//	updateFlag, errInfo := console_mysql.DeleteOauth2InfoInRedis("", "")
	//	if !updateFlag {
	//		utils.SystemLog(errInfo)
	//	}
	//} else if strings.Contains(pluginNameList, "goku-proxy_caching") {
	//	console_mysql.ClearRedisProxyCache("", 0)
	//} else if strings.Contains(pluginNameList, "goku-circuit_breaker") {
	//	console_mysql.ClearRedisCircuitBreaker("", 0)
	//}
	tableName := "goku_plugin"
	flag, result, err := console_mysql.BatchStopPlugin(pluginNameList)
	if flag {
		dao.UpdateTable(tableName)
	}
	return flag, result, err
}

// 批量关闭插件
func BatchStartPlugin(pluginNameList string) (bool, string, error) {
	tableName := "goku_plugin"
	flag, result, err := console_mysql.BatchStartPlugin(pluginNameList)
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.UpdatePluginTagByPluginName(pluginNameList)
	}
	return flag, result, err
}

// 更新插件检测状态
func EditPluginCheckStatus(pluginName string, isCheck int) (bool, string, error) {
	return console_mysql.EditPluginCheckStatus(pluginName, isCheck)
}
