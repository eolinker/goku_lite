package goku311

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/console/updater"
	updaterDao "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/updater"
)

//Version 版本号
const Version = "3.1.1"

type factory struct {
	version string
}

var updaterFactory = &factory{version: Version}

func init() {
	updater.Add(Version, updaterFactory)
}

func (f *factory) GetVersion() string {
	return f.version
}

func (f *factory) UpdateVersion() {
	return
}
func (f *factory) Exec() error {
	err := Exec()
	if err != nil {
		return err
	}
	return nil
}

//Exec 执行goku_node_info
func Exec() error {

	db := database.GetConnection()
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
	}
	if version := updaterDao.GetTableVersion("goku_monitor_module"); version != Version {
		err := createGokuMonitorModule(db)
		if err != nil {

			return err
		}
	}

	updaterDao.UpdateTableVersion("goku_node_info", Version)
	updaterDao.UpdateTableVersion("goku_monitor_module", Version)
	return nil
}
