package access_field

import "strings"

//AccessFieldKey access域key
type AccessFieldKey string

const (
	//RemoteAddr 记录客户端IP地址
	RemoteAddr = "$remote_addr"
	//HTTPXForwardedFor HTTP请求端真实IP
	HTTPXForwardedFor = "$http_x_forwarded_for"
	//Request 记录请求的方法、URL和协议（例如 POST /proxy HTTPS)
	Request = "$request"
	//StatusCode 状态码
	StatusCode = "$status_code"
	//BodyBytesSent 发送给客户端的字节数，不包括响应头的大小； 该变量与Apache模块mod_log_config里的“%B”参数兼容。
	BodyBytesSent = "$body_bytes_sent"
	//Msec 日志写入时间。单位为秒，精度是毫秒。
	Msec = "$msec"
	//HTTPReferer 记录从哪个页面链接访问过来的
	HTTPReferer = "$http_referer"
	//HTTPUserAgent 	记录客户端浏览器相关信息
	HTTPUserAgent = "$http_user_agent"
	//RequestTime 请求处理时间，单位为秒，精度毫秒； 从读入客户端的第一个字节开始，直到把最后一个字符发送给客户端后进行日志写入为止。
	RequestTime = "$request_time"
	//TimeIso8601 	ISO8601标准格式下的本地时间。
	TimeIso8601 = "$time_iso8601"
	//TimeLocal 通用日志格式下的本地时间。
	TimeLocal = "$time_local"
	//RequestID 请求id
	RequestID = "$request_id"
	//FinallyServer 最后一次转发的主机信息（IP端口或域名端口）
	FinallyServer = "$finally_server"
	//Balance 	负载信息
	Balance = "$balance"
	//Strategy 策略信息，包括策略名称和ID
	Strategy = "$strategy"
	//API 	API信息，包括 API名称和ID
	API = "$api"
	//Retry 重试信息
	Retry = "$retry"
	//Proxy 记录转发的方法、URL和协议（例如 POST /proxy HTTPS)
	Proxy = "$proxy"
	//ProxyStatusCode 转发状态码
	ProxyStatusCode = "$proxy_status_code"
	//Host 主机信息
	Host = "$host"
)

//Info 获取域信息
func (k AccessFieldKey) Info() string {
	key := strings.ToLower(string(k))
	v, has := infos[key]
	if has {
		return v
	}
	return "unknown"
}

//Key 获取key
func (k AccessFieldKey) Key() string {

	return strings.ToLower(string(k))
}

//Parse 解析
func Parse(key string) AccessFieldKey {
	return AccessFieldKey(strings.ToLower(key))
}

//CopyKey copy
func CopyKey() map[string]string {
	v := make(map[string]string)
	for key, value := range infos {
		v[key] = value
	}
	return v
}

//Has 判断是否存在域key
func Has(key string) bool {
	_, has := infos[AccessFieldKey(key).Key()]
	return has
}
