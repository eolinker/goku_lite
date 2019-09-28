package strategy

import (
	"errors"

	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//AddStrategyGroup 新建策略组分组
func AddStrategyGroup(groupName string) (bool, interface{}, error) {
	return console_mysql.AddStrategyGroup(groupName)
}

//EditStrategyGroup 修改策略组分组
func EditStrategyGroup(groupName string, groupID int) (bool, string, error) {
	return console_mysql.EditStrategyGroup(groupName, groupID)
}

// DeleteStrategyGroup 删除策略组分组
func DeleteStrategyGroup(groupID int) (bool, string, error) {
	flag := console_mysql.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	}
	flag, result, err := console_mysql.DeleteStrategyGroup(groupID)
	if flag {
		dao.UpdateTable("goku_gateway_strategy")
		dao.UpdateTable("goku_gateway_api")
		dao.UpdateTable("goku_conn_strategy_api")
		dao.UpdateTable("goku_conn_plugin_strategy")
		dao.UpdateTable("goku_conn_plugin_api")
	}
	return flag, result, err
}

//GetStrategyGroupList 获取策略组分组列表
func GetStrategyGroupList() (bool, []map[string]interface{}, error) {
	return console_mysql.GetStrategyGroupList()
}

//CheckIsOpenGroup 判断是否是开放分组
func CheckIsOpenGroup(groupID int) bool {
	return console_mysql.CheckIsOpenGroup(groupID)
}
