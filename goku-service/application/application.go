package application

import (
	goku_plugin "github.com/eolinker/goku-plugin"
	"net/http"
	"net/url"
	"time"
)

//IHttpApplication iHttpApplication
type IHttpApplication interface {
	Send(ctx goku_plugin.ContextAccess,Proto string, method string, path string, querys url.Values, header http.Header, body []byte, timeout time.Duration, retry int) (*http.Response, string, []string, error)
}
