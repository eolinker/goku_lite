package api

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	consolemysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//AddAPIGroup 新建接口分组
func AddAPIGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	return consolemysql.AddAPIGroup(groupName, projectID, parentGroupID)
}

//EditAPIGroup 修改接口分组
func EditAPIGroup(groupName string, groupID, projectID int) (bool, string, error) {
	return consolemysql.EditAPIGroup(groupName, groupID, projectID)
}

//DeleteAPIGroup 删除接口分组
func DeleteAPIGroup(projectID, groupID int) (bool, string, error) {
	flag, result, err := consolemysql.DeleteAPIGroup(projectID, groupID)
	if flag {
		dao.UpdateTable("goku_gateway_strategy")
		dao.UpdateTable("goku_gateway_api")
		dao.UpdateTable("goku_conn_strategy_api")
		dao.UpdateTable("goku_conn_plugin_strategy")
		dao.UpdateTable("goku_conn_plugin_api")
	}
	return flag, result, err
}

//GetAPIGroupList 获取接口分组列表
func GetAPIGroupList(projectID int) (bool, []map[string]interface{}, error) {
	return consolemysql.GetAPIGroupList(projectID)
}

//UpdateAPIGroupScript 更新接口分组脚本
func UpdateAPIGroupScript() bool {
	return consolemysql.UpdateAPIGroupScript()
}
