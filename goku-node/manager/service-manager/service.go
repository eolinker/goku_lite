package servicemanager

import (
	"encoding/json"
	log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
	dao_service "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-service"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
	"strings"
)

func init() {
	updater.Add(loadBalanceInfo, 2, "goku_service_config")
}

func loadBalanceInfo() {

	services, err := dao_service.GetAll()
	if err != nil {
		log.Info(err)
	}

	confs := make([]*discovery.Config, 0, len(services))
	for _, s := range services {
		confs = append(confs, toDiscoverConfig(s))
	}
	discovery.ResetAllServiceConfig(confs)
}

func toDiscoverConfig(e *entity.Service) *discovery.Config {
	c := &discovery.Config{
		Name:   e.Name,
		Driver: e.Driver,
		Config: e.Config,
		HealthCheckConfig: discovery.HealthCheckConfig{
			IsHealthCheck: e.HealthCheck,
			Url:           e.HealthCheckPath,
			Second:        e.HealthCheckPeriod,
			TimeOutMill:   e.HealthCheckTimeOut,
			StatusCode:    e.HealthCheckCode,
		},
	}

	if len(e.ClusterConfig) > 0 {
		var clusterConfigObj map[string]string
		err := json.Unmarshal([]byte(e.ClusterConfig), &clusterConfigObj)
		if err == nil {

			clusterName := node_common.ClusterName()
			if v, has := clusterConfigObj[clusterName]; has {
				if len(strings.Trim(v, " ")) > 0 {
					c.Config = v
				}
			}
		}

	}

	return c
}
