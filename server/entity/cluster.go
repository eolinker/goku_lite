package entity

import "fmt"

//const ClusterDefaultName = "default"

//Cluster cluster
type Cluster struct {
	ID    int    `json:"-" yaml:"-"`
	Name  string `json:"name" yaml:"name"`
	Title string `json:"title" yaml:"title"`
}

//ClusterInfo 集群信息
type ClusterInfo struct {
	ID    int          `json:"-" yaml:"-"`
	Name  string       `json:"name" yaml:"name"`
	Title string       `json:"title" yaml:"title"`
	Note  string       `json:"note" yaml:"note"`
	DB    ClusterDB    `json:"db" yaml:"db"`
	Redis CLusterRedis `json:"redis" yaml:"redis"`
}

//ClusterDB cluster db
type ClusterDB struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	UserName string `json:"userName" yaml:"userName"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

//CLusterRedis cluster redis
type CLusterRedis struct {
	Mode     string `json:"mode" yaml:"mode"`
	Addrs    string `json:"addrs" yaml:"addrs"`
	DbIndex  int    `json:"dbIndex" yaml:"dbIndex"`
	Masters  string `json:"masters" yaml:"masters"`
	Password string `json:"password" yaml:"password"`
}

//GetDriver 获取驱动
func (c *ClusterDB) GetDriver() string {
	return c.Driver
}

//GetSource 获取源字符串
func (c *ClusterDB) GetSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.Database)
}

//Cluster 获取cluster
func (c *ClusterInfo) Cluster() *Cluster {
	return &Cluster{
		ID:    c.ID,
		Name:  c.Name,
		Title: c.Title,
	}

}
