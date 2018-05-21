package controller

import (
	"goku-ce/server/conf"
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)
var oldPort = conf.GlobalConf.Port
var isStart = false
// 修改全局配置
func EditGlobalConfig(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	nameCookie,nameErr := httpRequest.Cookie("loginName")
	userCookie,userErr := httpRequest.Cookie("userToken")
	if nameErr != nil || userErr != nil{
		resultInfo.StatusCode = "100001"
		resultInfo.ResultType = "user"
	} else {
		flag := module.CheckLogin(userCookie.Value,nameCookie.Value)
		if !flag {
			resultInfo.StatusCode = "100001"
			resultInfo.ResultType = "user"
		} else {
			if isStart {
				oldPort = conf.GlobalConf.Port
				isStart = false
			}
			gatewayPort := httpRequest.PostFormValue("gatewayPort")
			flag := module.EditGlobalConfig(gatewayPort)
			if flag{
				resultInfo.StatusCode = "000000"
			}else {
				resultInfo.StatusCode = "220000"
			}
			resultInfo.ResultType = "global_conf"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 获取全局配置信息
func GetGlobalConfig(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	nameCookie,nameErr := httpRequest.Cookie("loginName")
	userCookie,userErr := httpRequest.Cookie("userToken")
	if nameErr != nil || userErr != nil{
		resultInfo.StatusCode = "100001"
		resultInfo.ResultType = "user"
	} else {
		flag := module.CheckLogin(userCookie.Value,nameCookie.Value)
		if !flag {
			resultInfo.StatusCode = "100001"
			resultInfo.ResultType = "user"
		} else {
			gatewayServiceStatus := utils.GetGatewayServiceStatus(oldPort)
			
			resultInfo.Result = map[string]interface{}{
				"gatewayPort": conf.GlobalConf.Port,
				"gatewayServiceStatus":gatewayServiceStatus,
			}
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "confInfo"
			resultInfo.ResultType = "global_conf"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}