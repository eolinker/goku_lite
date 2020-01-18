package backend

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/goku-service/application"
	"github.com/eolinker/goku-api-gateway/goku-service/balance"

	"time"

	"github.com/eolinker/goku-api-gateway/node/gateway/application/interpreter"
	"github.com/eolinker/goku-api-gateway/node/gateway/response"
)

//Proxy proxy
type Proxy struct {
	//step * config.APIStepConfig
	BalanceName string
	Balance     application.IHttpApplication
	HasBalance  bool
	Protocol    string

	Method  string
	Path    interpreter.Interpreter
	OrgPath string
	Decode  response.DecodeHandle

	RequestPath string

	Retry   int
	TimeOut time.Duration
}

//NewProxyBackendTarget 创建新的转发后端目标
func NewProxyBackendTarget(step *config.APIStepConfig, requestPath string, balanceTarget string) *Proxy {
	b := &Proxy{
		BalanceName: balanceTarget,
		Protocol:    step.Proto,
		Method:      strings.ToUpper(step.Method),
		Path:        interpreter.GenPath(step.Path),

		RequestPath: requestPath,
		Decode:      response.GetDecoder(step.Decode),

		TimeOut: time.Duration(step.TimeOut) * time.Millisecond,
		Retry:   step.Retry,
	}

	b.Balance, b.HasBalance = balance.GetByName(balanceTarget)

	return b
}

//Send send
func (b *Proxy) Send(ctx *common.Context, variables *interpreter.Variables) (*BackendResponse, error) {

	if !b.HasBalance {
		err := fmt.Errorf("get balance error:%s", b.BalanceName)
		return nil, err
	}

	path := b.Path.Execution(variables)

	// 不是restful时，将匹配路由之后对url拼接到path之后
	if len(variables.Restful) == 0 {
		orgRequestURL := ctx.RequestOrg.URL().Path
		lessPath := strings.TrimPrefix(orgRequestURL, b.RequestPath)
		lessPath = strings.TrimPrefix(lessPath, "/")
		if lessPath != "" {
			path = strings.TrimSuffix(path, "/")
			path = fmt.Sprint(path, "/", lessPath)
		}
	}

	method := b.Method
	if method == "FOLLOW" {
		method = ctx.ProxyRequest.Method
	}
	r, finalTargetServer, retryTargetServers, err := b.Balance.Send(ctx, b.Protocol, method, path, ctx.ProxyRequest.Querys(), ctx.ProxyRequest.Headers(), variables.Org, b.TimeOut, b.Retry)

	backendResponse := &BackendResponse{
		Method:     method,
		Protocol:   b.Protocol,
		StatusCode: 200,
		Status:     "200",
		//Response:           r,
		TargetURL:          path,
		FinalTargetServer:  finalTargetServer,
		RetryTargetServers: retryTargetServers,

		//Cookies:r.Cookies(),
	}
	if err != nil {
		backendResponse.StatusCode, backendResponse.Status = 503, "503"
		return backendResponse, err
	}
	backendResponse.Header = r.Header
	defer r.Body.Close()
	bd := r.Body
	if r.Header.Get("Content-Encoding") == "gzip" {
		bd, _ = gzip.NewReader(r.Body)
		r.Header.Del("Content-Encoding")
	}
	backendResponse.BodyOrg, err = ioutil.ReadAll(bd)
	if err != nil {
		return backendResponse, nil
	}

	if b.Decode != nil {
		rp, e := response.Decode(backendResponse.BodyOrg, b.Decode)
		if e != nil {
			backendResponse.Body = nil
		} else {
			backendResponse.Body = rp.Data
		}
	}

	return backendResponse, nil

}
