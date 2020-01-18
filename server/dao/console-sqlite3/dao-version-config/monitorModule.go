package dao_version_config

import (
	"fmt"
)

//GetMonitorModules 获取监控模块信息
func (d *VersionConfigDao)GetMonitorModules(status int, isAll bool) (map[string]string, error) {
	db := d.db
	sql := "SELECT `name`,`config` FROM goku_monitor_module %s;"
	if isAll {
		sql = fmt.Sprintf(sql, "")
	} else {
		sql = fmt.Sprintf(sql, fmt.Sprintf("WHERE moduleStatus = %d", status))
	}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	modules := make(map[string]string)
	for rows.Next() {
		var name, config string
		err = rows.Scan(&name, &config)
		if err != nil {
			return nil, err
		}
		modules[name] = config
	}
	return modules, nil
}
