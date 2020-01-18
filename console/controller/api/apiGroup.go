package api

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
	"github.com/eolinker/goku-api-gateway/console/module/project"
	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

const operationAPIGroup = "apiManagement"

//GroupHandlers 接口分组处理器
type GroupHandlers struct {
}

//Handlers 处理器
func (g *GroupHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":     factory.NewAccountHandleFunction(operationAPIGroup, true, AddAPIGroup),
		"/edit":    factory.NewAccountHandleFunction(operationAPIGroup, true, EditAPIGroup),
		"/delete":  factory.NewAccountHandleFunction(operationAPIGroup, true, DeleteAPIGroup),
		"/getList": factory.NewAccountHandleFunction(operationAPIGroup, false, GetAPIGroupList),
	}
}

//NewGroupHandlers new 接口分组处理器
func NewGroupHandlers() *GroupHandlers {
	return &GroupHandlers{}
}

//AddAPIGroup 新建接口分组
func AddAPIGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	groupName := httpRequest.PostFormValue("groupName")
	projectID := httpRequest.PostFormValue("projectID")
	parentGroupID := httpRequest.PostFormValue("parentGroupID")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"290006",
			"api", "[ERROR]Illegal groupName!", nil)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290001",
			"api", "[ERROR]Illegal projectID!", err)
		return

	}
	pgID, err := strconv.Atoi(parentGroupID)
	if err != nil && parentGroupID != "" {
		controller.WriteError(httpResponse,
			"290002",
			"api", "[ERROR]Illegal parentGroupID!", err)
		return

	}
	flag, err := project.CheckProjectIsExist(pjID)
	if !flag {

		controller.WriteError(httpResponse,
			"290005",
			"api", "[ERROR]The project does not exist", err)
		return

	}
	flag, result, err := api.AddAPIGroup(groupName, pjID, pgID)
	if !flag {

		controller.WriteError(httpResponse,
			"290000",
			"apiGroup",
			result.(string),
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "apiGroup", "groupID", result)
	return
}

//EditAPIGroup 修改接口分组
func EditAPIGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	groupName := httpRequest.PostFormValue("groupName")
	groupID := httpRequest.PostFormValue("groupID")
	projectID := httpRequest.PostFormValue("projectID")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"290006",
			"apiGroup", "[ERROR]Illegal groupName!", nil)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290004",
			"apiGroup", "[ERROR]Illegal groupID!", err)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290001",
			"apiGroup", "[ERROR]Illegal projectID!", err)
		return

	}
	flag, err := project.CheckProjectIsExist(pjID)
	if !flag {

		controller.WriteError(httpResponse,
			"290005",
			"apiGroup", "[ERROR]The project does not exist", err)
		return

	}
	flag, result, err := api.EditAPIGroup(groupName, gID, pjID)
	if !flag {

		controller.WriteError(httpResponse,
			"290000",
			"apiGroup", result, err)
	}
	controller.WriteResultInfo(httpResponse, "apiGroup", "", nil)
	return
}

//DeleteAPIGroup 删除接口分组
func DeleteAPIGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	groupID := httpRequest.PostFormValue("groupID")
	projectID := httpRequest.PostFormValue("projectID")

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290004",
			"apiGroup", "[ERROR]Illegal groupID!", err)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290001",
			"apiGroup", "[ERROR]Illegal projectID!", err)
		return

	}
	flag, result, err := api.DeleteAPIGroup(pjID, gID)
	if !flag {

		controller.WriteError(httpResponse,
			"290000",
			"apiGroup", result, err)
	}

	controller.WriteResultInfo(httpResponse, "apiGroup", "", nil)
	return
}

//GetAPIGroupList 获取接口分组列表
func GetAPIGroupList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	projectID := httpRequest.PostFormValue("projectID")
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"290001",
			"apiGroup", "[ERROR]Illegal projectID!", err)
		return

	}
	flag, result, err := api.GetAPIGroupList(pjID)
	if !flag {

		controller.WriteError(httpResponse,
			"290000",
			"apiGroup", "[ERROR]Empty api group list!", err)
		return

	}

	controller.WriteResultInfo(httpResponse, "apiGroup", "groupList", result)
	return
}
