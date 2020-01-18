package api

import (
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddAPI 新增接口
func AddAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType int) (bool, int, error) {

	flag, result, err := apiDao.AddAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType)

	return flag, result, err
}

//EditAPI 新增接口
func EditAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error) {
	flag, err := apiDao.EditAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID)

	return flag, err
}

//GetAPIInfo 获取接口信息
func GetAPIInfo(apiID int) (bool, *entity.API, error) {
	return apiDao.GetAPIInfo(apiID)
}

// GetAPIIDList 获取接口ID列表
func GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error) {
	return apiDao.GetAPIIDList(projectID, groupID, keyword, condition, ids)
}

// GetAPIList 获取接口列表
func GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error) {
	return apiDao.GetAPIList(projectID, groupID, keyword, condition, page, pageSize, ids)
}

//CheckURLIsExist 接口路径是否存在
func CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool {
	return apiDao.CheckURLIsExist(requestURL, requestMethod, projectID, apiID)
}

//CheckAPIIsExist 检查接口是否存在
func CheckAPIIsExist(apiID int) (bool, error) {
	return apiDao.CheckAPIIsExist(apiID)
}

//CheckAliasIsExist 检查接口是否存在
func CheckAliasIsExist(apiID int, alias string) bool {
	return apiDao.CheckAliasIsExist(apiID, alias)
}

//BatchEditAPIGroup 批量修改接口分组
func BatchEditAPIGroup(apiIDList []string, groupID int) (bool, string, error) {
	r, e := apiDao.BatchEditAPIGroup(apiIDList, groupID)

	return e == nil, r, e
}

//BatchEditAPIBalance 批量修改接口负载
func BatchEditAPIBalance(apiIDList []string, balance string) (string, error) {

	r, err := apiDao.BatchEditAPIBalance(apiIDList, balance)

	return r, err
}

//BatchDeleteAPI 批量删除接口
func BatchDeleteAPI(apiIDList string) (bool, string, error) {

	flag, result, err := apiDao.BatchDeleteAPI(apiIDList)

	return flag, result, err
}
