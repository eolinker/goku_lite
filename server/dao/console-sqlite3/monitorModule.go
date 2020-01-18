package console_sqlite3

import (
	SQL "database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//MonitorModulesDao MonitorModulesDao
type MonitorModulesDao struct {
	db *SQL.DB
}

//NewMonitorModulesDao MonitorModulesDao
func NewMonitorModulesDao() *MonitorModulesDao {
	return &MonitorModulesDao{}
}

//Create create
func (d *MonitorModulesDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.MonitorModulesDao = d
	return &i, nil
}

//GetMonitorModules 获取监控模块列表
func (d *MonitorModulesDao) GetMonitorModules() (map[string]*entity.MonitorModule, error) {
	db := d.db
	sql := "SELECT `name`,IFNULL(`config`,'{}'),`moduleStatus` FROM goku_monitor_module;"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	modules := make(map[string]*entity.MonitorModule)
	for rows.Next() {
		var module entity.MonitorModule
		err = rows.Scan(&module.Name, &module.Config, &module.ModuleStatus)
		if err != nil {
			return nil, err
		}

		modules[module.Name] = &module
	}
	return modules, nil
}

//SetMonitorModule 设置监控模块
func (d *MonitorModulesDao) SetMonitorModule(moduleName string, config string, moduleStatus int) error {
	db := d.db
	sql := "REPLACE INTO goku_monitor_module (`name`,`config`,`moduleStatus`) VALUES (?,?,?)"
	_, err := db.Exec(sql, moduleName, config, moduleStatus)
	if err != nil {
		return err
	}
	return nil
}

//CheckModuleStatus 检查模块状态
func (d *MonitorModulesDao) CheckModuleStatus(moduleName string) int {
	db := d.db
	status := 0
	sql := "SELECT moduleStatus FROM goku_monitor_module WHERE moduleName = ?"
	err := db.QueryRow(sql, moduleName).Scan(&status)
	if err != nil {
		return status
	}
	return status
}
