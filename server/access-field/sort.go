package accessfield

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

//Size 返回size
func Size() int {
	return size
}

//All all
func All() []AccessFieldKey {
	return all
}
