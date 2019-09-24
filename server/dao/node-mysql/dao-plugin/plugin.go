package dao_plugin

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func GetAll() (map[string]*entity.PluginInfo, error) {

	const sql = "SELECT P.`pluginName`,P.`pluginPriority`,IFNULL(P.`pluginConfig`,''),P.`isStop`,P.`pluginType` FROM `goku_plugin` P WHERE P.`isCheck` = TRUE AND P.`pluginStatus` = 1;"
	stmt, e := database.GetConnection().Prepare(sql)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	plugins := make(map[string]*entity.PluginInfo)
	for rows.Next() {
		p := new(entity.PluginInfo)
		err := rows.Scan(&p.Name, &p.Priority, &p.Config, &p.IsStop, &p.Type)
		if err != nil {
			continue
		}
		plugins[p.Name] = p
	}
	return plugins, nil

}
