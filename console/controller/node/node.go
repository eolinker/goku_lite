package node

import (
	"encoding/json"
	"errors"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	cluster2 "github.com/eolinker/goku-api-gateway/server/cluster"
	"github.com/eolinker/goku-api-gateway/utils"
)

// 新增节点信息
func AddNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

	cluserName := httpRequest.PostFormValue("cluster")

	clusterId, has := cluster2.GetId(cluserName)
	if !has {
		controller.WriteError(httpResponse, "340003", "", "[ERROR]Illegal cluster!", nil)
		return
	}

	//nodeNumber := rsa.CertConf["nodeNumber"].(int)

	nodeName := httpRequest.PostFormValue("nodeName")
	nodeIP := httpRequest.PostFormValue("nodeIP")
	nodePort := httpRequest.PostFormValue("nodePort")
	groupID := httpRequest.PostFormValue("groupID")
	gatewayPath := httpRequest.PostFormValue("gatewayPath")

	gID, err := strconv.Atoi(groupID)
	if err != nil && groupID != "" {
		controller.WriteError(httpResponse, "230015", "", "[ERROR]Illegal groupID!", err)
		return
	}

	flag := utils.ValidateRemoteAddr(nodeIP + ":" + nodePort)
	if !flag {
		controller.WriteError(httpResponse,
			"230006",
			"node", "[ERROR]Illegal remote address!",
			errors.New("Illegal remote address"))
		return
	}
	if gID != 0 {
		// 检查分组是否存在
		flag, err = node.CheckNodeGroupIsExist(gID)

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

	exits := node.CheckIsExistRemoteAddr(0, nodeIP, nodePort)

	if exits {
		controller.WriteError(httpResponse,
			"230005",
			"node",
			"[ERROR]The remote address is alreadey existed!",
			errors.New("The remote address is alreadey existed"))
		return
	}

	flag, result, err := node.AddNode(clusterId, nodeName, nodeIP, nodePort, gatewayPath, gID)

	if !flag {
		controller.WriteError(httpResponse,
			"330000",
			"node",
			result["error"].(string),
			err)
		return
	}

	res := map[string]interface{}{
		"nodeID":     result["nodeID"],
		"version":    result["version"],
		"statusCode": "000000",
		"type":       "node",
		"resultDesc": "",
	}
	data, _ := json.Marshal(res)

	httpResponse.Write(data)
}

// 修改节点信息
func EditNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

	nodeName := httpRequest.PostFormValue("nodeName")
	nodeIP := httpRequest.PostFormValue("nodeIP")
	nodePort := httpRequest.PostFormValue("nodePort")
	groupID := httpRequest.PostFormValue("groupID")
	nodeID := httpRequest.PostFormValue("nodeID")

	gatewayPath := httpRequest.PostFormValue("gatewayPath")
	// key := httpRequest.PostFormValue("key")

	flag := utils.ValidateRemoteAddr(nodeIP + ":" + nodePort)
	if !flag {

		controller.WriteError(httpResponse, "230006", "node", "[ERROR]Illegal remote address!", errors.New("[ERROR]Illegal remote address!"))
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
		flag, err = node.CheckNodeGroupIsExist(gID)

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

	exits := node.CheckIsExistRemoteAddr(id, nodeIP, nodePort)
	if exits {

		controller.WriteError(httpResponse, "230005", "node", "[ERROR]The remote address is existed!", nil)
		return

	}

	flag, result, _ := node.EditNode(nodeName, nodeIP, nodePort, gatewayPath, id, gID)

	if !flag {
		controller.WriteError(httpResponse, "330000", "node", result, nil)
		return
	}

	controller.WriteResultInfo(httpResponse, "node", "", nil)
	return
}

// 删除节点信息
func DeleteNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {

		return
	}

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
	flag, result, err := node.DeleteNode(id)
	if !flag {

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

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	cluserName := httpRequest.Form.Get("cluster")
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

	clusterID, has := cluster2.GetId(cluserName)
	if !has {
		controller.WriteError(httpResponse, "330003", "node", "[ERROR]The cluster dosen't exist!", nil)
		return
	}

	flag, result, err := node.GetNodeList(clusterID, gID, keyword)
	if !flag {
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

// 获取节点信息
func GetNodeInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationREAD)
	if e != nil {
		return
	}

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
	flag, result, err := node.GetNodeInfo(id)
	if !flag {

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

// 节点IP查重
func CheckIsExistRemoteAddr(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationREAD)
	if e != nil {
		return
	}

	nodeIP := httpRequest.PostFormValue("nodeIP")
	nodePort := httpRequest.PostFormValue("nodePort")

	flag := utils.ValidateRemoteAddr(nodeIP)

	if !flag {
		controller.WriteError(httpResponse,
			"230006",
			"node",
			"[ERROR]The remote address does not exist!",
			nil)
		return
	}

	flag = node.CheckIsExistRemoteAddr(0, nodePort, nodePort)
	if !flag {

		controller.WriteError(httpResponse,
			"330000",
			"node",
			"[ERROR]Remote address is existed!",
			nil)
		return

	}
	controller.WriteResultInfo(httpResponse, "node", "", nil)

	return
}

// 批量修改节点分组
func BatchEditNodeGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

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
	flag, result, err := node.BatchEditNodeGroup(nodeIDList, gID)
	if !flag {

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

// 批量删除节点
func BatchDeleteNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {

		return
	}

	nodeIDList := httpRequest.PostFormValue("nodeIDList")

	flag, result, err := node.BatchDeleteNode(nodeIDList)
	if !flag {

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
