package application

import (
	"fmt"
	goku_plugin "github.com/eolinker/goku-plugin"
	"net/http"
	"net/url"
	"time"
)

func request(ctx goku_plugin.ContextAccess,method string, backendDomain string, query url.Values, header http.Header, body []byte, timeout time.Duration) (*http.Response, error) {

	if backendDomain == "" {
		return nil, fmt.Errorf("invaild url")
	}
	u, err := url.ParseRequestURI(backendDomain)
	if err != nil {

		return nil, err
	}
	u.Query()
	req, err := NewRequest(method, u)
	if err != nil {

		return nil, err
	}

	queryDest := u.Query()
	if query != nil {
		for k, vs := range query {
			for _, v := range vs {
				queryDest.Add(k, v)
			}
		}
	}

	req.headers = header

	req.queryParams = queryDest

	req.SetRawBody(body)
	if timeout != 0 {
		req.SetTimeout(timeout)
	}
	return req.Send(ctx)
}
