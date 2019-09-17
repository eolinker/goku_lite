package dao_monitor

import (
	"fmt"
	"github.com/eolinker/goku/common/database"
	monitor_key "github.com/eolinker/goku/server/monitor/monitor-key"
	"strconv"
)

var (
	valueField []string

	sqlSave              string
	sqlSaveParameterSize = 0

	sqlSelectGateway          string
	sqlSelectAllGateway       string
	sqlSelectGatewayByHour    string
	sqlSelectGatewayAllByHour string

	sqlSelectStrategy         string
	sqlSelectStrategyWhere    string
	sqlSelectAllStrategy      string
	sqlSelectAllStrategyWhere string

	sqlSelectApi         string
	sqlSelectApiWhere    string
	sqlSelectAllApi      string
	sqlSelectAllApiWhere string

	sqlSelectApiOfStrategy         string
	sqlSelectApiOfStrategyWhere    string
	sqlSelectAllApiOfStrategy      string
	sqlSelectAllApiOfStrategyWhere string

	sqlSelectStrategyOfAPI         string
	sqlSelectStrategyOfAPIWhere    string
	sqlSelectAllStrategyOfAPI      string
	sqlSelectAllStrategyOfAPIWhere string

	sqlSelectStrategyByHour      _HourSqlBuild
	sqlSelectStrategyOfAPIByHour _HourSqlBuild
	sqlSelectApiByHour           _HourSqlBuild
	sqlSelectApiOfStrategyByHour _HourSqlBuild
)

type SCAN interface {
	Scan(dest ...interface{}) error
}

func init() {
	keys := monitor_key.Keys()
	fields := make([]string, len(keys))
	keyFields := make([]string, len(keys))
	for i := range keys {
		fields[i] = fmt.Sprintf("`%s`", keys[i].String())
		keyFields[i] = keys[i].String()
	}
	valueField = keyFields

	sqlSave, sqlSaveParameterSize = genSqlSave(fields)

	sqlSelectGateway = genSqlSelectGateway(fields, false)
	sqlSelectStrategy = genSqlSelectStrategy(fields, false)
	sqlSelectStrategyWhere = genSqlSelectStrategySearch(fields, false)

	sqlSelectAllGateway = genSqlSelectGateway(fields, true)

	sqlSelectApiOfStrategy = genSqlSelectApiOfStrategy(fields, false)
	sqlSelectApiOfStrategyWhere = genSqlSelectApiOfStrategySearch(fields, false)

	sqlSelectApi = genSqlSelectAPI(fields, false)
	sqlSelectApiWhere = genSqlSelectAPISearch(fields, false)

	sqlSelectStrategyOfAPI = genSqlSelectStrategyOfApi(fields, false)
	sqlSelectStrategyOfAPIWhere = genSqlSelectStrategyOfApiSearch(fields, false)

	sqlSelectGatewayByHour = genSqlSelectGatewayByHour(fields, false)
	sqlSelectStrategyByHour = initSqlSelectStrategyByHour(fields)
	sqlSelectStrategyOfAPIByHour = initSqlSelectStrategyOfApiByHour(fields)
	sqlSelectApiByHour = initSqlSelectAPIByHour(fields)
	sqlSelectApiOfStrategyByHour = initSqlSelectAPIOfStrategyByHour(fields)

	sqlSelectGatewayAllByHour = genSqlSelectGatewayByHour(fields, true)
	sqlSelectAllStrategy = genSqlSelectStrategy(fields, true)
	sqlSelectAllStrategyWhere = genSqlSelectStrategySearch(fields, true)
	sqlSelectAllApi = genSqlSelectAPI(fields, true)
	sqlSelectAllApiWhere = genSqlSelectAPISearch(fields, true)

	sqlSelectAllApiOfStrategy = genSqlSelectApiOfStrategy(fields, true)
	sqlSelectAllApiOfStrategyWhere = genSqlSelectApiOfStrategySearch(fields, true)
	sqlSelectAllStrategyOfAPI = genSqlSelectStrategyOfApi(fields, true)
	sqlSelectAllStrategyOfAPIWhere = genSqlSelectStrategyOfApiSearch(fields, true)

}

func Save(strategyId string, apiId int, clusterId int, hour int, now string, values map[string]string) error {

	parameter := make([]interface{}, 0, sqlSaveParameterSize)
	parameter = append(parameter, strategyId, apiId, clusterId, hour)
	for i := range valueField {
		v, has := values[valueField[i]]
		if has {
			vint, err := strconv.Atoi(v)
			if err != nil {
				parameter = append(parameter, 0)
			} else {
				parameter = append(parameter, vint)
			}
		} else {
			parameter = append(parameter, 0)
		}
	}

	parameter = append(parameter, now)
	stmt, e := database.GetConnection().Prepare(sqlSave)
	if e != nil {
		return e
	}
	defer stmt.Close()
	_, err := stmt.Exec(parameter...)
	if err != nil {
		return err
	}

	return nil

}

func GetGateway(clusterId, start, end int) (monitor_key.MonitorValues, error) {
	s := make([]interface{}, 0, 3)
	s = append(s, start, end)
	sql := sqlSelectAllGateway
	if clusterId != 0 {
		sql = sqlSelectGateway
		s = append(s, clusterId)
	}
	stmt, e := database.GetConnection().Prepare(sql)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	row := stmt.QueryRow(s...)
	values, _ := read(row)
	//if err!=nil{
	//	return nil,err
	//}
	return values, nil
}

func GetGatewayInfo() (nodeStartCount, nodeStopCount, projectCount, apiCount, strategyCount int, err error) {
	db := database.GetConnection()
	// 获取节点启动数量

	err = db.QueryRow("SELECT COUNT(0) FROM goku_node_info WHERE nodeStatus = 1;").Scan(&nodeStartCount)
	if err != nil {
		return
	}

	// 获取节点关闭数量

	err = db.QueryRow("SELECT COUNT(0) FROM goku_node_info WHERE nodeStatus = 0;").Scan(&nodeStopCount)
	if err != nil {
		return
	}
	// 获取项目数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_project;").Scan(&projectCount)
	if err != nil {
		return
	}

	// 获取api数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_api;").Scan(&apiCount)
	if err != nil {
		return
	}

	// 获取策略数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_strategy;").Scan(&strategyCount)
	if err != nil {
		return
	}
	return
}

type MonitorValueWidthStrategy struct {
	Id     string
	Name   string
	Value  monitor_key.MonitorValues
	Status int
	Hour   int
}
type MonitorValueWidthAPI struct {
	Id         int
	Name       string
	RequestURL string
	Value      monitor_key.MonitorValues
	Hour       int
}
