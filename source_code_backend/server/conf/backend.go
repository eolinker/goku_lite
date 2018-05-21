package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Backend struct {
	BackendList				[]*BackendInfo			`json:"backend" yaml:"backend"`
}

type BackendInfo struct {
	BackendID				int						`json:"backendID" yaml:"backend_id"`
	BackendName				string					`json:"backendName" yaml:"backend_name"` 
	BackendPath				string					`json:"backendPath" yaml:"backend_path"`
}

// 读入后端信息
func ParseBackendInfo(path string) ([]*BackendInfo,map[int]*BackendInfo) {
	backendInfo := make(map[int]*BackendInfo)
	backendList := make([]*BackendInfo,0)
	var backend Backend
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return backendList,backendInfo
	}

	err = yaml.Unmarshal(content,&backend)
	if err != nil {
		panic(err)
	}

	
	if len(backend.BackendList) != 0 {
		backendList = backend.BackendList
	}
	for _,b := range backend.BackendList {
		backendInfo[b.BackendID] = b
	}
	
	return backendList,backendInfo
}