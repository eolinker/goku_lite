package goku_node

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eolinker/goku/goku-node/common"

	balance_manager "github.com/eolinker/goku/goku-node/manager/balance-manager"
	entity "github.com/eolinker/goku/server/entity/node-entity"

	"strings"
)

// 创建转发请求
func CreateRequest(ctx *common.Context, apiInfo *entity.ApiExtend, timeout, retry int) (error, *http.Response) {

	app, has := balance_manager.Get(apiInfo.Target)
	if !has {
		err := fmt.Errorf("get balance error:%s", apiInfo.Target)
		return err, nil
	}
	rawbody, _ := ctx.ProxyRequest.RawBody()

	response, finalTargetServer, retryTargetServers, err := app.Send(apiInfo.Protocol, ctx.ProxyRequest.Method, ctx.ProxyRequest.TargetURL(), ctx.ProxyRequest.Querys(), ctx.ProxyRequest.Headers(), rawbody, time.Duration(timeout)*time.Millisecond, retry)

	ctx.SetRetryTargetServers(strings.Join(retryTargetServers, ","))
	ctx.SetFinalTargetServer(finalTargetServer)

	if err != nil {
		return err, nil
	}
	return nil, response

}
