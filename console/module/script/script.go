package script

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

// 新建项目
func RefreshApiInfo() bool {
	//return console_mysql.RefreshApiInfo()
	return true
}

// 新建项目
func RefreshGatewayAlertConfig() bool {
	return console_mysql.RefreshGatewayAlertConfig()
}

func UpdateTables(names []string) {
	for _, name := range names {
		dao.UpdateTable(name)
	}
}
