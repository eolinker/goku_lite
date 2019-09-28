package redismanager

import "github.com/go-redis/redis"

const (
	//RedisModeCluster cluster模式
	RedisModeCluster = "cluster"
	//RedisModeStand stand模式
	RedisModeStand = "stand"
)

//Redis redis
type Redis interface {
	redis.Cmdable
	GetConfig() RedisConfig
	//Foreach(fn func(client *localRedis.Client) error) error
	Nodes() []string
}

//RedisConfig redis config
type RedisConfig interface {
	GetMode() string
	GetAddrs() []string
	GetMasters() []string
	GetDbIndex() int
	GetPassword() string
}
