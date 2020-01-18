package node

import (
	"time"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
	"github.com/eolinker/goku-api-gateway/utils"
)

//AddNode 新增节点信息
func AddNode(clusterID int, nodeName, listenAddress, adminAddress, gatewayPath string, groupID int) (int64, string, string, error) {
	now := time.Now().Format("20060102150405")
	nodeKey := utils.Md5(utils.GetRandomString(16) + now)
	return nodeDao.AddNode(clusterID, nodeName, nodeKey, listenAddress, adminAddress, gatewayPath, groupID)
}

//EditNode 修改节点
func EditNode(nodeName, listenAddress, adminAddress, gatewayPath string, nodeID, groupID int) (string, error) {
	return nodeDao.EditNode(nodeName, listenAddress, adminAddress, gatewayPath, nodeID, groupID)
}

//DeleteNode 删除节点
func DeleteNode(nodeID int) (string, error) {
	return nodeDao.DeleteNode(nodeID)
}

//GetNodeInfo 获取节点信息
func GetNodeInfo(nodeID int) (*entity.Node, error) {
	node, e := nodeDao.GetNodeInfo(nodeID)
	if e != nil {
		return nil, e
	}
	ResetNodeStatus(node)
	return node, e
}

//GetNodeInfoByKey 获取节点信息
func GetNodeInfoByKey(nodeKey string) (*entity.Node, error) {
	node, e := nodeDao.GetNodeByKey(nodeKey)
	if e != nil {
		return nil, e
	}
	ResetNodeStatus(node)
	return node, e
}

// GetNodeList 获取节点列表
func GetNodeList(clusterID, groupID int, keyword string) ([]*entity.Node, error) {
	nodes, e := nodeDao.GetNodeList(clusterID, groupID, keyword)
	if e != nil {
		return nil, e
	}
	ResetNodeStatus(nodes...)
	return nodes, e
}

//GetAllNode 获取所有节点
func GetAllNode() ([]*entity.Node, error) {
	nodes, e := nodeDao.GetNodeInfoAll()
	if e != nil {
		return nil, e
	}

	return nodes, e
}

//BatchDeleteNode 批量删除节点
func BatchDeleteNode(nodeIDList string) (string, error) {
	nodeIDList, err := nodeDao.GetAvaliableNodeListFromNodeList(nodeIDList, 0)
	if err != nil {
		return err.Error(), err
	}
	if nodeIDList == "" {
		return "230013", err
	}
	return nodeDao.BatchDeleteNode(nodeIDList)
}

//BatchEditNodeGroup 批量修改节点分组
func BatchEditNodeGroup(nodeIDList string, groupID int) (string, error) {
	return nodeDao.BatchEditNodeGroup(nodeIDList, groupID)
}
