package cluster

import (
	"net/http"
	"regexp"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/pkg/errors"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/cluster"
)

const operationCluster = "nodeManagement"

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":        factory.NewAccountHandleFunction(operationCluster, true, AddCluster),
		"/edit":       factory.NewAccountHandleFunction(operationCluster, true, EditCluster),
		"/delete":     factory.NewAccountHandleFunction(operationCluster, true, DeleteCluster),
		"/list":       factory.NewAccountHandleFunction(operationCluster, false, GetClusterInfoList),
		"/simpleList": factory.NewAccountHandleFunction(operationCluster, false, GetClusterList),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

//GetClusterList 获取集群列表
func GetClusterList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	list, _ := cluster.GetClusters()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		list)
}

//GetCluster 获取集群信息
func GetCluster(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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

	result, _ := cluster.GetClusters()

	controller.WriteResultInfo(httpResponse,
		"cluster",
		"clusters",
		result)
}
