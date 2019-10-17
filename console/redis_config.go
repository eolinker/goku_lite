package console

import (
	"strings"

	"github.com/eolinker/goku-api-gateway/server/entity"
)

//ClusterRedisConfig 集群redis配置
type ClusterRedisConfig entity.CLusterRedis

//GetMode 获取redis启动模式
func (c ClusterRedisConfig) GetMode() string {
	return c.Mode
}

//GetAddrs 获取redis地址
func (c ClusterRedisConfig) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

//GetMasters 获取master
func (c ClusterRedisConfig) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

//GetDbIndex 获取数据库序号
func (c ClusterRedisConfig) GetDbIndex() int {
	return c.DbIndex
}

//GetPassword 获取密码
func (c ClusterRedisConfig) GetPassword() string {
	return c.Password
}
