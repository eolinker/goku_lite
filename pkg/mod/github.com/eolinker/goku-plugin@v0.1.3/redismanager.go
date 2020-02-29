package goku_plugin

type RedisManager interface {
	Default() Redis
	Get(name string) (redis Redis, has bool)
}
