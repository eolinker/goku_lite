package router

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
)

//IRouter iRouter
type IRouter interface {
	Router(ctx *common.Context)
}

//HandleFunc handlefunc
type HandleFunc func(ctx *common.Context)

//Router router
func (handlerFunc HandleFunc) Router(ctx *common.Context) {
	handlerFunc(ctx)
}

//APIRouter apiRouter
type APIRouter interface {
	AddRouter(method, path string, router IRouter)
	HandleFunc(method, path string, handler HandleFunc)
	AddNotFound(handle HandleFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request, ctx *common.Context)
}

//Factory factory
type Factory interface {
	New() APIRouter
}
