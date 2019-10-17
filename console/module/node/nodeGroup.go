package node

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//AddNodeGroup 新建节点分组
func AddNodeGroup(groupName string, clusterID int) (bool, interface{}, error) {
	return console_sqlite3.AddNodeGroup(groupName, clusterID)
}

//EditNodeGroup 修改节点分组信息
func EditNodeGroup(groupName string, groupID int) (bool, string, error) {
	return console_sqlite3.EditNodeGroup(groupName, groupID)
}

//DeleteNodeGroup 删除节点分组
func DeleteNodeGroup(groupID int) (bool, string, error) {
	return console_sqlite3.DeleteNodeGroup(groupID)
}

//GetNodeGroupInfo 获取节点分组信息
func GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error) {
	return console_sqlite3.GetNodeGroupInfo(groupID)
}

//GetNodeGroupList 获取节点分组列表
func GetNodeGroupList(clusterID int) (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetNodeGroupList(clusterID)
}

//CheckNodeGroupIsExist 检查节点分组是否存在
func CheckNodeGroupIsExist(groupID int) (bool, error) {
	return console_sqlite3.CheckNodeGroupIsExist(groupID)
}

//GetRunningNodeCount 获取分组内启动节点数量
func GetRunningNodeCount(groupID int) (bool, interface{}, error) {
	return console_sqlite3.GetRunningNodeCount(groupID)
}
