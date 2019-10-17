package api

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddAPI 新增接口
func AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType int) (bool, int, error) {

	flag, result, err := console_sqlite3.AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType)

	return flag, result, err
}

//EditAPI 新增接口
func EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error) {
	flag, err := console_sqlite3.EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID)

	return flag, err
}

//GetAPIInfo 获取接口信息
func GetAPIInfo(apiID int) (bool, *entity.API, error) {
	return console_sqlite3.GetAPIInfo(apiID)
}

// GetAPIIDList 获取接口ID列表
func GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error) {
	return console_sqlite3.GetAPIIDList(projectID, groupID, keyword, condition, ids)
}

// GetAPIList 获取接口列表
func GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error) {
	return console_sqlite3.GetAPIList(projectID, groupID, keyword, condition, page, pageSize, ids)
}

//CheckURLIsExist 接口路径是否存在
func CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool {
	return console_sqlite3.CheckURLIsExist(requestURL, requestMethod, projectID, apiID)
}

//CheckAPIIsExist 检查接口是否存在
func CheckAPIIsExist(apiID int) (bool, error) {
	return console_sqlite3.CheckAPIIsExist(apiID)
}

//BatchEditAPIGroup 批量修改接口分组
func BatchEditAPIGroup(apiIDList []string, groupID int) (bool, string, error) {
	r, e := console_sqlite3.BatchEditAPIGroup(apiIDList, groupID)

	return e == nil, r, e
}

//BatchEditAPIBalance 批量修改接口负载
func BatchEditAPIBalance(apiIDList []string, balance string) (string, error) {

	r, err := console_sqlite3.BatchEditAPIBalance(apiIDList, balance)

	return r, err
}

//BatchDeleteAPI 批量删除接口
func BatchDeleteAPI(apiIDList string) (bool, string, error) {

	flag, result, err := console_sqlite3.BatchDeleteAPI(apiIDList)

	return flag, result, err
}
