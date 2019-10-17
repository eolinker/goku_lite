package common

import "net/http"

//CookiesHandler cookies处理器
type CookiesHandler struct {
	req http.Request
}

//AddCookie 新增cookiess
func (cs *CookiesHandler) AddCookie(c *http.Cookie) {
	cs.req.AddCookie(c)
}

//Cookie 获取cookie
func (cs *CookiesHandler) Cookie(name string) (*http.Cookie, error) {
	return cs.req.Cookie(name)
}

//Cookies 获取cookies
func (cs *CookiesHandler) Cookies() []*http.Cookie {
	return cs.req.Cookies()
}

func newCookieHandle(header http.Header) *CookiesHandler {
	return &CookiesHandler{
		req: http.Request{Header: header},
	}
}
