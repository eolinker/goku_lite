package access_field

var (
	all = []AccessFieldKey{
		RequestId,
		Msec,
		TimeLocal,
		TimeIso8601,
		Strategy,
		Api,
		Host,
		Request,
		Balance,
		FinallyServer,
		Proxy,
		StatusCode,
		ProxyStatusCode,
		RequestTime,
		RemoteAddr,
		HttpXForwardedFor,
		Retry,
		BodyBytesSent,
		HttpReferer,
		HttpUserAgent,
	}
	size = len(all)
)

func Size() int {
	return size
}
func All() []AccessFieldKey {
	return all
}
