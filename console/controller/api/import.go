package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

const operationImportAMS = "apiManagement"

//ImportHandlers 导入handlers
type ImportHandlers struct {
}

//Handlers handler
func (i *ImportHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {

	return map[string]http.Handler{
		"/api":     factory.NewAccountHandleFunction(operationImportAMS, true, ImportAPIFromAms),
		"/group":   factory.NewAccountHandleFunction(operationImportAMS, true, ImportAPIGroupFromAms),
		"/project": factory.NewAccountHandleFunction(operationImportAMS, true, ImportProjectFromAms),
	}
}

//NewImportHandlers new导入处理器
func NewImportHandlers() *ImportHandlers {
	return &ImportHandlers{}
}

//ImportAPIGroupFromAms 导入分组
func ImportAPIGroupFromAms(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	contentType := httpRequest.Header.Get("Content-Type")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if !strings.Contains(contentType, "multipart/form-data") {
		controller.WriteError(httpResponse,
			"310001",
			"import",
			"[ERROR]Request Content-Type isn't multipart/form-data",
			nil)
		return
	}
	projectID := httpRequest.PostFormValue("projectID")

	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"310002",
			"import",
			"[ERROR]Illegal projectID!",
			err)
		return

	}
	file, _, err := httpRequest.FormFile("file")
	if err != nil {
		controller.WriteError(httpResponse,
			"310004",
			"import",
			"[ERROR]Param file does not exist!",
			err)
		return

	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		controller.WriteError(httpResponse,
			"310005",
			"import",
			"[ERROR]Fail to read file!",
			err)
		return

	}
	var groupInfo entity.AmsGroupInfo
	err = json.Unmarshal(body, &groupInfo)
	if err != nil {
		controller.WriteError(httpResponse,
			"310003",
			"import",
			"[ERROR]Fail to parse json!",
			err)
		return

	}
	if groupInfo.GroupName == "" {
		controller.WriteError(httpResponse,
			"310006",
			"import",
			"[ERROR]File type Error!",
			err)
		return

	}
	flag, _, err := api.ImportAPIGroupFromAms(pjID, userID, groupInfo)
	if !flag {

		controller.WriteError(httpResponse,
			"310000",
			"import",
			"[ERROR]Fail to import api group!",
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "import", "", nil)

}

//ImportProjectFromAms 导入项目
func ImportProjectFromAms(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	contentType := httpRequest.Header.Get("Content-Type")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if !strings.Contains(contentType, "multipart/form-data") {
		controller.WriteError(httpResponse,
			"310001",
			"import",
			"[ERROR]Request Content-Type isn't multipart/form-data",
			nil)
		return
	}
	file, _, err := httpRequest.FormFile("file")
	if err != nil {
		controller.WriteError(httpResponse,
			"310004",
			"import",
			"[ERROR]Param file does not exist!",
			err)
		return

	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		controller.WriteError(httpResponse,
			"310005",
			"import",
			"[ERROR]Fail to read file!",
			err)
		return

	}
	var projectInfo entity.AmsProject
	err = json.Unmarshal(body, &projectInfo)
	if err != nil {
		controller.WriteError(httpResponse,
			"310003",
			"import",
			"[ERROR]Fail to parse json!",
			err)
		return

	}
	if projectInfo.ProjectInfo.ProjectName == "" {
		controller.WriteError(httpResponse,
			"310006",
			"import",
			"[ERROR]File type Error!",
			nil)
		return

	}
	flag, _, err := api.ImportProjectFromAms(userID, projectInfo)
	if !flag {

		controller.WriteError(httpResponse,
			"310000",
			"import",
			"[ERROR]Fail to import project!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "import", "", nil)

	return
}

//ImportAPIFromAms 导入接口
func ImportAPIFromAms(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	contentType := httpRequest.Header.Get("Content-Type")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if !strings.Contains(contentType, "multipart/form-data") {
		controller.WriteError(httpResponse,
			"310001",
			"import",
			"[ERROR]Request Content-Type isn't multipart/form-data",
			nil)
		return

	}
	projectID := httpRequest.PostFormValue("projectID")
	groupID := httpRequest.PostFormValue("groupID")

	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"310002",
			"import",
			"[ERROR]Illegal projectID!",
			err)
		return

	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"310007",
			"import",
			"[ERROR]Illegal groupID",
			err)
		return

	}
	file, _, err := httpRequest.FormFile("file")
	if err != nil {
		controller.WriteError(httpResponse,
			"310004",
			"import",
			"[ERROR]Param file does not exist!",
			err)
		return

	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		controller.WriteError(httpResponse,
			"310005",
			"import",
			"[ERROR]Fail to read file!",
			err)
		return

	}
	apiList := make([]entity.AmsAPIInfo, 0)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &apiList)
	if err != nil {
		controller.WriteError(httpResponse,
			"310003",
			"import",
			"[ERROR]Fail to parse json!",
			err)
		return

	}
	if len(apiList) == 0 {
		controller.WriteError(httpResponse,
			"310006",
			"import",
			"[ERROR]File type Error!",
			nil)
		return

	}
	flag, _, err := api.ImportAPIFromAms(pjID, gID, userID, apiList)
	if !flag {

		controller.WriteError(httpResponse,
			"310000",
			"import",
			"[ERROR]Fail to import api!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "import", "", nil)

	return
}
