package api

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//AddAPIGroup 新建接口分组
func AddAPIGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	return console_sqlite3.AddAPIGroup(groupName, projectID, parentGroupID)
}

//EditAPIGroup 修改接口分组
func EditAPIGroup(groupName string, groupID, projectID int) (bool, string, error) {
	return console_sqlite3.EditAPIGroup(groupName, groupID, projectID)
}

//DeleteAPIGroup 删除接口分组
func DeleteAPIGroup(projectID, groupID int) (bool, string, error) {
	flag, result, err := console_sqlite3.DeleteAPIGroup(projectID, groupID)

	return flag, result, err
}

//GetAPIGroupList 获取接口分组列表
func GetAPIGroupList(projectID int) (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetAPIGroupList(projectID)
}
