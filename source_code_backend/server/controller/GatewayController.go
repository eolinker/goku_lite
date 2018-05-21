package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
	"goku-ce/server/conf"
)

// 新增网关
func AddGateway(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayName := httpRequest.PostFormValue("gatewayName")
			flag := module.AddGateway(gatewayName,gatewayAlias)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 修改网关信息
func EditGateway(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gatewayName := httpRequest.PostFormValue("gatewayName")
			oldGatewayAlias := httpRequest.PostFormValue("oldGatewayAlias")
			flag := module.EditGateway(gatewayName,gatewayAlias,oldGatewayAlias)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 删除网关
func DeleteGateway(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag := module.DeleteGateway(gatewayAlias)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取网关列表
func GetGatewayList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag,gatewayList := module.GetGatewayList()
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
				resultInfo.ResultKey = "gatewayList"
				resultInfo.Result = gatewayList
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取网关信息
func GetGatewayInfo(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag,gatewayInfo := module.GetGatewayInfo(gatewayAlias)
			if flag {
				url := "http://localhost:" + conf.GlobalConf.Port + "/goku/Count/getVisitCount"
				gatewayInfo["gatewayVisitCount"] = utils.GetVisitCount(url)
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
				resultInfo.ResultKey = "gatewayInfo"
				resultInfo.Result = gatewayInfo
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 检查网关是否存在
func CheckGatewayAliasIsExist(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag := module.CheckGatewayAliasIsExist(gatewayAlias)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway"
			} else {
				resultInfo.StatusCode = "130000"
				resultInfo.ResultType = "gateway"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

