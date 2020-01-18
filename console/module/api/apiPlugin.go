package api

//BatchEditAPIPluginStatus BatchEditAPIPluginStatus批量修改接口插件状态
func BatchEditAPIPluginStatus(connIDList, strategyID string, pluginStatus, userID int) (bool, string, error) {
	flag, result, err := apiPluginDao.BatchEditAPIPluginStatus(connIDList, strategyID, pluginStatus, userID)
	return flag, result, err
}

//BatchDeleteAPIPlugin 批量删除接口插件
func BatchDeleteAPIPlugin(connIDList, strategyID string) (bool, string, error) {
	flag, result, err := apiPluginDao.BatchDeleteAPIPlugin(connIDList, strategyID)

	return flag, result, err
}

//AddPluginToAPI 新增插件到接口
func AddPluginToAPI(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	flag, result, err := apiPluginDao.AddPluginToAPI(pluginName, config, strategyID, apiID, userID)

	return flag, result, err
}

//EditAPIPluginConfig 修改接口插件配置
func EditAPIPluginConfig(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	flag, result, err := apiPluginDao.EditAPIPluginConfig(pluginName, config, strategyID, apiID, userID)

	return flag, result, err
}

//GetAPIPluginList 获取接口插件列表
func GetAPIPluginList(apiID int, strategyID string) (bool, []map[string]interface{}, error) {
	return apiPluginDao.GetAPIPluginList(apiID, strategyID)
}

//GetPluginIndex 获取插件优先级
func GetPluginIndex(pluginName string) (bool, int, error) {
	return apiPluginDao.GetPluginIndex(pluginName)
}

//GetAPIPluginConfig 通过ApiID获取配置信息
func GetAPIPluginConfig(apiID int, strategyID, pluginName string) (bool, map[string]string, error) {
	return apiPluginDao.GetAPIPluginConfig(apiID, strategyID, pluginName)
}

//CheckPluginIsExistInAPI 检查策略组是否绑定插件
func CheckPluginIsExistInAPI(strategyID, pluginName string, apiID int) (bool, error) {
	return apiPluginDao.CheckPluginIsExistInAPI(strategyID, pluginName, apiID)
}

//GetAllAPIPluginInStrategy 获取策略组中所有接口插件列表
func GetAllAPIPluginInStrategy(strategyID string) (bool, []map[string]interface{}, error) {
	return apiPluginDao.GetAllAPIPluginInStrategy(strategyID)
}

// GetAPIPluginInStrategyByAPIID 获取策略组中所有接口插件列表
func GetAPIPluginInStrategyByAPIID(strategyID string, apiID int, keyword string, condition int) (bool, []map[string]interface{}, map[string]interface{}, error) {
	return apiPluginDao.GetAPIPluginInStrategyByAPIID(strategyID, apiID, keyword, condition)
}

//GetAPIPluginListWithNotAssignAPIList 获取没有绑定插件得接口列表
func GetAPIPluginListWithNotAssignAPIList(strategyID string) (bool, []map[string]interface{}, error) {
	return apiPluginDao.GetAPIPluginListWithNotAssignAPIList(strategyID)
}
