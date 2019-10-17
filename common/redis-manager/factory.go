package redis_manager

import (
	"fmt"

	"github.com/go-redis/redis"

	"sort"
)

const (
	_PoolSize = 2000
)

//Create 创建
func Create(config RedisConfig) Redis {

	switch config.GetMode() {
	case RedisModeCluster:
		{
			option := &redis.ClusterOptions{
				Addrs:          config.GetAddrs(),
				Password:       config.GetPassword(),
				PoolSize:       2000,
				ReadOnly:       true,
				RouteByLatency: true,
			}

			return &redisProxy{
				Cmdable: redis.NewClusterClient(option),
				config:  config,
			}
		}

	case RedisModeStand:
		{

			addrs := config.GetAddrs()

			option := redis.RingOptions{
				Addrs:    make(map[string]string),
				Password: config.GetPassword(),
				DB:       config.GetDbIndex(),

				PoolSize: _PoolSize,
			}
			sort.Strings(addrs)
			for i, addr := range addrs {
				option.Addrs[fmt.Sprintf("shad:%d", i)] = addr
			}

			return &redisProxy{
				Cmdable: redis.NewRing(&option),
				config:  config,
			}
		}
	}

	return nil
}
