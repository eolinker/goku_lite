package controller

import (
	"strconv"
	"goku-ce/server/module"
	"goku-ce/utils"
	"net/http"
)

// 新增分组
func AddApiGroup(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			groupName := httpRequest.PostFormValue("groupName")
			flag,id := module.AddApiGroup(gatewayAlias,groupName)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.Result = id
				resultInfo.ResultKey = "groupID"
			}else {
				resultInfo.StatusCode = "150000"
			}
			resultInfo.ResultType = "group"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 修改分组
func EditApiGroup(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			groupName := httpRequest.PostFormValue("groupName")
			groupID := httpRequest.PostFormValue("groupID")
			gID,err := strconv.Atoi(groupID)
			if err != nil {
				resultInfo.StatusCode = "150001"
			} else {
				flag := module.EditApiGroup(gatewayAlias,groupName,gID)
				if flag {
					resultInfo.StatusCode = "000000"
				}else {
					resultInfo.StatusCode = "150000"
				}
			}
			resultInfo.ResultType = "group"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}

// 删除分组
func DeleteApiGroup(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			gID,err := strconv.Atoi(groupID)
			if err != nil {
				resultInfo.StatusCode = "150001"
			} else {
				flag := module.DeleteApiGroup(gatewayAlias,gID)
				if flag {
					resultInfo.StatusCode = "000000"
				}else {
					resultInfo.StatusCode = "150000"
				}
			}
			resultInfo.ResultType = "group"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}


// 获取分组列表
func GetApiGroupList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			groupList := module.GetApiGroupList(gatewayAlias)
			resultInfo.StatusCode = "000000"
			resultInfo.Result = groupList
			resultInfo.ResultKey = "groupList"
			resultInfo.ResultType = "group"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1"))))
	return
}
