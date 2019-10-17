package entity

import (
	"strings"
)

//GetMode 获取redis使用模式
func (c CLusterRedis) GetMode() string {
	return c.Mode
}

//GetAddrs 获取地址
func (c CLusterRedis) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

//GetMasters 获取master
func (c CLusterRedis) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

//GetDbIndex 获取数据库序号
func (c CLusterRedis) GetDbIndex() int {
	return c.DbIndex
}

//GetPassword 获取密码
func (c CLusterRedis) GetPassword() string {
	return c.Password
}
