package redis_manager

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

var (
	def           Redis
	defaultConfig RedisConfig
	defLocker     sync.Locker
)

//SetDefault 设置默认redis
func SetDefault(r Redis) {
	def = r
}

//GetConnection 获取redis连接
func GetConnection() Redis {
	//if def != nil {
	return def
	//}
	//defLocker.Lock()
	//defer defLocker.Unlock()
	//
	//def = Create(defaultConfig)
	//return def
}

//CheckConnection 获取redis连接
func CheckConnection(mode, addrs, password, masters string, dbIndex int) bool {
	switch mode {
	case RedisModeCluster:
		{
			a := strings.Split(addrs, ",")
			option := &redis.ClusterOptions{
				Addrs:          a,
				Password:       password,
				PoolSize:       2000,
				ReadOnly:       true,
				RouteByLatency: true,
			}
			r := redis.NewClusterClient(option)
			defer r.Close()
			if _, err := r.Ping().Result(); err != nil {
				return false
			}
			return true
		}
	case RedisModeSentinel:
		{
			a := strings.Split(addrs, ",")
			option := redis.FailoverOptions{
				SentinelAddrs: a,
				MasterName:    masters,
				Password:      password,
				DB:            dbIndex,
				PoolSize:      _PoolSize,
			}
			r := redis.NewFailoverClient(&option)
			defer r.Close()
			if _, err := r.Ping().Result(); err != nil {
				return false
			}
			return true
		}
	case RedisModeStand:
		{

			a := strings.Split(addrs, ",")

			option := redis.RingOptions{
				Addrs:    make(map[string]string),
				Password: password,
				DB:       dbIndex,

				PoolSize: _PoolSize,
			}
			sort.Strings(a)
			for i, ad := range a {
				option.Addrs[fmt.Sprintf("shad:%d", i)] = ad
			}
			r := redis.NewRing(&option)
			defer r.Close()
			if _, err := r.Ping().Result(); err != nil {
				return false
			}
			return true
		}
	}
	return false

}
