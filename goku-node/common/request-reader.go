package common

import (
	"net/http"
	"net/url"
)

//RequestReader 请求header
type RequestReader struct {
	*Header
	*BodyRequestHandler
	req *http.Request
}

//Proto 获取协议
func (r *RequestReader) Proto() string {
	return r.req.Proto
}

//NewRequestReader 获取新请求header
func NewRequestReader(req *http.Request) *RequestReader {
	r := new(RequestReader)
	r.req = req
	r.ParseRequest()
	return r
}

//ParseRequest 解析请求
func (r *RequestReader) ParseRequest() {

	r.Header = NewHeader(r.req.Header)

	body := make([]byte, r.req.ContentLength, r.req.ContentLength)
	i, err := r.req.Body.Read(body)
	_ = r.req.Body.Close()
	if err != nil && int64(i) == r.req.ContentLength {
		r.BodyRequestHandler = NewBodyRequestHandler(r.req.Header.Get("Content-Type"), body)

	} else {
		r.BodyRequestHandler = NewBodyRequestHandler(r.req.Header.Get("Content-Type"), nil)
	}

	// todo
}

//Cookie 获取cookie
func (r *RequestReader) Cookie(name string) (*http.Cookie, error) {
	return r.req.Cookie(name)
}

//Cookies 获取cookies
func (r *RequestReader) Cookies() []*http.Cookie {
	return r.req.Cookies()
}

//Method 获取方法
func (r *RequestReader) Method() string {
	return r.req.Method
}

//URL 获取URL
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

//RemoteAddr 获取客户端地址
func (r *RequestReader) RemoteAddr() string {
	return r.req.RemoteAddr
}
