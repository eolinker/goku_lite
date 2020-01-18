package dao_version_config

import (
	"encoding/json"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"

	"github.com/eolinker/goku-api-gateway/config"
)

//GetDiscoverConfig 获取服务发现信息
func (d *VersionConfigDao)GetDiscoverConfig(clusters []*entity.Cluster) (map[string]map[string]*config.DiscoverConfig, error) {
	db := d.db
	sql := "SELECT `name`,`driver`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut` FROM goku_service_config"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	discoverMaps := make(map[string]map[string]*config.DiscoverConfig)
	for rows.Next() {
		var name, discoverConfig, clusterConfig, healthCheckPath, healthCheckCode, driver string
		var healthCheck bool
		var healthCheckPeriod, healthCheckTimeOut int
		err = rows.Scan(&name, &driver, &discoverConfig, &clusterConfig, &healthCheck, &healthCheckPath, &healthCheckPeriod, &healthCheckCode, &healthCheckTimeOut)

		configMap := make(map[string]string)
		if clusterConfig != "" {
			err := json.Unmarshal([]byte(clusterConfig), &configMap)
			if err != nil {
				return nil, err
			}
		}

		for _, c := range clusters {
			if _, ok := discoverMaps[c.Name]; !ok {
				discoverMaps[c.Name] = make(map[string]*config.DiscoverConfig)
			}
			if driver == "static" {
				discoverMaps[c.Name][name] = &config.DiscoverConfig{
					Name:   name,
					Driver: driver,
					HealthCheck: &config.HealthCheckConfig{
						IsHealthCheck: healthCheck,
						URL:           healthCheckPath,
						Second:        healthCheckPeriod,
						TimeOutMill:   healthCheckTimeOut,
						StatusCode:    healthCheckCode,
					},
				}
				continue
			}
			defaultConfig := discoverConfig
			if v, ok := configMap[c.Name]; ok {
				defaultConfig = v
			}
			discoverMaps[c.Name][name] = &config.DiscoverConfig{
				Name:   name,
				Driver: driver,
				Config: defaultConfig,
				HealthCheck: &config.HealthCheckConfig{
					IsHealthCheck: healthCheck,
					URL:           healthCheckPath,
					Second:        healthCheckPeriod,
					TimeOutMill:   healthCheckTimeOut,
					StatusCode:    healthCheckCode,
				},
			}
		}
	}
	return discoverMaps, nil
}
