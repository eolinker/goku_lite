package node

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/node"
	cluster2 "github.com/eolinker/goku/server/cluster"
)

func AddNodeGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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
	groupName := httpRequest.PostFormValue("groupName")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"280002",
			"nodeGroup",
			"[ERROR]Illegal groupName!",
			nil)
		return

	}
	flag, result, err := node.AddNodeGroup(groupName, clusterId)
	if !flag {

		controller.WriteError(httpResponse,
			"280000",
			"nodeGroup",
			result.(string),
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "node", "groupID", result)

	return
}

// 修改节点分组信息
func EditNodeGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

	groupName := httpRequest.PostFormValue("groupName")
	groupID := httpRequest.PostFormValue("groupID")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"280002",
			"nodeGroup",
			"[ERROR]Illegal groupName!",
			nil)
		return

	}
	id, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"280001",
			"nodeGroup",
			"[ERROR]Illegal groupID!",
			err)
		return

	}

	flag, result, err := node.EditNodeGroup(groupName, id)
	if !flag {

		controller.WriteError(httpResponse,
			"280000",
			"nodeGroup",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "nodeGroup", "", nil)
}

// 删除节点分组
func DeleteNodeGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationEDIT)
	if e != nil {
		return
	}

	groupID := httpRequest.PostFormValue("groupID")

	id, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"280001",
			"nodeGroup",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	flag, result, err := node.GetRunningNodeCount(id)
	if !flag {
		controller.WriteError(httpResponse,
			"280013",
			"nodeGroup",
			result.(string),
			err)
		return

	}
	if result.(int) > 0 {
		controller.WriteError(httpResponse,
			"280013",
			"nodeGroup",
			"[ERROR]Contains running nodes",
			err)
		return
	}
	flag, result, err = node.DeleteNodeGroup(id)
	if !flag {

		controller.WriteError(httpResponse,
			"280000",
			"nodeGroup",
			result.(string),
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "nodeGroup", "", nil)

}

// 获取节点分组信息
func GetNodeGroupInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationREAD)
	if e != nil {
		return
	}

	groupID := httpRequest.PostFormValue("groupID")

	id, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"280001",
			"nodeGroup",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	flag, result, err := node.GetNodeGroupInfo(id)
	if !flag {

		controller.WriteError(httpResponse,
			"280000",
			"nodeGroup",
			"[ERROR]The node group information does not exist!",
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "node", "groupInfo", result)
}

// 获取节点分组列表
func GetNodeGroupList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNode, controller.OperationREAD)
	if e != nil {
		return
	}

	cluserName := httpRequest.FormValue("cluster")
	clusterId, has := cluster2.GetId(cluserName)
	if !has {
		controller.WriteError(httpResponse,
			"280001",
			"nodeGroup",
			"[ERROR]Illegal cluster!",
			nil)
		return
	}
	flag, result, err := node.GetNodeGroupList(clusterId)
	if !flag {

		controller.WriteError(httpResponse,
			"280000",
			"nodeGroup",
			"[ERROR]Empty group list!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "node", "groupList", result)
}
