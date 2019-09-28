package monitorwrite

import (
	"time"

	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	gateway_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/gateway-manager"
	monitor_key "github.com/eolinker/goku-api-gateway/server/monitor/monitor-key"
)

var (
	cluster string
	saveCH  = make(chan *_Action, 100)
	writeCH = make(chan *_MonitorMap, 1)

	writeProid = 5
)

type _Action struct {
	StrategyID string
	APIID      string
	Keys       []monitor_key.MonitorKeyType
}
type _MonitorMap struct {
	now       time.Time
	strategys map[string]*_StrategyInfo
}
type _StrategyInfo struct {
	values map[string]monitor_key.MonitorValues
}

//InitMonitorWrite 初始化监控写入器
func InitMonitorWrite(clusterName string) {
	cluster = clusterName
	go saveLoop()
	go temporaryStorage()

}
func add(action *_Action) {
	saveCH <- action
}
func saveLoop() {
	for {
		select {
		case m := <-writeCH:
			{
				conn := redis_manager.GetConnection()
				pipeline := conn.Pipeline()
				tnow := m.now.Format("2006010215")

				strategyMapKey := monitor_key.StrategyMapKey(cluster, tnow)
				for strategyID, strategy := range m.strategys {
					pipeline.HSetNX(strategyMapKey, strategyID, 1)

					apisOfStrategyKey := monitor_key.APIMapKey(cluster, strategyID, tnow)

					for apiID, values := range strategy.values {

						pipeline.HSetNX(apisOfStrategyKey, apiID, 1)

						key := monitor_key.APIValueKey(cluster, strategyID, apiID, tnow)
						for k, v := range values {
							if v > 0 {
								pipeline.HIncrBy(key, monitor_key.ToString(k), v)
							}
						}
					}
				}
				pipeline.Exec()
				pipeline.Close()
			}
		}
	}
}
func temporaryStorage() {
	t := time.NewTicker(time.Duration(writeProid) * time.Second)
	defer t.Stop()

	ts := &_MonitorMap{
		strategys: make(map[string]*_StrategyInfo),
	}

	for {
		select {
		case <-t.C:

			if len(ts.strategys) > 0 {
				o := ts
				ts = &_MonitorMap{
					strategys: make(map[string]*_StrategyInfo),
				}
				o.now = time.Now()
				writeCH <- o
			}
		case action, ok := <-saveCH:
			if !ok {
				return
			}

			strategy, has := ts.strategys[action.StrategyID]
			if !has {
				strategy = &_StrategyInfo{
					values: make(map[string]monitor_key.MonitorValues),
				}
				ts.strategys[action.StrategyID] = strategy
			}

			apivalue, has := strategy.values[action.APIID]
			if !has {
				apivalue = monitor_key.MakeValue()
				strategy.values[action.APIID] = apivalue
			}

			for _, i := range action.Keys {
				apivalue.Add(i)
			}
		}
	}
}

//AddMonitor 新增监控
func AddMonitor(strategyID string, apiID string, proxyStatusCode int, gatewayStatusCode int) {

	keys := createField(proxyStatusCode, gatewayStatusCode)

	add(&_Action{
		StrategyID: strategyID,
		APIID:      apiID,
		Keys:       keys,
	})
	// proxyStatusCode == 0 没有进行转发

}

func createField(proxyStatusCode int, gatewayStatusCode int) []monitor_key.MonitorKeyType {
	fieldkeys := make([]monitor_key.MonitorKeyType, 0, 7)

	fieldkeys = append(fieldkeys, monitor_key.GatewayRequestCount)
	if proxyStatusCode != 0 {
		fieldkeys = append(fieldkeys, monitor_key.ProxyRequestCount)
		// 超时
		if proxyStatusCode == -1 {
			fieldkeys = append(fieldkeys, monitor_key.ProxyTimeoutCount)
			fieldkeys = append(fieldkeys, monitor_key.ProxyStatus5xxCount)
		} else {

			if proxyStatusCode > 199 && proxyStatusCode < 300 {
				fieldkeys = append(fieldkeys, monitor_key.ProxyStatus2xxCount)
			} else if proxyStatusCode > 399 && proxyStatusCode < 500 {
				fieldkeys = append(fieldkeys, monitor_key.ProxyStatus4xxCount)
			} else if proxyStatusCode > 499 && proxyStatusCode < 600 {
				fieldkeys = append(fieldkeys, monitor_key.ProxyStatus5xxCount)
			}
			if gateway_manager.IsSucess(proxyStatusCode) {
				fieldkeys = append(fieldkeys, monitor_key.ProxySuccessCount)
			}
		}
	}

	if gatewayStatusCode != 0 {

		if gatewayStatusCode > 199 && gatewayStatusCode < 300 {
			fieldkeys = append(fieldkeys, monitor_key.GatewayStatus2xxCount)

		} else if gatewayStatusCode > 399 && gatewayStatusCode < 500 {
			fieldkeys = append(fieldkeys, monitor_key.GatewayStatus4xxCount)

		} else if gatewayStatusCode > 499 && gatewayStatusCode < 600 {
			fieldkeys = append(fieldkeys, monitor_key.GatewayStatus5xxCount)
		}
		if gateway_manager.IsSucess(gatewayStatusCode) {
			fieldkeys = append(fieldkeys, monitor_key.GatewaySuccessCount)
		}
	}
	return fieldkeys
}
