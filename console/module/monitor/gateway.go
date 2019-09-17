package monitor

import (
	. "github.com/eolinker/goku/common/version"
	"github.com/eolinker/goku/server/cluster"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
	dao_monitor "github.com/eolinker/goku/server/dao/console-mysql/dao-monitor"
)

func GetGatewayMonitorSummaryByPeriod(clientId int, beginTime, endTime string, period int) (bool, *SystemInfo, error) {

	startHour, endHour := genHour(beginTime, endTime, period)

	values, e := dao_monitor.GetGateway(clientId, startHour, endHour)
	if e != nil {
		return false, nil, e
	}

	nodeStartCount, nodeStopCount, projectCount, apiCount, strategyCount, e := dao_monitor.GetGatewayInfo()
	if e != nil {
		return false, nil, e
	}
	info := new(SystemInfo)
	info.GatewayRequestInfo.read(values)
	info.ProxyInfo.read(values)

	info.BaseInfo.NodeCount = nodeStartCount + nodeStopCount
	info.BaseInfo.ProjectCount = projectCount
	info.BaseInfo.ApiCount = apiCount
	info.BaseInfo.StrategyCount = strategyCount
	activeRedisCount, redisErrorCount := console_mysql.GetRedisCount()
	info.BaseInfo.RedisCount = activeRedisCount + redisErrorCount
	info.BaseInfo.Version = Version
	info.BaseInfo.ClusterCount = cluster.GetClusterCount()
	//dao_monitor.GetGatewayMonitorSummaryByPeriod(beginTime, endTime, period)
	return true, info, nil

}
