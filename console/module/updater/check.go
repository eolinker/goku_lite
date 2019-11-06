package updater

import (
	"github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/updater"
)

//IsTableExist 检查table是否存在
func IsTableExist(name string) bool {
	return updater.IsTableExist(name)
}

//IsColumnExist 检查列是否存在
func IsColumnExist(name string, column string) bool {
	return updater.IsColumnExist(name, column)
}
