package cluster

import "github.com/eolinker/goku-api-gateway/server/entity"

var (
	byName       map[string]*entity.ClusterInfo
	byID         map[int]*entity.ClusterInfo
	clusterInfos []*entity.ClusterInfo
	nodes        []*entity.Cluster
)

//Init 初始化集群
func Init(cs []*entity.ClusterInfo) {
	nt := make(map[string]*entity.ClusterInfo)
	it := make(map[int]*entity.ClusterInfo)
	ns := make([]*entity.Cluster, 0, len(cs))
	cis := make([]*entity.ClusterInfo, 0, len(cs))
	for _, c := range cs {
		nt[c.Name] = c
		it[c.ID] = c
		ns = append(ns, c.Cluster())
		cis = append(cis, c)
	}

	byName = nt

	byID = it

	clusterInfos = cis
	nodes = ns

}

//Get 根据集群名获取集群信息
func Get(name string) (*entity.ClusterInfo, bool) {

	v, has := byName[name]
	return v, has
}

//GetAll 获取所有集群信息
func GetAll() []*entity.ClusterInfo {
	return clusterInfos
}

//GetList 获取集群列表
func GetList() []*entity.Cluster {
	return nodes
}

//GetID 通过集群名获取集群ID
func GetID(name string) (id int, has bool) {
	v, h := byName[name]
	if !h {
		return 0, false
	}
	return v.ID, true
}

//GetClusterCount 获取集群数量
func GetClusterCount() int {
	return len(nodes)
}

//GetByID 通过ID获取集群信息
func GetByID(id int) (*entity.ClusterInfo, bool) {
	v, h := byID[id]
	return v, h
}
