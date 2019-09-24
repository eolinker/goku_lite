package api

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

// 将接口加入策略组
func AddApiToStrategy(apiList []string, strategyID string) (bool, string, error) {
	name := "goku_conn_strategy_api"
	flag, result, err := console_mysql.AddApiToStrategy(apiList, strategyID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 重置目标地址
func SetTarget(apiId int, strategyID string, target string) (bool, string, error) {
	name := "goku_conn_strategy_api"
	flag, result, err := console_mysql.SetApiTargetOfStrategy(apiId, strategyID, target)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// BatchSetTarget 批量重置目标地址
func BatchSetTarget(apiIds []int, strategyID string, target string) (bool, string, error) {
	name := "goku_conn_strategy_api"
	flag, result, err := console_mysql.BatchSetApiTargetOfStrategy(apiIds, strategyID, target)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// GetAPIIDListFromStrategy 获取策略组接口ID列表
func GetAPIIDListFromStrategy(strategyID, keyword string, condition int, ids []int, balanceNames []string) (bool, []int, error) {
	return console_mysql.GetAPIIDListFromStrategy(strategyID, keyword, condition, ids, balanceNames)
}

// GetAPIListFromStrategy 获取策略组接口列表
func GetAPIListFromStrategy(strategyID, keyword string, condition, page, pageSize int, ids []int, balanceNames []string) (bool, []map[string]interface{}, int, error) {
	return console_mysql.GetAPIListFromStrategy(strategyID, keyword, condition, page, pageSize, ids, balanceNames)
}

// 检查插件是否添加进策略组
func CheckIsExistApiInStrategy(apiID int, strategyID string) (bool, string, error) {
	return console_mysql.CheckIsExistApiInStrategy(apiID, strategyID)
}

// GetAPIIDListNotInStrategyByProject 获取未被该策略组绑定的接口ID列表(通过项目)
func GetAPIIDListNotInStrategy(strategyID string, projectID, groupID int, keyword string) (bool, []int, error) {
	return console_mysql.GetAPIIDListNotInStrategy(strategyID, projectID, groupID, keyword)
}

// GetAPIListNotInStrategy 获取未被该策略组绑定的接口列表(通过项目)
func GetAPIListNotInStrategy(strategyID string, projectID, groupID, page, pageSize int, keyword string) (bool, []map[string]interface{}, int, error) {
	return console_mysql.GetAPIListNotInStrategy(strategyID, projectID, groupID, page, pageSize, keyword)
}

// 批量删除策略组接口
func BatchDeleteApiInStrategy(apiIDList, strategyID string) (bool, string, error) {
	name := "goku_conn_strategy_api"
	flag, result, err := console_mysql.BatchDeleteApiInStrategy(apiIDList, strategyID)
	if flag {
		dao.UpdateTable(name)
	}
	return flag, result, err
}
