package access_field

var (
	infos = map[string]string{
		RemoteAddr:        "记录客户端IP地址",
		HTTPXForwardedFor: "记录客户端IP地址(反向)",
		Request:           "记录请求的方法、URL和协议（例如 POST /proxy HTTPS)",
		StatusCode:        "记录请求状态",
		BodyBytesSent:     "发送给客户端的字节数，不包括响应头的大小； 该变量与Apache模块mod_log_config里的“%B”参数兼容。",
		Msec:              "日志写入时间。单位为秒，精度是毫秒。",
		HTTPReferer:       "记录从哪个页面链接访问过来的",
		HTTPUserAgent:     "记录客户端浏览器相关信息",
		RequestTime:       "请求处理时间，单位为秒，精度毫秒； 从读入客户端的第一个字节开始，直到把最后一个字符发送给客户端后进行日志写入为止。",
		TimeIso8601:       "ISO8601标准格式下的本地时间。",
		TimeLocal:         "通用日志格式下的本地时间。",
		RequestID:         "请求id",
		FinallyServer:     "最后一次转发的主机信息（IP端口或域名端口）",
		Balance:           "负载信息",
		Strategy:          "策略信息，包括策略名称和ID",
		API:               "API信息，包括 API名称和ID",
		Retry:             "重试信息",
		Proxy:             "记录转发的方法、URL和协议（例如 POST /proxy HTTPS)",
		ProxyStatusCode:   "转发状态码",
		Host:              "主机信息",
	}
)
