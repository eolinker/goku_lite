package gokunode

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-node/common"

	balance_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/balance-manager"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"

	"strings"
)

// CreateRequest 创建转发请求
func CreateRequest(ctx *common.Context, apiInfo *entity.APIExtend, timeout, retry int) (*http.Response, error) {

	app, has := balance_manager.Get(apiInfo.Target)
	if !has {
		err := fmt.Errorf("get balance error:%s", apiInfo.Target)
		return nil, err
	}
	rawbody, _ := ctx.ProxyRequest.RawBody()

	response, finalTargetServer, retryTargetServers, err := app.Send(apiInfo.Protocol, ctx.ProxyRequest.Method, ctx.ProxyRequest.TargetURL(), ctx.ProxyRequest.Querys(), ctx.ProxyRequest.Headers(), rawbody, time.Duration(timeout)*time.Millisecond, retry)

	ctx.SetRetryTargetServers(strings.Join(retryTargetServers, ","))
	ctx.SetFinalTargetServer(finalTargetServer)

	if err != nil {
		return nil, err
	}
	return response, nil

}
