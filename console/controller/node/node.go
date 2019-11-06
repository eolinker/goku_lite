package node

import (
	"encoding/json"
	"errors"
	"github.com/eolinker/goku-api-gateway/common/auto-form"

	"github.com/eolinker/goku-api-gateway/console/module/cluster"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/utils"
)

//AddNode 新增节点信息
func AddNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

	cluserName := httpRequest.PostFormValue("cluster")

	clusterID := cluster.GetClusterIDByName(cluserName)
	if clusterID == 0 {
		controller.WriteError(httpResponse, "340003", "", "[ERROR]Illegal cluster!", nil)
		return
	}

	//nodeNumber := rsa.CertConf["nodeNumber"].(int)
	type NodeParam struct {
		NodeName string `opt:"nodeName,require"`
		ListenAddress string `opt:"listenAddress,require"`
		AdminAddress string `opt:"adminAddress,require"`
		GroupID int `opt:"groupID,require"`
		Path string `opt:"gatewayPath"`
	}
	//
	//nodeName := httpRequest.PostFormValue("nodeName")
	//listenAddress := httpRequest.PostFormValue("listenAddress")
	//adminAddress := httpRequest.PostFormValue("adminAddress")
	//groupID := httpRequest.PostFormValue("groupID")
	//gatewayPath := httpRequest.PostFormValue("gatewayPath")
	//
	//gID, err := strconv.Atoi(groupID)
	//if err != nil && groupID != "" {
	//	controller.WriteError(httpResponse, "230015", "", "[ERROR]Illegal groupID!", err)
	//	return
	//}
	param:=new(NodeParam)
	err:=auto.SetValues(httpRequest.Form,param)
	if err!= nil{
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

	flag, result, err := node.AddNode(clusterID, param.NodeName, param.ListenAddress, param.AdminAddress, param.Path, param.GroupID)

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

	_,_=httpResponse.Write(data)
}

//EditNode 修改节点信息
func EditNode(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

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

	//exits := node.CheckIsExistRemoteAddr(id, listenAddress, adminAddress)
	//if exits {
	//
	//	controller.WriteError(httpResponse, "230005", "node", "[ERROR]The remote address is existed!", nil)
	//	return
	//
	//}

	flag, result, _ := node.EditNode(nodeName, listenAddress, adminAddress, gatewayPath, id, gID)

	if !flag {
		controller.WriteError(httpResponse, "330000", "node", result, nil)
		return
	}

	controller.WriteResultInfo(httpResponse, "node", "", nil)
	return
}

//DeleteNode 删除节点信息
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

	clusterID := cluster.GetClusterIDByName(cluserName)
	if clusterID == 0 {
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

//GetNodeInfo 获取节点信息
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
	  result, err := node.GetNodeInfo(id)
	if err!= nil {

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

//BatchDeleteNode 批量删除节点
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
