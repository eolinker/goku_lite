package project

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)
var (
	projectDao dao.ProjectDao
)

func init() {
	pdao.Need(&projectDao)
}
//AddProject 新建项目
func AddProject(projectName string) (bool, interface{}, error) {
	return projectDao.AddProject(projectName)
}

//EditProject 修改项目信息
func EditProject(projectName string, projectID int) (bool, string, error) {
	return projectDao.EditProject(projectName, projectID)
}

//DeleteProject 修改项目信息
func DeleteProject(projectID int) (bool, string, error) {
	flag, result, err := projectDao.DeleteProject(projectID)

	return flag, result, err
}

//BatchDeleteProject 批量删除项目
func BatchDeleteProject(projectIDList string) (bool, string, error) {
	flag, result, err := projectDao.BatchDeleteProject(projectIDList)
	return flag, result, err
}

//GetProjectInfo 获取项目信息
func GetProjectInfo(projectID int) (bool, entity.Project, error) {
	return projectDao.GetProjectInfo(projectID)
}

//GetProjectList 获取项目列表
func GetProjectList(keyword string) (bool, []*entity.Project, error) {
	return projectDao.GetProjectList(keyword)
}

//CheckProjectIsExist 检查项目是否存在
func CheckProjectIsExist(projectID int) (bool, error) {
	return projectDao.CheckProjectIsExist(projectID)
}

//GetAPIListFromProjectNotInStrategy 获取项目列表中没有被策略组绑定的接口
func GetAPIListFromProjectNotInStrategy() (bool, []map[string]interface{}, error) {
	return projectDao.GetAPIListFromProjectNotInStrategy()
}
