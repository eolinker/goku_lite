package middleware

import (
	"goku-ce-1.0/dao/cache"
	"github.com/codegangsta/inject"
	"github.com/farseer810/yawf"
	"log"
	"net/http"
)

func CleanupHandler(context yawf.Context, log *log.Logger) {
	defer func() {
		if err := recover(); err != nil {
			if log != nil {
				log.Printf("PANIC: %s\n", err)
			}

			val := context.Get(inject.InterfaceOf((*http.ResponseWriter)(nil)))
			res := val.Interface().(http.ResponseWriter)
			res.WriteHeader(http.StatusInternalServerError)

			res.Write([]byte("500 Internal Server Error"))
		}
		conn := cache.GetConnectionFromContext(context)
		if conn != nil {
			conn.Close()
		}
	}()
	context.Next()
}
