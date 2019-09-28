package auth

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//GetAuthStatus 获取认证信息
func GetAuthStatus(strategyID string) (bool, map[string]interface{}, error) {
	return console_mysql.GetAuthStatus(strategyID)
}

//GetAuthInfo 获取认证信息
func GetAuthInfo(strategyID string) (bool, map[string]interface{}, error) {
	return console_mysql.GetAuthInfo(strategyID)
}

//EditAuthInfo 编辑认证信息
func EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error) {
	flag, err := console_mysql.EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList,
		jwtCredentialList, oauth2CredentialList, delClientIDList)
	name := "goku_conn_plugin_strategy"
	if flag {
		dao.UpdateTable(name)
	}
	return flag, err
}
