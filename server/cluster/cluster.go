package cluster

import "github.com/eolinker/goku/server/entity"

var (
	byName       map[string]*entity.ClusterInfo
	byId         map[int]*entity.ClusterInfo
	clusterInfos []*entity.ClusterInfo
	nodes        []*entity.Cluster
)

func Init(cs []*entity.ClusterInfo) {
	nt := make(map[string]*entity.ClusterInfo)
	it := make(map[int]*entity.ClusterInfo)
	ns := make([]*entity.Cluster, 0, len(cs))
	cis := make([]*entity.ClusterInfo, 0, len(cs))
	for _, c := range cs {
		nt[c.Name] = c
		it[c.Id] = c
		ns = append(ns, c.Cluster())
		cis = append(cis, c)
	}

	byName = nt

	byId = it

	clusterInfos = cis
	nodes = ns

}
func Get(name string) (*entity.ClusterInfo, bool) {

	v, has := byName[name]
	return v, has
}

func GetAll() []*entity.ClusterInfo {
	return clusterInfos
}

func GetList() []*entity.Cluster {
	return nodes
}

func GetId(name string) (id int, has bool) {
	v, h := byName[name]
	if !h {
		return 0, false
	}
	return v.Id, true
}

func GetClusterCount() int {
	return len(nodes)
}

func GetById(id int) (*entity.ClusterInfo, bool) {
	v, h := byId[id]
	return v, h
}
