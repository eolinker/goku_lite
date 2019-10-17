package access_field

var (
	all = []AccessFieldKey{
		RequestID,
		Msec,
		TimeLocal,
		TimeIso8601,
		Strategy,
		API,
		Host,
		Request,
		Balance,
		FinallyServer,
		Proxy,
		StatusCode,
		ProxyStatusCode,
		RequestTime,
		RemoteAddr,
		HTTPXForwardedFor,
		Retry,
		BodyBytesSent,
		HTTPReferer,
		HTTPUserAgent,
	}
	size = len(all)
)

//Size 返回域数量
func Size() int {
	return size
}

//All 获取所有域key
func All() []AccessFieldKey {
	return all
}
