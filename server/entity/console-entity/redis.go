package entity

//RedisNode redis节点
type RedisNode struct {
	ID        int
	Server    string
	Password  string
	ClusterID int
}
