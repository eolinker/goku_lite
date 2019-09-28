package api

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//ImportAPIGroupFromAms 导入接口分组
func ImportAPIGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error) {
	flag, result, err := console_mysql.ImportAPIGroupFromAms(projectID, userID, groupInfo)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//ImportProjectFromAms 导入项目
func ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error) {
	flag, result, err := console_mysql.ImportProjectFromAms(userID, projectInfo)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//ImportAPIFromAms 导入接口
func ImportAPIFromAms(projectID, groupID, userID int, apiList []entity.AmsAPIInfo) (bool, string, error) {
	flag, result, err := console_mysql.ImportAPIFromAms(projectID, groupID, userID, apiList)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}
