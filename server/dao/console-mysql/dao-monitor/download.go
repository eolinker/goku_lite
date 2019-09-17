package dao_monitor

import (
	"github.com/eolinker/goku/common/database"
)

func GetStrategyByHour(clusterId int, hours []int) ([]*MonitorValueWidthStrategy, error) {

	stmt, e := database.GetConnection().Prepare(sqlSelectStrategyByHour.Build(hours))
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query(clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]*MonitorValueWidthStrategy, 0, 1000)

	for rows.Next() {
		v := new(MonitorValueWidthStrategy)

		values, err := read(rows, &v.Hour, &v.Id, &v.Name, &v.Status)
		if err != nil {
			return nil, err
		}

		v.Value = values

		all = append(all, v)
	}

	return all, nil
}

func GetAPIByHour(clusterId int, hours []int) ([]*MonitorValueWidthAPI, error) {

	stmt, e := database.GetConnection().Prepare(sqlSelectApiByHour.Build(hours))
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query(clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]*MonitorValueWidthAPI, 0, 1000)

	for rows.Next() {
		v := new(MonitorValueWidthAPI)

		values, err := read(rows, &v.Hour, &v.Id, &v.Name, &v.RequestURL)
		if err != nil {
			return nil, err
		}

		v.Value = values

		all = append(all, v)
	}

	return all, nil
}
func GetAPIOfStrategyByHour(clusterId int, strategyID string, hours []int) ([]*MonitorValueWidthAPI, error) {

	stmt, e := database.GetConnection().Prepare(sqlSelectApiOfStrategyByHour.Build(hours))
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query(strategyID, clusterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]*MonitorValueWidthAPI, 0, 1000)

	for rows.Next() {
		v := new(MonitorValueWidthAPI)

		values, err := read(rows, &v.Hour, &v.Id, &v.Name, &v.RequestURL)
		if err != nil {
			return nil, err
		}

		v.Value = values

		all = append(all, v)
	}

	return all, nil
}
func GetStrategyOfAPIHour(clusterId int, apiId int, hours []int) ([]*MonitorValueWidthStrategy, error) {

	stmt, e := database.GetConnection().Prepare(sqlSelectStrategyOfAPIByHour.Build(hours))
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query(apiId, clusterId, )
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]*MonitorValueWidthStrategy, 0, 1000)

	for rows.Next() {
		strategy := new(MonitorValueWidthStrategy)

		values, err := read(rows, &strategy.Hour, &strategy.Id, &strategy.Name, &strategy.Status)
		if err != nil {
			return nil, err
		}

		strategy.Value = values

		all = append(all, strategy)
	}

	return all, nil
}
