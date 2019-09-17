package application

import (
	"fmt"
	"github.com/eolinker/goku/utils"
	"net/http"
	"net/url"
	"time"
)

type Org struct {
	server string
}

// 忽略重试
func (app *Org) Send(proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error) {

	var response *http.Response = nil
	var err error = nil

	FinalTargetServer := ""
	RetryTargetServers := make([]string, 0, retry+1)

	path = utils.TrimPrefixAll(path, "/")

	for doTrice := retry + 1; doTrice > 0; doTrice-- {

		u := fmt.Sprintf("%s://%s/%s", proto, app.server, path)
		FinalTargetServer = app.server
		RetryTargetServers = append(RetryTargetServers, FinalTargetServer)
		response, err = request(method, u, querys, header, body, timeout)
		if err != nil {
			continue
		} else {
			return response, FinalTargetServer, RetryTargetServers, err
		}
	}

	return response, FinalTargetServer, RetryTargetServers, err
}

func NewOrg(server string) IHttpApplication {
	return &Org{
		server: server,
	}
}
