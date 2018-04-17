package module

import (
	"goku-ce-1.0/utils"
	"goku-ce-1.0/server/dao"
	"encoding/json"
)

// 新增api
func AddApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote string,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody int,gatewayRequestParam []utils.GatewayParam,constantResultParam []utils.ConstantMapping) (bool,int){
	if flag,_ := dao.CheckGatewayURLIsExist(gatewayID,gatewayRequestPath); flag == true{
		return false,0
	}
	var cacheJson utils.ApiCacheJson
	cacheJson.ApiName = apiName
	cacheJson.BackendID = backendID
	cacheJson.BackendPath = backendRequestPath
	cacheJson.BackendProtocol = backendProtocol
	cacheJson.BackendRequestType = backendRequestType
	cacheJson.BackendURI = backendRequestURI
	cacheJson.RequestParams = gatewayRequestParam
	cacheJson.ConstantParams = constantResultParam
	cacheJson.IsRequestBody = isRequestBody
	cacheJson.GatewayHashKey = gatewayHashKey
	cacheJson.GatewayProtocol = gatewayProtocol
	cacheJson.GatewayRequestBodyNote = gatewayRequestBodyNote
	cacheJson.GatewayRequestPath = gatewayRequestPath
	cacheJson.GatewayRequestType = gatewayRequestType
	apiCacheJson,_ := json.Marshal(cacheJson)
	apiString := string(apiCacheJson[:])

	var redisCacheJson utils.RedisCacheJson
	redisCacheJson.BackendPath = backendRequestPath
	redisCacheJson.BackendProtocol = backendProtocol
	redisCacheJson.BackendRequestType = backendRequestType
	redisCacheJson.BackendURI = backendRequestURI
	redisCacheJson.ConstantParams = constantResultParam
	redisCacheJson.GatewayHashKey = gatewayHashKey
	redisCacheJson.GatewayID = gatewayID
	redisCacheJson.GatewayRequestPath = gatewayRequestPath
	redisCacheJson.IsRequestBody = isRequestBody
	redisCacheJson.RequestParams = gatewayRequestParam
	redisJson,_ := json.Marshal(redisCacheJson)
	redisString := string(redisJson[:])
	return dao.AddApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,apiString,redisString,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody,gatewayRequestParam,constantResultParam)
}

// 修改api
func EditApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote string,apiID,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody int,gatewayRequestParam []utils.GatewayParam,constantResultParam []utils.ConstantMapping) (bool,int){
	if flag,id := dao.CheckGatewayURLIsExist(gatewayID,gatewayRequestPath); flag == true && id != apiID{
		return false,0
	}
	var cacheJson utils.ApiCacheJson
	cacheJson.ApiName = apiName
	cacheJson.BackendID = backendID
	cacheJson.BackendPath = backendRequestPath
	cacheJson.BackendProtocol = backendProtocol
	cacheJson.BackendRequestType = backendRequestType
	cacheJson.BackendURI = backendRequestURI
	cacheJson.RequestParams = gatewayRequestParam
	cacheJson.ConstantParams = constantResultParam
	cacheJson.IsRequestBody = isRequestBody
	cacheJson.GatewayHashKey = gatewayHashKey
	cacheJson.GatewayProtocol = gatewayProtocol
	cacheJson.GatewayRequestBodyNote = gatewayRequestBodyNote
	cacheJson.GatewayRequestPath = gatewayRequestPath
	cacheJson.GatewayRequestType = gatewayRequestType
	apiCacheJson,_ := json.Marshal(cacheJson)
	apiString := string(apiCacheJson[:])

	var redisCacheJson utils.RedisCacheJson
	redisCacheJson.BackendPath = backendRequestPath
	redisCacheJson.BackendProtocol = backendProtocol
	redisCacheJson.BackendRequestType = backendRequestType
	redisCacheJson.BackendURI = backendRequestURI
	redisCacheJson.ConstantParams = constantResultParam
	redisCacheJson.GatewayHashKey = gatewayHashKey
	redisCacheJson.GatewayID = gatewayID
	redisCacheJson.GatewayRequestPath = gatewayRequestPath
	redisCacheJson.IsRequestBody = isRequestBody
	redisCacheJson.RequestParams = gatewayRequestParam
	redisJson,_ := json.Marshal(redisCacheJson)
	redisString := string(redisJson[:])
	return dao.EditApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,apiString,redisString,apiID,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody,gatewayRequestParam,constantResultParam)
}

// 彻底删除Api
func DeleteApi(apiID,gatewayID int,gatewayHashKey string) bool{
	return dao.DeleteApi(apiID,gatewayID,gatewayHashKey)
}

// 获取api列表并按照名称排序
func GetApiListOrderByName(groupID int) (bool,[]*utils.ApiInfo){
	return dao.GetApiListOrderByName(groupID)
}

func GetApi(apiID int) (bool,utils.ApiInfo){
	return dao.GetApi(apiID)
}

// 获取所有API列表并依据接口名称排序
func GetAllApiListOrderByName(gatewayID int) (bool,[]*utils.ApiInfo){
	return dao.GetAllApiListOrderByName(gatewayID)
}

//搜索api
func SearchApi(tips string,gatewayID int) (bool,[]*utils.ApiInfo){
	return dao.SearchApi(tips,gatewayID)
}

// 获取网关接口信息，方便更新网关服务器的Redis数据
func GetRedisApiList(gatewayID int) (bool,[]*utils.ApiInfo){
	return dao.GetRedisApiList(gatewayID)
}

// 获取接口的核心信息
func GetRedisApi(apiID int) (bool,utils.ApiInfo){
	return dao.GetRedisApi(apiID)
}

// 查重
func CheckGatewayURLIsExist(gatewayID int,gatewayURI string) bool{
	flag,_ := dao.CheckGatewayURLIsExist(gatewayID,gatewayURI)
	return flag
}	
