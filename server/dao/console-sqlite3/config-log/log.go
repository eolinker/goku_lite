package config_log

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/config-log"
)

const sqlSelect = "SELECT `name`,`enable`,`dir`,`file`,`level`,`period`,`expire`,`fields` FROM `goku_config_log` WHERE `name` = ? LIMIT 1;"
const sqlInsert = "REPLACE INTO `goku_config_log`(`name`,`enable`,`dir`,`file`,`level`,`period`,`expire`,`fields`)VALUES(?,?,?,?,?,?,?,?);"

//ConfigLogDao ConfigLogDao
type ConfigLogDao struct {
	db *sql.DB
}

//NewConfigLogDao new ConfigLogDao
func NewConfigLogDao() *ConfigLogDao {
	return &ConfigLogDao{}
}

//Create create
func (d *ConfigLogDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db
	var i dao.ConfigLogDao = d
	return &i, nil
}

//Get get
func (d *ConfigLogDao) Get(name string) (*entity.LogConfig, error) {
	stmt, e := d.db.Prepare(sqlSelect)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	ent := &entity.LogConfig{}
	err := stmt.QueryRow(name).Scan(
		&ent.Name,
		&ent.Enable,
		&ent.Dir,
		&ent.File,
		&ent.Level,
		&ent.Period,
		&ent.Expire,
		&ent.Fields,
	)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

//Set set
func (d *ConfigLogDao) Set(ent *entity.LogConfig) error {
	stmt, e := d.db.Prepare(sqlInsert)
	if e != nil {
		return e
	}
	defer stmt.Close()
	_, err := stmt.Exec(
		ent.Name,
		ent.Enable,
		ent.Dir,
		ent.File,
		ent.Level,
		ent.Period,
		ent.Expire,
		ent.Fields,
	)

	return err
}
