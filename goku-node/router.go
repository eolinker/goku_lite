package gokunode

import (
	"fmt"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_log "github.com/eolinker/goku-api-gateway/goku-node/access-log"
	"github.com/eolinker/goku-api-gateway/goku-node/handler"
	"github.com/eolinker/goku-api-gateway/goku-node/plugin-flow"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
	monitor_write "github.com/eolinker/goku-api-gateway/server/monitor/monitor-write"

	gateway_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/gateway-manager"
	strategy_api_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/strategy-api-manager"
	strategy_api_plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/strategy-api-plugin-manager"
	strategy_plugin_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/strategy-plugin-manager"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"github.com/eolinker/goku-api-gateway/goku-node/visit"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
	// "time"
)

// Router 路由
type Router struct {
	mu map[string]http.HandlerFunc
}

func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path := r.URL.Path
	h, has := mux.mu[path]
	if has {
		h.ServeHTTP(w, r)
		return
	}
	ServeHTTP(w, r)

}

// NewRouter 创建新路由
func NewRouter() http.Handler {

	r := &Router{
		mu: make(map[string]http.HandlerFunc),
	}

	hs := handler.Handler()

	for _, h := range hs {
		r.mu[h.Pattern] = h.HandlerFunc
	}

	return r
}

var systemRequestPath = []string{"/oauth2/token", "/oauth2/authorize", "/oauth2/verify"}

//ServeHTTP httpHandle
func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
		}
	}()

	timeStart := time.Now()

	logFields := make(log.Fields)

	// 记录访问次数
	requestID := GetRandomString(16)

	ctx := common.NewContext(req, requestID, w)
	proxyStatusCode := 0

	log.Debug(requestID, " url: ", req.URL.String())
	log.Debug(requestID, " header: ", ctx.RequestOrg.Header.String())
	rawBody, err := ctx.RequestOrg.RawBody()
	if err == nil {
		log.Debug(requestID, " body: ", string(rawBody))
	}

	defer func() {
		n, status := ctx.Finish()

		if ctx.ProxyResponseHandler != nil {
			proxyStatusCode = ctx.ProxyResponseHandler.StatusCode()
		}
		logFields[access_field.RequestID] = requestID
		logFields[access_field.StatusCode] = status
		logFields[access_field.HTTPUserAgent] = fmt.Sprint("\"", req.UserAgent(), "\"")
		logFields[access_field.HTTPReferer] = req.Referer()
		logFields[access_field.RequestTime] = time.Since(timeStart)
		logFields[access_field.Request] = fmt.Sprint("\"", req.Method, " ", req.URL.Path, " ", req.Proto, "\"")
		logFields[access_field.BodyBytesSent] = n
		logFields[access_field.Host] = req.Host
		access_log.Log(logFields)
		log.WithFields(logFields).Info()

		for _, path := range systemRequestPath {
			if path == req.URL.Path {
				return
			}
		}
		apiID := strconv.Itoa(ctx.ApiID())

		monitor_write.AddMonitor(ctx.StrategyId(), apiID, proxyStatusCode, ctx.StatusCode())
	}()

	remoteAddr := Intercept(req.RemoteAddr, ":")
	logFields[access_field.RemoteAddr] = remoteAddr

	if realIP := ctx.GetHeader("X-Real-Ip"); realIP == "" {
		ctx.ProxyRequest.SetHeader("X-Real-Ip", remoteAddr)
		logFields[access_field.HTTPXForwardedFor] = remoteAddr
	} else {
		logFields[access_field.HTTPXForwardedFor] = realIP
	}

	// 匹配URI前执行函数
	var isBefor bool
	start := time.Now()
	isBefor = pluginflow.BeforeMatch(ctx)
	log.Info(requestID, " BeforeMatch plugin duration:", time.Since(start))
	if !isBefor {
		log.Info(requestID, " stop by BeforeMatch plugin")
		return
	}

	var timeout, retryCount int

	strategyID, ok := retrieveStrategyID(ctx)
	if !ok {
		return
	}

	logFields[access_field.Strategy] = fmt.Sprintf("\"%s %s\"", strategyID, ctx.StrategyName())

	requestPath := req.URL.Path
	requestMenthod := ctx.Request().Method()

	var handleFunc []*entity.PluginHandlerExce
	apiInfo, splitURL, param, ok := strategy_api_manager.CheckAPIFromStrategy(strategyID, requestPath, req.Method)
	if ok {
		ctx.SetAPIID(apiInfo.APIID)
		retryCount = apiInfo.RetryCount
		//ctx.IsMatch = true
		timeout = apiInfo.Timeout

		ctx.ProxyRequest.SetTargetServer(fmt.Sprintf("%s://%s", apiInfo.Protocol, apiInfo.Target))

		targetURL := apiInfo.TargetURL + requestPath

		if apiInfo.StripPrefix {
			targetURL = apiInfo.TargetURL + splitURL
		}
		if apiInfo.StripSlash {
			targetURL = node_common.FilterSlash(targetURL)
		}
		if !apiInfo.IsFollow {
			ctx.ProxyRequest.Method = strings.ToUpper(apiInfo.TargetMethod)
		}
		targetURL = node_common.MatchRestful(targetURL, param)

		ctx.ProxyRequest.SetTargetURL(targetURL)

		handleFunc, _ = strategy_api_plugin_manager.GetPluginsOfAPI(strategyID, apiInfo.APIID)

	} else {
		handleFunc, _ = strategy_plugin_manager.GetPluginsOfStrategy(strategyID)
	}
	start = time.Now()
	isAccess, _ := pluginflow.AccessFunc(ctx, handleFunc)
	log.Info(requestID, " Access plugin duration:", time.Since(start))
	if !isAccess {

		// todo
		return
	}

	if apiInfo == nil {
		log.Info(requestID, " URL dose not exist!")
		ctx.SetStatus(404, "404")
		ctx.SetBody([]byte("[ERROR]URL dose not exist!"))

		return
	}
	logFields[access_field.API] = fmt.Sprintf("\"%d %s\"", apiInfo.APIID, apiInfo.APIName)
	logFields[access_field.Proxy] = fmt.Sprintf("\"%s %s %s\"", ctx.ProxyRequest.Method, ctx.ProxyRequest.TargetURL(), apiInfo.Protocol)
	logFields[access_field.Balance] = apiInfo.Target
	start = time.Now()
	response, err := CreateRequest(ctx, apiInfo, timeout, retryCount)
	log.Info(requestID, " Proxy request duration:", time.Since(start))
	if err != nil {
		log.Warn(err.Error())
	}
	logFields[access_field.FinallyServer] = ctx.FinalTargetServer()
	logFields[access_field.Retry] = ctx.RetryTargetServers()
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
		logFields[access_field.ProxyStatusCode] = response.StatusCode
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
	isProxy, _ := pluginflow.ProxyFunc(ctx, handleFunc)
	log.Info(requestID, " Proxy plugin Duration:", time.Since(start))
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
