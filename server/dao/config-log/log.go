package config_log

import (
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/config-log"
)

const sqlSelect  = "SELECT `name`,`enable`,`dir`,`file`,`level`,`period`,`expire`,`fields` FROM `goku_config_log` WHERE `name` = ? LIMIT 1;"
const sqlInsert  = "INSERT  INTO `goku_config_log`(`name`,`enable`,`dir`,`file`,`level`,`period`,`expire`,`fields`)VALUES(?,?,?,?,?,?,?,?)ON DUPLICATE KEY UPDATE `enable`=VALUES(`enable`),`dir`=VALUES(`dir`),`file`=VALUES(`file`),`level`=VALUES(`level`),`period`=VALUES(`period`),`expire`=VALUES(`expire`),`fields`=VALUES(`fields`);"
func Get(name string) (*entity.LogConfig ,error){
	stmt, e := database.GetConnection().Prepare(sqlSelect)
	if e!=nil{
		return nil,e
	}
	ent:=&entity.LogConfig{}
	err:=stmt.QueryRow(name).Scan(
		&ent.Name,
		&ent.Enable,
		&ent.Dir,
		&ent.File,
		&ent.Level,
		&ent.Period,
		&ent.Expire,
		&ent.Fields,
		)
	if err!=nil{
		return nil,err
	}
	return ent,nil
}

func Set(ent *entity.LogConfig)error{
	stmt, e := database.GetConnection().Prepare(sqlInsert)
	if e!=nil{
		return e
	}
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
