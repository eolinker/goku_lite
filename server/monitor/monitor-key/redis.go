package monitor_key

import (
	"bytes"
)

func StrategyMapKey(cluster, now string) string {
	key := splicing("monitor-strategy:", cluster, ":", now)
	//fmt.Println("StrategyMapKey:",key)
	return key
}
func APiMapKey(cluster, strategyId, now string) string {
	key := splicing("monitor-api:", cluster, ":", strategyId, ":", now)
	//fmt.Println("APiMapKey:",key)
	return key
}
func ApiValueKey(cluster, strategyId string, apiId string, now string) string {
	key := splicing("monitor-value:", cluster, ":", strategyId, ":", apiId, ":", now)
	//fmt.Println("ApiValueKey:",key)
	return key
}

func splicing(args ...string) string {

	l := 0
	for _, arg := range args {

		l += len(arg)
	}
	buf := make([]byte, 0, l)
	b := bytes.NewBuffer(buf)

	for _, arg := range args {
		b.WriteString(arg)
	}
	return b.String()
}
