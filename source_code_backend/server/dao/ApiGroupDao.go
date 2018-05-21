package dao

import (
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
)

// 新增分组
func AddApiGroup(groupConfPath,groupName string) (bool,int) {
	_,group,cancelDefaultGroup := conf.ParseApiGroupInfo(groupConfPath)
	maxID := 0
	for id,_ := range group {
		if id > maxID {
			maxID = id
		} 
	}
	groupID := maxID + 1
	group[groupID] = &conf.GroupInfo{
		GroupID : groupID,
		GroupName : groupName,
	}

	groupList := make([]*conf.GroupInfo,0)
	for _,value := range group {
		groupList = append(groupList,value)
	}
	groupConf := conf.Group{}
	groupConf.CancelDefaultGroup = cancelDefaultGroup
	groupConf.GroupList = groupList

	content, err :=  yaml.Marshal(groupConf)
	if err != nil {
		panic(err);
	}
	conf.WriteConfigToFile(groupConfPath,content)
	return true,groupID
}

// 修改分组
func EditApiGroup(groupConfPath,groupName string,groupID int) (bool) {
	_,group,cancelDefaultGroup := conf.ParseApiGroupInfo(groupConfPath)
	value,ok := group[groupID]
	if !ok {
		return false
	} else {
		value.GroupName = groupName
	}
	groupList := make([]*conf.GroupInfo,0)
	for _,v := range group {
		groupList = append(groupList,v)
	}
	groupConf := conf.Group{}
	groupConf.CancelDefaultGroup = cancelDefaultGroup
	groupConf.GroupList = groupList

	content, err :=  yaml.Marshal(groupConf)
	if err != nil {
		return false
	}
	conf.WriteConfigToFile(groupConfPath,content)
	return true
}

// 删除分组
func DeleteApiGroup(groupConfPath string,groupID int) (bool) {
	_,group,cancelDefaultGroup := conf.ParseApiGroupInfo(groupConfPath)
	_,ok := group[groupID]
	if !ok {
		return false
	} else {
		delete(group,groupID)
	}
	
	if groupID == 0 {
		cancelDefaultGroup = true
	}
	groupList := make([]*conf.GroupInfo,0)
	for _,v := range group {
		groupList = append(groupList,v)
	}
	groupConf := conf.Group{}
	groupConf.CancelDefaultGroup = cancelDefaultGroup
	groupConf.GroupList = groupList

	content, err :=  yaml.Marshal(groupConf)

	if err != nil {
		return false
	}
	conf.WriteConfigToFile(groupConfPath,content)
	return true
}


// 获取分组列表
func GetApiGroupList(groupConfPath string) []*conf.GroupInfo{
	groupList,_,_ := conf.ParseApiGroupInfo(groupConfPath)
	return groupList
}

// 获取api分组信息
func GetApiGroupInfo(groupConfPath string,groupID int) (bool,*conf.GroupInfo){
	_,group,_ := conf.ParseApiGroupInfo(groupConfPath)
	value,ok := group[groupID]
	if !ok {
		return false,&conf.GroupInfo{}
	} 

	return true,value
}

// 获取接口数量
func GetApiGroupCount(groupConfPath string) int {
	groupList,_,_ := conf.ParseApiGroupInfo(groupConfPath)
	return len(groupList)
}

