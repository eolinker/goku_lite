package console_mysql

import (
	SQL "database/sql"
	"fmt"

	"github.com/eolinker/goku-api-gateway/common/database"
)

func getCountSQL(sql string, args ...interface{}) int {
	var count int
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s) A", sql)
	err := database.GetConnection().QueryRow(countSQL, args...).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count
}

func getPageSQL(sql string, orderBy, orderType string, page, pageSize int, args ...interface{}) (*SQL.Rows, error) {
	pageSQL := fmt.Sprintf("%s ORDER BY %s %s LIMIT ?,?", sql, orderBy, orderType)
	args = append(args, (page-1)*pageSize, pageSize)
	return database.GetConnection().Query(pageSQL, args...)
}
