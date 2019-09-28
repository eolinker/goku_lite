package node

import (
	"github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//AddNodeGroup 新建节点分组
func AddNodeGroup(groupName string, clusterID int) (bool, interface{}, error) {
	return consolemysql.AddNodeGroup(groupName, clusterID)
}

//EditNodeGroup 修改节点分组信息
func EditNodeGroup(groupName string, groupID int) (bool, string, error) {
	return consolemysql.EditNodeGroup(groupName, groupID)
}

//DeleteNodeGroup 删除节点分组
func DeleteNodeGroup(groupID int) (bool, string, error) {
	return consolemysql.DeleteNodeGroup(groupID)
}

//GetNodeGroupInfo 获取节点分组信息
func GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error) {
	return consolemysql.GetNodeGroupInfo(groupID)
}

//GetNodeGroupList 获取节点分组列表
func GetNodeGroupList(clusterID int) (bool, []map[string]interface{}, error) {
	return consolemysql.GetNodeGroupList(clusterID)
}

//CheckNodeGroupIsExist 检查节点分组是否存在
func CheckNodeGroupIsExist(groupID int) (bool, error) {
	return consolemysql.CheckNodeGroupIsExist(groupID)
}

//GetRunningNodeCount 获取分组内启动节点数量
func GetRunningNodeCount(groupID int) (bool, interface{}, error) {
	return consolemysql.GetRunningNodeCount(groupID)
}
