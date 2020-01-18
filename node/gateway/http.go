package gateway

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/eolinker/goku-api-gateway/diting"
	goku_labels "github.com/eolinker/goku-api-gateway/goku-labels"
	"github.com/eolinker/goku-api-gateway/node/monitor"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_log "github.com/eolinker/goku-api-gateway/goku-node/access-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"github.com/eolinker/goku-api-gateway/node/utils"
	fields "github.com/eolinker/goku-api-gateway/server/access-field"
)

var systemRequestPath = map[string]bool{"/oauth2/token": true, "/oauth2/authorize": true, "/oauth2/verify": true}

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

	delay := time.Since(timeStart)

	ctx.LogFields[fields.RequestID] = requestID
	ctx.LogFields[fields.StatusCode] = status
	ctx.LogFields[fields.HTTPUserAgent] = fmt.Sprint("\"", req.UserAgent(), "\"")
	ctx.LogFields[fields.HTTPReferer] = req.Referer()
	ctx.LogFields[fields.RequestTime] = delay
	ctx.LogFields[fields.Request] = fmt.Sprint("\"", req.Method, " ", req.URL.String(), " ", req.Proto, "\"")
	ctx.LogFields[fields.BodyBytesSent] = n
	ctx.LogFields[fields.Host] = req.Host
	access_log.Log(ctx.LogFields)
	log.WithFields(ctx.LogFields).Info()

	// /oauth2的path不参与统计
	if systemRequestPath[req.URL.Path] {
		return
	}
	// 监控计数
	labels := make(diting.Labels)

	labels[goku_labels.API] = strconv.Itoa(ctx.ApiID())
	labels[goku_labels.Strategy] = ctx.StrategyId()
	labels[goku_labels.Status] = strconv.Itoa(status)
	monitor.APIMonitor.Observe(float64(delay/time.Millisecond), labels)

}
