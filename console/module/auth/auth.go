package auth

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)
var (
	authDao dao.AuthDao
)

func init() {
	pdao.Need(&authDao)
}
//GetAuthStatus 获取认证状态
func GetAuthStatus(strategyID string) (bool, map[string]interface{}, error) {
	return authDao.GetAuthStatus(strategyID)
}

//GetAuthInfo 获取认证信息
func GetAuthInfo(strategyID string) (bool, map[string]interface{}, error) {
	return authDao.GetAuthInfo(strategyID)
}

//EditAuthInfo 编辑认证信息
func EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error) {
	flag, err := authDao.EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList,
		jwtCredentialList, oauth2CredentialList, delClientIDList)

	return flag, err
}
