package accessfield

var (
	defaultFields = []AccessFieldKey{
		RequestID,

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
	}

	defaultSet = make(map[AccessFieldKey]bool)
)

func init() {
	for _, k := range defaultFields {
		defaultSet[k] = true
	}
}

//IsDefault 是否有默认值
func IsDefault(k AccessFieldKey) bool {
	_, has := defaultSet[k]
	return has

}

//Default 获取默认的字段
func Default() []AccessFieldKey {
	return defaultFields

}
