package alert

import (
	"strconv"

	log "github.com/eolinker/goku-api-gateway/goku-log"
)

// 记录告警日志
func AlertLog(requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList string, responseStatus int, host, strategyID, strategyName, requestID string) {

	fields := log.Fields{
		"request_id":          requestID,
		"strategy_name":       strategyName,
		"strategy_id":         strategyID,
		"node_host":           host,
		"request_method":      requestMethod,
		"request_url":         requestURL,
		"target_method":       proxyMethod,
		"target_server":       targetServer,
		"target_url":          targetURL,
		"request_query":       queryParamList,
		"request_header":      headerList,
		"request_form_param":  formParamList,
		"response_statusCode": strconv.Itoa(responseStatus),
		"response_header":     responseHeaderList,
	}

	log.WithFields(fields).Warning("alert")

	return
}
