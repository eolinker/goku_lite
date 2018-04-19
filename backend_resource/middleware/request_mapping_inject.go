package middleware

import (
	"goku-ce-1.0/dao"
	"net/http"
	"fmt"
	"strings"
	"github.com/farseer810/yawf"
)

var (
	methodIndicator = map[string]string{"POST": "0", "GET": "1", "PUT": "2", "DELETE": "3", "HEAD": "4",
		"OPTIONS": "5", "PATCH": "6"}
)

func isURIMatched(context yawf.Context, incomingURI, testURI string) bool {
	isMatched := incomingURI == testURI
	return isMatched
}

//注入请求映射
func InjectRequestMapping(httpRequest *http.Request, context yawf.Context,
	httpResponse http.ResponseWriter, headers yawf.Headers) (bool, string) {
	var domain, method, scheme, gatewayHashkey, requestURL string
	// TODO: 0 for http, 1 for https
	scheme = "0"
	fmt.Println(httpRequest.RemoteAddr)
	method = methodIndicator[httpRequest.Method]
	if method == "" {
		httpResponse.WriteHeader(404)
		return false, "empty method"
	}

	domain = httpRequest.Host
	requestURL = httpRequest.RequestURI
	fmt.Println(domain)
	fmt.Println(requestURL)
	if len(requestURL) <= 1 {
		httpResponse.WriteHeader(404)
		return false, "lack url"
	}
	// 获取请求路径中的网关别名
	requestInfo := strings.Split(requestURL,"/")
	if len(requestInfo) < 3{
		httpResponse.WriteHeader(404)
		return false, "lack strategyKey"
	}
	gatewayAlias := requestInfo[1]
	strategyKey := requestInfo[2]
	
	// 通过网关别名获取网关hashKey
	gatewayHashkey = dao.GetGatewayHashKey(context,gatewayAlias)
	if gatewayHashkey == "" {
		httpResponse.WriteHeader(404)
		return false, "error gatewayAlias"
	}
	fmt.Println(gatewayHashkey)
	paths := dao.GetAllAPIPaths(context, gatewayHashkey)
	fmt.Println(paths)
	var matchedURI string
	gatewayLen := len(requestInfo[1]) + len(requestInfo[2]) + 2
	flag := false
	for _, uri := range paths {
		if uri[0:4] != scheme+":"+method+":" {
			continue
		}
		if isURIMatched(context, requestURL[gatewayLen:], uri[4:]) {
			matchedURI = uri
			flag = true
		}
	}
	if !flag {
		httpResponse.WriteHeader(404)
		return false, "error request method!"
	}
	if matchedURI == "" {
		httpResponse.WriteHeader(404)
		return false, "url is not exist!"
	}
	info := dao.GetMapping(context, gatewayHashkey, matchedURI)
	info.StrategyKey = strategyKey
	context.Map(info)
	return true, ""
}
