package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/utils"
)

// 用户登录
func Login(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	loginName := httpRequest.PostFormValue("loginName")
	loginPassword := httpRequest.PostFormValue("loginPassword")
	loginPassword = utils.Md5(loginPassword)
	flag := module.Login(loginName,loginPassword)
	if flag {
		userCookie := &http.Cookie{Name: "userToken", Value: utils.Md5(loginName+loginPassword), Path: "/", MaxAge: 86400}
		nameCookie := &http.Cookie{Name: "loginName", Value: loginName, Path: "/", MaxAge: 86400}
		http.SetCookie(httpResponse,userCookie)
		http.SetCookie(httpResponse,nameCookie)
		resultInfo.StatusCode = "000000"
		resultInfo.ResultType = "guest"
	} else {
		resultInfo.StatusCode = "100000"
		resultInfo.ResultType = "guest"
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}
