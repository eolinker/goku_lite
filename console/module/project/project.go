package project

import (
	"github.com/eolinker/goku/server/dao"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
	entity "github.com/eolinker/goku/server/entity/console-entity"
)

// 新建项目
func AddProject(projectName string) (bool, interface{}, error) {
	return console_mysql.AddProject(projectName)
}

// 修改项目信息
func EditProject(projectName string, projectID int) (bool, string, error) {
	return console_mysql.EditProject(projectName, projectID)
}

// 修改项目信息
func DeleteProject(projectID int) (bool, string, error) {
	flag, result, err := console_mysql.DeleteProject(projectID)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 批量删除项目
func BatchDeleteProject(projectIDList string) (bool, string, error) {
	flag, result, err := console_mysql.BatchDeleteProject(projectIDList)
	if flag {
		name := "goku_gateway_api"
		dao.UpdateTable(name)
	}
	return flag, result, err
}

// 获取项目信息
func GetProjectInfo(projectID int) (bool, entity.Project, error) {
	return console_mysql.GetProjectInfo(projectID)
}

// 获取项目列表
func GetProjectList(keyword string) (bool, []*entity.Project, error) {
	return console_mysql.GetProjectList(keyword)
}

// 检查项目是否存在
func CheckProjectIsExist(projectID int) (bool, error) {
	return console_mysql.CheckProjectIsExist(projectID)
}

// 获取项目列表中没有被策略组绑定的接口
func GetApiListFromProjectNotInStrategy() (bool, []map[string]interface{}, error) {
	return console_mysql.GetApiListFromProjectNotInStrategy()
}
