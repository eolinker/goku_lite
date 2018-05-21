package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)

// 编辑鉴权信息
func EditAuth(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			auth := httpRequest.PostFormValue("auth")
			basicUserName := httpRequest.PostFormValue("basicUserName")
			basicUserPassword := httpRequest.PostFormValue("basicUserPassword")
			apiKey := httpRequest.PostFormValue("apiKey")

			flag := module.EditAuth(gatewayAlias,strategyID,auth,basicUserName,basicUserPassword,apiKey)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "180000"
			}		
			resultInfo.ResultType = "auth"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取鉴权信息
func GetAuthInfo(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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

			result := module.GetAuthInfo(gatewayAlias,strategyID)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "authInfo"
			resultInfo.Result = result
			resultInfo.ResultType = "auth"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

