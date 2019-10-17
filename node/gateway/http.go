package gateway

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_log "github.com/eolinker/goku-api-gateway/goku-node/access-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/node/utils"
	fields "github.com/eolinker/goku-api-gateway/server/access-field"
)

var systemRequestPath = []string{"/oauth2/token", "/oauth2/authorize", "/oauth2/verify"}

//HTTPHandler httpHandler
type HTTPHandler struct {
	router *Before
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	timeStart := time.Now()
	// 记录访问次数
	requestID := utils.GetRandomString(16)

	ctx := common.NewContext(req, requestID, w)

	log.Debug(requestID, " url: ", ctx.Request().URL().String())
	log.Debug(requestID, " header: ", ctx.RequestOrg.Header.String())

	if log.GetLogger().Level == log.DebugLevel {
		rawBody, err := ctx.RequestOrg.RawBody()
		if err == nil {
			log.Debug(requestID, " body: ", string(rawBody))
		}
	}
	remoteAddr := utils.Intercept(req.RemoteAddr, ":")
	ctx.LogFields[fields.RemoteAddr] = remoteAddr

	if realIP := ctx.GetHeader("X-Real-Ip"); realIP == "" {
		ctx.ProxyRequest.SetHeader("X-Real-Ip", remoteAddr)
		ctx.LogFields[fields.HTTPXForwardedFor] = remoteAddr
	} else {
		ctx.LogFields[fields.HTTPXForwardedFor] = realIP
	}

	h.router.Router(w, req, ctx)

	n, status := ctx.Finish()

	//proxyStatusCode := 0
	//if ctx.ProxyResponseHandler != nil {
	//	proxyStatusCode = ctx.ProxyResponseHandler.StatusCode()
	//}

	ctx.LogFields[fields.RequestID] = requestID
	ctx.LogFields[fields.StatusCode] = status
	ctx.LogFields[fields.HTTPUserAgent] = fmt.Sprint("\"", req.UserAgent(), "\"")
	ctx.LogFields[fields.HTTPReferer] = req.Referer()
	ctx.LogFields[fields.RequestTime] = time.Since(timeStart)
	ctx.LogFields[fields.Request] = fmt.Sprint("\"", req.Method, " ", req.URL.Path, " ", req.Proto, "\"")
	ctx.LogFields[fields.BodyBytesSent] = n
	ctx.LogFields[fields.Host] = req.Host
	access_log.Log(ctx.LogFields)
	log.WithFields(ctx.LogFields).Info()

	//for _, path := range systemRequestPath {
	//	if path == req.URL.Path {
	//		return
	//	}
	//}

}
