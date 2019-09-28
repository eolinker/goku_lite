package accessfield

import "strings"

//AccessFieldKey access日志域的键名
type AccessFieldKey string

const (
	//RemoteAddr remote_addr
	RemoteAddr        = "$remote_addr"
	//HTTPXForwardedFor http_x_forwarded_for
	HTTPXForwardedFor = "$http_x_forwarded_for"
	//Request request
	Request       = "$request"
	//StatusCode status_code
	StatusCode      = "$status_code"
	//BodyBytesSent body_bytes_sent
	BodyBytesSent   = "$body_bytes_sent"
	//Msec msec
	Msec            = "$msec"
	//HTTPReferer http_referer
	HTTPReferer     = "$http_referer"
	//HTTPUserAgent http_user_agent
	HTTPUserAgent   = "$http_user_agent"
	//RequestTime request_time
	RequestTime     = "$request_time"
	//TimeIso8601 time_iso8601
	TimeIso8601     = "$time_iso8601"
	//TimeLocal time_local
	TimeLocal       = "$time_local"
	//RequestID request_id
	RequestID       = "$request_id"
	//FinallyServer finally_server
	FinallyServer   = "$finally_server"
	//Balance balance
	Balance         = "$balance"
	//Strategy strategy
	Strategy        = "$strategy"
	//API api
	API             = "$api"
	//Retry retry
	Retry           = "$retry"
	//Proxy proxy
	Proxy           = "$proxy"
	//ProxyStatusCode proxy_status_code
	ProxyStatusCode = "$proxy_status_code"
	//Host host
	Host            = "$host"
)

//Info info
func (k AccessFieldKey) Info() string {
	key := strings.ToLower(string(k))
	v, has := infos[key]
	if has {
		return v
	}
	return "unknown"
}

//Key key
func (k AccessFieldKey) Key() string {

	return strings.ToLower(string(k))
}

//Parse parse
func Parse(key string) AccessFieldKey {
	return AccessFieldKey(strings.ToLower(key))
}

//CopyKey copy key
func CopyKey() map[string]string {
	v := make(map[string]string)
	for key, value := range infos {
		v[key] = value
	}
	return v
}

//Has 判断key是否存在
func Has(key string) bool {
	_, has := infos[AccessFieldKey(key).Key()]
	return has
}
