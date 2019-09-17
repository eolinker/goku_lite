package console_mysql

import (
	"encoding/json"
	database2 "github.com/eolinker/goku/common/database"
	"github.com/eolinker/goku/common/general"
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/server/entity/console-entity"
	"io"
	"strings"
	"time"
)

func init() {
	general.RegeditLater(CreateTable)
}

func SaveMemoryInfo(server string, used int, peak int) int64 {
	db := database2.GetConnection()

	stmt, err := db.Prepare("INSERT INTO `goku_redis_memory`(used,peak,server,datetime) VALUES(?,?,?,?)")
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	defer stmt.Close()
	datetime := time.Now().Format("2006-01-02 15:04:05")
	ret, err := stmt.Exec(used, peak, server, datetime)
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	id, err := ret.LastInsertId()
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	return id
}

func SaveInfoCommand(server string, info map[string]interface{}) int64 {

	datetime := time.Now().Format("2006-01-02 15:04:05")
	jsonByte, err := json.Marshal(info)
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	db := database2.GetConnection()
	stmt, err := db.Prepare("INSERT INTO `goku_redis_info`(server,info,datetime) VALUES(?,?,?)")
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	defer stmt.Close()
	ret, err := stmt.Exec(server, string(jsonByte), datetime)
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	id, err := ret.LastInsertId()
	if err != nil {
		log.Info(err.Error())
		return 0
	}
	return id
}

func SaveServer(serverId string, password string, clusterId int, status int) {
	db := database2.GetConnection()

	stmt, err := db.Prepare("INSERT INTO `goku_redis_server`(`server`,`password`,`clusterId`,`status`) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE `password`=VALUES(`password`),`status` =VALUES(`status`)")
	if err != nil {
		log.Info(err.Error())
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(serverId, password, clusterId, status)
	if err != nil {
		log.Info(err.Error())
		return
	}
}
func RemoveServer(serverId string, clusterId int) {
	db := database2.GetConnection()

	stmt, err := db.Prepare("DELETE from  `goku_redis_server` WHERE `server`=? and `clusterID` = ?;")
	if err != nil {
		log.Info(err.Error())
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(serverId, clusterId)
	if err != nil {
		log.Info(err.Error())
		return
	}
}

func GetRedisServerByStatus(status int) ([]*entity.RedisNode, error) {
	sql := "SELECT `server`,`clusterID` FROM `goku_redis_server` WHERE `status` = ?"
	smt, err := database2.GetConnection().Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := smt.Query(status)
	if err != nil {
		return nil, err
	}
	servers := make([]*entity.RedisNode, 0)
	for rows.Next() {
		node := new(entity.RedisNode)
		err := rows.Scan(&node.Server, &node.ClusterId)
		if err != nil {
			return nil, err
		}
		servers = append(servers, node)
	}
	return servers, nil
}
func GetServers(clusterId int) ([]*entity.RedisNode, error) {

	db := database2.GetConnection()
	rows, err := db.Query("SELECT `server`,`password`,`clusterID` from `goku_redis_server` WHERE `clusterID`=?;", clusterId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	servers := make([]*entity.RedisNode, 0)
	for rows.Next() {
		node := new(entity.RedisNode)
		err := rows.Scan(&node.Server, &node.Password, &node.ClusterId)

		if err != nil {
			return nil, err
		}
		servers = append(servers, node)
	}
	return servers, nil
}

func SetRedisNodeStatus(server string, clusterId int, status int) {
	db := database2.GetConnection()
	stmt, err := db.Prepare("UPDATE  `goku_redis_server` SET `status` = ? WHERE `server` = ? AND `clusterID`=?;")
	if err != nil {
		log.Info(err.Error())
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(status, server, clusterId)
	if err != nil {
		log.Info(err.Error())
		return
	}
}

func GetInfo(serverId string) (map[string]interface{}, error) {
	var info string
	db := database2.GetConnection()
	err := db.QueryRow("SELECT `info` FROM `goku_redis_info` WHERE server=? ORDER BY datetime DESC LIMIT 1", serverId).Scan(&info)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	jsonMap := make(map[string]interface{})
	jsonErr := json.Unmarshal([]byte(info), &jsonMap)

	if jsonErr != nil {
		log.Info(jsonErr.Error())
		return nil, jsonErr
	}
	return jsonMap, nil
}

func GetMemoryInfo(serverId, fromDate, toDate string) ([]map[string]interface{}, error) {
	db := database2.GetConnection()
	sql := "SELECT used,peak,datetime FROM `goku_redis_memory` WHERE server=? AND datetime>=? AND datetime<=?"
	rows, err := db.Query(sql, serverId, fromDate, toDate)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	defer rows.Close()
	var ret []map[string]interface{}
	for rows.Next() {
		var (
			used     string
			peak     string
			datetime string
		)
		if err := rows.Scan(&used, &peak, &datetime); err != nil {
			if err != io.EOF {
				log.Info(err.Error())
			}
			continue
		}
		ret = append(ret, map[string]interface{}{"used": used, "peak": peak, "datetime": datetime})
	}
	return ret, nil
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
