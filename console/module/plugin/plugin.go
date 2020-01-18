package plugin

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)
var (
	pluginDao dao.PluginDao
)

func init() {
	pdao.Need(&pluginDao)
}
//GetPluginInfo 获取插件配置信息
func GetPluginInfo(pluginName string) (bool, *entity.Plugin, error) {
	return pluginDao.GetPluginInfo(pluginName)
}

//GetPluginList 获取插件列表
func GetPluginList(keyword string, condition int) (bool, []*entity.Plugin, error) {
	return pluginDao.GetPluginList(keyword, condition)
}

//AddPlugin 新增插件信息
func AddPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error) {
	return pluginDao.AddPlugin(pluginName, pluginConfig, pluginDesc, version, pluginPriority, isStop, pluginType)
}

//EditPlugin 修改插件信息
func EditPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error) {
	return pluginDao.EditPlugin(pluginName, pluginConfig, pluginDesc, version, pluginPriority, isStop, pluginType)
}

//DeletePlugin 删除插件信息
func DeletePlugin(pluginName string) (bool, string, error) {
	return pluginDao.DeletePlugin(pluginName)
}

//CheckIndexIsExist 判断插件ID是否存在
func CheckIndexIsExist(pluginName string, pluginPriority int) (bool, error) {
	return pluginDao.CheckIndexIsExist(pluginName, pluginPriority)
}

//GetPluginConfig 获取插件配置及插件信息
func GetPluginConfig(pluginName string) (bool, string, error) {
	return pluginDao.GetPluginConfig(pluginName)
}

//CheckNameIsExist 检查插件名称是否存在
func CheckNameIsExist(pluginName string) (bool, error) {
	return pluginDao.CheckNameIsExist(pluginName)
}

//EditPluginStatus 修改插件开启状态
func EditPluginStatus(pluginName string, pluginStatus int) (bool, error) {
	flag, err := pluginDao.EditPluginStatus(pluginName, pluginStatus)

	return flag, err
}

//GetPluginListByPluginType 获取不同类型的插件列表
func GetPluginListByPluginType(pluginType int) (bool, []map[string]interface{}, error) {
	return pluginDao.GetPluginListByPluginType(pluginType)
}

//BatchStopPlugin 批量关闭插件
func BatchStopPlugin(pluginNameList string) (bool, string, error) {
	flag, result, err := pluginDao.BatchStopPlugin(pluginNameList)

	return flag, result, err
}

//BatchStartPlugin 批量关闭插件
func BatchStartPlugin(pluginNameList string) (bool, string, error) {
	flag, result, err := pluginDao.BatchStartPlugin(pluginNameList)

	return flag, result, err
}

//EditPluginCheckStatus 更新插件检测状态
func EditPluginCheckStatus(pluginName string, isCheck int) (bool, string, error) {
	return pluginDao.EditPluginCheckStatus(pluginName, isCheck)
}
