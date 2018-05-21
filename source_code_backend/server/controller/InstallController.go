package controller

import (
	"net/http"
	"goku-ce/server/module"
	"goku-ce/server/conf"
	"goku-ce/utils"
)

// 检查是否已安装
func CheckIsInstall(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	flag:= utils.CheckFileIsExist(utils.ConfFilepath)
	if flag {
		resultInfo.StatusCode = "000000"
	}else {
		resultInfo.StatusCode = "200000"
	}
	resultInfo.ResultType = "install"
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 安装
func Install(httpResponse http.ResponseWriter,httpRequest *http.Request) {
	resultInfo := utils.ResultInfo{}
	port := httpRequest.PostFormValue("port")
	loginName := httpRequest.PostFormValue("userName")
	loginPassword := httpRequest.PostFormValue("userPassword")
	gatewayConfPath := httpRequest.PostFormValue("gatewayConfPath")
	flag := module.Install(port,loginName,utils.Md5(loginPassword),gatewayConfPath)
	if flag{
		resultInfo.StatusCode = "000000"
		utils.ReloadConf()
		conf.ParseAdminConfig()
	} else {
		resultInfo.StatusCode = "200000"
	}
			
	resultInfo.ResultType = "install"
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}