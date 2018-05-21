package module

import (
	"goku-ce/server/dao"
)

func EditGlobalConfig(gatewayPort string) bool {
	return dao.EditGlobalConfig(gatewayPort)
}