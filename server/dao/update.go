package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/common/database"
)

//GetLastUpdateOfAPI 获取最后更新的接口记录
func GetLastUpdateOfAPI(tables ...string) (time.Time, error) {
	t := time.Time{}
	var updateTime string
	tb := make([]string, len(tables))
	tv := make([]interface{}, len(tables))
	for i, table := range tables {
		tb[i] = "?"
		tv[i] = table
	}
	tablesStr := strings.Join(tb, ",")
	db := database.GetConnection()

	sql := fmt.Sprintf("SELECT updateTime FROM goku_table_update_record WHERE name IN (%s) ORDER BY updateTime desc LIMIT 1;", tablesStr)

	err := db.QueryRow(sql, tv...).Scan(&updateTime)
	if err != nil {
		return t, err
	}
	t, _ = time.ParseInLocation("2006-01-02 15:04:05", updateTime, time.Local)
	return t, nil
}

//UpdateTable 更新goku_table_update_record的updateTime字段
func UpdateTable(name string) error {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO `goku_table_update_record` (`name`,`updateTime`) VALUES (?,?) ON DUPLICATE KEY UPDATE `updateTime` = VALUES(`updateTime`)"
	_, err := db.Exec(sql, name, now)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
