package api

import (
	"github.com/eolinker/goku/server/dao"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
)

// 新建接口分组
func AddApiGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	return console_mysql.AddApiGroup(groupName, projectID, parentGroupID)
}

// 修改接口分组
func EditApiGroup(groupName string, groupID, projectID int) (bool, string, error) {
	return console_mysql.EditApiGroup(groupName, groupID, projectID)
}

// 删除接口分组
func DeleteApiGroup(projectID, groupID int) (bool, string, error) {
	flag, result, err := console_mysql.DeleteApiGroup(projectID, groupID)
	if flag {
		dao.UpdateTable("goku_gateway_strategy")
		dao.UpdateTable("goku_gateway_api")
		dao.UpdateTable("goku_conn_strategy_api")
		dao.UpdateTable("goku_conn_plugin_strategy")
		dao.UpdateTable("goku_conn_plugin_api")
	}
	return flag, result, err
}

// 获取接口分组列表
func GetApiGroupList(projectID int) (bool, []map[string]interface{}, error) {
	return console_mysql.GetApiGroupList(projectID)
}

func UpdateApiGroupScript() bool {
	return console_mysql.UpdateApiGroupScript()
}
