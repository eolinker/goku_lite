package api

//AddAPIToStrategy 将接口加入策略组
func AddAPIToStrategy(apiList []string, strategyID string) (bool, string, error) {
	flag, result, err := apiStrategyDao.AddAPIToStrategy(apiList, strategyID)

	return flag, result, err
}

//SetTarget 重置目标地址
func SetTarget(apiID int, strategyID string, target string) (bool, string, error) {
	flag, result, err := apiStrategyDao.SetAPITargetOfStrategy(apiID, strategyID, target)

	return flag, result, err
}

// BatchSetTarget 批量重置目标地址
func BatchSetTarget(apiIds []int, strategyID string, target string) (bool, string, error) {
	flag, result, err := apiStrategyDao.BatchSetAPITargetOfStrategy(apiIds, strategyID, target)

	return flag, result, err
}

// GetAPIIDListFromStrategy 获取策略组接口ID列表
func GetAPIIDListFromStrategy(strategyID, keyword string, condition int, ids []int, balanceNames []string) (bool, []int, error) {
	return apiStrategyDao.GetAPIIDListFromStrategy(strategyID, keyword, condition, ids, balanceNames)
}

// GetAPIListFromStrategy 获取策略组接口列表
func GetAPIListFromStrategy(strategyID, keyword string, condition, page, pageSize int, ids []int, balanceNames []string) (bool, []map[string]interface{}, int, error) {
	return apiStrategyDao.GetAPIListFromStrategy(strategyID, keyword, condition, page, pageSize, ids, balanceNames)
}

//CheckIsExistAPIInStrategy 检查插件是否添加进策略组
func CheckIsExistAPIInStrategy(apiID int, strategyID string) (bool, string, error) {
	return apiStrategyDao.CheckIsExistAPIInStrategy(apiID, strategyID)
}

// GetAPIIDListNotInStrategy 获取未被该策略组绑定的接口ID列表(通过项目)
func GetAPIIDListNotInStrategy(strategyID string, projectID, groupID int, keyword string) (bool, []int, error) {
	return apiStrategyDao.GetAPIIDListNotInStrategy(strategyID, projectID, groupID, keyword)
}

// GetAPIListNotInStrategy 获取未被该策略组绑定的接口列表(通过项目)
func GetAPIListNotInStrategy(strategyID string, projectID, groupID, page, pageSize int, keyword string) (bool, []map[string]interface{}, int, error) {
	return apiStrategyDao.GetAPIListNotInStrategy(strategyID, projectID, groupID, page, pageSize, keyword)
}

//BatchDeleteAPIInStrategy 批量删除策略组接口
func BatchDeleteAPIInStrategy(apiIDList, strategyID string) (bool, string, error) {
	flag, result, err := apiStrategyDao.BatchDeleteAPIInStrategy(apiIDList, strategyID)

	return flag, result, err
}
