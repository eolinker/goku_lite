package access_field

import "strings"

type AccessFieldKey string

const (
	RemoteAddr        = "$remote_addr"
	HttpXForwardedFor = "$http_x_forwarded_for"
	Request           = "$request"
	StatusCode        = "$status_code"
	BodyBytesSent     = "$body_bytes_sent"
	Msec            = "$msec"
	HttpReferer     = "$http_referer"
	HttpUserAgent   = "$http_user_agent"
	RequestTime     = "$request_time"
	TimeIso8601     = "$time_iso8601"
	TimeLocal       = "$time_local"
	RequestId       = "$request_id"
	FinallyServer   = "$finally_server"
	Balance         = "$balance"
	Strategy        = "$strategy"
	Api             = "$api"
	Retry           = "$retry"
	Proxy           = "$proxy"
	ProxyStatusCode = "$proxy_status_code"
	Host 			= "$host"
)

func (k AccessFieldKey) Info() string {
	key := strings.ToLower(string(k))
	v, has := infos[key]
	if has {
		return v
	}
	return "unknown"
}
func (k AccessFieldKey) Key() string {

	return strings.ToLower(string(k))
}


func Parse(key string) AccessFieldKey {
	return AccessFieldKey(strings.ToLower(key))
}
func CopyKey()map[string]string{
	v:=make(map[string]string)
	for key,value:=range infos{
		v[key]=value
	}
	return v
}
func Has(key string)  bool {
	_,has:=infos[AccessFieldKey(key).Key()]
	return has
}
