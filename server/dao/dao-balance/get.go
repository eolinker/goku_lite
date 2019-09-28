package daobalance

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity-service"
	"strings"
)

//Get 获取负载信息
func Get(name string) (*entity.Balance, error) {
	const sql = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,''),A.`balanceDesc`,A.`updateTime`,A.`createTime` FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`NAME` WHERE A.`balanceName`= ?;"
	db := database.GetConnection()
	v := new(entity.Balance)
	err := db.QueryRow(sql, name).Scan(&v.Name, &v.ServiceName, &v.ServiceDriver, &v.AppName, &v.Static, &v.StaticCluster, &v.Desc, &v.UpdateTime, &v.CreateTime)
	if err != nil {
		return nil, err
	}

	return v.Type(), nil
}

//GetAll 获取负载列表
func GetAll() ([]*entity.Balance, error) {
	const sql = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,''),A.`balanceDesc`,A.`updateTime`,A.`createTime` FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`name` ORDER BY `updateTime` DESC;"
	db := database.GetConnection()
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

//Search 查询
func Search(keyword string) ([]*entity.Balance, error) {
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
	db := database.GetConnection()
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
