package application

import (
	"fmt"
	"github.com/eolinker/goku/goku-service/common"
	"github.com/eolinker/goku/goku-service/health"
	"github.com/eolinker/goku/utils"
	"net/http"
	"net/url"
	"time"
)

type Application struct {
	service            *common.Service
	healthCheckHandler health.CheckHandler
}

func NewApplication(service *common.Service, healthCheckHandler health.CheckHandler) *Application {
	return &Application{
		service:            service,
		healthCheckHandler: healthCheckHandler,
	}

}
func (app *Application) Send(proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error) {

	var response *http.Response = nil
	var err error = nil

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

		FinalTargetServer = fmt.Sprintf("%s:%d", instance.IP, instance.Port)
		RetryTargetServers = append(RetryTargetServers, FinalTargetServer)
		u := fmt.Sprintf("%s://%s:%d/%s", proto, instance.IP, instance.Port, path)
		response, err = request(method, u, querys, header, body, timeout)

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
