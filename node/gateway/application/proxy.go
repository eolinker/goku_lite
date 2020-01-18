package application

import (
	"fmt"
	"strings"

	"github.com/eolinker/goku-api-gateway/goku-service/application"

	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/node/gateway/application/backend"
	"github.com/eolinker/goku-api-gateway/node/gateway/application/interpreter"

	"github.com/eolinker/goku-api-gateway/node/gateway/response"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
)

//DefaultApplication default application
type DefaultApplication struct {
	output        response.Encoder
	backend       *backend.Proxy
	static        *staticeResponse
	balanceTarget string
}

//NewDefaultApplication create new default application
func NewDefaultApplication(apiContent *config.APIContent, target string) *DefaultApplication {

	app := &DefaultApplication{
		backend:       nil,
		static:        nil,
		balanceTarget: target,
		output:        response.GetEncoder(apiContent.OutPutEncoder),
	}
	if len(apiContent.Steps) == 1 {
		step := apiContent.Steps[0]
		app.backend = backend.NewProxyBackendTarget(step, apiContent.RequestURL, target)
	}
	if apiContent.StaticResponse != "" {
		staticResponseStrategy := config.Parse(apiContent.StaticResponseStrategy)
		app.static = newStaticeResponse(apiContent.StaticResponse, staticResponseStrategy)
	}

	return app
}

//Execute execute
func (app *DefaultApplication) Execute(ctx *common.Context) {

	ctx.LogFields[access_field.Balance] = app.balanceTarget

	if app.backend != nil {
		orgBody, _ := ctx.ProxyRequest.RawBody()

		variables := interpreter.NewVariables(orgBody, nil, ctx.ProxyRequest.Headers(), ctx.ProxyRequest.Cookies(), ctx.RestfulParam, ctx.ProxyRequest.Querys(), 1)

		r, err := app.backend.Send(ctx, variables)
		if r != nil {

			ctx.ProxyRequest.Method = r.Method
			ctx.ProxyRequest.SetTargetURL(r.TargetURL)

			ctx.SetRetryTargetServers(strings.Join(r.RetryTargetServers, ","))
			ctx.SetFinalTargetServer(r.FinalTargetServer)

			ctx.LogFields[access_field.FinallyServer] = ctx.FinalTargetServer()
			ctx.LogFields[access_field.Retry] = ctx.RetryTargetServers()
			ctx.LogFields[access_field.Proxy] = fmt.Sprintf("\"%s %s %s\"", r.Method, application.URLPath(r.TargetURL, ctx.ProxyRequest.Querys()), r.Protocol)

		}
		if err != nil {
			ctx.SetStatus(504, "504")
			ctx.SetBody([]byte("[ERROR]Fail to get response after proxy!"))
			log.Warn(err)
			return
		}

		ctx.LogFields[access_field.ProxyStatusCode] = r.StatusCode

		body, err := app.output.Encode(r.Body, r.BodyOrg)
		if err != nil {
			body = r.BodyOrg
		}
		ctx.SetProxyResponseHandler(common.NewResponseReader(r.Header, r.StatusCode, r.Status, body))

		return

	}
	if app.static != nil {
		app.static.Do(ctx)
		return
	}

	ctx.SetStatus(504, "504")
	ctx.SetBody([]byte("[ERROR]Fail to get response after proxy!"))

}
