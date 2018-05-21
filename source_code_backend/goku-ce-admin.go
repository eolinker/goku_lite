package main

import (
	"fmt"
	_ "goku-ce/server/conf"
	"net/http"
	"goku-ce/server/controller"
)

func main() {
	// 游客
	http.HandleFunc("/Web/Guest/login",controller.Login)
	
	// 用户
	http.HandleFunc("/Web/User/checkLogin",controller.CheckLogin)
	http.HandleFunc("/Web/User/logout",controller.Logout)

	// 网关
	http.HandleFunc("/Web/Gateway/addGateway",controller.AddGateway)
	http.HandleFunc("/Web/Gateway/editGateway",controller.EditGateway)
	http.HandleFunc("/Web/Gateway/deleteGateway",controller.DeleteGateway)
	http.HandleFunc("/Web/Gateway/getGatewayList",controller.GetGatewayList)
	http.HandleFunc("/Web/Gateway/getGateway",controller.GetGatewayInfo)
	http.HandleFunc("/Web/Gateway/checkGatewayAliasIsExist",controller.CheckGatewayAliasIsExist)

	//网关服务
	http.HandleFunc("/Web/GatewayService/restart",controller.RestartGatewayService)
	http.HandleFunc("/Web/GatewayService/reload",controller.ReloadGatewayService)
	http.HandleFunc("/Web/GatewayService/stop",controller.StopGatewayService)
	http.HandleFunc("/Web/GatewayService/start",controller.StartGatewayService)

	// 后端服务
	http.HandleFunc("/Web/Backend/addBackend",controller.AddBackend)
	http.HandleFunc("/Web/Backend/editBackend",controller.EditBackend)
	http.HandleFunc("/Web/Backend/deleteBackend",controller.DeleteBackend)
	http.HandleFunc("/Web/Backend/getBackendList",controller.GetBackendList)
	http.HandleFunc("/Web/Backend/getBackend",controller.GetBackendInfo)

	// API分组
	http.HandleFunc("/Web/ApiGroup/addGroup",controller.AddApiGroup)
	http.HandleFunc("/Web/ApiGroup/editGroup",controller.EditApiGroup)
	http.HandleFunc("/Web/ApiGroup/deleteGroup",controller.DeleteApiGroup)
	http.HandleFunc("/Web/ApiGroup/getGroupList",controller.GetApiGroupList)

	// API
	http.HandleFunc("/Web/Api/addApi",controller.AddApi)
	http.HandleFunc("/Web/Api/editApi",controller.EditApi)
	http.HandleFunc("/Web/Api/deleteApi",controller.DeleteApi)
	http.HandleFunc("/Web/Api/getApi",controller.GetApiInfo)
	http.HandleFunc("/Web/Api/searchApi",controller.SearchApi)
	http.HandleFunc("/Web/Api/getAllApiList",controller.GetAllApiList)
	http.HandleFunc("/Web/Api/getApiList",controller.GetApiListByGroup)
	http.HandleFunc("/Web/Api/checkApiURLIsExist",controller.CheckApiURLIsExist)

	// 策略组
	http.HandleFunc("/Web/Strategy/addStrategy",controller.AddStrategy)
	http.HandleFunc("/Web/Strategy/editStrategy",controller.EditStrategy)
	http.HandleFunc("/Web/Strategy/deleteStrategy",controller.DeleteStrategy)
	http.HandleFunc("/Web/Strategy/getStrategyList",controller.GetStrategyList)
	http.HandleFunc("/Web/Strategy/getSimpleStrategyList",controller.GetSimpleStrategyList)

	// 流控
	http.HandleFunc("/Web/RateLimit/addRateLimit",controller.AddRateLimit)
	http.HandleFunc("/Web/RateLimit/editRateLimit",controller.EditRateLimit)
	http.HandleFunc("/Web/RateLimit/deleteRateLimit",controller.DeleteRateLimit)
	http.HandleFunc("/Web/RateLimit/getRateLimitInfo",controller.GetRateLimitInfo)
	http.HandleFunc("/Web/RateLimit/getRateLimitList",controller.GetRateLimitList)

	// 鉴权
	http.HandleFunc("/Web/Auth/editAuth",controller.EditAuth)
	http.HandleFunc("/Web/Auth/getAuthInfo",controller.GetAuthInfo)

	// 黑白名单
	http.HandleFunc("/Web/IP/editGatewayIPList",controller.EditGatewayIPList)
	http.HandleFunc("/Web/IP/editStrategyIPList",controller.EditStrategyIPList)
	http.HandleFunc("/Web/IP/getGatewayIPList",controller.GetGatewayIPList)
	http.HandleFunc("/Web/IP/getStrategyIPList",controller.GetStrategyIPList)

	// 安装
	http.HandleFunc("/Web/Install/install",controller.Install)
	http.HandleFunc("/Web/Install/checkIsInstall",controller.CheckIsInstall)

	// 全局配置
	http.HandleFunc("/Web/Global/getConfInfo",controller.GetGlobalConfig)
	http.HandleFunc("/Web/Global/editConfInfo",controller.EditGlobalConfig)

	fmt.Println("Listen: 9900")
	err := http.ListenAndServe(":9900", nil)  
	
	if err != nil {  
		panic(err)  
	}  
}