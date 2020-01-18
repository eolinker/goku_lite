package goku_handler

import (
	"context"
	"fmt"
	"net/http"
)

//ContextKey contextKey
type userIDKey struct {
}

var _userIDKey = userIDKey{}

//UserIDFromRequest 从request中读取用户ID
func UserIDFromRequest(req *http.Request) int {
	ctx := req.Context()
	p, ok := ctx.Value(_userIDKey).(int)
	if !ok {
		fmt.Println("error")
	}
	return p
}

//SetUserIDToRequest 设置userID到request
func SetUserIDToRequest(req *http.Request, userID int) *http.Request {
	rctx := req.Context()
	rctx = context.WithValue(rctx, _userIDKey, userID)
	req = req.WithContext(rctx)
	return req
}
