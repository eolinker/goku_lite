package dao_version_config

import (
	"strconv"

	"github.com/eolinker/goku-api-gateway/config"
)

var autoAuthNames = map[string]string{
	"goku-oauth2_auth": "Oauth2",
	"goku-apikey_auth": "Apikey",
	"goku-basic_auth":  "Basic",
	"goku-jwt_auth":    "Jwt",
}

//GetAPIsOfStrategy 获取策略内接口数据
func (d *VersionConfigDao)GetAPIsOfStrategy() (map[string][]*config.APIOfStrategy, error) {
	db := d.db
	sql := "SELECT goku_conn_strategy_api.apiID,IFNULL(goku_conn_strategy_api.target,''),goku_conn_strategy_api.strategyID FROM goku_conn_strategy_api;"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	apiPlugins, err := d.GetAPIPlugins()
	if err != nil {
		return nil, err
	}
	apiMaps := make(map[string][]*config.APIOfStrategy)
	for rows.Next() {
		var apiID int
		var balanceName, strategyID string
		err = rows.Scan(&apiID, &balanceName, &strategyID)
		if err != nil {
			return nil, err
		}
		if _, ok := apiMaps[strategyID]; !ok {
			apiMaps[strategyID] = make([]*config.APIOfStrategy, 0, 20)
		}
		ap := make([]*config.PluginConfig, 0)
		key := strategyID + ":" + strconv.Itoa(apiID)
		if v, ok := apiPlugins[key]; ok {
			ap = v
		}
		apiMaps[strategyID] = append(apiMaps[strategyID], &config.APIOfStrategy{
			ID:      apiID,
			Balance: balanceName,
			Plugins: ap,
		})
	}
	return apiMaps, nil
}

//GetStrategyConfig 获取策略配置
func (d *VersionConfigDao)GetStrategyConfig() (string, []*config.StrategyConfig, error) {
	db := d.db
	sql := "SELECT strategyID,strategyName,enableStatus,strategyType FROM goku_gateway_strategy"

	rows, err := db.Query(sql)
	if err != nil {
		return "", nil, err
	}
	defer rows.Close()
	strategyConfigs := make([]*config.StrategyConfig, 0, 20)
	strategyPlugins, authMaps, err := d.GetStrategyPlugins()
	if err != nil {
		return "", nil, err
	}
	apiOfStrategy, err := d.GetAPIsOfStrategy()
	if err != nil {
		return "", nil, err
	}
	openStrategy := ""
	for rows.Next() {
		var strategyConfig config.StrategyConfig
		var strategyType int
		err = rows.Scan(&strategyConfig.ID, &strategyConfig.Name, &strategyConfig.Enable, &strategyType)
		if err != nil {
			return "", nil, err
		}
		if _, ok := strategyPlugins[strategyConfig.ID]; ok {
			strategyConfig.Plugins = strategyPlugins[strategyConfig.ID]
		}
		if _, ok := authMaps[strategyConfig.ID]; ok {
			strategyConfig.AUTH = authMaps[strategyConfig.ID]
		}
		if _, ok := apiOfStrategy[strategyConfig.ID]; ok {
			strategyConfig.APIS = apiOfStrategy[strategyConfig.ID]
		}
		if strategyType == 1 {
			openStrategy = strategyConfig.ID
		}
		strategyConfigs = append(strategyConfigs, &strategyConfig)
	}
	return openStrategy, strategyConfigs, nil

}
