package controller

import (
	"strconv"
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)

// 新增流量限制
func AddRateLimit(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			period := httpRequest.PostFormValue("period")

			startTime := httpRequest.PostFormValue("startTime")
			endTime := httpRequest.PostFormValue("endTime")
			priority := httpRequest.PostFormValue("priority")
			limit := httpRequest.PostFormValue("limit")
			st,_ := strconv.Atoi(startTime)
			et,_ := strconv.Atoi(endTime)
			pr,_ := strconv.Atoi(priority)
			count,_ := strconv.Atoi(limit)


			allow := httpRequest.PostFormValue("allow")
			isAllow := false
			if allow == "true" {
				isAllow = true
			}

			flag := module.AddRateLimit(gatewayAlias,strategyID,period,st,et,pr,count,isAllow)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "200000"
			}
			resultInfo.ResultType = "rateLimit"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 修改流量限制
func EditRateLimit(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			period := httpRequest.PostFormValue("period")

			startTime := httpRequest.PostFormValue("startTime")
			endTime := httpRequest.PostFormValue("endTime")
			priority := httpRequest.PostFormValue("priority")
			limitID := httpRequest.PostFormValue("limitID")
			limit := httpRequest.PostFormValue("limit")
			st,_ := strconv.Atoi(startTime)
			et,_ := strconv.Atoi(endTime)
			pr,_ := strconv.Atoi(priority)
			count,_ := strconv.Atoi(limit)
			id,_ := strconv.Atoi(limitID)


			allow := httpRequest.PostFormValue("allow")
			isAllow := false
			if allow == "true" {
				isAllow = true
			}

			flag := module.EditRateLimit(gatewayAlias,strategyID,period,id,st,et,pr,count,isAllow)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "200000"
			}
			resultInfo.ResultType = "rateLimit"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 删除流量限制
func DeleteRateLimit(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			limitID := httpRequest.PostFormValue("limitID")
			id,_ := strconv.Atoi(limitID)

			flag := module.DeleteRateLimit(gatewayAlias,strategyID,id)
			if flag {
				resultInfo.StatusCode = "000000"
			} else {
				resultInfo.StatusCode = "200000"
			}
			resultInfo.ResultType = "rateLimit"
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取流量限制列表
func GetRateLimitInfo(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			limitID := httpRequest.PostFormValue("limitID")
			lID,_ := strconv.Atoi(limitID)

			flag,result := module.GetRateLimitInfo(gatewayAlias,strategyID,lID)
			if flag{
				resultInfo.StatusCode = "000000"
				resultInfo.ResultKey = "limitList"
				resultInfo.Result = result
			} else {
				resultInfo.StatusCode = "200000"
			}
			
			resultInfo.ResultType = "rateLimit"
			
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 获取流量限制列表
func GetRateLimitList(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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

			result := module.GetRateLimitList(gatewayAlias,strategyID)
			resultInfo.StatusCode = "000000"
			resultInfo.ResultKey = "limitList"
			resultInfo.ResultType = "rateLimit"
			resultInfo.Result = result
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}