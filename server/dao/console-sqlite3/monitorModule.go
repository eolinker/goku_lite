package console_sqlite3

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//GetMonitorModules 获取监控模块列表
func GetMonitorModules() (map[string]*entity.MonitorModule, error) {
	db := database.GetConnection()
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
func SetMonitorModule(moduleName string, config string, moduleStatus int) error {
	db := database.GetConnection()
	sql := "REPLACE INTO goku_monitor_module (`name`,`config`,`moduleStatus`) VALUES (?,?,?)"
	_, err := db.Exec(sql, moduleName, config, moduleStatus)
	if err != nil {
		return err
	}
	return nil
}
