package middleware

import (
	"goku-ce/goku"
	"net/http"
	"encoding/json"
	"time"
)

func GetVisitCount(res http.ResponseWriter, req *http.Request,param goku.Params,context *goku.Context) {
	visitCount,_ := json.Marshal(map[string]interface{}{
		"gatewaySuccessCount": context.VisitCount.SuccessCount.GetCount(),
		"gatewayFailureCount": context.VisitCount.FailureCount.GetCount(),
		"gatewayDayCount": context.VisitCount.TotalCount.GetCount(),
		"gatewayMinuteCount": context.VisitCount.CurrentCount.GetCount(),
		"lastUpdateTime":time.Now().Format("15:04:05"),
	})
	res.Write(visitCount)
	return
}