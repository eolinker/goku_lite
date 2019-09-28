package redispluginproxy

import (
	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	"github.com/eolinker/goku-plugin"
)

//Create 创建redisManager
func Create() goku_plugin.RedisManager {

	return &RedisManager{
		def: &RedisProxy{
			redisClient: redis_manager.GetConnection(),
		},
	}

}

//RedisManager redisManager
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
