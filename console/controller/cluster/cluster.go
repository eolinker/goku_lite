package cluster

import (
	"net/http"
	"regexp"

	"github.com/pkg/errors"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/cluster"
)

//GetClusterList 获取集群列表
func GetClusterList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	list, _ := cluster.GetClusters()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		list)
}

//GetCluster 获取集群信息
func GetCluster(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")

	result, err := cluster.GetCluster(name)
	if err != nil {
		controller.WriteError(httpResponse, "370000", "cluster", "[ERROR]The cluster does not exist", err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusterInfo",
		result)
}

//AddCluster 新增集群
func AddCluster(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	title := httpRequest.Form.Get("title")
	note := httpRequest.Form.Get("note")
	match, err := regexp.MatchString(`^[a-zA-Z][a-zA-z0-9_]*$`, name)
	if err != nil || !match {
		controller.WriteError(httpResponse, "370001", "cluster", "[ERROR]Illegal cluster name", err)
		return
	}
	if cluster.CheckClusterNameIsExist(name) {
		controller.WriteError(httpResponse, "370003", "cluster", "[ERROR]Cluster name already exists", errors.New("[error]cluster name already exists"))
		return
	}
	err = cluster.AddCluster(name, title, note)
	if err != nil {
		controller.WriteError(httpResponse, "370000", "cluster", err.Error(), err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"",
		nil)
}

//EditCluster 新增集群
func EditCluster(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	title := httpRequest.Form.Get("title")
	note := httpRequest.Form.Get("note")
	match, err := regexp.MatchString(`^[a-zA-Z][a-zA-z0-9_]*$`, name)
	if err != nil || !match {
		controller.WriteError(httpResponse, "370001", "cluster", "[ERROR]Illegal cluster name", err)
		return
	}
	err = cluster.EditCluster(name, title, note)
	if err != nil {
		controller.WriteError(httpResponse, "370000", "cluster", err.Error(), err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"",
		nil)
}

//DeleteCluster 新增集群
func DeleteCluster(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	name := httpRequest.Form.Get("name")
	match, err := regexp.MatchString(`^[a-zA-Z][a-zA-z0-9_]*$`, name)
	if err != nil || !match {
		controller.WriteError(httpResponse, "370001", "cluster", "[ERROR]Illegal cluster name", err)
		return
	}
	nodeCount := cluster.GetClusterNodeCount(name)
	if nodeCount > 0 {
		controller.WriteError(httpResponse, "370002", "cluster", "[ERROR]There are nodes in the cluster", errors.New("[error]there are nodes in the cluster"))
		return
	}
	err = cluster.DeleteCluster(name)
	if err != nil {
		controller.WriteError(httpResponse, "370000", "cluster", err.Error(), err)
		return
	}

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"",
		nil)
}

//GetClusterInfoList 获取集群信息列表
func GetClusterInfoList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}
	result, _ := cluster.GetClusters()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		result)
}
