package admin

import (
	"net/http"
	"strings"
)


//GetIPPort 获取客户端IP和instance
func GetInstanceAndIP(r *http.Request) (string, string, error) {
	ip := r.RemoteAddr
	ip = ip[:strings.Index(ip, ":")]
	if realIP := strings.TrimSpace(r.Header.Get("X-Real-Ip")); realIP != "" {
		ip = realIP
	}
	r.ParseForm()
	instance := r.FormValue("instance")
	return ip, instance, nil
}