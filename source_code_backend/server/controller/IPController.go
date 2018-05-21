package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)

// 修改网关黑白名单
func EditGatewayIPList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayAlias := httpRequest.PostFormValue("gatewayAlias")
			ipLimitType := httpRequest.PostFormValue("ipLimitType")
			ipWhiteList := httpRequest.PostFormValue("ipWhiteList")
			ipBlackList := httpRequest.PostFormValue("ipBlackList")

			flag := module.EditGatewayIPList(gatewayAlias,ipLimitType,ipWhiteList,ipBlackList)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "170000"
			}
			resultInfo.ResultType = "ip"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 修改策略组黑白名单
func EditStrategyIPList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayAlias := httpRequest.PostFormValue("gatewayAlias")
			strategyID := httpRequest.PostFormValue("strategyID")
			ipLimitType := httpRequest.PostFormValue("ipLimitType")
			ipWhiteList := httpRequest.PostFormValue("ipWhiteList")
			ipBlackList := httpRequest.PostFormValue("ipBlackList")

			flag := module.EditStrategyIPList(gatewayAlias,strategyID,ipLimitType,ipWhiteList,ipBlackList)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "170000"
			}
			resultInfo.ResultType = "ip"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取网关黑白名单
func GetGatewayIPList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayAlias := httpRequest.PostFormValue("gatewayAlias")

			result := module.GetGatewayIPList(gatewayAlias)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "ipList"
			resultInfo.Result = result			
			resultInfo.ResultType = "ip"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取策略组黑白名单
func GetStrategyIPList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayAlias := httpRequest.PostFormValue("gatewayAlias")
			strategyID := httpRequest.PostFormValue("strategyID")
			result := module.GetStrategyIPList(gatewayAlias,strategyID)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "ipList"
			resultInfo.Result = result			
			resultInfo.ResultType = "ip"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}