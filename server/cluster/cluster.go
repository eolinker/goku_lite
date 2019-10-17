package cluster

import "github.com/eolinker/goku-api-gateway/server/entity"

var (
	byName       map[string]*entity.ClusterInfo
	byID         map[int]*entity.ClusterInfo
	clusterInfos []*entity.ClusterInfo
	nodes        []*entity.Cluster
)

//Init 初始化cluster
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

//Get 通过名字获取集群信息
func Get(name string) (*entity.ClusterInfo, bool) {

	v, has := byName[name]
	return v, has
}

//GetID 通过名字获取集群ID
func GetID(name string) (id int, has bool) {
	v, h := byName[name]
	if !h {
		return 0, false
	}
	return v.ID, true
}
