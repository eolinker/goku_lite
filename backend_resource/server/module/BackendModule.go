package module

import (
	"goku-ce-1.0/utils"
	"goku-ce-1.0/server/dao"
)

// 获取环境列表
func GetBackendList(gatewayID int) (bool,[]*utils.BackendInfo){
	return dao.GetBackendList(gatewayID)
}

// 添加环境
func AddBackend(gatewayID int,backendName ,backendURI string) (bool,int){
	return dao.AddBackend(gatewayID,backendName,backendURI)
}

func DeleteBackend(gatewayID,backendID int) bool{
	return dao.DeleteBackend(gatewayID,backendID)
}

func EditBackend(backendID, gatewayID int,backendName,backendURI,gatewayHashKey string) bool{
	return dao.EditBackend(backendID,gatewayID,backendName,backendURI,gatewayHashKey)
}

// 获取网关信息
func GetBackendInfo(backendID int) (bool,utils.BackendInfo){
	return dao.GetBackendInfo(backendID)
}