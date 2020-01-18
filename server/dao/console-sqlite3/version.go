package console_sqlite3

import (
	SQL "database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/eolinker/goku-api-gateway/server/dao"

	"github.com/eolinker/goku-api-gateway/config"
)

//VersionDao VersionDao
type VersionDao struct {
	db *SQL.DB
}

//NewVersionDao new VersionDao
func NewVersionDao() *VersionDao {
	return &VersionDao{}
}

//Create create
func (d *VersionDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.VersionDao = d

	return &i, nil
}

//GetVersionList 获取版本列表
func (d *VersionDao) GetVersionList(keyword string) ([]config.VersionConfig, error) {
	db := d.db
	rule := make([]string, 0, 2)
	if keyword != "" {
		rule = append(rule, "V.name LIKE '%"+keyword+"%' OR V.version LIKE '%"+keyword+"%' OR V.remark LIKE '%"+keyword+"%'")
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}

	sql := "SELECT V.versionID,V.name,V.version,V.remark,V.createTime,V.publishTime,CASE WHEN V.versionID = G.versionID THEN 1 ELSE 0 END AS publishStatus FROM goku_gateway_version_config V LEFT JOIN goku_gateway G ON V.versionID = G.versionID %s ORDER BY publishStatus DESC,V.createTime DESC"
	rows, err := db.Query(fmt.Sprintf(sql, ruleStr))
	if err != nil {
		return make([]config.VersionConfig, 0), err
	}
	defer rows.Close()
	configList := make([]config.VersionConfig, 0, 10)
	for rows.Next() {
		var config config.VersionConfig
		err = rows.Scan(&config.VersionID, &config.Name, &config.Version, &config.Remark, &config.CreateTime, &config.PublishTime, &config.PublishStatus)
		if err != nil {
			return configList, err
		}

		configList = append(configList, config)
	}
	return configList, nil
}

//AddVersionConfig 新增版本配置
func (d *VersionDao) AddVersionConfig(name, version, remark, config, balanceConfig, discoverConfig, now string, userID int) (int, error) {
	db := d.db
	sql := "INSERT INTO goku_gateway_version_config (`name`,`version`,`remark`,`createTime`,`updateTime`,`publishTime`,`config`,`balanceConfig`,`discoverConfig`) VALUES (?,?,?,?,?,?,?,?,?)"
	result, err := db.Exec(sql, name, version, remark, now, now, now, config, balanceConfig, discoverConfig)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), nil
}

//EditVersionBasicConfig 修改版本基本配置
func (d *VersionDao) EditVersionBasicConfig(name, version, remark string, userID, versionID int) error {
	db := d.db
	sql := "UPDATE goku_gateway_version_config SET `name` = ?,`version` = ?,`remark` = ?,`updaterID` = ? WHERE versionID = ?"
	_, err := db.Exec(sql, name, version, remark, userID, versionID)
	if err != nil {
		return err
	}

	return nil
}

//BatchDeleteVersionConfig 批量删除版本配置
func (d *VersionDao) BatchDeleteVersionConfig(ids []int, publishID int) error {
	db := d.db
	s := ""
	idCount := len(ids)
	for i, id := range ids {
		if id == publishID {
			continue
		}
		s = s + strconv.Itoa(id)
		if i < idCount-1 {
			s = s + ","
		}
	}
	sql := fmt.Sprintf("DELETE FROM goku_gateway_version_config WHERE versionID IN (%s)", s)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

//PublishVersion 发布版本
func (d *VersionDao) PublishVersion(id, userID int, now string) error {
	db := d.db
	sql := "UPDATE goku_gateway SET versionID = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	sql = "UPDATE goku_gateway_version_config SET publishTime = ? WHERE versionID = ?"
	_, err = db.Exec(sql, now, id)
	if err != nil {
		return err
	}
	return nil
}

//GetVersionConfigCount 获取版本配置数量
func (d *VersionDao) GetVersionConfigCount() int {
	db := d.db
	sql := "SELECT COUNT(*) FROM goku_gateway_version_config"
	var count int
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

//GetPublishVersionID 获取发布版本ID
func (d *VersionDao) GetPublishVersionID() int {
	db := d.db
	sql := "SELECT versionID FROM goku_gateway"
	var id int
	err := db.QueryRow(sql).Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

//GetVersionConfig 获取当前版本配置
func (d *VersionDao) GetVersionConfig() (*config.GokuConfig, map[string]map[string]*config.BalanceConfig, map[string]map[string]*config.DiscoverConfig, error) {
	db := d.db
	sql := "SELECT IFNULL(goku_gateway_version_config.config,'{}'),IFNULL(goku_gateway_version_config.balanceConfig,'{}'),IFNULL(goku_gateway_version_config.discoverConfig,'{}') FROM goku_gateway_version_config INNER JOIN goku_gateway ON goku_gateway.versionID = goku_gateway_version_config.versionID"
	var cf, bf, df string

	err := db.QueryRow(sql).Scan(&cf, &bf, &df)
	if err != nil {
		return nil, nil, nil, err
	}
	var c config.GokuConfig
	bc := make(map[string]map[string]*config.BalanceConfig)
	dc := make(map[string]map[string]*config.DiscoverConfig)
	err = json.Unmarshal([]byte(cf), &c)
	if cf != "" {
		if err != nil {
			return nil, nil, nil, err
		}
	}
	if bf != "" {
		err = json.Unmarshal([]byte(bf), &bc)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if df != "" {
		err = json.Unmarshal([]byte(df), &dc)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return &c, bc, dc, nil
}
