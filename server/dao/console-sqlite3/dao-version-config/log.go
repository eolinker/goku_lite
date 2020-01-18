package dao_version_config

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/config"
)

//GetLogInfo 获取日志信息
func (d *VersionConfigDao)GetLogInfo() (*config.LogConfig, *config.AccessLogConfig, error) {
	db := d.db
	sql := "SELECT `name`,`enable`,`dir`,`file`,`period`,IFNULL(`level`,''),IFNULL(`fields`,''),`expire` FROM goku_config_log;"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	var logCf *config.LogConfig
	var accessCf *config.AccessLogConfig
	for rows.Next() {
		var name, dir, file, level, fields, period string
		var enable, expire int
		err = rows.Scan(&name, &enable, &dir, &file, &period, &level, &fields, &expire)
		if err != nil {
			return nil, nil, err
		}
		if name == "console" {
			continue
		} else if name == "access" {
			tmp := make([]map[string]interface{}, 0)
			err = json.Unmarshal([]byte(fields), &tmp)
			if err != nil {
				return nil, nil, err
			}
			fields := make([]string, 0)
			for _, t := range tmp {
				fields = append(fields, t["name"].(string))
			}
			accessCf = &config.AccessLogConfig{
				Name:   name,
				Enable: enable,
				Dir:    dir,
				File:   file,
				Period: period,
				Expire: expire,
				Fields: fields,
			}
		} else if name == "node" {
			logCf = &config.LogConfig{
				Name:   name,
				Enable: enable,
				Dir:    dir,
				File:   file,
				Period: period,
				Level:  level,
				Expire: expire,
			}
		}
	}
	return logCf, accessCf, nil
}
