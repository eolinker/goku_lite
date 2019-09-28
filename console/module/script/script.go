package script

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	consolemysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//RefreshAPIInfo 刷新接口信息
func RefreshAPIInfo() bool {
	//return consolemysql.RefreshAPIInfo()
	return true
}

//RefreshGatewayAlertConfig 刷新网关告警配置
func RefreshGatewayAlertConfig() bool {
	return consolemysql.RefreshGatewayAlertConfig()
}

//UpdateTables 更新表
func UpdateTables(names []string) {
	for _, name := range names {
		dao.UpdateTable(name)
	}
}
