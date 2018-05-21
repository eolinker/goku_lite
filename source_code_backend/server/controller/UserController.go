package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)


// 检查用户登录
func CheckLogin(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			resultInfo.StatusCode = "000000"
			resultInfo.ResultType = "user"
		}	
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
} 

// 用户注销
func Logout(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	nameCookie := http.Cookie{Name: "loginName", Path: "/", MaxAge: -1}
	userCookie := http.Cookie{Name: "userToken", Path: "/", MaxAge: -1}
	http.SetCookie(httpResponse, &nameCookie)
	http.SetCookie(httpResponse, &userCookie)
	resultInfo.StatusCode = "000000"
	resultInfo.ResultType = "user"
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}
