package request

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	// "fmt"
	"time"
)

const requestVersion = "2.0"

//Request request
type Request struct {
	client  *http.Client
	method  string
	URL     string
	headers map[string][]string
	body    []byte

	queryParams map[string][]string

	timeout int
}

//NewRequest 创建新请求
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

//SetHeader 设置请求头
func (rq *Request) SetHeader(key string, values ...string) {
	if len(values) > 0 {
		rq.headers[key] = values[:]
	} else {
		delete(rq.headers, key)
	}
}

//Headers 获取请求头
func (rq *Request) Headers() map[string][]string {
	headers := make(map[string][]string)
	for key, values := range rq.headers {
		headers[key] = values[:]
	}
	return headers
}

//SetQueryParam 设置Query参数
func (rq *Request) SetQueryParam(key string, values ...string) {
	if len(values) > 0 {
		rq.queryParams[key] = values[:]
	} else {
		delete(rq.queryParams, key)
	}
}

//SetTimeout 设置请求超时时间
func (rq *Request) SetTimeout(timeout int) {
	rq.timeout = timeout
}

//GetTimeout 获取请求超时时间
func (rq *Request) GetTimeout() int {
	return rq.timeout
}

//Send 发送请求
func (rq *Request) Send() (*http.Response, error) {
	// now := time.Now()
	req, err := rq.parseBody()
	// fmt.Println("Parse body",time.Since(now))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "gzip")
	// now = time.Now()
	req.Header = parseHeaders(rq.headers)

	rq.client.Timeout = time.Duration(rq.timeout) * time.Millisecond

	httpResponse, err := rq.client.Do(req)

	if err != nil {
		return nil, err
	}
	return httpResponse, nil

}

//QueryParams 获取query参数
func (rq *Request) QueryParams() map[string][]string {
	params := make(map[string][]string)
	for key, values := range rq.queryParams {
		params[key] = values[:]
	}
	return params
}

//URLPath 获取完整的URL路径
func (rq *Request) URLPath() string {
	if len(rq.queryParams) > 0 {
		return rq.URL + "?" + parseParams(rq.queryParams).Encode()
	}
	return rq.URL
}

//SetURL 设置URL
func (rq *Request) SetURL(url string) {
	rq.URL = url
}

//SetRawBody 设置源数据
func (rq *Request) SetRawBody(body []byte) {
	rq.body = body
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
		h.Add("User-Agent", "goku-requests/"+requestVersion)
	}
	return h
}

// 解析请求体
func (rq *Request) parseBody() (req *http.Request, err error) {
	var body io.Reader = nil
	if len(rq.body) > 0 {
		body = bytes.NewBuffer(rq.body)

	}
	req, err = http.NewRequest(rq.method, rq.URLPath(), body)
	return
	//if rq.method == "GET" || rq.method == "TRACE" {
	//	req, err = http.NewRequest(rq.method, rq.URLPath(), nil)
	//}
	//
	//if len(rq.body) > 0 {
	//
	//	if rq.isJSON {
	//		if _, ok := rq.headers["Content-Type"]; !ok {
	//			rq.headers["Content-Type"] = []string{"application/json"}
	//		}
	//		req, err = http.NewRequest(rq.method, rq.URLPath(),
	//			strings.NewReader(string(rq.body)))
	//	} else {
	//		var body *bytes.Buffer
	//		body = bytes.NewBuffer(rq.body)
	//
	//	}
	//} else if len(rq.files) > 0 {
	//	body := new(bytes.Buffer)
	//	writer := multipart.NewWriter(body)
	//	var part io.Writer
	//	for fieldname, file := range rq.files {
	//		part, err = writer.CreateFormFile(fieldname, file.filename)
	//		if err != nil {
	//			return
	//		}
	//		_, err = part.Write(file.data)
	//		if err != nil {
	//			return
	//		}
	//	}
	//	for fieldname, values := range rq.formParams {
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
	//	rq.headers["Content-Type"] = []string{writer.FormDataContentType()}
	//	req, err = http.NewRequest(rq.method, rq.URLPath(), body)
	//} else {
	//	rq.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	//	req, err = http.NewRequest(rq.method, rq.URLPath(),
	//		strings.NewReader(parseParams(rq.formParams).Encode()))
	//}
	//return
}

//parseParams 解析参数
func parseParams(params map[string][]string) url.Values {
	v := url.Values{}
	for key, values := range params {
		for _, value := range values {
			v.Add(key, value)
		}
	}
	return v
}

//parseURL 解析URL
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
