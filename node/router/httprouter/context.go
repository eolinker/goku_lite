package httprouter

import (
	"context"
	"net/http"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
)

type contextKey struct {
}

//ContextKey contextKey
var ContextKey = contextKey{}

//ContextFromRequest ParamsFromContext pulls the URL parameters from a request context,
// or returns nil if none are present.
//
// This is only present from go 1.7.
func ContextFromRequest(req *http.Request) *common.Context {
	ctx := req.Context()
	p, _ := ctx.Value(ContextKey).(*common.Context)

	return p
}

//SetContextToRequest setContextToRequest
func SetContextToRequest(req *http.Request, ctx *common.Context) *http.Request {
	rctx := req.Context()
	rctx = context.WithValue(rctx, ContextKey, ctx)
	req = req.WithContext(rctx)
	return req
}
