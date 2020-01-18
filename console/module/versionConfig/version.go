package versionConfig

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"

	"github.com/eolinker/goku-api-gateway/ksitigarbha"

	"github.com/eolinker/goku-api-gateway/config"
)

var (
	versionDao       dao.VersionDao
	versionConfigDao dao.VersionConfigDao
	clusterDao       dao.ClusterDao
	authNames        = map[string]string{
		"Oauth2": "goku-oauth2_auth",
		"Apikey": "goku-apikey_auth",
		"Basic":  "goku-basic_auth",
		"Jwt":    "goku-jwt_auth",
	}
)

func init() {
	pdao.Need(&versionConfigDao, &versionDao, &clusterDao)
}

//GetVersionList 获取版本列表
func GetVersionList(keyword string) ([]config.VersionConfig, error) {
	return versionDao.GetVersionList(keyword)
}

//AddVersionConfig 新增版本配置
func AddVersionConfig(name, version, remark, now string, userID int) (int, error) {
	config, balanceConfig, discoverConfig := buildVersionConfig(version)
	return versionDao.AddVersionConfig(name, version, remark, config, balanceConfig, discoverConfig, now, userID)
}
func EditVersionBasicConfig(name, version, remark string, userID, versionID int) error {
	return versionDao.EditVersionBasicConfig(name, version, remark, userID, versionID)
}

//BatchDeleteVersionConfig 批量删除版本配置
func BatchDeleteVersionConfig(ids []int) error {
	publishID := versionDao.GetPublishVersionID()
	return versionDao.BatchDeleteVersionConfig(ids, publishID)
}

//PublishVersion 发布版本
func PublishVersion(id, userID int, now string) error {
	err := versionDao.PublishVersion(id, userID, now)
	if err == nil {
		load()
	}
	return err
}

//GetVersionConfigCount 获取版本配置数量
func GetVersionConfigCount() int {
	return versionDao.GetVersionConfigCount()
}

func getRedisConfig(clusters []*entity.Cluster) map[string]interface{} {
	redisConfig := map[string]interface{}{}
	for _, c := range clusters {
		redisConfig[c.Name] = c.Redis
	}
	return redisConfig
}

func buildVersionConfig(v string) (string, string, string) {
	clusters, err := clusterDao.GetClusters()
	if err != nil {
		return "", "", ""
	}
	discoverMap, err := versionConfigDao.GetDiscoverConfig(clusters)
	if err != nil {
		return "", "", ""
	}
	balanceMap, err := versionConfigDao.GetBalances(clusters)
	if err != nil {
		return "", "", ""
	}
	openStrategy, strategyConfigs, err := versionConfigDao.GetStrategyConfig()
	if err != nil {
		return "", "", ""
	}
	apiContents, err := versionConfigDao.GetAPIContent()
	if err != nil {
		return "", "", ""
	}
	plugins, err := versionConfigDao.GetGlobalPlugin()
	if err != nil {
		return "", "", ""
	}
	logCf, accessCf, err := versionConfigDao.GetLogInfo()
	if err != nil {
		return "", "", ""
	}

	g, _ := versionConfigDao.GetGatewayBasicConfig()
	routers, _ := versionConfigDao.GetRouterRules(1)
	ms := make(map[string]string)
	modules, _ := versionConfigDao.GetMonitorModules(1, false)
	if modules != nil {
		for key, config := range modules {
			module, has := ksitigarbha.GetMonitorModuleModel(key)
			if has {
				ms[module.GetName()] = config
			}
		}
	}

	c := config.GokuConfig{
		Version:             v,
		Plugins:             *plugins,
		APIS:                apiContents,
		Strategy:            strategyConfigs,
		AnonymousStrategyID: openStrategy,
		AuthPlugin:          authNames,
		Log:                 logCf,
		AccessLog:           accessCf,
		MonitorModules:      ms,
		Routers:             routers,
		GatewayBasicInfo:    g,
		RedisConfig:         getRedisConfig(clusters),
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
