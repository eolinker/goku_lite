package versionConfig

import (
	"context"
	"encoding/json"
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

func (c *versionConfig) getConfig(cluster string, version string) []byte {
	c.lock.RLock()
	v, ok := c.config[cluster]
	c.lock.RUnlock()

	if ok {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

		r := v.GetWidthContext(ctx, version)
		if r != nil {
			return r.([]byte)
		}
	}
	data, _ := json.Marshal(config.GokuConfig{
		Version: version,
		Cluster: cluster,
	})
	return data
}

//func (c *versionConfig) getVersion() string {
//	c.lock.RLock()
//	defer c.lock.RUnlock()
//	return c.version
//}

//GetVersionConfig 获取版本配置
func GetVersionConfig(cluster, version string) []byte {
	return vc.getConfig(cluster, version)
}

func (c *versionConfig) reset(clusters []*entity.Cluster, gokuConfig *config.GokuConfig, balanceConfig map[string]map[string]*config.BalanceConfig, discoverConfig map[string]map[string]*config.DiscoverConfig) {
	newConfig := make(map[string][]byte)
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
		configByte, _ := json.Marshal(&config.GokuConfig{
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
		})
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
