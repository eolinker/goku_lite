package console

import (
	"github.com/eolinker/goku-api-gateway/server/entity"
	"strings"
)

//ClusterRedisConfig 集群redis配置
type ClusterRedisConfig entity.CLusterRedis

//GetMode 获取redis模式
func (c ClusterRedisConfig) GetMode() string {
	return c.Mode
}

//GetAddrs 获取redis地址
func (c ClusterRedisConfig) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

//GetMasters getMasters
func (c ClusterRedisConfig) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

//GetDbIndex 获取redis数据库序号
func (c ClusterRedisConfig) GetDbIndex() int {
	return c.DbIndex
}

//GetPassword 获取redis密码
func (c ClusterRedisConfig) GetPassword() string {
	return c.Password
}
