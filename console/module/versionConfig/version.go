package versionConfig

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/ksitigarbha"

	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	dao_version_config2 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-version-config"

	"github.com/eolinker/goku-api-gateway/config"
)

var authNames = map[string]string{
	"Oauth2": "goku-oauth2_auth",
	"Apikey": "goku-apikey_auth",
	"Basic":  "goku-basic_auth",
	"Jwt":    "goku-jwt_auth",
}

//GetVersionList 获取版本列表
func GetVersionList(keyword string) ([]config.VersionConfig, error) {
	return console_sqlite3.GetVersionList(keyword)
}

//AddVersionConfig 新增版本配置
func AddVersionConfig(name, version, remark, now string) (int, error) {
	config, balanceConfig, discoverConfig := buildVersionConfig(version)
	return console_sqlite3.AddVersionConfig(name, version, remark, config, balanceConfig, discoverConfig, now)
}

//BatchDeleteVersionConfig 批量删除版本配置
func BatchDeleteVersionConfig(ids []int) error {
	publishID := console_sqlite3.GetPublishVersionID()
	return console_sqlite3.BatchDeleteVersionConfig(ids, publishID)
}

//PublishVersion 发布版本
func PublishVersion(id int, now string) error {
	err := console_sqlite3.PublishVersion(id, now)
	if err == nil {
		load()
	}
	return err
}

//GetVersionConfigCount 获取版本配置数量
func GetVersionConfigCount() int {
	return console_sqlite3.GetVersionConfigCount()
}

func buildVersionConfig(v string) (string, string, string) {
	clusters, err := console_sqlite3.GetClusters()
	if err != nil {
		return "", "", ""
	}
	discoverMap, err := dao_version_config2.GetDiscoverConfig(clusters)
	if err != nil {
		return "", "", ""
	}
	balanceMap, err := dao_version_config2.GetBalances(clusters)
	if err != nil {
		return "", "", ""
	}
	openStrategy, strategyConfigs, err := dao_version_config2.GetStrategyConfig()
	if err != nil {
		return "", "", ""
	}
	apiContents, err := dao_version_config2.GetAPIContent()
	if err != nil {
		return "", "", ""
	}
	plugins, err := dao_version_config2.GetGlobalPlugin()
	if err != nil {
		return "", "", ""
	}
	logCf, accessCf, err := dao_version_config2.GetLogInfo()
	if err != nil {
		return "", "", ""
	}
	ms := make(map[string]string)
	modules, _ := dao_version_config2.GetMonitorModules(1, false)
	if modules != nil {
		for key, config := range modules {
			module ,has:= ksitigarbha.GetMonitorModuleModel(key)
			if has{
				ms[module.GetName()] = config
			}
		}
	}

	c := config.GokuConfig{
		Version:             v,
		Plugins:             plugins,
		APIS:                apiContents,
		Strategy:            strategyConfigs,
		AnonymousStrategyID: openStrategy,
		AuthPlugin:          authNames,
		Log:                 logCf,
		AccessLog:           accessCf,
		MonitorModules:      ms,
	}

	cByte, err := json.Marshal(c)
	if err != nil {
		return "", "", ""
	}
	bByte, err := json.Marshal(balanceMap)
	if err != nil {
		return "", "", ""
	}
	dByte, err := json.Marshal(discoverMap)
	if err != nil {
		return "", "", ""
	}

	return string(cByte), string(bByte), string(dByte)
}
