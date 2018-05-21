package conf

import (
	"sort"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Group struct {
	GroupList 				[]*GroupInfo				`json:"group" yaml:"group"`
	CancelDefaultGroup			bool						`json:"cancel_default_group" yaml:"cancel_default_group"`
}

type GroupInfo struct {
	GroupID					int						`json:"groupID" yaml:"group_id"`
	GroupName				string					`json:"groupName" yaml:"group_name"`
}

// 按照 Person.Age 从大到小排序
type GroupSlice []*GroupInfo
 
func (a GroupSlice) Len() int {    // 重写 Len() 方法
    return len(a)
}
func (a GroupSlice) Swap(i, j int){     // 重写 Swap() 方法
    a[i], a[j] = a[j], a[i]
}
func (a GroupSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[i].GroupID > a[j].GroupID
}

// 读入后端信息
func ParseApiGroupInfo(path string) ([]*GroupInfo,map[int]*GroupInfo,bool) {
	groupInfo := make(map[int]*GroupInfo)
	groupList := make([]*GroupInfo,0)
	var group Group
	content,err := ioutil.ReadFile(path)
	if err != nil {
		groupList = append(groupList,&GroupInfo{
			GroupID:0,
			GroupName:"默认分组",
		})
		groupInfo[0] = &GroupInfo{
			GroupID:0,
			GroupName:"默认分组",
		}
		return groupList,groupInfo,false
	}

	err = yaml.Unmarshal(content,&group)
	if err != nil {
		panic(err)
	}

	for _,g := range group.GroupList {
		groupInfo[g.GroupID] = g
	}
	

	if len(group.GroupList) != 0 {
		groupList = group.GroupList
	}
	_,ok := groupInfo[0]
	if !ok && !group.CancelDefaultGroup {
		groupList = append(groupList,&GroupInfo{
			GroupID:0,
			GroupName:"默认分组",
		})
		groupInfo[0] = &GroupInfo{
			GroupID:0,
			GroupName:"默认分组",
		}
	}

	
	sort.Sort(sort.Reverse(GroupSlice(groupList)))
	return groupList,groupInfo,group.CancelDefaultGroup
}
