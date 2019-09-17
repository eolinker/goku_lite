package dao_strategy

import (
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

func GetApiPlugin() ([]*entity.StrategyApiPlugin, error) {
	//const sql="SELECT A.`apiID`,A.`strategyID`,A.`pluginName`,A.`pluginConfig` FROM `goku_conn_plugin_api` A INNER JOIN `goku_conn_strategy_api` SA ON A.`strategyID` =  SA.`strategyID` AND  A.`apiID` = SA.`apiID` INNER JOIN `goku_gateway_strategy` S ON S.`enableStatus` =1  AND A.`strategyID` = S.`strategyID` INNER JOIN `goku_gateway_api` API ON API.`apiID` = A.`apiID` INNER JOIN `goku_plugin` P ON P.`isCheck` = TRUE AND P.`pluginStatus` = 1 AND A.`pluginName` = P.`pluginName` ;"
	const sql = "SELECT A.`apiID`,A.`strategyID`,A.`pluginName`,A.`pluginConfig`,IFNULL(A.`updateTag`,'') FROM `goku_conn_plugin_api` A INNER JOIN `goku_conn_strategy_api` SA ON A.`strategyID` =  SA.`strategyID` AND  A.`apiID` = SA.`apiID` INNER JOIN `goku_gateway_strategy` S ON A.`strategyID` = S.`strategyID` INNER JOIN `goku_gateway_api` API ON API.`apiID` = A.`apiID` INNER JOIN `goku_plugin` P ON P.`isCheck` = TRUE AND P.`pluginStatus` = 1 AND A.`pluginStatus` = 1 AND A.`pluginName` = P.`pluginName` ;"

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

	saps := make([]*entity.StrategyApiPlugin, 0, 200)

	for rows.Next() {
		sap := new(entity.StrategyApiPlugin)

		err := rows.Scan(
			&sap.ApiId,
			&sap.StrategyID,
			&sap.PluginName,
			&sap.PluginConfig,
			&sap.UpdateTag,
		)
		if err != nil {
			continue
		}
		saps = append(saps, sap)
	}

	return saps, nil
}
