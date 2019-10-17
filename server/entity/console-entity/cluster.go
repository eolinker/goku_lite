package entity

//Cluster 集群配置
type Cluster struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	NodeCount int    `json:"nodeCount"`
}
