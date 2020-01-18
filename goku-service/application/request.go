package application

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"

	goku_plugin "github.com/eolinker/goku-plugin"

	"github.com/eolinker/goku-api-gateway/diting"
	goku_labels "github.com/eolinker/goku-api-gateway/goku-labels"
	"github.com/eolinker/goku-api-gateway/node/monitor"

	// "fmt"
	"time"
)

//Version 版本号
var Version = "2.0"

var skipCertificate = 0

//SetSkipCertificate 设置跳过证书
func SetSkipCertificate(skip int) {
	skipCertificate = skip
}

//Request request
type Request struct {
	client  *http.Client
	method  string
	URL     string
	headers map[string][]string
	body    []byte

	queryParams map[string][]string

	timeout time.Duration
}

//NewRequest 创建新请求
func NewRequest(method string, URL *url.URL) (*Request, error) {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" &&
		method != "HEAD" && method != "OPTIONS" && method != "PATCH" {
		return nil, errors.New("Unsupported Request Method")
	}
	return newRequest(method, URL)
}

//URLPath urlPath
func URLPath(url string, query url.Values) string {
	if len(query) < 1 {
		return url
	}
	return url + "?" + query.Encode()
}

func newRequest(method string, URL *url.URL) (*Request, error) {
	var urlPath string
	queryParams := make(map[string][]string)
	for key, values := range URL.Query() {
		queryParams[key] = values
	}
	urlPath = URL.Scheme + "://" + URL.Host + URL.Path
	tp := http.DefaultTransport
	if skipCertificate == 1 {
		tp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	r := &Request{
		client:      &http.Client{Transport: tp},
		method:      method,
		URL:         urlPath,
		headers:     make(map[string][]string),
		queryParams: queryParams,
	}
	return r, nil
}

//SetHeader 设置请求头
func (r *Request) SetHeader(key string, values ...string) {
	if len(values) > 0 {
		r.headers[key] = values[:]
	} else {
		delete(r.headers, key)
	}
}

//Headers 获取请求头
func (r *Request) Headers() map[string][]string {
	headers := make(map[string][]string)
	for key, values := range r.headers {
		headers[key] = values[:]
	}
	return headers
}

//SetQueryParam 设置Query参数
func (r *Request) SetQueryParam(key string, values ...string) {
	if len(values) > 0 {
		r.queryParams[key] = values[:]
	} else {
		delete(r.queryParams, key)
	}
}

//SetTimeout 设置请求超时时间
func (r *Request) SetTimeout(timeout time.Duration) {
	r.timeout = timeout
}

//// 获取请求超时时间
//func (r *Request) GetTimeout() time.Duration {
//	return r.timeout
//}

//Send 发送请求
func (r *Request) Send(ctx goku_plugin.ContextAccess) (*http.Response, error) {
	// now := time.Now()
	req, err := r.parseBody()
	if err != nil {
		return nil, err
	}
	status := 0
	start := time.Now()
	defer func() {
		delay := time.Since(start)
		labels := make(diting.Labels)

		labels[goku_labels.Proto] = req.Proto
		labels[goku_labels.Host] = req.Host
		labels[goku_labels.Path] = req.URL.Path
		labels[goku_labels.Method] = req.Method
		labels[goku_labels.API] = strconv.Itoa(ctx.ApiID())
		labels[goku_labels.Strategy] = ctx.StrategyId()
		labels[goku_labels.Status] = strconv.Itoa(status)
		monitor.ProxyMonitor.Observe(float64(delay/time.Millisecond), labels)
	}()
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header = parseHeaders(r.headers)

	r.client.Timeout = r.timeout

	httpResponse, err := r.client.Do(req)

	if err != nil {
		if netErr, ok := err.(net.Error); ok {
			if netErr.Timeout() {
				status = 504
			} else {
				status = 503
			}
		} else {
			status = 503
		}
		return nil, err
	}
	status = httpResponse.StatusCode
	return httpResponse, nil

}

//QueryParams 获取query参数
func (r *Request) QueryParams() map[string][]string {
	params := make(map[string][]string)
	for key, values := range r.queryParams {
		params[key] = values[:]
	}
	return params
}

//URLPath 获取完整的URL路径
func (r *Request) URLPath() string {
	if len(r.queryParams) > 0 {
		return r.URL + "?" + parseParams(r.queryParams).Encode()
	}
	return r.URL
}

//SetURL 设置URL
func (r *Request) SetURL(url string) {
	r.URL = url
}

//SetRawBody 设置源数据
func (r *Request) SetRawBody(body []byte) {
	r.body = body
}

// 解析请求头
func parseHeaders(headers map[string][]string) http.Header {
	h := http.Header{}
	for key, values := range headers {
		for _, value := range values {
			h.Add(key, value)
		}
	}

	_, hasAccept := h["Accept"]
	if !hasAccept {
		h.Add("Accept", "*/*")
	}
	_, hasAgent := h["User-Agent"]
	if !hasAgent {
		h.Add("User-Agent", "goku-requests/"+Version)
	}
	return h
}

// 解析请求体
func (r *Request) parseBody() (req *http.Request, err error) {
	var body io.Reader
	if len(r.body) > 0 {
		body = bytes.NewBuffer(r.body)

	}
	req, err = http.NewRequest(r.method, r.URLPath(), body)
	return

}

// 解析参数
func parseParams(params map[string][]string) url.Values {
	v := url.Values{}
	for key, values := range params {
		for _, value := range values {
			v.Add(key, value)
		}
	}
	return v
}

// 解析URL
func parseURL(urlPath string) (URL *url.URL, err error) {
	URL, err = url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	if URL.Scheme != "http" && URL.Scheme != "https" {
		urlPath = "http://" + urlPath
		URL, err = url.Parse(urlPath)
		if err != nil {
			return nil, err
		}

		if URL.Scheme != "http" && URL.Scheme != "https" {
			return nil, errors.New("[package requests] only HTTP and HTTPS are accepted")
		}
	}
	return
}
