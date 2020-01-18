package node

import (
	"encoding/json"
	"errors"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/common/auto-form"
	"github.com/eolinker/goku-api-gateway/console/module/cluster"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/utils"
)

const operationNode = "nodeManagement"

//Handlers Handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":            factory.NewAccountHandleFunction(operationNode, true, AddNode),
		"/edit":           factory.NewAccountHandleFunction(operationNode, true, EditNode),
		"/delete":         factory.NewAccountHandleFunction(operationNode, true, DeleteNode),
		"/getInfo":        factory.NewAccountHandleFunction(operationNode, false, GetNodeInfo),
		"/getList":        factory.NewAccountHandleFunction(operationNode, false, GetNodeList),
		"/batchEditGroup": factory.NewAccountHandleFunction(operationNode, true, BatchEditNodeGroup),
		"/batchDelete":    factory.NewAccountHandleFunction(operationNode, true, BatchDeleteNode),
	}
}

//NewNodeHandlers new nodeHandlers
func NewNodeHandlers() *Handlers {
	return &Handlers{}
}

//AddNode 新增节点信息
func AddNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	cluserName := httpRequest.PostFormValue("cluster")

	clusterID := cluster.GetClusterIDByName(cluserName)
	if clusterID == 0 {
		controller.WriteError(httpResponse, "340003", "", "[ERROR]Illegal cluster!", nil)
		return
	}

	//nodeNumber := rsa.CertConf["nodeNumber"].(int)
	type NodeParam struct {
		NodeName      string `opt:"nodeName,require"`
		ListenAddress string `opt:"listenAddress,require"`
		AdminAddress  string `opt:"adminAddress,require"`
		GroupID       int    `opt:"groupID,require"`
		Path          string `opt:"gatewayPath"`
	}

	param := new(NodeParam)
	err := auto.SetValues(httpRequest.Form, param)
	if err != nil {
		controller.WriteError(httpResponse, "230015", "", "[ERROR]", err)
		return
	}
	if !utils.ValidateRemoteAddr(param.ListenAddress) {
		controller.WriteError(httpResponse,
			"230006",
			"node", "[ERROR]Illegal listenAddress!",
			errors.New("illegal listenAddress"))
		return
	}
	if !utils.ValidateRemoteAddr(param.AdminAddress) {
		controller.WriteError(httpResponse,
			"230007",
			"node", "[ERROR]Illegal listenAddress!",
			errors.New("illegal listenAddress"))
		return
	}
	if param.GroupID != 0 {
		// 检查分组是否存在
		flag, err := node.CheckNodeGroupIsExist(param.GroupID)

		if !flag {
			controller.WriteError(
				httpResponse,
				"230014",
				"node",
				"[ERROR]The node group does not exist!",
				err)
			return
		}
	}

	id, v, result, err := node.AddNode(clusterID, param.NodeName, param.ListenAddress, param.AdminAddress, param.Path, param.GroupID)

	if err != nil {
		controller.WriteError(httpResponse,
			"330000",
			"node",
			result,
			err)
		return
	}

	res := map[string]interface{}{
		"nodeID":     id,
		"version":    v,
		"statusCode": "000000",
		"type":       "node",
		"resultDesc": "",
	}
	data, _ := json.Marshal(res)

	_, _ = httpResponse.Write(data)
}

//EditNode 修改节点信息
func EditNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	nodeName := httpRequest.PostFormValue("nodeName")
	listenAddress := httpRequest.PostFormValue("listenAddress")
	adminAddress := httpRequest.PostFormValue("adminAddress")
	groupID := httpRequest.PostFormValue("groupID")
	nodeID := httpRequest.PostFormValue("nodeID")

	gatewayPath := httpRequest.PostFormValue("gatewayPath")
	// key := httpRequest.PostFormValue("key")

	if !utils.ValidateRemoteAddr(listenAddress) {
		controller.WriteError(httpResponse,
			"230006",
			"node", "[ERROR]Illegal listenAddress!",
			errors.New("illegal listenAddress"))
		return
	}
	if !utils.ValidateRemoteAddr(adminAddress) {
		controller.WriteError(httpResponse,
			"230007",
			"node", "[ERROR]Illegal listenAddress!",
			errors.New("illegal listenAddress"))
		return
	}

	id, err := strconv.Atoi(nodeID)
	if err != nil {

		controller.WriteError(httpResponse, "230001", "node", "[ERROR]Illegal nodeID!", err)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil && groupID != "" {
		controller.WriteError(httpResponse, "230015", "node", "[ERROR]Illegal groupID!", err)
		return
	}

	if gID != 0 {
		// 检查分组是否存在
		flag, err := node.CheckNodeGroupIsExist(gID)

		if !flag {
			controller.WriteError(
				httpResponse,
				"230014",
				"node",
				"[ERROR]The node group does not exist!",
				err)
			return
		}
	}

	result, err := node.EditNode(nodeName, listenAddress, adminAddress, gatewayPath, id, gID)

	if err != nil {
		controller.WriteError(httpResponse, "330000", "node", result, nil)
		return
	}

	controller.WriteResultInfo(httpResponse, "node", "", nil)
	return
}

//DeleteNode 删除节点信息
func DeleteNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	nodeID := httpRequest.PostFormValue("nodeID")

	id, err := strconv.Atoi(nodeID)
	if err != nil {
		controller.WriteError(httpResponse,
			"230001",
			"node",
			"[ERROR]Illegal nodeID!",
			err)
		return

	}
	result, err := node.DeleteNode(id)
	if err != nil {

		controller.WriteError(httpResponse,
			"330000",
			"node",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "node", "", nil)
}

// GetNodeList 获取节点列表
func GetNodeList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	clusterName := httpRequest.Form.Get("cluster")
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")

	gID, err := strconv.Atoi(groupID)
	if err != nil && groupID != "" {
		if groupID != "" {
			controller.WriteError(httpResponse, "330002", "node", "[ERROR]Illegal groupID!", nil)
			return
		}
		gID = -1
	}

	clusterID := cluster.GetClusterIDByName(clusterName)
	if clusterID == 0 {
		controller.WriteError(httpResponse, "330003", "node", "[ERROR]The cluster dosen't exist!", nil)
		return
	}

	result, err := node.GetNodeList(clusterID, gID, keyword)
	if err != nil {
		controller.WriteError(httpResponse,
			"330000",
			"node",
			"[ERROR]Empty node list!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "node", "nodeList", result)
	// controller.WriteResultInfo(httpResponse, "nodeList", result)

}

//GetNodeInfo 获取节点信息
func GetNodeInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	nodeID := httpRequest.PostFormValue("nodeID")

	id, err := strconv.Atoi(nodeID)
	if err != nil {

		log.Info("[ERROR] ", httpRequest.RequestURI, ":", err.Error())
		controller.WriteError(httpResponse,
			"230001",
			"node",
			"[ERROR]Illegal nodeID!",
			err)
		return
	}
	result, err := node.GetNodeInfo(id)
	if err != nil {

		controller.WriteError(httpResponse,
			"330000",
			"node",
			"[ERROR]The node does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "node", "nodeInfo", result)

	return
}

//BatchEditNodeGroup 批量修改节点分组
func BatchEditNodeGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	nodeIDList := httpRequest.PostFormValue("nodeIDList")
	groupID := httpRequest.PostFormValue("groupID")

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"230015",
			"node",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	result, err := node.BatchEditNodeGroup(nodeIDList, gID)
	if err != nil {

		controller.WriteError(httpResponse,
			"330000",
			"node",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "node", "", nil)

	return
}

//BatchDeleteNode 批量删除节点
func BatchDeleteNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	nodeIDList := httpRequest.PostFormValue("nodeIDList")

	result, err := node.BatchDeleteNode(nodeIDList)
	if err != nil {

		if result == "230013" {
			controller.WriteError(httpResponse,
				"230013",
				"node",
				"[ERROR]Can not find the avaliable node",
				err)
			return

		}
		controller.WriteError(httpResponse,
			"330000",
			"node",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "node", "", nil)

	return
}
