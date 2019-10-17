package cluster

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddCluster 新增集群
func AddCluster(name, title, note string) error {
	return console_sqlite3.AddCluster(name, title, note)
}

//EditCluster 修改集群信息
func EditCluster(name, title, note string) error {
	return console_sqlite3.EditCluster(name, title, note)
}

//DeleteCluster 删除集群
func DeleteCluster(name string) error {
	return console_sqlite3.DeleteCluster(name)
}

//GetClusters 获取集群列表
func GetClusters() ([]*entity.Cluster, error) {
	return console_sqlite3.GetClusters()
}

//GetCluster 获取集群信息
func GetCluster(name string) (*entity.Cluster, error) {
	return console_sqlite3.GetCluster(name)
}

//GetClusterNodeCount 获取集群节点数量
func GetClusterNodeCount(name string) int {
	return console_sqlite3.GetClusterNodeCount(name)
}

//CheckClusterNameIsExist 判断集群名称是否存在
func CheckClusterNameIsExist(name string) bool {
	return console_sqlite3.CheckClusterNameIsExist(name)
}

//GetClusterIDByName 获取集群节点数量
func GetClusterIDByName(name string) int {
	return console_sqlite3.GetClusterIDByName(name)
}
