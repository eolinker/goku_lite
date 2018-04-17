package middleware

import (
	"goku-ce-1.0/dao"
	"net/http"
	"strings"
	"time"
	"fmt"
	"github.com/farseer810/yawf"
	"encoding/base64"
)

func GatewayAuth(httpRequest *http.Request,context yawf.Context, httpResponse http.ResponseWriter,headers yawf.Headers) (bool, string) {
	var gatewayHashkey string
	t1 := time.Now()
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

	authInfo := dao.GetAuthInfo(context,gatewayHashkey,strategyKey)
	if authInfo.AuthType == 0{
		authStr := []byte(authInfo.UserName + ":" + authInfo.UserPassword)
		authorization := "Basic " + base64.StdEncoding.EncodeToString(authStr)
		fmt.Println(authorization)
		fmt.Println(headers["Authorization"])
		// basic鉴权
		if authorization != headers["Authorization"] {
			httpResponse.WriteHeader(406)
			return false, "鉴权失败"
		}
	}else if authInfo.AuthType == 1{
		// apiKey鉴权
		if authInfo.ApiKey != headers["Apikey"] {
			httpResponse.WriteHeader(401)
			return false, "apiKey有误"
		}
	}
	fmt.Println("auth time:",time.Since(t1))
	return true,""
}


