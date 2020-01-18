package gateway

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	v "github.com/eolinker/goku-api-gateway/common/version"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

var (
	gatewayDao dao.GatewayDao
	pluginDao  dao.PluginDao
	clusterDao dao.ClusterDao
)

func init() {
	pdao.Need(&gatewayDao, &pluginDao, &clusterDao)
}

//BaseGatewayInfo 网关基本配置
type BaseGatewayInfo struct {
	NodeCount      int    `json:"nodeCount"`
	NodeStartCount int    `json:"nodeStartCount"`
	NodeStopCount  int    `json:"nodeStopCount"`
	ProjectCount   int    `json:"projectCount"`
	APICount       int    `json:"apiCount"`
	StrategyCount  int    `json:"strategyCount"`
	PluginCount    int    `json:"pluginCount"`
	ClusterCount   int    `json:"clusterCount"`
	Version        string `json:"version"`
}

//SystemInfo 系统配置
type SystemInfo struct {
	BaseInfo BaseGatewayInfo `json:"baseInfo"`
}

//GetGatewayConfig 获取网关配置
func GetGatewayConfig() (map[string]interface{}, error) {
	return gatewayDao.GetGatewayConfig()
}

//EditGatewayBaseConfig 编辑网关基本配置
func EditGatewayBaseConfig(successCode string, nodeUpdatePeriod, monitorUpdatePeriod, timeout int) (bool, string, error) {
	config := entity.GatewayBasicConfig{
		SuccessCode:         successCode,
		NodeUpdatePeriod:    nodeUpdatePeriod,
		MonitorUpdatePeriod: monitorUpdatePeriod,
		MonitorTimeout:      timeout,
	}
	flag, result, err := gatewayDao.EditGatewayBaseConfig(config)
	return flag, result, err
}

//GetGatewayMonitorSummaryByPeriod 获取监控summary
func GetGatewayMonitorSummaryByPeriod() (bool, *SystemInfo, error) {

	nodeStartCount, nodeStopCount, projectCount, apiCount, strategyCount, e := gatewayDao.GetGatewayInfo()
	if e != nil {
		return false, nil, e
	}
	info := new(SystemInfo)
	info.BaseInfo.PluginCount = pluginDao.GetPluginCount()
	info.BaseInfo.NodeCount = nodeStartCount + nodeStopCount
	info.BaseInfo.ProjectCount = projectCount
	info.BaseInfo.APICount = apiCount
	info.BaseInfo.StrategyCount = strategyCount
	info.BaseInfo.Version = v.Version
	info.BaseInfo.ClusterCount = clusterDao.GetClusterCount()

	return true, info, nil

}
