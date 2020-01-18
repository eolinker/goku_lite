package dao_version_config

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/config"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//GetBalances 获取balance信息
func (d *VersionConfigDao)GetBalances(clusters []*entity.Cluster) (map[string]map[string]*config.BalanceConfig, error) {
	db := d.db
	sql := "SELECT goku_balance.balanceName,goku_balance.static,goku_balance.staticCluster,goku_balance.serviceName,goku_balance.appName,goku_service_config.driver FROM goku_balance INNER JOIN goku_service_config ON goku_service_config.`name` = goku_balance.serviceName"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	balanceMaps := make(map[string]map[string]*config.BalanceConfig)
	for rows.Next() {
		var balanceName, static, staticCluster, serviceName, appName, driver string
		err = rows.Scan(&balanceName, &static, &staticCluster, &serviceName, &appName, &driver)
		staticMap := make(map[string]string)
		if staticCluster != "" {
			err := json.Unmarshal([]byte(staticCluster), &staticMap)
			if err != nil {
				return nil, err
			}
		}

		for _, c := range clusters {
			if _, ok := balanceMaps[c.Name]; !ok {
				balanceMaps[c.Name] = make(map[string]*config.BalanceConfig)
			}
			if driver != "static" {
				balanceMaps[c.Name][balanceName] = &config.BalanceConfig{
					Name:         balanceName,
					DiscoverName: serviceName,
					Config:       appName,
				}
				continue
			}
			staticBalance := static
			if v, ok := staticMap[c.Name]; ok {
				staticBalance = v
			}

			balanceMaps[c.Name][balanceName] = &config.BalanceConfig{
				Name:         balanceName,
				DiscoverName: serviceName,
				Config:       staticBalance,
			}
		}

	}
	return balanceMaps, nil
}
