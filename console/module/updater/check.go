package updater

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)
var (
	updaterDao dao.UpdaterDao
)

func init() {
	pdao.Need(&updaterDao)
}
//IsTableExist 检查table是否存在
func IsTableExist(name string) bool {
	return updaterDao.IsTableExist(name)
}

//IsColumnExist 检查列是否存在
func IsColumnExist(name string, column string) bool {
	return updaterDao.IsColumnExist(name, column)
}
