package dao_monitor

import (
	"fmt"
	"strconv"
	"strings"
)

// sava
func genSqlSave(fields []string) (string, int) {

	insertField := make([]string, 0, len(fields)+5)

	insertField = append(insertField, "`strategyID`")
	insertField = append(insertField, "`apiID`")
	insertField = append(insertField, "`clusterID`")
	insertField = append(insertField, "`hour`")

	insertField = append(insertField, fields...)

	insertField = append(insertField, "`updateTime`")
	sqlSaveParameterSize := len(insertField)
	args := make([]string, sqlSaveParameterSize)
	for i := range args {
		args[i] = "?"
	}

	build := strings.Builder{}
	build.WriteString("INSERT INTO `goku_monitor_cluster`(")
	build.WriteString(strings.Join(insertField, ","))
	build.WriteString(")VALUES(")
	build.WriteString(strings.Join(args, ","))
	build.WriteString(") ON DUPLICATE KEY UPDATE ")
	for i := range fields {
		build.WriteString(fields[i])
		build.WriteString(" = ")
		build.WriteString("VALUES(")
		build.WriteString(fields[i])
		build.WriteString("),")
	}
	build.WriteString("`updateTIme` = VALUES(`updateTime`)")
	return build.String(), sqlSaveParameterSize
}

// 全局
func genSqlSelectGateway(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT ")
	build.WriteString(strings.Join(sums, ","))
	build.WriteString(" FROM `goku_monitor_cluster` M")
	if isAll {
		build.WriteString(" WHERE M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString(" WHERE M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}

	return build.String()
}

// // strategy
// func genSqlSelectStrategyCount(isAll bool) string {

// 	build := strings.Builder{}
// 	build.WriteString("SELECT")
// 	build.WriteString("\nCOUNT(*)")
// 	build.WriteString("\nFROM `goku_monitor_cluster` M ")
// 	build.WriteString("\nRIGHT JOIN `goku_gateway_strategy` S ON M.`strategyID` = S.`strategyID`")
// 	if isAll {
// 		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ?")
// 	} else {
// 		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
// 	}
// 	return build.String()
// }

// // strategy
// func genSqlSelectStrategyCountSearch(isAll bool) string {

// 	build := strings.Builder{}
// 	build.WriteString("SELECT")
// 	build.WriteString("\nCOUNT(*)")
// 	build.WriteString("\nFROM `goku_monitor_cluster` M ")
// 	build.WriteString("\nRIGHT JOIN `goku_gateway_strategy` S ON M.`strategyID` = S.`strategyID`")
// 	if isAll {
// 		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ?")
// 	} else {
// 		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
// 	}
// 	build.WriteString("\nWHERE S.`strategyName` LIKE ? OR S.`strategyID` LIKE ?")
// 	return build.String()
// }

// strategy
func genSqlSelectStrategy(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nS.`strategyId` AS `strategyId`,")
	build.WriteString("\nS.`strategyName` AS `strategyName`,")
	build.WriteString("\nS.`enableStatus` AS `enableStatus`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_monitor_cluster` M ")
	build.WriteString("\nRIGHT JOIN `goku_gateway_strategy` S ON M.`strategyID` = S.`strategyID`")
	if isAll {
		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}
	// build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	build.WriteString("\nGROUP by `strategyId`,`strategyName`,`enableStatus`")

	return build.String()
}

// strategy
func genSqlSelectStrategySearch(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nS.`strategyId` AS `strategyId`,")
	build.WriteString("\nS.`strategyName` AS `strategyName`,")
	build.WriteString("\nS.`enableStatus` AS `enableStatus`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_monitor_cluster` M ")
	build.WriteString("\nRIGHT JOIN `goku_gateway_strategy` S ON M.`strategyID` = S.`strategyID`")
	if isAll {
		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}
	// build.WriteString("\nAND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	build.WriteString("\nWHERE S.`strategyName` LIKE ? OR S.`strategyID` LIKE ?")
	build.WriteString("\nGROUP by `strategyId`,`strategyName`,`enableStatus`")

	return build.String()
}

// strategy
func genSqlSelectAPICount(isAll bool) string {

	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nCOUNT(*)")
	build.WriteString("\nFROM  `goku_gateway_api` A")
	build.WriteString("\nLEFT JOIN  `goku_monitor_cluster` M ON M.`apiID` = A.`apiID`")
	if isAll {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}
	// build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	build.WriteString("\nGROUP by `apiID`,`apiName` ,`requestURL`")
	return build.String()
}

// API
func genSqlSelectAPI(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nA.`apiID` AS `apiID`,")
	build.WriteString("\nA.`apiName` AS `apiName`,")
	build.WriteString("\nA.`requestURL` AS `requestURL`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM  `goku_gateway_api` A")
	build.WriteString("\nLEFT JOIN  `goku_monitor_cluster` M ON M.`apiID` = A.`apiID`")
	if isAll {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}
	// build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	build.WriteString("\nGROUP by `apiID`,`apiName` ,`requestURL`")

	return build.String()

}

// API
func genSqlSelectAPISearch(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nA.`apiID` AS `apiID`,")
	build.WriteString("\nA.`apiName` AS `apiName`,")
	build.WriteString("\nA.`requestURL` AS `requestURL`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM  `goku_gateway_api` A")
	build.WriteString("\nLEFT JOIN  `goku_monitor_cluster` M ON M.`apiID` = A.`apiID`")
	if isAll {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}
	// build.WriteString(" AND `hour`>= ? AND `hour` <= ? AND `clusterID` = ?")
	build.WriteString("\nWHERE A.`apiName` like ? OR A.`requestURL` like ? \n")
	build.WriteString("\nGROUP by `apiID`,`apiName` ,`requestURL`")

	return build.String()

}

// 通过apiID 查 策略监控
func genSqlSelectStrategyOfApi(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nS.`strategyId` AS `strategyId`,")
	build.WriteString("\nS.`strategyName` AS `strategyName`,")
	build.WriteString("\nS.`enableStatus` AS `enableStatus`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_conn_strategy_api` SAPI ")
	build.WriteString("\nINNER JOIN `goku_gateway_strategy` S ON S.`strategyID` = SAPI.`strategyID` AND SAPI.`apiID` = ?")
	build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M  ON M.`strategyID` = SAPI.`strategyID`")

	build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ? AND SAPI.`apiID` = M.`apiID`")

	if !isAll {
		build.WriteString(" AND M.`clusterID` = ?")
	}
	build.WriteString(" GROUP by `strategyId`,M.`apiID`,`strategyName`,`enableStatus`")

	return build.String()
} // 通过apiID 查 策略监控
func genSqlSelectStrategyOfApiSearch(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nS.`strategyId` AS `strategyId`,")
	build.WriteString("\nS.`strategyName` AS `strategyName`,")
	build.WriteString("\nS.`enableStatus` AS `enableStatus`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_conn_strategy_api` SAPI ")
	build.WriteString("\nINNER JOIN `goku_gateway_strategy` S ON S.`strategyID` = SAPI.`strategyID` AND SAPI.`apiID` = ?")
	build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M  ON M.`strategyID` = SAPI.`strategyID` AND SAPI.`apiID` = M.`apiID")
	build.WriteString(" AND M.`hour`>= ? AND M.`hour` <= ?`")
	// build.WriteString(" AND `hour`>= ? AND `hour` <= ? AND `clusterID` = ?")
	if !isAll {
		build.WriteString(" AND M.`clusterID` = ?\n")
	}
	build.WriteString("\nWHERE S.`strategyName` LIKE ? OR S.`strategyID` LIKE ?\n")

	build.WriteString(" GROUP by `strategyId`,M.`apiID`,`strategyName`,`enableStatus`")

	return build.String()
}

// 通过strategyId 查 API监控
func genSqlSelectApiOfStrategy(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nA.`apiID` AS `apiID`,")
	build.WriteString("\nA.`apiName` AS `apiName`,")
	build.WriteString("\nA.`requestURL` AS `requestURL`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_conn_strategy_api` SAPI\n")
	build.WriteString("\nINNER JOIN `goku_gateway_api` A ON A.`apiID`=SAPI.`apiID` AND SAPI.`strategyID`= ? \n")
	if isAll {
		build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = SAPI.`apiID` AND  M.`hour`>= ? AND M.`hour` <= ?\n")
	} else {
		build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = SAPI.`apiID` AND  M.`hour`>= ? AND M.`hour` <= ? AND `clusterID` = ?\n")
	}

	build.WriteString("\nGROUP BY `apiID`,`apiName` ,`requestURL`")

	return build.String()
}

// 通过strategyId 查 API监控
func genSqlSelectApiOfStrategySearch(fields []string, isAll bool) string {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nA.`apiID` AS `apiID`,")
	build.WriteString("\nA.`apiName` AS `apiName`,")
	build.WriteString("\nA.`requestURL` AS `requestURL`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ",\n"))
	build.WriteString("\nFROM `goku_conn_strategy_api` SAPI\n")
	build.WriteString("\nINNER JOIN `goku_gateway_api` A ON A.`apiID`=SAPI.`apiID` AND SAPI.`strategyID`= ? \n")
	if isAll {
		build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = SAPI.`apiID`\n")
	} else {
		build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = SAPI.`apiID`  AND  M.`hour`>= ? AND M.`hour` <= ? AND `clusterID` = ?\n")
	}
	// build.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = SAPI.`apiID`  AND  M.`hour`>= ? AND M.`hour` <= ? AND `clusterID` = ?\n")
	build.WriteString("\nWHERE A.`apiName` like ? OR A.`requestURL` like ? \n")

	build.WriteString("\nGROUP BY `apiID`,`apiName` ,`requestURL`")

	return build.String()
}

// 按小时取全局数据
func genSqlSelectGatewayByHour(fields []string, isAll bool) string {
	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])
	}
	build := strings.Builder{}
	build.WriteString("SELECT")
	build.WriteString("\nM.`hour` AS `hour`,")
	build.WriteString("\n")
	build.WriteString(strings.Join(sums, ","))
	build.WriteString("\nFROM `goku_monitor_cluster` M")
	if isAll {
		build.WriteString("\nWHERE M.`hour`>= ? AND M.`hour` <= ?")
	} else {
		build.WriteString("\nWHERE M.`hour`>= ? AND M.`hour` <= ? AND M.`clusterID` = ?")
	}

	build.WriteString("\nGROUP by `hour`")

	return build.String()
}

type _HourSqlBuild struct {
	start string
	end   string
}

func (b *_HourSqlBuild) Build(hours []int) string {

	build := strings.Builder{}
	build.WriteString(b.start)
	build.WriteString("\nJOIN (")
	for i := range hours {
		if i > 0 {
			build.WriteString(" UNION ALL ")
		}
		build.WriteString("SELECT ")
		build.WriteString(strconv.Itoa(hours[i]))
		build.WriteString(" AS `hour`")
	}
	//build.WriteString("select 2019060921 AS `hour` UNION ALL select 2019060920")
	build.WriteString(") T ON 1=1")
	build.WriteString(b.end)

	return build.String()
}

func initSqlSelectStrategyByHour(fields []string) _HourSqlBuild {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	start := strings.Builder{}
	start.WriteString("SELECT")
	start.WriteString("\nT.`hour`,")
	start.WriteString("\nS.`strategyId` AS `strategyId`,")
	start.WriteString("\nS.`strategyName` AS `strategyName`,")
	start.WriteString("\nS.`enableStatus` AS `enableStatus`,")

	start.WriteString("\n")
	start.WriteString(strings.Join(sums, ",\n"))
	start.WriteString("\nFROM `goku_gateway_strategy` S")

	//build.WriteString("\nJOIN (select 2019060921 AS `hour` UNION ALL select 2019060920) T ON 1=1")
	end := strings.Builder{}
	end.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`strategyID` = S.`strategyID`")
	end.WriteString("\nAND T.`hour` = M.`hour` AND M.`clusterID` = ?")
	end.WriteString("\nGROUP by `strategyId`,`strategyName`,`enableStatus`,`hour`")

	return _HourSqlBuild{
		start: start.String(),
		end:   end.String(),
	}
}
func initSqlSelectStrategyOfApiByHour(fields []string) _HourSqlBuild {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}

	start := strings.Builder{}
	start.WriteString("SELECT")
	start.WriteString("\nT.`hour`,")
	start.WriteString("\nS.`strategyId` AS `strategyId`,")
	start.WriteString("\nS.`strategyName` AS `strategyName`,")
	start.WriteString("\nS.`enableStatus` AS `enableStatus`,")
	start.WriteString("\n")
	start.WriteString(strings.Join(sums, ",\n"))
	start.WriteString("\nFROM `goku_conn_strategy_api` SAPI ")
	start.WriteString("\nINNER JOIN `goku_gateway_strategy` S ON S.`strategyID` = SAPI.`strategyID` AND SAPI.`apiID` = ?")
	//build.WriteString("\nJOIN (select 2019060921 AS `hour` UNION ALL select 2019060920) T ON 1=1")
	end := strings.Builder{}
	end.WriteString("\nLEFT JOIN `goku_monitor_cluster` M  ON M.`strategyID` = SAPI.`strategyID`")
	end.WriteString("\nAND T.`hour` = M.`hour` AND `clusterID` = ?")
	end.WriteString("\nGROUP by `strategyId`,`strategyName`,`enableStatus`,`hour`")
	return _HourSqlBuild{
		start: start.String(),
		end:   end.String(),
	}
}

func initSqlSelectAPIByHour(fields []string) _HourSqlBuild {

	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	start := strings.Builder{}
	start.WriteString("SELECT")
	start.WriteString("\nT.`hour`,")
	start.WriteString("\nA.`apiID` AS `apiID`,")
	start.WriteString("\nA.`apiName` AS `apiName`,")
	start.WriteString("\nA.`requestURL` AS `requestURL`,")
	start.WriteString("\n")
	start.WriteString(strings.Join(sums, ",\n"))
	start.WriteString("\nFROM  `goku_gateway_api` A")

	//build.WriteString("\nJOIN (select 2019060921 AS `hour` UNION ALL select 2019060920) T ON 1=1")
	end := strings.Builder{}
	end.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = A.`apiID`")
	end.WriteString("\nAND T.`hour` = M.`hour` AND M.`clusterID` = ?")
	end.WriteString("\nGROUP by `apiID`,`apiName` ,`requestURL`,`hour`")

	return _HourSqlBuild{
		start: start.String(),
		end:   end.String(),
	}
}
func initSqlSelectAPIOfStrategyByHour(fields []string) _HourSqlBuild {
	sums := make([]string, len(fields))

	for i := range fields {
		sums[i] = fmt.Sprintf("IFNULL(SUM(M.%s), 0) AS %s", fields[i], fields[i])

	}
	start := strings.Builder{}
	start.WriteString("SELECT")
	start.WriteString("\nT.`hour`,")
	start.WriteString("\nA.`apiID` AS `apiID`,")
	start.WriteString("\nA.`apiName` AS `apiName`,")
	start.WriteString("\nA.`requestURL` AS `requestURL`,")
	start.WriteString("\n")
	start.WriteString(strings.Join(sums, ",\n"))
	start.WriteString("\nFROM `goku_conn_strategy_api` SAPI ")
	start.WriteString("\nINNER JOIN `goku_gateway_api` A ON A.`apiID` = SAPI.`apiID` AND SAPI.`strategyID` = ?")
	//build.WriteString("\nJOIN (select 2019060921 AS `hour` UNION ALL select 2019060920) T ON 1=1")
	end := strings.Builder{}
	end.WriteString("\nLEFT JOIN `goku_monitor_cluster` M ON M.`apiID` = A.`apiID`")
	end.WriteString("\nAND T.`hour` = M.`hour` AND M.`clusterID` = ?")
	end.WriteString("\nGROUP by `apiID`,`apiName` ,`requestURL`,`hour`")

	return _HourSqlBuild{
		start: start.String(),
		end:   end.String(),
	}
}
