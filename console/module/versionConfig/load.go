package versionConfig

import (
	"context"
	"sync"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"

	"github.com/eolinker/goku-api-gateway/common/telegraph"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"

	"github.com/eolinker/goku-api-gateway/config"
)

type versionConfig struct {
	config map[string]*telegraph.Telegraph
	lock   sync.RWMutex
}

var (
	vc *versionConfig
)

func init() {
	vc = &versionConfig{
		config: make(map[string]*telegraph.Telegraph),
		lock:   sync.RWMutex{},
	}
}

//InitVersionConfig 初始化版本配置
func InitVersionConfig() {
	load()
}
func (c *versionConfig) GetV(cluster string) *telegraph.Telegraph {
	c.lock.RLock()
	v, ok := c.config[cluster]
	c.lock.RUnlock()
	if !ok {
		c.lock.Lock()
		v, ok = c.config[cluster]
		if !ok {
			v = telegraph.NewTelegraph("", nil)
			c.config[cluster] = v
		}
		c.lock.Unlock()
	}
	return v
}
func (c *versionConfig) getConfig(ctx context.Context, cluster string, version string) (*config.GokuConfig, error) {

	v := c.GetV(cluster)

	r, err := v.GetWidthContext(ctx, version)
	if err != nil {
		return nil, err
	}
	return r.(*config.GokuConfig), err

}

//func (c *versionConfig) getVersion() string {
//	c.lock.RLock()
//	defer c.lock.RUnlock()
//	return c.version
//}

//GetVersionConfig 获取版本配置
func GetVersionConfig(ctx context.Context, cluster, version string) (*config.GokuConfig, error) {
	return vc.getConfig(ctx, cluster, version)
}

func (c *versionConfig) reset(clusters []*entity.Cluster, gokuConfig *config.GokuConfig, balanceConfig map[string]map[string]*config.BalanceConfig, discoverConfig map[string]map[string]*config.DiscoverConfig) {
	newConfig := make(map[string]*config.GokuConfig)
	now := time.Now().Format("20060102150405")
	for _, cl := range clusters {
		bf := make(map[string]*config.BalanceConfig)
		if v, ok := balanceConfig[cl.Name]; ok {
			bf = v
		}
		df := make(map[string]*config.DiscoverConfig)
		if v, ok := discoverConfig[cl.Name]; ok {
			df = v
		}
		configByte := &config.GokuConfig{
			Version:             now,
			Cluster:             cl.Name,
			DiscoverConfig:      df,
			Balance:             bf,
			Plugins:             gokuConfig.Plugins,
			APIS:                gokuConfig.APIS,
			Strategy:            gokuConfig.Strategy,
			AuthPlugin:          gokuConfig.AuthPlugin,
			AnonymousStrategyID: gokuConfig.AnonymousStrategyID,
			Log:                 gokuConfig.Log,
			AccessLog:           gokuConfig.AccessLog,
			MonitorModules:      gokuConfig.MonitorModules,
		}
		newConfig[cl.Name] = configByte
	}
	c.lock.Lock()

	for name, cf := range vc.config {
		if _, has := newConfig[name]; !has {
			cf.Close()
			delete(vc.config, name)
		}
	}
	for name, cs := range newConfig {
		cf, has := vc.config[name]
		if !has {
			cf = telegraph.NewTelegraph(now, cs)
			vc.config[name] = cf
		} else {
			cf.Set(now, cs)
		}
	}
	c.lock.Unlock()
}

func load() {
	clusters, err := console_sqlite3.GetClusters()
	if err != nil {
		return
	}
	cf, bf, df, err := console_sqlite3.GetVersionConfig()
	if err != nil {
		log.Warn("load config error:", err)
		return
	}
	vc.reset(clusters, cf, bf, df)
}
