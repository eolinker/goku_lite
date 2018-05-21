package controller

import(
	"goku-ce/server/conf"
	"goku-ce/utils"
	"goku-ce/server/module"
	"net/http"
	"fmt"
)

// 重启网关后端服务
func RestartGatewayService(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			fmt.Println(oldPort)
			flag = utils.RestartGatewayService(oldPort)
			if flag {
				oldPort = conf.GlobalConf.Port
				isStart = true
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway_service"
			} else {
				resultInfo.StatusCode = "230000"
				resultInfo.ResultType = "gateway_service"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 关闭后端服务
func StopGatewayService(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag = utils.StopGatewayService(oldPort,true)
			if flag {
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway_service"
			} else {
				resultInfo.StatusCode = "230000"
				resultInfo.ResultType = "gateway_service"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}

// 关闭后端服务
func StartGatewayService(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag = utils.StartGatewayService(conf.GlobalConf.Port)
			if flag {
				isStart = true
				oldPort = conf.GlobalConf.Port
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway_service"
			} else {
				resultInfo.StatusCode = "230000"
				resultInfo.ResultType = "gateway_service"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}


// 重载网关后端服务
func ReloadGatewayService(httpResponse http.ResponseWriter,httpRequest *http.Request) {
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
			flag = utils.StopGatewayService(oldPort,false)
			if flag {
				oldPort = conf.GlobalConf.Port
				isStart = true
				resultInfo.StatusCode = "000000"
				resultInfo.ResultType = "gateway_service"
			} else {
				resultInfo.StatusCode = "230000"
				resultInfo.ResultType = "gateway_service"
			}
		}
	}
	httpResponse.Write([]byte(utils.String(utils.GetResultInfo(resultInfo.StatusCode, resultInfo.ResultType,resultInfo.ResultKey,resultInfo.Result,"1")))) 
	return
}



