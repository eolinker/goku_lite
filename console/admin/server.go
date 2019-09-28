package admin

import (
	"net/http"
)

//StartServer 开启admin服务
func StartServer(bind string) error {
	handler := router()
	return http.ListenAndServe(bind, handler)
}
