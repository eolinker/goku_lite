package monitor

import (
	"fmt"

	monitor_key "github.com/eolinker/goku-api-gateway/server/monitor/monitor-key"
)

type GatewayRequestInfo struct {
	GatewayRequestCount   int64   `json:"gatewayRequestCount"`
	GatewaySuccessCount   int64   `json:"gatewaySuccessCount"`
	GatewayStatus2xxCount int64   `json:"gatewayStatus2xxCount"`
	GatewayStatus4xxCount int64   `json:"gatewayStatus4xxCount"`
	GatewayStatus5xxCount int64   `json:"gatewayStatus5xxCount"`
	GatewaySuccessRate    float64 `json:"-"`
	GatewaySuccessRateStr string  `json:"gatewaySuccessRate"`
}

func (i *GatewayRequestInfo) read(values monitor_key.MonitorValues) {
	i.GatewayRequestCount = values.Get(monitor_key.GatewayRequestCount)
	i.GatewaySuccessCount = values.Get(monitor_key.GatewaySuccessCount)
	i.GatewayStatus2xxCount = values.Get(monitor_key.GatewayStatus2xxCount)
	i.GatewayStatus4xxCount = values.Get(monitor_key.GatewayStatus4xxCount)
	i.GatewayStatus5xxCount = values.Get(monitor_key.GatewayStatus5xxCount)
	if i.GatewayRequestCount != 0 {
		i.GatewaySuccessRate = float64(i.GatewaySuccessCount) / float64(i.GatewayRequestCount)
	}
	i.GatewaySuccessRateStr = fmt.Sprintf("%.2f%%", i.GatewaySuccessRate*100)
}

type ProxyInfo struct {
	ProxyRequestCount   int64   `json:"proxyRequestCount"`
	ProxySuccessCount   int64   `json:"proxySuccessCount"`
	ProxyStatus2xxCount int64   `json:"proxyStatus2xxCount"`
	ProxyStatus4xxCount int64   `json:"proxyStatus4xxCount"`
	ProxyStatus5xxCount int64   `json:"proxyStatus5xxCount"`
	ProxyTimeoutCount   int64   `json:"proxyTimeoutCount"`
	ProxySuccessRate    float64 `json:"-"`
	ProxyTimeoutRate    float64 `json:"-"`
	ProxySuccessRateStr string  `json:"proxySuccessRate"`
	ProxyTimeoutRateStr string  `json:"proxyTimeoutRate"`
}

func (i *ProxyInfo) read(values monitor_key.MonitorValues) {
	i.ProxyRequestCount = values.Get(monitor_key.ProxyRequestCount)
	i.ProxySuccessCount = values.Get(monitor_key.ProxySuccessCount)
	i.ProxyStatus2xxCount = values.Get(monitor_key.ProxyStatus2xxCount)
	i.ProxyStatus4xxCount = values.Get(monitor_key.ProxyStatus4xxCount)
	i.ProxyStatus5xxCount = values.Get(monitor_key.ProxyStatus5xxCount)
	i.ProxyTimeoutCount = values.Get(monitor_key.ProxyTimeoutCount)

	if i.ProxyRequestCount != 0 {
		i.ProxySuccessRate = float64(i.ProxySuccessCount) / float64(i.ProxyRequestCount)
		i.ProxyTimeoutRate = float64(i.ProxyTimeoutCount) / float64(i.ProxyRequestCount)
	}
	i.ProxyTimeoutRateStr = fmt.Sprintf("%.2f%%", i.ProxyTimeoutRate*100)
	i.ProxySuccessRateStr = fmt.Sprintf("%.2f%%", i.ProxySuccessRate*100)
}

type BaseGatewayInfo struct {
	NodeCount      int    `json:"nodeCount"`
	NodeStartCount int    `json:"-"`
	NodeStopCount  int    `json:"-"`
	ProjectCount   int    `json:"projectCount"`
	ApiCount       int    `json:"apiCount"`
	StrategyCount  int    `json:"strategyCount"`
	Version        string `json:"version"`
	ClusterCount   int    `json:"clusterCount"`
	RedisCount     int    `json:"redisCount"`
}
type SystemInfo struct {
	GatewayRequestInfo GatewayRequestInfo `json:"gatewayRequestInfo"`
	ProxyInfo          ProxyInfo          `json:"proxyRequestInfo"`
	BaseInfo           BaseGatewayInfo    `json:"baseInfo"`
}

type Info struct {
	GatewayRequestInfo
	ProxyInfo
}

func (i *Info) Get(key string) interface{} {
	switch key {
	case "gatewayRequestCount":
		return i.GatewayRequestCount
	case "gatewaySuccessCount":
		return i.GatewaySuccessCount
	case "gatewayStatus2xxCount":
		return i.GatewayStatus2xxCount
	case "gatewayStatus4xxCount":
		return i.GatewayStatus4xxCount
	case "gatewayStatus5xxCount":
		return i.GatewayStatus5xxCount
	case "gatewaySuccessRate":
		return i.GatewaySuccessRateStr
	case "proxyRequestCount":
		return i.ProxyRequestCount
	case "proxySuccessCount":
		return i.ProxySuccessCount
	case "proxyStatus2xxCount":
		return i.ProxyStatus2xxCount
	case "proxyStatus4xxCount":
		return i.ProxyStatus4xxCount
	case "proxyStatus5xxCount":
		return i.ProxyStatus5xxCount
	case "proxyTimeoutCount":
		return i.ProxyTimeoutCount
	case "proxySuccessRate":
		return i.ProxySuccessRateStr
	case "proxyTimeoutRate":
		return i.ProxyTimeoutRateStr
	}
	return ""
}
func (i *Info) Value(key string) int64 {
	switch key {
	case "gatewayRequestCount":
		return i.GatewayRequestCount
	case "gatewaySuccessCount":
		return i.GatewaySuccessCount
	case "gatewayStatus2xxCount":
		return i.GatewayStatus2xxCount
	case "gatewayStatus4xxCount":
		return i.GatewayStatus4xxCount
	case "gatewayStatus5xxCount":
		return i.GatewayStatus5xxCount
	case "gatewaySuccessRate":
		return int64(i.GatewaySuccessRate * 10000)
	case "proxyRequestCount":
		return i.ProxyRequestCount
	case "proxySuccessCount":
		return i.ProxySuccessCount
	case "proxyStatus2xxCount":
		return i.ProxyStatus2xxCount
	case "proxyStatus4xxCount":
		return i.ProxyStatus4xxCount
	case "proxyStatus5xxCount":
		return i.ProxyStatus5xxCount
	case "proxyTimeoutCount":
		return i.ProxyTimeoutCount
	case "proxySuccessRate":
		return int64(i.ProxySuccessRate * 10000)
	case "proxyTimeoutRate":
		return int64(i.ProxyTimeoutRate * 10000)
	}
	return 0
}

func (i *Info) read(values monitor_key.MonitorValues) {
	i.GatewayRequestInfo.read(values)
	i.ProxyInfo.read(values)
}

type APIInfo struct {
	Info
	Id   int    `json:"apiID"`
	Name string `json:"apiName"`
	URL  string `json:"requestURL"`
}

func (i *APIInfo) Get(key string) interface{} {
	switch key {
	case "id":
		return i.Id
	case "name":
		return i.Name
	case "url":
		return i.URL
	}
	return i.Info.Get(key)
}

type StrategyInfo struct {
	Info
	Id     string `json:"strategyID"`
	Name   string `json:"strategyName"`
	Status string `json:"-"`
}

func (s *StrategyInfo) Get(key string) interface{} {

	switch key {
	case "id":
		return s.Id
	case "name":
		return s.Name
	case "status":
		return s.Status
	}
	return s.Info.Get(key)

}

type StrategyInfoList []*StrategyInfo

func (l StrategyInfoList) Value(i int, key string) int64 {
	return l[i].Value(key)
}

func (l StrategyInfoList) Len() int {
	return len(l)
}

func (l StrategyInfoList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type APIInfoList []*APIInfo

func (l APIInfoList) Value(i int, key string) int64 {
	return l[i].Value(key)
}

func (l APIInfoList) Len() int {
	return len(l)
}

func (l APIInfoList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
