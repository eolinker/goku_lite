package node

//AddNodeGroup 新建节点分组
func AddNodeGroup(groupName string, clusterID int) (bool, interface{}, error) {
	return nodeGroupDao.AddNodeGroup(groupName, clusterID)
}

//EditNodeGroup 修改节点分组信息
func EditNodeGroup(groupName string, groupID int) (bool, string, error) {
	return nodeGroupDao.EditNodeGroup(groupName, groupID)
}

//DeleteNodeGroup 删除节点分组
func DeleteNodeGroup(groupID int) (bool, string, error) {
	return nodeGroupDao.DeleteNodeGroup(groupID)
}

//GetNodeGroupInfo 获取节点分组信息
func GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error) {
	return nodeGroupDao.GetNodeGroupInfo(groupID)
}

//GetNodeGroupList 获取节点分组列表
func GetNodeGroupList(clusterID int) (bool, []map[string]interface{}, error) {
	return nodeGroupDao.GetNodeGroupList(clusterID)
}

//CheckNodeGroupIsExist 检查节点分组是否存在
func CheckNodeGroupIsExist(groupID int) (bool, error) {
	return nodeGroupDao.CheckNodeGroupIsExist(groupID)
}

//GetRunningNodeCount 获取分组内启动节点数量
func GetRunningNodeCount(groupID int) (bool, interface{}, error) {
	return nodeGroupDao.GetRunningNodeCount(groupID)
}
