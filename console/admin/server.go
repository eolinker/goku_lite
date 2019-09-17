package admin

import (
	"net/http"
)

func StartServer(bind string) error {
	handler := router()
	return http.ListenAndServe(bind, handler)
}
