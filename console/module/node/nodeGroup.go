package node

import (
	"github.com/eolinker/goku/server/dao/console-mysql"
)

// 新建节点分组
func AddNodeGroup(groupName string, clusterId int) (bool, interface{}, error) {
	return console_mysql.AddNodeGroup(groupName, clusterId)
}

// 修改节点分组信息
func EditNodeGroup(groupName string, groupID int) (bool, string, error) {
	return console_mysql.EditNodeGroup(groupName, groupID)
}

// 删除节点分组
func DeleteNodeGroup(groupID int) (bool, string, error) {
	return console_mysql.DeleteNodeGroup(groupID)
}

// 获取节点分组信息
func GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error) {
	return console_mysql.GetNodeGroupInfo(groupID)
}

// 获取节点分组列表
func GetNodeGroupList(clusterId int) (bool, []map[string]interface{}, error) {
	return console_mysql.GetNodeGroupList(clusterId)
}

// 检查节点分组是否存在
func CheckNodeGroupIsExist(groupID int) (bool, error) {
	return console_mysql.CheckNodeGroupIsExist(groupID)
}

// 获取分组内启动节点数量
func GetRunningNodeCount(groupID int) (bool, interface{}, error) {
	return console_mysql.GetRunningNodeCount(groupID)
}
