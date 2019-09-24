package dao_balance

import (
	"github.com/eolinker/goku-api-gateway/common/database"
)

func GetBalanceNames() (bool, []string, error) {
	db := database.GetConnection()
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
