package common

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//RequestReader 请求reader
type RequestReader struct {
	*Header
	*BodyRequestHandler
	req *http.Request
}

//Proto 获取协议
func (r *RequestReader) Proto() string {
	return r.req.Proto
}

//NewRequestReader 创建RequestReader
func NewRequestReader(req *http.Request) *RequestReader {
	r := new(RequestReader)
	r.req = req
	r.ParseRequest()
	return r
}

//ParseRequest 解析请求
func (r *RequestReader) ParseRequest() {

	r.Header = NewHeader(r.req.Header)
	body, err := ioutil.ReadAll(r.req.Body)
	_ = r.req.Body.Close()
	if err != nil {
		r.BodyRequestHandler = NewBodyRequestHandler(r.req.Header.Get("Content-Type"), nil)
	} else {
		r.BodyRequestHandler = NewBodyRequestHandler(r.req.Header.Get("Content-Type"), body)
	}
}

//Cookie 获取cookie
func (r *RequestReader) Cookie(name string) (*http.Cookie, error) {
	return r.req.Cookie(name)
}

//Cookies 获取cookies
func (r *RequestReader) Cookies() []*http.Cookie {
	return r.req.Cookies()
}

//Method 获取请求方式
func (r *RequestReader) Method() string {
	return r.req.Method
}

//URL url
func (r *RequestReader) URL() *url.URL {
	return r.req.URL
}

//RequestURI 获取请求URI
func (r *RequestReader) RequestURI() string {
	return r.req.RequestURI
}

//Host 获取host
func (r *RequestReader) Host() string {
	return r.req.Host
}

//RemoteAddr 远程地址
func (r *RequestReader) RemoteAddr() string {
	return r.req.RemoteAddr
}
