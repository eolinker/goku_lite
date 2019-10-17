package api

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//ImportAPIGroupFromAms 导入接口分组
func ImportAPIGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error) {
	flag, result, err := console_sqlite3.ImportAPIGroupFromAms(projectID, userID, groupInfo)

	return flag, result, err
}

//ImportProjectFromAms 导入项目
func ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error) {
	flag, result, err := console_sqlite3.ImportProjectFromAms(userID, projectInfo)

	return flag, result, err
}

//ImportAPIFromAms 导入接口
func ImportAPIFromAms(projectID, groupID, userID int, apiList []entity.AmsAPIInfo) (bool, string, error) {
	flag, result, err := console_sqlite3.ImportAPIFromAms(projectID, groupID, userID, apiList)

	return flag, result, err
}
