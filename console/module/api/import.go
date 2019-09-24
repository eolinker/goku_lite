package api

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

func ImportApiGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error) {
	flag, result, err := console_mysql.ImportApiGroupFromAms(projectID, userID, groupInfo)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 导入项目
func ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error) {
	flag, result, err := console_mysql.ImportProjectFromAms(userID, projectInfo)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 导入接口
func ImportApiFromAms(projectID, groupID, userID int, apiList []entity.AmsApiInfo) (bool, string, error) {
	flag, result, err := console_mysql.ImportApiFromAms(projectID, groupID, userID, apiList)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}
