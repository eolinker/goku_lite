package entity

import "strings"

//Cluster 集群配置
type Cluster struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Title     string        `json:"title"`
	Note      string        `json:"note"`
	NodeCount int           `json:"nodeCount"`
	Redis     *ClusterRedis `json:"redis"`
}

//ClusterRedis 集群redis配置
type ClusterRedis struct {
	Mode     string `json:"mode" yaml:"mode"`
	Addrs    string `json:"addrs" yaml:"addrs"`
	DbIndex  int    `json:"dbIndex" yaml:"dbIndex"`
	Masters  string `json:"masters" yaml:"masters"`
	Password string `json:"password" yaml:"password"`
}

//GetMode 获取redis模式
func (c ClusterRedis) GetMode() string {
	return c.Mode
}

//GetAddrs 获取地址
func (c ClusterRedis) GetAddrs() []string {
	return strings.Split(c.Addrs, ",")
}

//GetMasters 或者masters
func (c ClusterRedis) GetMasters() []string {
	return strings.Split(c.Masters, ",")
}

//GetDbIndex 获取所有
func (c ClusterRedis) GetDbIndex() int {
	return c.DbIndex
}

//GetPassword 获取密码
func (c ClusterRedis) GetPassword() string {
	return c.Password
}
