package redis_manager

import "github.com/eolinker/goku/common/redis"

const (
	RedisModeCluster  = "cluster"
	RedisModeSentinel = "sentinel"
	RedisModeStand    = "stand"
)


type Redis interface {
	redis.Cmdable
	GetConfig() RedisConfig
	//Foreach(fn func(client *localRedis.Client) error) error
	Nodes()[]string
}

type RedisConfig interface {
	GetMode() string
	GetAddrs() []string
	GetMasters() []string
	GetDbIndex() int
	GetPassword() string
}
