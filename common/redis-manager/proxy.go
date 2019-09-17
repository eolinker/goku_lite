package redis_manager

import (
	"github.com/go-redis/redis"
)

type redisProxy struct {
	redis.Cmdable
	config RedisConfig
}

func (p *redisProxy) Nodes() []string {

	ch := make(chan string, 1)
	switch p.config.GetMode() {
	case RedisModeCluster:
		{
			conn := p.Cmdable.(*redis.ClusterClient)
			go func(ch chan string) {
				conn.ForEachMaster(func(client *redis.Client) error {
					ch <- client.Options().Addr
					return nil
				})
				close(ch)
			}(ch)

		}

	case RedisModeStand:
		{

			conn := p.Cmdable.(*redis.Ring)
			go func(ch chan string) {
				conn.ForEachShard(func(client *redis.Client) error {
					ch <- client.Options().Addr
					return nil
				})
				close(ch)
			}(ch)
		}
	}

	nodes := make([]string, 0, 10)
	for addr := range ch {
		nodes = append(nodes, addr)
	}
	return nodes
}

func (p *redisProxy) Foreach(fn func(client *redis.Client) error) error {
	switch p.config.GetMode() {
	case RedisModeCluster:
		{
			conn := p.Cmdable.(*redis.ClusterClient)
			return conn.ForEachMaster(fn)
		}

	case RedisModeStand:
		{

			conn := p.Cmdable.(*redis.Ring)
			return conn.ForEachShard(fn)
		}
	}

	panic("未知错误")

}

func (p *redisProxy) GetConfig() RedisConfig {
	return p.config
}
