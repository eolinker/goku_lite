package cluster

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)
var (
	clusterDao dao.ClusterDao
)

func init() {
	pdao.Need(&clusterDao)
}
//AddCluster 新增集群
func AddCluster(name, title, note string) error {
	return clusterDao.AddCluster(name, title, note)
}

//EditCluster 修改集群信息
func EditCluster(name, title, note string) error {
	return clusterDao.EditCluster(name, title, note)
}

//DeleteCluster 删除集群
func DeleteCluster(name string) error {
	return clusterDao.DeleteCluster(name)
}

//GetClusters 获取集群列表
func GetClusters() ([]*entity.Cluster, error) {
	return clusterDao.GetClusters()
}

//GetCluster 获取集群信息
func GetCluster(name string) (*entity.Cluster, error) {
	return clusterDao.GetCluster(name)
}

//GetClusterNodeCount 获取集群节点数量
func GetClusterNodeCount(name string) int {
	return clusterDao.GetClusterNodeCount(name)
}

//CheckClusterNameIsExist 判断集群名称是否存在
func CheckClusterNameIsExist(name string) bool {
	return clusterDao.CheckClusterNameIsExist(name)
}

//GetClusterIDByName 获取集群节点数量
func GetClusterIDByName(name string) int {
	return clusterDao.GetClusterIDByName(name)
}
