package application

import (
	"fmt"
	goku_plugin "github.com/eolinker/goku-plugin"
	"net/http"
	"net/url"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"github.com/eolinker/goku-api-gateway/goku-service/health"
	"github.com/eolinker/goku-api-gateway/utils"
)

//Application 应用
type Application struct {
	service            *common.Service
	healthCheckHandler health.CheckHandler
}

//NewApplication 创建Application
func NewApplication(service *common.Service, healthCheckHandler health.CheckHandler) *Application {
	return &Application{
		service:            service,
		healthCheckHandler: healthCheckHandler,
	}

}

//Send send
func (app *Application) Send(ctx goku_plugin.ContextAccess,proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error) {

	var response *http.Response
	var err error

	FinalTargetServer := ""
	RetryTargetServers := make([]string, 0, retry+1)

	lastIndex := -1
	path = utils.TrimPrefixAll(path, "/")
	for doTrice := retry + 1; doTrice > 0; doTrice-- {
		instance, index, has := app.service.Next(lastIndex)
		lastIndex = index
		if !has {
			return nil, FinalTargetServer, RetryTargetServers, fmt.Errorf("not found instance for app:%s", app.service.Name)
		}

		FinalTargetServer = instance.IP
		if instance.Port != 0 {
			FinalTargetServer = fmt.Sprintf("%s:%d", instance.IP, instance.Port)
		}

		RetryTargetServers = append(RetryTargetServers, FinalTargetServer)
		u := fmt.Sprintf("%s://%s/%s", proto, FinalTargetServer, path)
		response, err = request(ctx,method, u, querys, header, body, timeout)

		if err != nil {
			if app.healthCheckHandler.IsNeedCheck() {
				app.healthCheckHandler.Check(instance)
			}
		} else {
			return response, FinalTargetServer, RetryTargetServers, err
		}

	}

	return response, FinalTargetServer, RetryTargetServers, err
}
