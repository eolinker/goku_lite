package application

import (
	"net/http"
	"net/url"
	"time"
)

type IHttpApplication interface {
	Send(Proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error)
}
