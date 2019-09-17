package common

import (
	"net/http"
	"strconv"

	goku_plugin "github.com/eolinker/goku-plugin"
)

var _ goku_plugin.ContextProxy = (*Context)(nil)

type Context struct {
	w http.ResponseWriter
	*CookiesHandler
	*PriorityHeader
	*StatusHandler
	*StoreHandler
	RequestOrg           *RequestReader
	ProxyRequest         *Request
	ProxyResponseHandler *ResponseReader
	Body                 []byte
	strategyId           string
	strategyName         string
	apiId                int
	requestId            string
	finalTargetServer    string
	retryTargetServers   string
}

func (ctx *Context) FinalTargetServer() string {
	return ctx.finalTargetServer
}
func (ctx *Context) SetFinalTargetServer(finalTargetServer string) {
	ctx.finalTargetServer = finalTargetServer
}
func (ctx *Context) RetryTargetServers() string {
	return ctx.retryTargetServers
}
func (ctx *Context) SetRetryTargetServers(retryTargetServers string) {
	ctx.retryTargetServers = retryTargetServers
}

func (ctx *Context) Finish() (n int,statusCode int){

	header := ctx.PriorityHeader.header

	statusCode = ctx.StatusHandler.code
	if statusCode == 0 {
		statusCode = 504
	}

	bodyAllowed := true
	switch {
	case statusCode >= 100 && statusCode <= 199:
		bodyAllowed = false
		break
	case statusCode == 204:
		bodyAllowed = false
		break
	case statusCode == 304:
		bodyAllowed = false
		break
	}

	if ctx.PriorityHeader.appendHeader != nil {
		for k, vs := range ctx.PriorityHeader.appendHeader.header {
			for _, v := range vs {
				header.Add(k, v)
			}
		}
	}

	if ctx.PriorityHeader.setHeader != nil {
		for k, vs := range ctx.PriorityHeader.setHeader.header {
			header.Del(k)
			for _, v := range vs {
				header.Add(k, v)
			}
		}
	}

	for k, vs := range ctx.PriorityHeader.header {
		if k == "Content-Length" && bodyAllowed {
			vs = []string{strconv.Itoa(len(string(ctx.Body)))}
		}
		for _, v := range vs {
			ctx.w.Header().Add(k, v)
		}
	}

	ctx.w.WriteHeader(statusCode)

	if !bodyAllowed {
		return 0,statusCode
	}
	n, _ = ctx.w.Write(ctx.Body)
	return n,statusCode
}
func (ctx *Context) RequestId() string {
	return ctx.requestId
}

func NewContext(r *http.Request, requestId string, w http.ResponseWriter) *Context {
	requestreader := NewRequestReader(r)
	return &Context{
		CookiesHandler:       newCookieHandle(r.Header),
		PriorityHeader:       NewPriorityHeader(),
		StatusHandler:        NewStatusHandler(),
		StoreHandler:         NewStoreHandler(),
		RequestOrg:           requestreader,
		ProxyRequest:         NewRequest(requestreader),
		ProxyResponseHandler: nil,
		requestId:            requestId,
		w:                    w,
	}
}
func (ctx *Context) SetProxyResponse(response *http.Response) {

	ctx.ProxyResponseHandler = newResponseReader(response)
	if ctx.ProxyResponseHandler != nil {
		ctx.Body = ctx.ProxyResponseHandler.body
		ctx.SetStatus(ctx.ProxyResponseHandler.StatusCode(), ctx.ProxyResponseHandler.Status())
		ctx.header = ctx.ProxyResponseHandler.header
	}

}
func (ctx *Context) Write(w http.ResponseWriter) {
	if ctx.StatusCode() == 0 {
		ctx.SetStatus(200, "200 ok")
	}
	if ctx.Body != nil {
		w.Write(ctx.Body)
	}

	w.WriteHeader(ctx.StatusCode())

}

func (ctx *Context) GetBody() []byte {
	return ctx.Body
}
func (ctx *Context) SetBody(data []byte) {
	ctx.Body = data
}

func (ctx *Context) ProxyResponse() goku_plugin.ResponseReader {
	return ctx.ProxyResponseHandler
}

func (ctx *Context) StrategyId() string {
	return ctx.strategyId
}
func (ctx *Context) SetStrategyId(strategyId string) {
	ctx.strategyId = strategyId
}
func (ctx *Context) StrategyName() string {
	return ctx.strategyName
}
func (ctx *Context) SetStrategyName(strategyName string) {
	ctx.strategyName = strategyName
}
func (ctx *Context) ApiID() int {
	return ctx.apiId
}
func (ctx *Context) SetApiID(apiId int) {
	ctx.apiId = apiId
}
func (ctx *Context) Request() goku_plugin.RequestReader {
	return ctx.RequestOrg
}

func (ctx *Context) Proxy() goku_plugin.Request {
	return ctx.ProxyRequest
}
