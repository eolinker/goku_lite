package common

import "net/http"

type CookiesHandler struct {
	req http.Request
}

func (cs *CookiesHandler) AddCookie(c *http.Cookie) {
	cs.req.AddCookie(c)
}

func (cs *CookiesHandler) Cookie(name string) (*http.Cookie, error) {
	return cs.req.Cookie(name)
}

func (cs *CookiesHandler) Cookies() []*http.Cookie {
	return cs.req.Cookies()
}

func newCookieHandle(header http.Header) *CookiesHandler {
	return &CookiesHandler{
		req: http.Request{Header: header},
	}
}
