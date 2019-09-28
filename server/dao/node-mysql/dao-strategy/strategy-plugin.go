package daostrategy

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

// GetAllStrategyPluginList 获取策略组插件列表
func GetAllStrategyPluginList() ([]*entity.StrategyPluginItem, error) {

	sql := "SELECT  A.`strategyID`,A.`pluginName`,A.`pluginConfig`,IFNULL(A.`updateTag`,'') FROM `goku_conn_plugin_strategy` A INNER JOIN `goku_gateway_strategy` S ON  A.`strategyID` = S.`strategyID` INNER JOIN `goku_plugin` P ON P.`isCheck` = TRUE AND P.`pluginStatus` = 1 AND A.`pluginName` = P.`pluginName` WHERE A.`pluginStatus` = 1;"

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

	strategyList := make([]*entity.StrategyPluginItem, 0, 100)
	for rows.Next() {
		s := new(entity.StrategyPluginItem)
		err := rows.Scan(
			&s.StrategyID,
			&s.PluginName,
			&s.PluginConfig,
			&s.UpdateTag,
		)
		if err != nil {
			continue
		}
		strategyList = append(strategyList, s)
	}
	return strategyList, nil

}
