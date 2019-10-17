package strategy

import (
	"errors"

	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//AddStrategyGroup 新建策略组分组
func AddStrategyGroup(groupName string) (bool, interface{}, error) {
	return console_sqlite3.AddStrategyGroup(groupName)
}

//EditStrategyGroup 修改策略组分组
func EditStrategyGroup(groupName string, groupID int) (bool, string, error) {
	return console_sqlite3.EditStrategyGroup(groupName, groupID)
}

// DeleteStrategyGroup 删除策略组分组
func DeleteStrategyGroup(groupID int) (bool, string, error) {
	flag := console_sqlite3.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	}
	flag, result, err := console_sqlite3.DeleteStrategyGroup(groupID)

	return flag, result, err
}

//GetStrategyGroupList 获取策略组分组列表
func GetStrategyGroupList() (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetStrategyGroupList()
}

//CheckIsOpenGroup 判断是否是开放分组
func CheckIsOpenGroup(groupID int) bool {
	return console_sqlite3.CheckIsOpenGroup(groupID)
}
