package application

import (
	"fmt"
	goku_plugin "github.com/eolinker/goku-plugin"
	"net/http"
	"net/url"
	"time"

	"github.com/eolinker/goku-api-gateway/utils"
)

//Org org
type Org struct {
	server string
}

//Send 请求发送，忽略重试
func (app *Org) Send(ctx goku_plugin.ContextAccess,proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error) {

	var response *http.Response
	var err error

	FinalTargetServer := ""
	RetryTargetServers := make([]string, 0, retry+1)

	path = utils.TrimPrefixAll(path, "/")

	for doTrice := retry + 1; doTrice > 0; doTrice-- {

		u := fmt.Sprintf("%s://%s/%s", proto, app.server, path)
		FinalTargetServer = app.server
		RetryTargetServers = append(RetryTargetServers, FinalTargetServer)
		response, err = request(ctx,method, u, querys, header, body, timeout)
		if err != nil {
			continue
		} else {
			return response, FinalTargetServer, RetryTargetServers, err
		}
	}

	return response, FinalTargetServer, RetryTargetServers, err
}

//NewOrg 创建新的IHttpApplication
func NewOrg(server string) IHttpApplication {
	return &Org{
		server: server,
	}
}
