package controller

import (
	"strconv"
	"goku-ce/server/module"
	"goku-ce/utils"
	"net/http"
)

// 新增后端
func AddBackend(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			backendName := httpRequest.PostFormValue("backendName")
			backendPath := httpRequest.PostFormValue("backendPath")
			flag,id := module.AddBackend(gatewayAlias,backendName,backendPath)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultKey = "backendID"
				resultInfo.Result = id
			}else {
				resultInfo.StatusCode = "140000"
			}

			resultInfo.ResultType = "backend"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 修改后端信息
func EditBackend(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			backendName := httpRequest.PostFormValue("backendName")
			backendPath := httpRequest.PostFormValue("backendPath")
			backendID := httpRequest.PostFormValue("backendID")
			bID,err := strconv.Atoi(backendID)
			if err != nil {
				resultInfo.StatusCode = "140001"
			} else {
				flag := module.EditBackend(gatewayAlias,backendName,backendPath,bID)
				if flag {
					resultInfo.StatusCode = "000000"
				}else {
					resultInfo.StatusCode = "140000"
				}
			}
			
			resultInfo.ResultType = "backend"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 删除后端信息
func DeleteBackend(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			backendID := httpRequest.PostFormValue("backendID")
			bID,err := strconv.Atoi(backendID)
			if err != nil {
				resultInfo.StatusCode = "140001"
			} else {
				flag := module.DeleteBackend(gatewayAlias,bID)
				if flag {
					resultInfo.StatusCode = "000000"
				}else {
					resultInfo.StatusCode = "140000"
				}
			}
			resultInfo.ResultType = "backend"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取后端列表
func GetBackendList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			backednList := module.GetBackendList(gatewayAlias)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultType = "backend"
			resultInfo.ResultKey = "backendList"
			resultInfo.Result = backednList
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取后端信息
func GetBackendInfo(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			backendID := httpRequest.PostFormValue("backendID")
			bID,err := strconv.Atoi(backendID)
			if err != nil {

			}
			flag,backednInfo := module.GetBackendInfo(gatewayAlias,bID)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.Result = backednInfo
				resultInfo.ResultKey = "backendInfo"
			} else {
				resultInfo.StatusCode = "140000"
			}
			resultInfo.ResultType = "backend"
			
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}