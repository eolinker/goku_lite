package cluster

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"

	"github.com/eolinker/goku-api-gateway/console/controller"
)

//GetVersionList 获取版本列表
func GetVersionList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
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
	id, err := versionConfig.AddVersionConfig(name, version, remark, now)
	if err != nil {
		controller.WriteError(httpResponse, "380000", "versionConfig", err.Error(), err)
		return
	}

	if p == 1 {
		versionConfig.PublishVersion(id, now)
	}
	controller.WriteResultInfo(httpResponse,
		"versionConfig",
		"",
		nil)
	return
}

//BatchDeleteVersionConfig 批量删除版本配置
func BatchDeleteVersionConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
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
	id, err := strconv.Atoi(versionID)
	if err != nil {
		controller.WriteError(httpResponse, "380002", "versionConfig", "[ERROR]Illegal versionID", err)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	err = versionConfig.PublishVersion(id, now)
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
