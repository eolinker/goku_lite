package redis_plugin_proxy

import (
	"github.com/eolinker/goku-plugin"
	"github.com/eolinker/goku-api-gateway/common/redis-manager"
)

func Create() goku_plugin.RedisManager {

	return &RedisManager{
		def: &RedisProxy{
			redisClient: redis_manager.GetConnection(),
		},
	}

}

type RedisManager struct {
	def goku_plugin.Redis
}

func (m *RedisManager) Default() goku_plugin.Redis {
	return m.def
}

func (m *RedisManager) Get(name string) (redis goku_plugin.Redis, has bool) {
	panic("not implement")
}
