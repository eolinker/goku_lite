package cluster

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"

	"github.com/eolinker/goku-api-gateway/console/controller"
)

const operationVersion = "versionManagement"

//VersionHandlers 版本处理器
type VersionHandlers struct {
}

//Handlers handlers
func (h *VersionHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":        factory.NewAccountHandleFunction(operationVersion, true, AddVersionConfig),
		"/basic/edit": factory.NewAccountHandleFunction(operationVersion, true, EditVersionBasicConfig),
		"/delete":     factory.NewAccountHandleFunction(operationVersion, true, BatchDeleteVersionConfig),
		"/getList":    factory.NewAccountHandleFunction(operationVersion, false, GetVersionList),
		"/publish":    factory.NewAccountHandleFunction(operationVersion, true, PublishVersion),
	}
}

//NewVersionHandlers new versionHandlers
func NewVersionHandlers() *VersionHandlers {
	return &VersionHandlers{}
}

//GetVersionList 获取版本列表
func GetVersionList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	keyword := httpRequest.Form.Get("keyword")
	result, _ := versionConfig.GetVersionList(keyword)
	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"configList",
		result)
}

//AddVersionConfig 新增版本配置
func AddVersionConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	version := httpRequest.Form.Get("version")
	remark := httpRequest.Form.Get("remark")
	publish := httpRequest.Form.Get("publish")
	p, err := strconv.Atoi(publish)
	if err != nil && publish != "" {
		controller.WriteError(httpResponse, "380003", "versionConfig", "[ERROR]Illegal publish", err)
		return
	}
	//count := cluster.GetVersionConfigCount()
	now := time.Now().Format("2006-01-02 15:04:05")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	id, err := versionConfig.AddVersionConfig(name, version, remark, now, userID)
	if err != nil {
		controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
		return
	}

	if p == 1 {
		versionConfig.PublishVersion(id, userID, now)
	}
	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"",
		nil)
	return
}

//EditVersionBasicConfig 新增版本配置
func EditVersionBasicConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	version := httpRequest.Form.Get("version")
	remark := httpRequest.Form.Get("remark")
	versionID := httpRequest.Form.Get("versionID")
	id, err := strconv.Atoi(versionID)
	if err != nil {
		controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
		return
	}

	userID := goku_handler.UserIDFromRequest(httpRequest)
	err = versionConfig.EditVersionBasicConfig(name, version, remark, userID, id)
	if err != nil {
		controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"",
		nil)
	return
}

//BatchDeleteVersionConfig 批量删除版本配置
func BatchDeleteVersionConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	ids := httpRequest.Form.Get("ids")
	idList := make([]int, 0, 10)
	err := json.Unmarshal([]byte(ids), &idList)
	if err != nil {
		controller.WriteError(httpResponse, "380001", "versionConfig", "[ERROR]Illegal ids", err)
		return
	}
	if len(idList) > 0 {
		err = versionConfig.BatchDeleteVersionConfig(idList)
		if err != nil {
			controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
			return
		}
	}

	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"",
		nil)
	return
}

//PublishVersion 发布版本
func PublishVersion(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpRequest.ParseForm()
	versionID := httpRequest.Form.Get("versionID")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	id, err := strconv.Atoi(versionID)
	if err != nil {
		controller.WriteError(httpResponse, "380002", "versionConfig", "[ERROR]Illegal versionID", err)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	err = versionConfig.PublishVersion(id, userID, now)
	if err != nil {
		controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"",
		nil)
	return
}
