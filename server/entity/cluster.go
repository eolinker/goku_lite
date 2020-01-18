package entity

import (
	"fmt"
)

//ClusterDB 集群DB配置
type ClusterDB struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	UserName string `json:"userName" yaml:"userName"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Path     string `json:"path" yaml:"path"`
}

//GetDriver 获取驱动名称
func (c *ClusterDB) GetDriver() string {
	return c.Driver
}

//GetSource 获取连接字符串
func (c *ClusterDB) GetSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.Database)
}
