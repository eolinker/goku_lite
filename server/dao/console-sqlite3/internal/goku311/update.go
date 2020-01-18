package goku311

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/updater"
)

//Version 版本号
const Version = "3.1.1"

//DBDriver dbDriver
const DBDriver = "sqlite3"

//RegisterUpdate RegisterUpdate
func RegisterUpdate() {
	pdao.RegisterDBBuilder(DBDriver, new(factory))
}

type factory struct {
}

func (f *factory) Build(db *sql.DB) error {
	return Exec(db)
}

//Exec 执行goku_node_info
func Exec(db *sql.DB) error {

	updaterDao := updater.NewUpdaterDaoWidthDB(db)

	existed := updaterDao.IsTableExist("goku_table_version")
	if !existed {
		err := createGokuTableVersion(db)
		if err != nil {

			return err
		}
	}
	if version := updaterDao.GetTableVersion("goku_node_info"); version != Version {
		err := createGokuNodeInfo(db)
		if err != nil {

			return err
		}
		updaterDao.UpdateTableVersion("goku_node_info", Version)
	}
	if version := updaterDao.GetTableVersion("goku_monitor_module"); version != Version {
		err := createGokuMonitorModule(db)
		if err != nil {
			return err
		}
		updaterDao.UpdateTableVersion("goku_monitor_module", Version)
	}

	updaterDao.SetGokuVersion(Version)

	return nil
}
