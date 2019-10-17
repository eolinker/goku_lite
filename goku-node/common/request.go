package common

import "net/url"

//Request 转发内容
type Request struct {
	*Header
	*CookiesHandler
	*BodyRequestHandler
	querys       url.Values
	targetURL    string
	targetServer string
	Method       string
}

//TargetURL 获取转发url
func (r *Request) TargetURL() string {
	return r.targetURL
}

//SetTargetURL 设置转发URL
func (r *Request) SetTargetURL(targetURL string) {
	r.targetURL = targetURL
}

//TargetServer 获取转发服务器地址
func (r *Request) TargetServer() string {
	return r.targetServer
}

//SetTargetServer 设置最终转发地址
func (r *Request) SetTargetServer(targetServer string) {
	r.targetServer = targetServer
}

//Querys 获取query参数
func (r *Request) Querys() url.Values {
	return r.querys
}

//NewRequest 创建请求
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
