package auth

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//GetAuthStatus 获取认证状态
func GetAuthStatus(strategyID string) (bool, map[string]interface{}, error) {
	return console_sqlite3.GetAuthStatus(strategyID)
}

//GetAuthInfo 获取认证信息
func GetAuthInfo(strategyID string) (bool, map[string]interface{}, error) {
	return console_sqlite3.GetAuthInfo(strategyID)
}

//EditAuthInfo 编辑认证信息
func EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error) {
	flag, err := console_sqlite3.EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList,
		jwtCredentialList, oauth2CredentialList, delClientIDList)

	return flag, err
}
