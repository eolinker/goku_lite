package common

import (
	"net/http"
	"net/url"
)

type RequestReader struct {
	*Header
	*BodyRequestHandler
	req *http.Request
}

func (r *RequestReader) Proto() string {
	return r.req.Proto
}

func NewRequestReader(req *http.Request) *RequestReader {
	r := new(RequestReader)
	r.req = req
	r.ParseRequest()
	return r
}
func (r *RequestReader) ParseRequest() () {

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

func (r *RequestReader) Cookie(name string) (*http.Cookie, error) {
	return r.req.Cookie(name)
}

func (r *RequestReader) Cookies() []*http.Cookie {
	return r.req.Cookies()
}

func (r *RequestReader) Method() string {
	return r.req.Method
}

func (r *RequestReader) URL() *url.URL {
	return r.req.URL
}

func (r *RequestReader) RequestURI() string {
	return r.req.RequestURI
}

func (r *RequestReader) Host() string {
	return r.req.Host
}

func (r *RequestReader) RemoteAddr() string {
	return r.req.RemoteAddr
}
