package console

import (
	"github.com/eolinker/goku/server/entity"
	"strings"
)

type ClusterRedisConfig entity.CLusterRedis

func (c ClusterRedisConfig) GetMode() string {
	return c.Mode
}

func (c ClusterRedisConfig) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

func (c ClusterRedisConfig) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

func (c ClusterRedisConfig) GetDbIndex() int {
	return c.DbIndex
}

func (c ClusterRedisConfig) GetPassword() string {
	return c.Password
}
