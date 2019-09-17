package access_field

var(
	defaultFields =[]AccessFieldKey{
		RequestId,

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
	}

	defaultSet =make(map[AccessFieldKey]bool)

)

func init() {
	for _,k:=range defaultFields{
		defaultSet[k] = true
	}
}
func IsDefault(k AccessFieldKey) bool {
	_,has:= defaultSet[k]
	return has

}
func Default() []AccessFieldKey {
	return defaultFields

}