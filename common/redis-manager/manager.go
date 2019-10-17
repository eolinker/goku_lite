package redis_manager

import (
	"sync"
)

var (
	redisOfCluster   = make(map[string]Redis)
	redisConfCluster = make(map[string]RedisConfig)

	locker sync.RWMutex
)

//InitRedisOfCluster 初始化集群的redis
func InitRedisOfCluster(rs map[string]RedisConfig) {
	locker.Lock()
	defer locker.Unlock()

	redisConfCluster = rs

}

func get(name string) (Redis, bool) {
	locker.RLock()
	defer locker.RUnlock()

	r, h := redisOfCluster[name]

	return r, h
}

//Get 获取配置
func Get(name string) (Redis, bool) {
	r, has := get(name)

	if has {
		return r, r != nil
	}

	locker.Lock()
	defer locker.Unlock()
	r, has = redisOfCluster[name]
	if has {
		return r, has
	}
	c, h := redisConfCluster[name]
	if h {
		r = Create(c)
		redisOfCluster[name] = r
		return r, h
	}

	redisOfCluster[name] = nil

	return nil, false
}
