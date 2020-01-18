package redis_plugin_proxy

import (
	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	goku_plugin "github.com/eolinker/goku-plugin"
)

//Create 创建RedisManager
func Create() goku_plugin.RedisManager {
	if redis_manager.GetConnection() == nil {
		return nil
	}
	return &RedisManager{
		def: &RedisProxy{
			redisClient: redis_manager.GetConnection(),
		},
	}
}

//RedisManager RedisManager
type RedisManager struct {
	def goku_plugin.Redis
}

//Default default
func (m *RedisManager) Default() goku_plugin.Redis {
	return m.def
}

//Get get
func (m *RedisManager) Get(name string) (redis goku_plugin.Redis, has bool) {
	panic("not implement")
}
