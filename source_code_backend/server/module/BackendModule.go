package module

import (
	"goku-ce/server/dao"
	"goku-ce/server/conf"
)

// 新增后端
func AddBackend(gatewayAlias,backendName,backendPath string) (bool,int) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.AddBackend(gateway["backendConfPath"],backendName,backendPath)
	}else {
		return false,0
	}
}

// 修改后端信息
func EditBackend(gatewayAlias,backendName,backendPath string,backendID int) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditBackend(gateway["backendConfPath"],backendName,backendPath,gatewayAlias,backendID)
	}else {
		return false
	}
}

// 删除后端信息
func DeleteBackend(gatewayAlias string,backendID int) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.DeleteBackend(gateway["backendConfPath"],backendID)
	}else {
		return false
	}
}

// 获取后端列表
func GetBackendList(gatewayAlias string) ([]*conf.BackendInfo){
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetBackendList(gateway["backendConfPath"])
	}else {
		return make([]*conf.BackendInfo,0)
	}
}

// 获取后端信息
func GetBackendInfo(gatewayAlias string,backendiID int) (bool,*conf.BackendInfo) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetBackendInfo(gateway["backendConfPath"],backendiID)
	}else {
		return false,&conf.BackendInfo{}
	}
}