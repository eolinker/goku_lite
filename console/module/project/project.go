package project

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddProject 新建项目
func AddProject(projectName string) (bool, interface{}, error) {
	return console_mysql.AddProject(projectName)
}

//EditProject 修改项目信息
func EditProject(projectName string, projectID int) (bool, string, error) {
	return console_mysql.EditProject(projectName, projectID)
}

//DeleteProject 修改项目信息
func DeleteProject(projectID int) (bool, string, error) {
	flag, result, err := console_mysql.DeleteProject(projectID)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//BatchDeleteProject 批量删除项目
func BatchDeleteProject(projectIDList string) (bool, string, error) {
	flag, result, err := console_mysql.BatchDeleteProject(projectIDList)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

//GetProjectInfo 获取项目信息
func GetProjectInfo(projectID int) (bool, entity.Project, error) {
	return console_mysql.GetProjectInfo(projectID)
}

//GetProjectList 获取项目列表
func GetProjectList(keyword string) (bool, []*entity.Project, error) {
	return console_mysql.GetProjectList(keyword)
}

//CheckProjectIsExist 检查项目是否存在
func CheckProjectIsExist(projectID int) (bool, error) {
	return console_mysql.CheckProjectIsExist(projectID)
}

//GetAPIListFromProjectNotInStrategy 获取项目列表中没有被策略组绑定的接口
func GetAPIListFromProjectNotInStrategy() (bool, []map[string]interface{}, error) {
	return console_mysql.GetAPIListFromProjectNotInStrategy()
}
