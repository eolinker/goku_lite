package application

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	// "fmt"
	"time"
)

var Version string = "2.0"

type Request struct {
	client  *http.Client
	method  string
	URL     string
	headers map[string][]string
	body    []byte

	queryParams map[string][]string

	timeout time.Duration
}

// 创建新请求
func NewRequest(method string, URL *url.URL) (*Request, error) {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" &&
		method != "HEAD" && method != "OPTIONS" && method != "PATCH" {
		return nil, errors.New("Unsupported Request Method")
	}
	return newRequest(method, URL)
}

//
func newRequest(method string, URL *url.URL) (*Request, error) {
	var urlPath string
	queryParams := make(map[string][]string)
	for key, values := range URL.Query() {
		queryParams[key] = values
	}
	urlPath = URL.Scheme + "://" + URL.Host + URL.Path
	r := &Request{
		client:      &http.Client{},
		method:      method,
		URL:         urlPath,
		headers:     make(map[string][]string),
		queryParams: queryParams,
	}
	return r, nil
}

// 设置请求头
func (this *Request) SetHeader(key string, values ...string) {
	if len(values) > 0 {
		this.headers[key] = values[:]
	} else {
		delete(this.headers, key)
	}
}

// 获取请求头
func (this *Request) Headers() map[string][]string {
	headers := make(map[string][]string)
	for key, values := range this.headers {
		headers[key] = values[:]
	}
	return headers
}

// 设置Query参数
func (this *Request) SetQueryParam(key string, values ...string) {
	if len(values) > 0 {
		this.queryParams[key] = values[:]
	} else {
		delete(this.queryParams, key)
	}
}

// 设置请求超时时间
func (this *Request) SetTimeout(timeout time.Duration) {
	this.timeout = timeout
}

//// 获取请求超时时间
//func (this *Request) GetTimeout() time.Duration {
//	return this.timeout
//}

// 发送请求
func (this *Request) Send() (*http.Response, error) {
	// now := time.Now()
	req, err := this.parseBody()
	// fmt.Println("Parse body",time.Since(now))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "gzip")
	// now = time.Now()
	req.Header = parseHeaders(this.headers)

	this.client.Timeout = this.timeout

	httpResponse, err := this.client.Do(req)

	if err != nil {
		return nil, err
	}
	return httpResponse, nil

}

// 获取query参数
func (this *Request) QueryParams() map[string][]string {
	params := make(map[string][]string)
	for key, values := range this.queryParams {
		params[key] = values[:]
	}
	return params
}

// 获取完整的URL路径
func (this *Request) UrlPath() string {
	if len(this.queryParams) > 0 {
		return this.URL + "?" + parseParams(this.queryParams).Encode()
	} else {
		return this.URL
	}
}

// 设置URL
func (this *Request) SetURL(url string) {
	this.URL = url
}

// 设置源数据
func (this *Request) SetRawBody(body []byte) {
	this.body = body
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
func (this *Request) parseBody() (req *http.Request, err error) {
	var body io.Reader = nil
	if len(this.body) > 0 {
		body = bytes.NewBuffer(this.body)

	}
	req, err = http.NewRequest(this.method, this.UrlPath(), body)
	return
	//if this.method == "GET" || this.method == "TRACE" {
	//	req, err = http.NewRequest(this.method, this.UrlPath(), nil)
	//}
	//
	//if len(this.body) > 0 {
	//
	//	if this.isJSON {
	//		if _, ok := this.headers["Content-Type"]; !ok {
	//			this.headers["Content-Type"] = []string{"application/json"}
	//		}
	//		req, err = http.NewRequest(this.method, this.UrlPath(),
	//			strings.NewReader(string(this.body)))
	//	} else {
	//		var body *bytes.Buffer
	//		body = bytes.NewBuffer(this.body)
	//
	//	}
	//} else if len(this.files) > 0 {
	//	body := new(bytes.Buffer)
	//	writer := multipart.NewWriter(body)
	//	var part io.Writer
	//	for fieldname, file := range this.files {
	//		part, err = writer.CreateFormFile(fieldname, file.filename)
	//		if err != nil {
	//			return
	//		}
	//		_, err = part.Write(file.data)
	//		if err != nil {
	//			return
	//		}
	//	}
	//	for fieldname, values := range this.formParams {
	//		temp := make(map[string][]string)
	//		temp[fieldname] = values
	//		value := parseParams(temp).Encode()
	//		err = writer.WriteField(fieldname, value)
	//		if err != nil {
	//			return
	//		}
	//	}
	//	err = writer.Close()
	//	if err != nil {
	//		return
	//	}
	//	this.headers["Content-Type"] = []string{writer.FormDataContentType()}
	//	req, err = http.NewRequest(this.method, this.UrlPath(), body)
	//} else {
	//	this.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	//	req, err = http.NewRequest(this.method, this.UrlPath(),
	//		strings.NewReader(parseParams(this.formParams).Encode()))
	//}
	//return
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
