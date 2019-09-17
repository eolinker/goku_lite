package goku_node

import (
	"fmt"

	log "github.com/eolinker/goku/goku-log"
	access_log "github.com/eolinker/goku/goku-node/access-log"
	"github.com/eolinker/goku/goku-node/handler"
	"github.com/eolinker/goku/goku-node/plugin-flow"
	. "github.com/eolinker/goku/server/access-field"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku/goku-node/common"
	monitor_write "github.com/eolinker/goku/server/monitor/monitor-write"

	gateway_manager "github.com/eolinker/goku/goku-node/manager/gateway-manager"
	strategy_api_manager "github.com/eolinker/goku/goku-node/manager/strategy-api-manager"
	strategy_api_plugin_manager "github.com/eolinker/goku/goku-node/manager/strategy-api-plugin-manager"
	strategy_plugin_manager "github.com/eolinker/goku/goku-node/manager/strategy-plugin-manager"
	node_common "github.com/eolinker/goku/goku-node/node-common"
	"github.com/eolinker/goku/goku-node/visit"
	entity "github.com/eolinker/goku/server/entity/node-entity"
	// "time"
)

type Router struct {
	mu map[string]http.HandlerFunc
}

func (mux*Router) ServeHTTP(w http.ResponseWriter, r*http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path:=r.URL.Path
	h, has:= mux.mu[path]
	if has{
		h.ServeHTTP(w, r)
		return
	}
	ServeHTTP(w,r)

}

func NewRouter() http.Handler {

	r:=&Router{
		mu :make(map[string]http.HandlerFunc),
	}


	hs:= handler.Handler()

	for _,h:=range hs{
		r.mu[h.Pattern] = h.HandlerFunc
	}

	return  r
}

var systemRequestPath = []string{"/oauth2/token", "/oauth2/authorize", "/oauth2/verify"}

func   ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
		}
	}()


	timeStart:=time.Now()

	logFields:=make(log.Fields)

	// 记录访问次数
	requestID := GetRandomString(16)


	ctx := common.NewContext(req, requestID, w)
	proxyStatusCode := 0

	log.Debug(requestID," url: ",req.URL.String())
	log.Debug(requestID," header: ",ctx.RequestOrg.Header.String())
	 rawBody ,err:=ctx.RequestOrg.RawBody()
	 if err==nil{
		 log.Debug(requestID," body: ",string(rawBody))
	 }

	defer func() {
		n,status:=ctx.Finish()

		if ctx.ProxyResponseHandler != nil {
			proxyStatusCode = ctx.ProxyResponseHandler.StatusCode()
		}
		logFields[RequestId]= requestID
		logFields[StatusCode] = status
		logFields[HttpUserAgent] = fmt.Sprint("\"",req.UserAgent(),"\"")
		logFields[HttpReferer] = req.Referer()
		logFields[RequestTime] = time.Since(timeStart)
		logFields[Request] = fmt.Sprint("\"",req.Method," ",req.URL.Path," ",req.Proto,"\"")
		logFields[BodyBytesSent] = n
		logFields[Host] = req.Host
		access_log.Log(logFields)
		log.WithFields(logFields).Info()

		for _, path := range systemRequestPath {
			if path == req.URL.Path {
				return
			}
		}
		apiId := strconv.Itoa(ctx.ApiID())

		monitor_write.AddMonitor(ctx.StrategyId(), apiId, proxyStatusCode, ctx.StatusCode())
	}()

	remoteAddr := Intercept(req.RemoteAddr, ":")
	logFields[RemoteAddr]= remoteAddr

	if realIp := ctx.GetHeader("X-Real-Ip") ;realIp == ""{
		ctx.ProxyRequest.SetHeader("X-Real-Ip", remoteAddr)
		logFields[HttpXForwardedFor] = remoteAddr
	}else{
		logFields[HttpXForwardedFor] = realIp
	}

	// 匹配URI前执行函数
	var isBefor bool
	start := time.Now()
	isBefor = plugin_flow.BeforeMatch(ctx)
	log.Info(requestID," BeforeMatch plugin duration:",time.Since(start))
	if !isBefor {
		log.Info(requestID," stop by BeforeMatch plugin")
		return
	}

	var timeout, retryCount int


	strategyID, ok := retrieveStrategyID(ctx)
	if !ok {
		return
	}

	logFields[Strategy] = fmt.Sprintf("\"%s %s\"", strategyID,ctx.StrategyName())

	requestPath := req.URL.Path
	requestMenthod := ctx.Request().Method()

	var handleFunc []*entity.PluginHandlerExce
	apiInfo, splitURL, param, ok := strategy_api_manager.CheckApiFromStrategy(strategyID, requestPath, req.Method)
	if ok {
		ctx.SetApiID(apiInfo.ApiID)
		retryCount = apiInfo.RetryCount
		//ctx.IsMatch = true
		timeout = apiInfo.Timeout


		ctx.ProxyRequest.SetTargetServer(fmt.Sprintf("%s://%s", apiInfo.Protocol, apiInfo.Target))

		targetUrl := apiInfo.TargetURL + requestPath

		if apiInfo.StripPrefix {
			targetUrl = apiInfo.TargetURL + splitURL
		}
		if apiInfo.StripSlash {
			targetUrl = node_common.FilterSlash(targetUrl)
		}
		if !apiInfo.IsFollow {
			ctx.ProxyRequest.Method = strings.ToUpper(apiInfo.TargetMethod)
		}
		targetUrl = node_common.MatchRestful(targetUrl, param)

		ctx.ProxyRequest.SetTargetURL(targetUrl)

		handleFunc, _ = strategy_api_plugin_manager.GetPluginsOfApi(strategyID, apiInfo.ApiID)

	} else {
		handleFunc, _ = strategy_plugin_manager.GetPluginsOfStrategy(strategyID)
	}
	start = time.Now()
	isAccess, _ := plugin_flow.AccessFunc(ctx, handleFunc)
	log.Info(requestID," Access plugin duration:",time.Since(start))
	if !isAccess {

		// todo
		return
	}

	if apiInfo == nil {
		log.Info(requestID," URL dose not exist!")
		ctx.SetStatus(404, "404")
		ctx.SetBody([]byte("[ERROR]URL dose not exist!"))
		
		return
	}
	logFields[Api] = fmt.Sprintf("\"%d %s\"",apiInfo.ApiID,apiInfo.ApiName)
	logFields[Proxy] = fmt.Sprintf("\"%s %s %s\"",ctx.ProxyRequest.Method,ctx.ProxyRequest.TargetURL(),apiInfo.Protocol)
	logFields[Balance] = apiInfo.Target
	start = time.Now()
	err, response := CreateRequest(ctx, apiInfo, timeout, retryCount)
	log.Info(requestID," Proxy request duration:",time.Since(start))
	if err != nil {
		 log.Warn(err.Error())
	}
	logFields[FinallyServer] = ctx.FinalTargetServer()
	logFields[Retry] = ctx.RetryTargetServers()
	ctx.SetProxyResponse(response)
	form, _ := ctx.RequestOrg.BodyForm()
	if response == nil {
		proxyStatusCode = -1
		// ctx.Proxy
		go visit.UpdateProxyFailureCount(apiInfo,
			requestMenthod, ctx.ProxyRequest.Method,
			ctx.RequestOrg.Headers(),
			ctx.RequestOrg.URL().Query(),
			form,
			504,
			make(map[string][]string),

			ctx)
	} else {
		logFields[ProxyStatusCode] = response.StatusCode
		// w.WriteHeader(ctx.ProxyStatusCode)
		if !gateway_manager.IsSucess(ctx.StatusCode()) {
			go visit.UpdateProxyFailureCount(
				apiInfo,
				requestMenthod,
				ctx.ProxyRequest.Method,
				ctx.RequestOrg.Headers(),
				ctx.RequestOrg.URL().Query(),
				form,
				ctx.ProxyResponseHandler.StatusCode(),
				ctx.ProxyResponseHandler.Headers(),
				ctx)
		}
	}

	start = time.Now()
	isProxy, _ := plugin_flow.ProxyFunc(ctx, handleFunc)
	log.Info(requestID," Proxy plugin Duration:",time.Since(start))
	if !isProxy {
		return
	}
	// 默认返回插件未开启时，直接返回回复相关内容
	if response != nil {

	} else {
		ctx.SetStatus(504, "504")
		ctx.SetBody([]byte("[ERROR]Fail to get response after proxy!"))
	}

	//logFunc(ctx, handleFunc)

	return
}
