package module

import (
	"goku-ce-1.0/utils"
	"goku-ce-1.0/server/dao"
)


// 添加分组
func AddGroup(gatewayID ,parentGroupID int,groupName string) (bool,int){
	if parentGroupID == 0{
		return dao.AddGroup(gatewayID,groupName)
	}else{
		return dao.AddChildGroup(gatewayID,parentGroupID,groupName)
	}
}

// 删除网关api分组
func DeleteGroup(groupID int) bool{
	return dao.DeleteGroup(groupID)
}

// 获取网关分组列表
func GetGroupList(gatewayID int) (bool,[]*utils.GroupInfo){
	return dao.GetGroupList(gatewayID)
}

// 修改分组信息
func EditGroup(groupID,parentGroupID int,groupName string) bool{
	return dao.EditGroup(groupID,parentGroupID,groupName)
}

// 获取分组名称
func GetGroupName(groupID int) (bool,string){
	return dao.GetGroupName(groupID)
}

// 获取网关分组列表
func GetGroupListByKeyword(keyword string,gatewayID int) (bool,[]*utils.GroupInfo){
	return dao.GetGroupListByKeyword(keyword,gatewayID)
}