package api

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddAPI 新增接口
func AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID int) (bool, int, error) {
	name := "goku_gateway_api"

	flag, result, err := console_mysql.AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, managerID, userID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//EditAPI 新增接口
func EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error) {
	name := "goku_gateway_api"
	flag, err := console_mysql.EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, err
}

//GetAPIInfo 获取接口信息
func GetAPIInfo(apiID int) (bool, *entity.API, error) {
	return console_mysql.GetAPIInfo(apiID)
}

// GetAPIIDList 获取接口ID列表
func GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error) {
	return console_mysql.GetAPIIDList(projectID, groupID, keyword, condition, ids)
}

// GetAPIList 获取接口列表
func GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error) {
	return console_mysql.GetAPIList(projectID, groupID, keyword, condition, page, pageSize, ids)
}

//CheckURLIsExist 接口路径是否存在
func CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool {
	return console_mysql.CheckURLIsExist(requestURL, requestMethod, projectID, apiID)
}

//CheckAPIIsExist 检查接口是否存在
func CheckAPIIsExist(apiID int) (bool, error) {
	return console_mysql.CheckAPIIsExist(apiID)
}

//BatchEditAPIGroup 批量修改接口分组
func BatchEditAPIGroup(apiIDList []string, groupID int) (bool, string, error) {
	r, e := console_mysql.BatchEditAPIGroup(apiIDList, groupID)

	return e == nil, r, e
}

//BatchEditAPIBalance 批量修改接口负载
func BatchEditAPIBalance(apiIDList []string, balance string) (string, error) {

	r, err := console_mysql.BatchEditAPIBalance(apiIDList, balance)
	if err != nil {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return r, err
}

//BatchDeleteAPI 批量删除接口
func BatchDeleteAPI(apiIDList string) (bool, string, error) {

	flag, result, err := console_mysql.BatchDeleteAPI(apiIDList)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}
