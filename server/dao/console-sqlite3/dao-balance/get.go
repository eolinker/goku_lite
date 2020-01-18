package dao_balance

import (
	sql2 "database/sql"
	"fmt"
	"strings"

	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity-service"
)

//GetBalanceNames 获取负载名称列表
func (b *BalanceDao) GetBalanceNames() (bool, []string, error) {
	db := b.db
	sql := "SELECT balanceName FROM goku_balance ;"

	rows, err := db.Query(sql)
	if err != nil {
		return false, nil, err
	}
	defer rows.Close()
	//获取记录列

	if _, err = rows.Columns(); err != nil {
		return false, nil, err
	}
	balanceList := make([]string, 0)
	for rows.Next() {
		balanceName := ""
		err = rows.Scan(&balanceName)
		if err != nil {
			return false, nil, err
		}
		balanceList = append(balanceList, balanceName)
	}
	return true, balanceList, nil

}

//Get 根据负载名获取负载配置
func (b *BalanceDao) Get(name string) (*entity.Balance, error) {
	const sql = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,''),A.`balanceDesc`,A.`updateTime`,A.`createTime` FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`NAME` WHERE A.`balanceName`= ?;"
	db := b.db
	v := new(entity.Balance)
	err := db.QueryRow(sql, name).Scan(&v.Name, &v.ServiceName, &v.ServiceDriver, &v.AppName, &v.Static, &v.StaticCluster, &v.Desc, &v.UpdateTime, &v.CreateTime)
	if err != nil {
		return nil, err
	}

	return v.Type(), nil
}

//GetAll 获取所有负载配置
func (b *BalanceDao) GetAll() ([]*entity.Balance, error) {
	const sql = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,''),A.`balanceDesc`,A.`updateTime`,A.`createTime` FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`name` ORDER BY A.`updateTime` DESC;"
	db := b.db
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make([]*entity.Balance, 0, 20)
	for rows.Next() {
		v := new(entity.Balance)
		err := rows.Scan(&v.Name, &v.ServiceName, &v.ServiceDriver, &v.AppName, &v.Static, &v.StaticCluster, &v.Desc, &v.UpdateTime, &v.CreateTime)
		if err != nil {
			return nil, err
		}
		r = append(r, v.Type())
	}
	return r, nil
}

//Search 关键字获取负载列表
func (b *BalanceDao) Search(keyword string) ([]*entity.Balance, error) {
	const sqlTpl = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,''),A.`balanceDesc`,A.`updateTime`,A.`createTime` FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`name` %s ORDER BY `updateTime` DESC;"

	where := ""
	args := make([]interface{}, 0, 3)
	keywordvalue := strings.Trim(keyword, "%")
	if keywordvalue != "" {
		where = "WHERE A.`balanceName` LIKE ? OR A.`serviceName` LIKE ? OR B.`driver` LIKE ?"
		kp := fmt.Sprint("%", keywordvalue, "%")
		args = append(args, kp, kp, kp)
	}
	sql := fmt.Sprintf(sqlTpl, where)
	db := b.db
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make([]*entity.Balance, 0, 20)
	for rows.Next() {
		v := new(entity.Balance)
		err := rows.Scan(&v.Name, &v.ServiceName, &v.ServiceDriver, &v.AppName, &v.Static, &v.StaticCluster, &v.Desc, &v.UpdateTime, &v.CreateTime)
		if err != nil {
			return nil, err
		}
		r = append(r, v.Type())
	}
	return r, nil
}

//GetUseBalanceNames 获取使用的负载名称列表
func (b *BalanceDao) GetUseBalanceNames() (map[string]int, error) {
	const sql = "SELECT `balanceName` as `name` FROM goku_gateway_api UNION SELECT `target` as `name` FROM goku_conn_strategy_api;"
	db := b.db
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make(map[string]int)
	for rows.Next() {
		var balanceName sql2.NullString
		err := rows.Scan(&balanceName)
		if err != nil {
			return nil, err
		}
		if balanceName.Valid == false || balanceName.String == "" {
			continue
		}
		if _, ok := r[balanceName.String]; !ok {
			r[balanceName.String] = 0
		}
		r[balanceName.String] = r[balanceName.String] + 1
	}
	return r, nil
}
