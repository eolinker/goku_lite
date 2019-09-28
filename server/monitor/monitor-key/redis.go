package monitorkey

import (
	"bytes"
)

//StrategyMapKey 策略字典Key
func StrategyMapKey(cluster, now string) string {
	key := splicing("monitor-strategy:", cluster, ":", now)
	//fmt.Println("StrategyMapKey:",key)
	return key
}

//APIMapKey 接口字典key
func APIMapKey(cluster, strategyID, now string) string {
	key := splicing("monitor-api:", cluster, ":", strategyID, ":", now)
	//fmt.Println("APIMapKey:",key)
	return key
}

//APIValueKey api value key
func APIValueKey(cluster, strategyID string, apiID string, now string) string {
	key := splicing("monitor-value:", cluster, ":", strategyID, ":", apiID, ":", now)
	//fmt.Println("APIValueKey:",key)
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
