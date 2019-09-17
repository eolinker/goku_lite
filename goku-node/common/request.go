package common

import "net/url"

// 转发内容
type Request struct {
	*Header
	*CookiesHandler
	*BodyRequestHandler
	querys       url.Values
	targetURL    string
	targetServer string
	Method       string
}

func (r *Request) TargetURL() string {
	return r.targetURL
}
func (r *Request) SetTargetURL(targetURL string) {
	r.targetURL = targetURL
}
func (r *Request) TargetServer() string {
	return r.targetServer
}
func (r *Request) SetTargetServer(targetServer string) {
	r.targetServer = targetServer
}
func (r *Request) Querys() url.Values {
	return r.querys
}

func NewRequest(r *RequestReader) *Request {
	if r == nil {
		return nil
	}
	header := r.Headers()
	return &Request{
		Method:             r.Method(),
		Header:             NewHeader(header),
		CookiesHandler:     newCookieHandle(header),
		BodyRequestHandler: r.BodyRequestHandler.Clone(),
		querys:             r.URL().Query(),
	}
}
