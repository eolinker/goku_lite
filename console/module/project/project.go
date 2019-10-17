package project

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddProject 新建项目
func AddProject(projectName string) (bool, interface{}, error) {
	return console_sqlite3.AddProject(projectName)
}

//EditProject 修改项目信息
func EditProject(projectName string, projectID int) (bool, string, error) {
	return console_sqlite3.EditProject(projectName, projectID)
}

//DeleteProject 修改项目信息
func DeleteProject(projectID int) (bool, string, error) {
	flag, result, err := console_sqlite3.DeleteProject(projectID)

	return flag, result, err
}

//BatchDeleteProject 批量删除项目
func BatchDeleteProject(projectIDList string) (bool, string, error) {
	flag, result, err := console_sqlite3.BatchDeleteProject(projectIDList)
	return flag, result, err
}

//GetProjectInfo 获取项目信息
func GetProjectInfo(projectID int) (bool, entity.Project, error) {
	return console_sqlite3.GetProjectInfo(projectID)
}

//GetProjectList 获取项目列表
func GetProjectList(keyword string) (bool, []*entity.Project, error) {
	return console_sqlite3.GetProjectList(keyword)
}

//CheckProjectIsExist 检查项目是否存在
func CheckProjectIsExist(projectID int) (bool, error) {
	return console_sqlite3.CheckProjectIsExist(projectID)
}

//GetAPIListFromProjectNotInStrategy 获取项目列表中没有被策略组绑定的接口
func GetAPIListFromProjectNotInStrategy() (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetAPIListFromProjectNotInStrategy()
}
