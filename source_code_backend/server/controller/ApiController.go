package controller

import (
	"goku-ce/server/conf"
	"strconv"
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
	"encoding/json"
)

// 新增接口
func AddApi(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			apiName := httpRequest.PostFormValue("apiName")
			requestURL := httpRequest.PostFormValue("requestURL")
			requestMethod := httpRequest.PostFormValue("requestMethod")
			proxyURL := httpRequest.PostFormValue("proxyURL")
			proxyMethod := httpRequest.PostFormValue("proxyMethod")

			groupID := httpRequest.PostFormValue("groupID")
			backendID := httpRequest.PostFormValue("backendID")
			gID,groupIDErr := strconv.Atoi(groupID)
			bID := -1
			var backendIDErr error
			if backendID != "" {
				bID,backendIDErr = strconv.Atoi(backendID)
			}


			isRaw := httpRequest.PostFormValue("isRaw")
			ir := false
			if isRaw == "true" {
				ir = true
			}
			follow := httpRequest.PostFormValue("follow")
			fl := false
			if follow == "true" {
				fl = true
			}

			param := httpRequest.PostFormValue("proxyParams")
			constantParam := httpRequest.PostFormValue("constantParams")
			var params []*conf.Param
			var constantParams []*conf.ConstantParam
			if param != "" {
				err := json.Unmarshal([]byte(param),&params)
				if err != nil {
					panic(err)
				}
			}
			
			if constantParam != ""{
				err := json.Unmarshal([]byte(constantParam),&constantParams)
				if err != nil {
					panic(err)
				}
			}

			if groupIDErr != nil{
				resultInfo.StatusCode = "190002"

			} else if backendIDErr != nil {
				resultInfo.StatusCode = "190003"

			} else {
				flag := module.CheckApiURLIsExist(gatewayAlias,requestURL,requestURL,follow,-1)
				if flag {
					resultInfo.StatusCode = "190005"
				} else {
					flag,id := module.AddApi(gatewayAlias,apiName,requestURL,requestMethod,proxyURL,proxyMethod,gID,bID,fl,ir,params,constantParams)
					if flag {
						resultInfo.StatusCode = "000000"
						resultInfo.Result = id
						resultInfo.ResultKey = "apiID"
					} else {
						resultInfo.StatusCode = "190000"
					}
				}
				resultInfo.ResultType = "api"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 修改接口
func EditApi(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			apiName := httpRequest.PostFormValue("apiName")
			requestURL := httpRequest.PostFormValue("requestURL")
			requestMethod := httpRequest.PostFormValue("requestMethod")
			proxyURL := httpRequest.PostFormValue("proxyURL")
			proxyMethod := httpRequest.PostFormValue("proxyMethod")

			apiID := httpRequest.PostFormValue("apiID")
			groupID := httpRequest.PostFormValue("groupID")
			backendID := httpRequest.PostFormValue("backendID")
			aID,apiIDErr := strconv.Atoi(apiID)
			gID,groupIDErr := strconv.Atoi(groupID)
			bID := -1
			var backendIDErr error
			if backendID != "" {
				bID,backendIDErr = strconv.Atoi(backendID)
			}
			

			isRaw := httpRequest.PostFormValue("isRaw")
			ir := false
			if isRaw == "true" {
				ir = true
			}

			follow := httpRequest.PostFormValue("follow")
			fl := false
			if follow == "true" {
				fl = true
			}


			param := httpRequest.PostFormValue("proxyParams")
			constantParam := httpRequest.PostFormValue("constantParams")
			var params []*conf.Param
			var constantParams []*conf.ConstantParam
			if param != "" {
				err := json.Unmarshal([]byte(param),&params)
				if err != nil {
					panic(err)
				}
			} 
			if constantParam != ""{
				err := json.Unmarshal([]byte(constantParam),&constantParams)
				if err != nil {
					panic(err)
				}
			} 
			if apiIDErr != nil {
				resultInfo.StatusCode = "190001"
			} else if groupIDErr != nil{
				resultInfo.StatusCode = "190002"

			} else if backendIDErr != nil {
				resultInfo.StatusCode = "190003"
			} else {
				flag := module.CheckApiURLIsExist(gatewayAlias,requestURL,requestURL,follow,aID)
				if flag {
					resultInfo.StatusCode = "190005"
				} else {
					flag:= module.EditApi(gatewayAlias,apiName,requestURL,requestMethod,proxyURL,proxyMethod,aID,gID,bID,fl,ir,params,constantParams)
					if flag {
						resultInfo.StatusCode = "000000"
					} else {
						resultInfo.StatusCode = "190000"
					}
				}
			}
			resultInfo.ResultType = "api"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

func DeleteApi(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			apiID := httpRequest.PostFormValue("apiID")
			aID,err := strconv.Atoi(apiID)
			if err != nil {
				resultInfo.StatusCode = "190001"
			} else {
					flag:= module.DeleteApi(gatewayAlias,aID)
				if flag {
					resultInfo.StatusCode = "000000"
				} else {
					resultInfo.StatusCode = "190000"
				}
			}
			resultInfo.ResultType = "api"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 获取接口详情
func GetApiInfo(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			apiID := httpRequest.PostFormValue("apiID")
			aID,err := strconv.Atoi(apiID)
			if err != nil {
				resultInfo.StatusCode = "190001"
			} else {
				flag,result := module.GetApiInfo(gatewayAlias,aID)
				if flag {
					resultInfo.StatusCode = "000000"
					resultInfo.Result = result
					resultInfo.ResultKey = "apiInfo"
				} else {
					resultInfo.StatusCode = "190000"
				}
			}
			resultInfo.ResultType = "api"
			
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 获取接口列表
func GetAllApiList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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

			result := module.GetAllApiList(gatewayAlias)
			resultInfo.StatusCode = "000000"
			resultInfo.Result = result
			resultInfo.ResultKey = "apis"
			resultInfo.ResultType = "api"
			resultByte,_  :=json.Marshal(map[string]interface{}{
				"type":resultInfo.ResultType,
				"statusCode":resultInfo.StatusCode,
				"apiList":result["apiList"],
				"gatewayInfo":result["gatewayInfo"],
			})
			httpResponse.Write(resultByte)
			return
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}
	

func GetApiListByGroup(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			groupID := httpRequest.PostFormValue("groupID")
			gID,_ := strconv.Atoi(groupID)

			result := module.GetApiListByGroup(gatewayAlias,gID)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "apis"
			resultInfo.Result = result
			resultInfo.ResultType = "api"

			resultByte,_  :=json.Marshal(map[string]interface{}{
				"type":resultInfo.ResultType,
				"statusCode":resultInfo.StatusCode,
				"apiList":result["apiList"],
				"gatewayInfo":result["gatewayInfo"],
			})
			httpResponse.Write(resultByte)
			return
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 请求路径及请求方式查重
func CheckApiURLIsExist(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			requestURL := httpRequest.PostFormValue("requestURL")
			requestMethod := httpRequest.PostFormValue("requestMethod")
			follow := httpRequest.PostFormValue("follow")
			apiID :=httpRequest.PostFormValue("apiID")
			id,err := strconv.Atoi(apiID)
			if err != nil {
				resultInfo.StatusCode = "190001"
			} else {
				flag := module.CheckApiURLIsExist(gatewayAlias,requestURL,requestMethod,follow,id)
				if flag {
					resultInfo.StatusCode = "000000"
				} else {
					resultInfo.StatusCode = "190000"
				}
			}
			
			
			resultInfo.ResultType = "api"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 搜索接口
func SearchApi(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			keyword := httpRequest.PostFormValue("keyword")
			result := module.SearchApi(gatewayAlias,keyword)
			resultInfo.StatusCode = "000000"
			resultInfo.Result = result
			resultInfo.ResultKey = "apiList"	
			resultInfo.ResultType = "api"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

