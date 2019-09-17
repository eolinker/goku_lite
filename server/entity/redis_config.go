package entity

import (
	"strings"
)

func (c CLusterRedis) GetMode() string {
	return c.Mode
}

func (c CLusterRedis) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

func (c CLusterRedis) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

func (c CLusterRedis) GetDbIndex() int {
	return c.DbIndex
}

func (c CLusterRedis) GetPassword() string {
	return c.Password
}
