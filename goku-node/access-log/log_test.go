package access_log

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	"testing"
)

func Test(t *testing.T) {

	dir := "/Users/huangmengzhu/test/log"
	file := "access.log"
	period := log.PeriodDay

	SetFields(access_field.All())
	log.SetOutPut(true, dir, file, period)
	demoCtx := Fields{
		"$remote_addr":          "192.168.0.1",
		"$http_x_forwarded_for": "192.168.0.99",
		//"$remote_user":"",
		"$request":         "\"GET /kingsword\"",
		"$status":          200,
		"$body_bytes_sent": 300,
		"$bytes_sent":      500,
		//"$msec":"日志写入时间。单位为秒，精度是毫秒。",
		//"$http_referer":"",
		"$http_user_agent": "\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36\"",
		"$request_length":  100,
		"$request_time":    100,
		//"$time_iso8601":"ISO8601标准格式下的本地时间。",
		//"$time_local":"通用日志格式下的本地时间。",
		"$requestId":      "xxffsdffadf",
		"$finally_server": "127.0.0.1:8080",
		"$balance":        "Static_Load",
		"$strategy":       "FKdCm2",
		"$api":            "\"1657 kingsword\"",
		"$retry":          "10.1.0.1:80,10.1.0.2:800",
		"$proxy":          "\"POST /proxy HTTPS\"",
		"$proxy_status":   200,
	}

	Log(demoCtx)

	writer.Close()

	t.Log("xx")
}
