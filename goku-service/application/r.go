package application

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func request(method string, backendDomain string, query url.Values, header http.Header, body []byte, timeout time.Duration) (*http.Response, error) {

	if backendDomain == "" {
		return nil, fmt.Errorf("invaild url")
	}

	u, err := url.ParseRequestURI(backendDomain)
	if err != nil {

		return nil, err
	}
	req, err := NewRequest(method, u)
	if err != nil {

		return nil, err
	}

	req.headers = header

	req.queryParams = query

	req.SetRawBody(body)
	if timeout != 0 {
		req.SetTimeout(timeout)
	}
	return req.Send()
}
