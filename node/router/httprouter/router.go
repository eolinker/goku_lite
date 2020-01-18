package httprouter

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/module/httprouter"
	"github.com/eolinker/goku-api-gateway/node/router"
)

//HTTPRouter httpRouter
type HTTPRouter struct {
}

//Factory factory
func Factory() router.Factory {
	return &HTTPRouter{}
}

//New new
func (*HTTPRouter) New() router.APIRouter {
	return new(Engine)
}

//Engine engine
type Engine struct {
	router httprouter.Router
}

//AddRouter addRouter
func (r *Engine) AddRouter(method, path string, router router.IRouter) {

	r.router.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
			ctx := ContextFromRequest(req)
			ctx.RestfulParam = make(map[string]string)
			for _, param := range params {

				ctx.RestfulParam[param.Key] = param.Value
			}
			router.Router(ctx)
		})
}

//HandleFunc handleFunc
func (r *Engine) HandleFunc(method, path string, handler router.HandleFunc) {
	r.router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := ContextFromRequest(req)
		handler(ctx)
	})

}

func (r *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request, ctx *common.Context) {

	req = SetContextToRequest(req, ctx)

	r.router.ServeHTTP(w, req)
}

//AddNotFound addNotFound
func (r *Engine) AddNotFound(handler router.HandleFunc) {

	r.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := ContextFromRequest(req)
		handler(ctx)
	})

}
