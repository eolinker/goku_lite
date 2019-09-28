package consolemysql

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/server/entity"
	jsoniter "github.com/json-iterator/go"
)

//InsertClusters 插入集群信息
func InsertClusters(cs []*entity.ClusterInfo) error {
	db := database.GetConnection()
	stmt, e := db.Prepare("INSERT INTO `goku_cluster`(`name`,`title`,`note`,`db`,`redis`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `title`=VALUES(`title`),`note`=VALUES(`note`),`db`=VALUES(`db`),`redis`=VALUES(`redis`)")
	if e != nil {
		return e
	}
	for _, c := range cs {
		db, _ := jsoniter.MarshalToString(c.DB)
		redis, _ := jsoniter.MarshalToString(c.Redis)
		stmt.Exec(c.Name, c.Title, c.Note, db, redis)
	}

	return nil
}

//LoadClusters 加载集群列表
func LoadClusters() (map[string]*entity.ClusterInfo, error) {
	db := database.GetConnection()
	sql := "select `id`, `name`,`title`,`note`,`db`,`redis` from `goku_cluster`;"
	stmt, e := db.Prepare(sql)
	if e != nil {
		return nil, e
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	vs := make(map[string]*entity.ClusterInfo)
	for rows.Next() {
		v := new(entity.ClusterInfo)
		db := ""
		redis := ""
		err := rows.Scan(&v.ID, &v.Name, &v.Title, &v.Note, &db, &redis)
		if err != nil {
			return nil, nil
		}

		err = jsoniter.UnmarshalFromString(db, &v.DB)
		if err != nil {
			return nil, nil
		}

		err = jsoniter.UnmarshalFromString(redis, &v.Redis)
		if err != nil {
			return nil, nil
		}
		vs[v.Name] = v
	}
	return vs, nil
}
