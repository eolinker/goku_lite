package dao

import (
	"strings"
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
	"time"
)

// 新增接口
func AddApi(apiConfPath,apiName,requestURL,requestMethod,proxyURL,proxyMethod string,groupID,backendID int,follow,isRaw bool,param []*conf.Param,constantParam []*conf.ConstantParam) (bool,int) {
	apiList,_,_ := conf.ParseApiInfo(apiConfPath)
	now := time.Now().Format("2006-01-02 15:04:05")
	api := &conf.ApiInfo{
		ApiName : apiName,
		GroupID : groupID,
		RequestURL : requestURL,
		RequestMethod : strings.Split(requestMethod,","),
		BackendID : backendID,
		ProxyURL : proxyURL,
		ProxyMethod : proxyMethod,
		IsRaw : isRaw,
		ProxyParams : param,
		ConstantParams : constantParam,
		UpdateTime:now,
		CreateTime:now,
		Follow:follow,
	}
	apiList = append(apiList,api)
	content,err := yaml.Marshal(conf.Api{
		ApiList : apiList,
	})
	if err != nil {
		panic(err)
	}
	conf.WriteConfigToFile(apiConfPath,content)
	return true,len(apiList)
}

// 修改接口
func EditApi(apiConfPath,apiName,requestURL,requestMethod,proxyURL,proxyMethod string,apiID,groupID,backendID int,follow,isRaw bool,param []*conf.Param,constantParam []*conf.ConstantParam) (bool) {
	apis,api,_ := conf.ParseApiInfo(apiConfPath)
	_,ok := api[apiID] 
	if !ok {
		return false
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	api[apiID].ApiName = apiName
	api[apiID].GroupID = groupID
	api[apiID].RequestURL = requestURL
	api[apiID].RequestMethod = strings.Split(requestMethod,",")
	api[apiID].BackendID = backendID
	api[apiID].ProxyURL = proxyURL
	api[apiID].ProxyMethod = proxyMethod
	api[apiID].IsRaw = isRaw
	api[apiID].ProxyParams = param
	api[apiID].ConstantParams = constantParam
	api[apiID].UpdateTime = now
	api[apiID].Follow = follow

	apiList := make([]*conf.ApiInfo,0)
	for i:=0 ; i<len(apis);i++ {
		apiList = append(apiList,api[i+1])
	}
	content,err := yaml.Marshal(conf.Api{
		ApiList : apiList,
	})
	if err != nil {
		panic(err)
	}
	conf.WriteConfigToFile(apiConfPath,content)
	return true
}

func DeleteApi(apiConfPath string,apiID int) (bool) {
	apis,api,_ := conf.ParseApiInfo(apiConfPath)
	_,ok := api[apiID] 
	if !ok {
		return false
	}
	delete(api,apiID)
	apiList := make([]*conf.ApiInfo,0)
	for i:=0 ; i<len(apis);i++ {
		if i + 1 != apiID {
			apiList = append(apiList,api[i+1])
		}
	}
	content,err := yaml.Marshal(conf.Api{
		ApiList : apiList,
	})
	if err != nil {
		panic(err)
	}
	conf.WriteConfigToFile(apiConfPath,content)
	return true
}

// 获取接口详情
func GetApiInfo(apiConfPath string,apiID int) (bool,map[string]interface{}) {
	_,api,_ := conf.ParseApiInfo(apiConfPath)
	
	value,ok := api[apiID] 
	if !ok {
		return false,nil
	}
	return true,map[string]interface{}{
		"apiID": apiID,
		"apiName": value.ApiName,
		"backendID": value.BackendID,
		"proxyParams" : value.ProxyParams,
		"constantParams": value.ConstantParams,
		"requestURL": value.RequestURL,
		"requestMethod" : strings.Join(value.RequestMethod,","),
		"proxyURL" : value.ProxyURL,
		"proxyMethod": value.ProxyMethod,
		"isRaw":value.IsRaw,
		"groupID":value.GroupID,
		"follow":value.Follow,
		"gatewayInfo": map[string]string{
			"gatewayHost" : conf.GlobalConf.Host,
			"gatewayPort" : conf.GlobalConf.Port,
		},
	}
}

// 获取接口列表
func GetAllApiList(apiConfPath string) map[string]interface{}{
	_,_,apiList := conf.ParseApiInfo(apiConfPath)
	apis := make([]map[string]interface{},0)
	for _,api := range apiList {
		apis = append(apis,api)
	}
	return map[string]interface{}{
		"apiList":apis,
		"gatewayInfo" : map[string]string{
			"gatewayHost" : conf.GlobalConf.Host,
			"gatewayPort" : conf.GlobalConf.Port,
		},
	}
}

func GetApiListByGroup(apiConfPath string,groupID int) map[string]interface{}{
	_,_,apiList := conf.ParseApiInfo(apiConfPath)
	apis := make([]map[string]interface{},0)
	for _,api := range apiList {
		if api["groupID"] == groupID {
			apis = append(apis,api)
		}
	}
	return map[string]interface{}{
		"apiList":apis,
		"gatewayInfo" : map[string]string{
			"gatewayHost" : conf.GlobalConf.Host,
			"gatewayPort" : conf.GlobalConf.Port,
		},
	}
}

// 获取接口数量
func GetApiCount(apiConfPath string) int {
	apiList,_,_ := conf.ParseApiInfo(apiConfPath)
	return len(apiList)
}

// 批量修改接口分组
func DeleteApiOfGroup(apiConfPath string,groupID int) bool {
	_,apis,_ := conf.ParseApiInfo(apiConfPath)
	apiList := make([]*conf.ApiInfo,0)
	for _,value := range apis {
		if value.GroupID != groupID {
			apiList = append(apiList,value)
		}
	}
	content,err := yaml.Marshal(conf.Api{
		ApiList : apiList,
	})
	if err != nil {
		panic(err)
	}
	conf.WriteConfigToFile(apiConfPath,content)
	return true
}

// 请求路径及请求方式查重
func CheckApiURLIsExist(apiConfPath,requestURL,requestMethod,follow string,apiID int) bool{
	_,_,apis := conf.ParseApiInfo(apiConfPath)
	method := strings.Split(requestMethod,",")
	for _,value := range apis {
		if value["requestURL"]== requestURL {
			if apiID == value["apiID"]{
				return false
			}
			if value["follow"] == false {
				apiRequestMethod := strings.Split(value["requestURL"].(string),",")
				for _,m := range method {
					for _,am := range apiRequestMethod {
						if am == m {
							return true
						} else {
							return false
						}
					}
				}
			} else if follow == "true" {
				return false
			} else if value["follow"] == true {
				return false
			}
		}
	}
	return false
}

// 搜索接口
func SearchApi(apiConfPath,keyword string) []map[string]interface{}{
	_,_,apiList := conf.ParseApiInfo(apiConfPath)
	apis := make([]map[string]interface{},0)
	for _,api := range apiList {
		if strings.Contains(api["apiName"].(string),keyword) || strings.Contains(api["requestURL"].(string),keyword) {
			apis = append(apis,api)
		}
	}
	return apis
}