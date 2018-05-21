package dao

import (
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
)

// 新增后端
func AddBackend(backendConfPath,backendName,backendPath string) (bool,int) {
	_,backend := conf.ParseBackendInfo(backendConfPath)
	maxID := 0
	for id,_ := range backend {
		if id > maxID {
			maxID = id
		} 
	}
	backendID := maxID + 1
	backend[backendID] = &conf.BackendInfo{
		BackendID : backendID,
		BackendName : backendName,
		BackendPath : backendPath,
	}

	backendList := make([]*conf.BackendInfo,0)
	for _,value := range backend {
		backendList = append(backendList,value)
	}

	content, err :=  yaml.Marshal(conf.Backend{
		BackendList: backendList,
	})
	if err != nil {
		panic(err);
	}

	conf.WriteConfigToFile(backendConfPath,content)
	return true,backendID
}

// 修改后端信息
func EditBackend(backendConfPath,backendName,backendPath,gatewayAlias string,backendID int) (bool) {
	_,backend := conf.ParseBackendInfo(backendConfPath)
	value,ok := backend[backendID]
	if !ok {
		return false
	} else {
		value.BackendName = backendName
		value.BackendPath = backendPath
	}

	backendList := make([]*conf.BackendInfo,0)
	for _,value := range backend {
		backendList = append(backendList,value)
	}

	content, err :=  yaml.Marshal(conf.Backend{
		BackendList: backendList,
	})
	if err != nil {
		panic(err);
	}
	conf.WriteConfigToFile(backendConfPath,content)
	return true
}

// 删除后端信息
func DeleteBackend(backendConfPath string,backendID int) (bool) {
	_,backend := conf.ParseBackendInfo(backendConfPath)
	_,ok := backend[backendID]
	if !ok {
		return false
	} else {
		delete(backend,backendID)
	}

	backendList := make([]*conf.BackendInfo,0)
	for _,value := range backend {
		backendList = append(backendList,value)
	}

	content, err :=  yaml.Marshal(conf.Backend{
		BackendList: backendList,
	})
	if err != nil {
		panic(err);
	}
	conf.WriteConfigToFile(backendConfPath,content)
	return true
}

// 获取后端信息
func GetBackendInfo(backendConfPath string,backendID int) (bool,*conf.BackendInfo){
	_,backend := conf.ParseBackendInfo(backendConfPath)
	value,ok := backend[backendID]
	if !ok {
		return false,&conf.BackendInfo{}
	} 

	return true,value
}

// 获取后端列表
func GetBackendList(backendConfPath string) ([]*conf.BackendInfo) {
	backendList,_ := conf.ParseBackendInfo(backendConfPath)
	return backendList
}
