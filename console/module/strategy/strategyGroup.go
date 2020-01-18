package strategy

import (
	"errors"
)

//AddStrategyGroup 新建策略组分组
func AddStrategyGroup(groupName string) (bool, interface{}, error) {
	return strategyGroupDao.AddStrategyGroup(groupName)
}

//EditStrategyGroup 修改策略组分组
func EditStrategyGroup(groupName string, groupID int) (bool, string, error) {
	return strategyGroupDao.EditStrategyGroup(groupName, groupID)
}

// DeleteStrategyGroup 删除策略组分组
func DeleteStrategyGroup(groupID int) (bool, string, error) {
	flag := strategyGroupDao.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	}
	flag, result, err := strategyGroupDao.DeleteStrategyGroup(groupID)

	return flag, result, err
}

//GetStrategyGroupList 获取策略组分组列表
func GetStrategyGroupList() (bool, []map[string]interface{}, error) {
	return strategyGroupDao.GetStrategyGroupList()
}

//CheckIsOpenGroup 判断是否是开放分组
func CheckIsOpenGroup(groupID int) bool {
	return strategyGroupDao.CheckIsOpenGroup(groupID)
}
