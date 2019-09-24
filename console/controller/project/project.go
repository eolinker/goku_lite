package project

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/project"
)

// 新建项目
func AddProject(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	projectName := httpRequest.PostFormValue("projectName")
	if projectName == "" {
		controller.WriteError(httpResponse,
			"270002",
			"project",
			"[ERROR]Illegal projectName!",
			nil)
		return
	}
	flag, result, err := project.AddProject(projectName)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			result.(string),
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "project", "projectID", result)

	return
}

// 修改项目信息
func EditProject(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	projectName := httpRequest.PostFormValue("projectName")
	projectID := httpRequest.PostFormValue("projectID")
	if projectName == "" {
		controller.WriteError(httpResponse,
			"270002",
			"project",
			"[ERROR]Illegal projectName!",
			nil)
		return
	}
	id, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"270001",
			"project",
			"[ERROR]Illegal projectID!",
			err)
		return

	}
	flag, result, err := project.EditProject(projectName, id)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "project", "", nil)

	return
}

// 删除项目信息
func DeleteProject(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	projectID := httpRequest.PostFormValue("projectID")

	id, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"270001",
			"project",
			"[ERROR]Illegal projectID!",
			err)
		return

	}
	flag, result, err := project.DeleteProject(id)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "project", "", nil)

	return
}

// 删除项目信息
func BatchDeleteProject(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	projectIDList := httpRequest.PostFormValue("projectIDList")

	flag, result, err := project.BatchDeleteProject(projectIDList)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "project", "", nil)
	return
}

// 获取项目信息
func GetProjectInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	projectID := httpRequest.PostFormValue("projectID")

	id, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"270001",
			"project",
			"[ERROR]Illegal projectID!",
			err)
		return
	}
	flag, result, err := project.GetProjectInfo(id)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			"[ERROR]The project information does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "project", "projectInfo", result)

	return
}

// 获取项目列表
func GetProjectList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	keyword := httpRequest.FormValue("keyword")

	flag, result, err := project.GetProjectList(keyword)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			"[ERROR]Empty project list!",
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "project", "projectList", result)

	return
}

// 获取项目列表中没有被策略组绑定的接口
func GetApiListFromProjectNotInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := project.GetApiListFromProjectNotInStrategy()
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"project",
			"[ERROR]Empty project list!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "project", "projectList", result)

}
