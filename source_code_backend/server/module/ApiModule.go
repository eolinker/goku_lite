package module

import (
	"goku-ce/server/dao"
	"goku-ce/server/conf"
)

// 新增接口
func AddApi(gatewayAlias,apiName,requestURL,requestMethod,proxyURL,proxyMethod string,groupID,backendID int,follow,isRaw bool,param []*conf.Param,constantParam []*conf.ConstantParam) (bool,int) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.AddApi(gateway["apiConfPath"],apiName,requestURL,requestMethod,proxyURL,proxyMethod,groupID,backendID,follow,isRaw,param,constantParam)
	}else {
		return false,0
	}
}

// 修改接口
func EditApi(gatewayAlias,apiName,requestURL,requestMethod,proxyURL,proxyMethod string,apiID,groupID,backendID int,follow,isRaw bool,param []*conf.Param,constantParam []*conf.ConstantParam) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditApi(gateway["apiConfPath"],apiName,requestURL,requestMethod,proxyURL,proxyMethod,apiID,groupID,backendID,follow,isRaw,param,constantParam)
	}else {
		return false
	}
}

func DeleteApi(gatewayAlias string,apiID int) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.DeleteApi(gateway["apiConfPath"],apiID)
	}else {
		return false
	}
}

// 获取接口详情
func GetApiInfo(gatewayAlias string,apiID int) (bool,map[string]interface{}) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		flag,result := dao.GetApiInfo(gateway["apiConfPath"],apiID)
		if !flag{
			return false,make(map[string]interface{})
		} 
		flag,groupInfo := dao.GetApiGroupInfo(gateway["apiGroupConfPath"],result["groupID"].(int))
		if !flag {
			result["groupName"] = "默认分组"
		} else {
			result["groupName"] = groupInfo.GroupName
		}
		
		flag,backendInfo := dao.GetBackendInfo(gateway["backendConfPath"],result["backendID"].(int))
		if !flag {
			result["backendName"] = ""
			result["backendPath"] = ""
		}else {
			result["backendName"] = backendInfo.BackendName
			result["backendPath"] = backendInfo.BackendPath
		}
		return true,result
	}else {
		return false,make(map[string]interface{})
	}
}

// 获取接口列表
func GetAllApiList(gatewayAlias string) map[string]interface{}{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetAllApiList(gateway["apiConfPath"])
	}else {
		return make(map[string]interface{})
	}
}

func GetApiListByGroup(gatewayAlias string,groupID int) map[string]interface{}{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetApiListByGroup(gateway["apiConfPath"],groupID)
	}else {
		return make(map[string]interface{})
	}
}

// 请求路径及请求方式查重
func CheckApiURLIsExist(gatewayAlias,requestURL,requestMethod,follow string,apiID int) bool{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.CheckApiURLIsExist(gateway["apiConfPath"],requestURL,requestMethod,follow,apiID)
	}else {
		return false
	}
}

// 搜索接口
func SearchApi(gatewayAlias,keyword string) []map[string]interface{}{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.SearchApi(gateway["apiConfPath"],keyword)
	}else {
		return make([]map[string]interface{},0)
	}
}
