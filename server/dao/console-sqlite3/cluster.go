package console_sqlite3

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//ClusterDao ClusterDao
type ClusterDao struct {
	db *sql.DB
}

//NewClusterDao new ClusterDao
func NewClusterDao() *ClusterDao {
	return &ClusterDao{}
}

//Create create
func (d *ClusterDao) Create(db *sql.DB) (interface{}, error) {

	d.db = db

	var i dao.ClusterDao = d

	return &i, nil
}

//AddCluster 新增集群
func (d *ClusterDao) AddCluster(name, title, note string) error {
	db := d.db
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
func (d *ClusterDao) EditCluster(name, title, note string) error {
	db := d.db
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
func (d *ClusterDao) DeleteCluster(name string) error {
	db := d.db
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
func (d *ClusterDao) GetClusterCount() int {
	db := d.db
	var count int
	sql := "SELECT COUNT(*) FROM goku_cluster;"
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetClusterNodeCount 获取集群节点数量
func (d *ClusterDao) GetClusterNodeCount(name string) int {
	db := d.db
	var count int
	sql := "SELECT COUNT(*) FROM goku_node_info INNER JOIN goku_cluster ON goku_node_info.clusterID = goku_clutser.id WHERE goku_clutser.`name` = ?;"
	err := db.QueryRow(sql, name).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetClusterIDByName 通过集群名称获取集群ID
func (d *ClusterDao) GetClusterIDByName(name string) int {
	db := d.db
	var id int
	sql := "SELECT `id` FROM goku_cluster WHERE `name` = ?"
	err := db.QueryRow(sql, name).Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

//GetClusterByID 获取集群信息
func (d *ClusterDao) GetClusterByID(id int) (*entity.Cluster, error) {
	db := d.db
	sql := "SELECT `id`,`name`,`title`,`note` FROM goku_cluster WHERE `id` = ?"
	var cluster entity.Cluster
	err := db.QueryRow(sql, id).Scan(&cluster.ID, &cluster.Name, &cluster.Title, &cluster.Note)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

//GetClusters 获取集群列表
func (d *ClusterDao) GetClusters() ([]*entity.Cluster, error) {
	db := d.db
	sql := "SELECT `id`,`name`,`title`,`note`,count(I.`nodeID`) as num FROM `goku_cluster` C left join `goku_node_info` I on c.id = I.`clusterID` group by `id`,`name`,`title`,`note`;"
	rows, err := db.Query(sql)
	if err != nil {
		return []*entity.Cluster{}, err
	}
	clusters := make([]*entity.Cluster, 0, 10)
	defer rows.Close()
	for rows.Next() {
		var cluster entity.Cluster
		err = rows.Scan(&cluster.ID, &cluster.Name, &cluster.Title, &cluster.Note, &cluster.NodeCount)
		if err != nil {
			return []*entity.Cluster{}, err
		}
		clusters = append(clusters, &cluster)
	}
	return clusters, nil
}

//GetCluster 获取集群信息
func (d *ClusterDao) GetCluster(name string) (*entity.Cluster, error) {
	db := d.db
	sql := "SELECT `id`,`name`,`title`,`note` FROM goku_cluster WHERE `name` = ?"
	var cluster entity.Cluster
	err := db.QueryRow(sql, name).Scan(&cluster.ID, &cluster.Name, &cluster.Title, &cluster.Note)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

//CheckClusterNameIsExist 判断集群名称是否存在
func (d *ClusterDao) CheckClusterNameIsExist(name string) bool {
	db := d.db
	sql := "SELECT `name` FROM goku_cluster WHERE `name` = ?"
	var clusterName string
	err := db.QueryRow(sql, name).Scan(&clusterName)
	if err != nil {
		return false
	}
	return true
}
