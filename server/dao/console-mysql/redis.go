package consolemysql

import (
	database2 "github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/common/general"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/server/entity/console-entity"
	"strings"
)

func init() {
	general.RegeditLater(CreateTable)
}

//GetServers 获取redis服务列表
func GetServers(clusterID int) ([]*entity.RedisNode, error) {

	db := database2.GetConnection()
	rows, err := db.Query("SELECT `server`,`password`,`clusterID` from `goku_redis_server` WHERE `clusterID`=?;", clusterID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	servers := make([]*entity.RedisNode, 0)
	for rows.Next() {
		node := new(entity.RedisNode)
		err := rows.Scan(&node.Server, &node.Password, &node.ClusterID)

		if err != nil {
			return nil, err
		}
		servers = append(servers, node)
	}
	return servers, nil
}

// GetRedisCount 获取redis数量
func GetRedisCount() (int, int) {
	sql := "SELECT status,COUNT(*) FROM goku_redis_server GROUP BY status;"
	rows, err := database2.GetConnection().Query(sql)
	if err != nil {
		return 0, 0
	}
	var normalCount, errorCount int
	defer rows.Close()
	for rows.Next() {
		var status, count int
		err := rows.Scan(&status, &count)
		if err != nil {
			return 0, 0
		}
		if status == 0 {
			normalCount = count
		} else {
			errorCount = count
		}
	}
	return normalCount, errorCount
}

//CreateTable 创建表
func CreateTable() error {
	sqlDatas := []string{

		`CREATE TABLE  IF NOT EXISTS '''goku_redis_info''' (
  '''id''' int(11) unsigned NOT NULL AUTO_INCREMENT,
  '''server''' varchar(20) NOT NULL COMMENT 'ip:port',
  '''info''' text NOT NULL COMMENT 'info,json',
  '''datetime''' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY ('''id'''),
  KEY '''server''' ('''server''')
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS '''goku_redis_memory''' (
  '''id''' int(11) unsigned NOT NULL AUTO_INCREMENT,
  '''server''' varchar(20) DEFAULT NULL COMMENT 'ip:port',
  '''used''' int(11) DEFAULT NULL,
  '''peak''' int(11) DEFAULT NULL,
  '''datetime''' timestamp NULL DEFAULT NULL,
  PRIMARY KEY ('''id'''),
  KEY '''server''' ('''server''')
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE  IF NOT EXISTS '''goku_redis_server''' (
  '''id''' INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  '''server''' VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'ip:port',
  '''password''' VARCHAR(20) DEFAULT NULL,
  '''clusterID''' INT(11) DEFAULT '1',
  '''status''' INT(11) DEFAULT NULL,
  PRIMARY KEY ('''id'''),
  UNIQUE KEY '''server''' ('''server''','''clusterID'''),
  KEY '''clusterID''' ('''clusterID'''),
  KEY '''status''' ('''status''')
) ENGINE=INNODB AUTO_INCREMENT=224 DEFAULT CHARSET=utf8;`}
	for _, sql := range sqlDatas {
		sqlData := strings.ReplaceAll(sql, "'''", "`")
		db := database2.GetConnection()
		_, err := db.Exec(sqlData)
		if err != nil {
			log.Info(err.Error())
			return err
		}
	}

	return nil
}
