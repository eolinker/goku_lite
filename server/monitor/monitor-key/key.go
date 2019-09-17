package monitor_key

type MonitorKeyType int

const (
	GatewayRequestCount MonitorKeyType = iota
	GatewaySuccessCount
	GatewayStatus2xxCount
	GatewayStatus4xxCount
	GatewayStatus5xxCount
	ProxyRequestCount
	ProxySuccessCount
	ProxyStatus2xxCount
	ProxyStatus4xxCount
	ProxyStatus5xxCount
	ProxyTimeoutCount

	MonitorKeyTypeSize = int(ProxyTimeoutCount) + 1
)

var (
	keys []MonitorKeyType
)

func init() {
	ks := make([]MonitorKeyType, MonitorKeyTypeSize)

	for i := range ks {
		ks[i] = MonitorKeyType(i)
	}
	keys = ks
}
func Keys() []MonitorKeyType {
	return keys
}

func ToString(key int) string {
	return MonitorKeyType(key).String()
}
func (t MonitorKeyType) String() string {
	switch t {
	case GatewayRequestCount:
		return "gatewayRequestCount"
	case GatewaySuccessCount:
		return "gatewaySuccessCount"
	case GatewayStatus2xxCount:
		return "gatewayStatus2xxCount"
	case GatewayStatus4xxCount:
		return "gatewayStatus4xxCount"
	case GatewayStatus5xxCount:
		return "gatewayStatus5xxCount"
	case ProxyRequestCount:
		return "proxyRequestCount"
	case ProxySuccessCount:
		return "proxySuccessCount"
	case ProxyStatus2xxCount:
		return "proxyStatus2xxCount"
	case ProxyStatus4xxCount:
		return "proxyStatus4xxCount"
	case ProxyStatus5xxCount:
		return "proxyStatus5xxCount"
	case ProxyTimeoutCount:
		return "proxyTimeoutCount"
	}
	return ""

}
