package console_sqlite3

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddCluster 新增集群
func AddCluster(name, title, note string) error {
	db := database.GetConnection()
	sql := "INSERT INTO goku_cluster (`name`,`title`,`note`) VALUES (?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, title, note)
	return err
}

//EditCluster 修改集群信息
func EditCluster(name, title, note string) error {
	db := database.GetConnection()
	sql := "UPDATE goku_cluster SET `title` = ?,`note` = ? WHERE `name` = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, note, name)
	return err
}

//DeleteCluster 删除集群
func DeleteCluster(name string) error {
	db := database.GetConnection()
	sql := "DELETE FROM goku_cluster WHERE `name` = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	return err
}

//GetClusterCount 获取集群数量
func GetClusterCount() int {
	db := database.GetConnection()
	var count int
	sql := "SELECT COUNT(*) FROM goku_cluster;"
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetClusterNodeCount 获取集群节点数量
func GetClusterNodeCount(name string) int {
	db := database.GetConnection()
	var count int
	sql := "SELECT COUNT(*) FROM goku_node_info INNER JOIN goku_cluster ON goku_node_info.clusterID = goku_clutser.id WHERE goku_clutser.`name` = ?;"
	err := db.QueryRow(sql, name).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetClusterIDByName 通过集群名称获取集群ID
func GetClusterIDByName(name string) int {
	db := database.GetConnection()
	var id int
	sql := "SELECT `id` FROM goku_cluster WHERE `name` = ?"
	err := db.QueryRow(sql, name).Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

//GetClusters 获取集群列表
func GetClusters() ([]*entity.Cluster, error) {
	db := database.GetConnection()
	sql := "SELECT `id`,`name`,`title`,`note` FROM goku_cluster"
	rows, err := db.Query(sql)
	if err != nil {
		return []*entity.Cluster{}, err
	}
	clusters := make([]*entity.Cluster, 0, 10)
	defer rows.Close()
	for rows.Next() {
		var cluster entity.Cluster
		err = rows.Scan(&cluster.ID, &cluster.Name, &cluster.Title, &cluster.Note)
		if err != nil {
			return []*entity.Cluster{}, err
		}
		sql = "SELECT COUNT(*) FROM goku_node_info WHERE clusterID = ?;"
		err = db.QueryRow(sql, cluster.ID).Scan(&cluster.NodeCount)
		if err != nil {
			return []*entity.Cluster{}, err
		}
		clusters = append(clusters, &cluster)
	}
	return clusters, nil
}

//GetCluster 获取集群信息
func GetCluster(name string) (*entity.Cluster, error) {
	db := database.GetConnection()
	sql := "SELECT `id`,`name`,`title`,`note` FROM goku_cluster WHERE `name` = ?"
	var cluster entity.Cluster
	err := db.QueryRow(sql, name).Scan(&cluster.ID, &cluster.Name, &cluster.Title, &cluster.Note)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

//CheckClusterNameIsExist 判断集群名称是否存在
func CheckClusterNameIsExist(name string) bool {
	db := database.GetConnection()
	sql := "SELECT `name` FROM goku_cluster WHERE `name` = ?"
	var clusterName string
	err := db.QueryRow(sql, name).Scan(&clusterName)
	if err != nil {
		return false
	}
	return true
}
