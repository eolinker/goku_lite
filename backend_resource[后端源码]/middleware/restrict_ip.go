package middleware

import (
	"goku-ce-1.0/dao"
	"goku-ce-1.0/utils"
	"fmt"
	"net/http"
	"strings"
	"github.com/farseer810/yawf"
	"time"
)

func IPValve(httpRequest *http.Request, context yawf.Context,
	httpResponse http.ResponseWriter, headers yawf.Headers) (bool, string) {
		t1 := time.Now()
	var info *utils.IPListInfo
	var gatewayHashkey string
	// 获取请求路径中的网关别名
	requestInfo := strings.Split(httpRequest.RequestURI,"/")
	gatewayAlias := requestInfo[1]
	strategyKey := requestInfo[2]
	// 通过网关别名获取网关hashKey
	gatewayHashkey = dao.GetGatewayHashKey(context,gatewayAlias)
	if gatewayHashkey == "" {
		httpResponse.WriteHeader(404)
		return false, "error gatewayAlias"
	}

	remoteAddr := httpRequest.RemoteAddr
	remoteIP := InterceptIP(remoteAddr, ":")

	info = dao.GetIPList(context, gatewayHashkey,strategyKey)
	if info == nil{
		return true,""
	}
	chooseType := info.ChooseType

	if chooseType == 1 {
		for _, ipList := range info.IPList {
			if ipList == remoteIP {
				fmt.Println(remoteIP)
				httpResponse.WriteHeader(403)
				return false, "illegal IP"
			}
		}
	} else if chooseType == 2 {
		for _, ipList := range info.IPList {
			if ipList == remoteIP {
				return true, ""
			}
		}
		fmt.Println(remoteIP)
		httpResponse.WriteHeader(403)
		return false, "illegal IP"
	}
	fmt.Println("restrictIP time:",time.Since(t1))
	return true, ""
}

func InterceptIP(str, substr string) string {
	result := strings.Index(str, substr)
	var rs string
	if result > 7 {
		rs = str[:result]
	}
	return rs
}
