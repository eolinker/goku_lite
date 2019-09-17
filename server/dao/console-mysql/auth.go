package console_mysql

import (
	"encoding/json"
	log "github.com/eolinker/goku/goku-log"
	"time"

	database2 "github.com/eolinker/goku/common/database"
	"github.com/eolinker/goku/utils"
)

type basicAuthConf struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type APIKeyConf struct {
	APIKey         string `json:"Apikey"`
	TokenPlace     string `json:"tokenPlace"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type OAuth2GlobalConf struct {
	oauth2Conf
	Oauth2CredentialList []*oauth2Credential `json:"oauth2CredentialList"`
}

type oauth2Conf struct {
	Scopes                        []string `json:"scopes"`                        //scopes = { required = false, type = "array" },
	MandatoryScope                bool     `json:"mandatoryScope"`                //mandatory_scope = { required = true, type = "boolean", default = false, func = check_mandatory_scope },
	TokenExpiration               int      `json:"tokenExpiration"`               //token_expiration = { required = true, type = "number", default = 7200 },
	EnableAuthorizationCode       bool     `json:"enableAuthorizationCode"`       //enable_authorization_code = { required = true, type = "boolean", default = false },
	EnableImplicitGrant           bool     `json:"enableImplicitGrant"`           //enable_implicit_grant = { required = true, type = "boolean", default = false },
	EnableClientCredentials       bool     `json:"enableClientCredentials"`       //enable_client_credentials = { required = true, type = "boolean", default = false },
	HideCredentials               bool     `json:"hideCredentials"`               //hide_credentials = { type = "boolean", default = false },
	AcceptHttpIfAlreadyTerminated bool     `json:"acceptHttpIfAlreadyTerminated"` //accept_http_if_already_terminated = { required = false, type = "boolean", default = false },
	RefreshTokenTTL               int      `json:"refreshTokenTTL"`               //refresh_token_ttl = {required = true, type = "number", default = 1209600} -- original hardcoded value - 14 days
}

type jwtCredential struct {
	ISS          string `json:"iss"`
	Secret       string `json:"secret"`
	RsaPublicKey string `json:"rsaPublicKey"`
	Algorithm    string `json:"algorithm"`
	Remark       string `json:"remark"`
}

type jwtConf struct {
	SignatureIsBase64 bool            `json:"signatureIsBase64"`
	ClaimsToVerify    []string        `json:"claimsToVerify"`
	RunOnPreflight    bool            `json:"runOnPreflight"`
	JwtCredentials    []jwtCredential `json:"jwtCredentials"`
	HideCredentials   bool            `json:"hideCredentials"`
}

type oauth2Credential struct {
	CredentialID string `json:"credentialID"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RedirectURI  string `json:"redirectURI"`
	Remark       string `json:"remark"`
}

// 获取认证信息
func GetAuthStatus(strategyID string) (bool, map[string]interface{}, error) {
	db := database2.GetConnection()
	var basicStatus, apikeyStatus, jwtStatus, oAuthStatus int
	sql := `SELECT IF(goku_plugin.pluginStatus = 0,0,goku_conn_plugin_strategy.pluginStatus) AS pluginStatus FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_strategy.pluginName WHERE goku_conn_plugin_strategy.pluginName = ? AND goku_conn_plugin_strategy.strategyID = ?;`
	err := db.QueryRow(sql, "goku-basic_auth", strategyID).Scan(&basicStatus)
	if err != nil {

	}
	sql = `SELECT IF(goku_plugin.pluginStatus = 0,0,goku_conn_plugin_strategy.pluginStatus) AS pluginStatus FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_strategy.pluginName WHERE goku_conn_plugin_strategy.pluginName = ? AND goku_conn_plugin_strategy.strategyID = ?;`
	err = db.QueryRow(sql, "goku-apikey_auth", strategyID).Scan(&apikeyStatus)
	if err != nil {

	}
	sql = `SELECT IF(goku_plugin.pluginStatus = 0,0,goku_conn_plugin_strategy.pluginStatus) AS pluginStatus FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_strategy.pluginName WHERE goku_conn_plugin_strategy.pluginName = ? AND goku_conn_plugin_strategy.strategyID = ?;`
	err = db.QueryRow(sql, "goku-jwt_auth", strategyID).Scan(&jwtStatus)
	if err != nil {

	}
	sql = `SELECT IF(goku_plugin.pluginStatus = 0,0,goku_conn_plugin_strategy.pluginStatus) AS pluginStatus FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_strategy.pluginName WHERE goku_conn_plugin_strategy.pluginName = ? AND goku_conn_plugin_strategy.strategyID = ?;`
	err = db.QueryRow(sql, "goku-oauth2_auth", strategyID).Scan(&oAuthStatus)
	if err != nil {

	}
	authInfo := map[string]interface{}{
		"basicAuthStatus": basicStatus,
		"apiKeyStatus":    apikeyStatus,
		"jwtStatus":       jwtStatus,
		"oAuthStatus":     oAuthStatus,
	}
	return true, authInfo, err
}

// 获取认证信息
func GetAuthInfo(strategyID string) (bool, map[string]interface{}, error) {
	db := database2.GetConnection()
	var strategyName, auth string
	sql := "SELECT auth,strategyName FROM goku_gateway_strategy WHERE strategyID = ?;"
	err := db.QueryRow(sql, strategyID).Scan(&auth, &strategyName)
	if err != nil {
		return false, make(map[string]interface{}), err
	}
	basicAuthList := make([]basicAuthConf, 0)
	apiKeyList := make([]APIKeyConf, 0)
	jwtConf := jwtConf{JwtCredentials: make([]jwtCredential, 0)}
	oauth2Conf := OAuth2GlobalConf{}

	var basicConfig, apiKeyConfig, jwtConfig, oAuthConfig string
	sql = `SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`
	err = db.QueryRow(sql, "goku-basic_auth", strategyID).Scan(&basicConfig)
	if err == nil {
		if basicConfig != "" {
			json.Unmarshal([]byte(basicConfig), &basicAuthList)
			if err != nil {
				return false, make(map[string]interface{}), err
			}
		}
	}
	sql = `SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`
	err = db.QueryRow(sql, "goku-apikey_auth", strategyID).Scan(&apiKeyConfig)
	if err == nil {
		if apiKeyConfig != "" {
			err = json.Unmarshal([]byte(apiKeyConfig), &apiKeyList)
			if err != nil {
				return false, make(map[string]interface{}), err
			}
		}
	}
	sql = `SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`
	err = db.QueryRow(sql, "goku-jwt_auth", strategyID).Scan(&jwtConfig)
	if err == nil {
		if jwtConfig != "" {
			err = json.Unmarshal([]byte(jwtConfig), &jwtConf)
			if err != nil {
				return false, make(map[string]interface{}), err
			}
		}
	}

	sql = `SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`
	err = db.QueryRow(sql, "goku-oauth2_auth", strategyID).Scan(&oAuthConfig)
	if err == nil {
		if oAuthConfig != "" {
			err = json.Unmarshal([]byte(oAuthConfig), &oauth2Conf)
			if err != nil {
				return false, make(map[string]interface{}), err
			}
		}
	}
	if oauth2Conf.Oauth2CredentialList == nil {
		oauth2Conf.Oauth2CredentialList = make([]*oauth2Credential, 0)
	}

	authInfo := map[string]interface{}{
		"strategyID":           strategyID,
		"strategyName":         strategyName,
		"auth":                 auth,
		"basicAuthList":        basicAuthList,
		"apiKeyList":           apiKeyList,
		"jwtCredentialList":    jwtConf.JwtCredentials,
		"oauth2CredentialList": oauth2Conf.Oauth2CredentialList,
	}
	return true, authInfo, nil
}

// 编辑认证信息
func EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE goku_gateway_strategy SET strategyName = ? WHERE strategyID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(strategyName, strategyID)
	if err != nil {
		return false, err
	}
	// 设置basic信息
	Tx, _ := db.Begin()
	_, err = Tx.Exec("UPDATE goku_conn_plugin_strategy SET pluginConfig = ?,updateTime = ? WHERE strategyID = ? AND pluginName = ? AND pluginStatus = 1;", basicAuthList, now, strategyID, "goku-basic_auth")
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	_, err = Tx.Exec("UPDATE goku_conn_plugin_strategy SET pluginConfig = ?,updateTime = ? WHERE strategyID = ? AND pluginName = ? AND pluginStatus = 1;", apikeyList, now, strategyID, "goku-apikey_auth")
	if err != nil {
		Tx.Rollback()
		return false, err
	}

	// 获取jwt配置信息
	var jwtConfInfo string
	err = Tx.QueryRow(`SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`, "goku-jwt_auth", strategyID).Scan(&jwtConfInfo)
	if err != nil {
	}
	jwtConf := &jwtConf{}
	jwtList := make([]jwtCredential, 0)
	if jwtCredentialList != "" {
		err = json.Unmarshal([]byte(jwtCredentialList), &jwtList)
		if err != nil {
			Tx.Rollback()
			return false, err
		}
	} else {

		sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
		_, err = Tx.Exec(sql, now, strategyID)
		if err != nil {
			Tx.Rollback()
			return false, err
		}
		Tx.Commit()
		return true, nil
	}

	err = json.Unmarshal([]byte(jwtConfInfo), &jwtConf)
	if err != nil && jwtConfInfo != "" {
		Tx.Rollback()
		return false, err
	}
	jwtConf.JwtCredentials = jwtList
	jwtJson, err := json.Marshal(jwtConf)
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	_, err = Tx.Exec("UPDATE goku_conn_plugin_strategy SET pluginConfig = ?,updateTime = ? WHERE strategyID = ? AND pluginName = ? AND pluginStatus = 1 ", jwtJson, now, strategyID, "goku-jwt_auth")
	if err != nil {
		Tx.Rollback()
		return false, err
	}

	var oAuthConf string
	// 获取oAuth配置信息
	err = Tx.QueryRow(`SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`, "goku-oauth2_auth", strategyID).Scan(&oAuthConf)
	if err != nil {

		sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
		_, err = Tx.Exec(sql, now, strategyID)
		if err != nil {
			Tx.Rollback()
			return false, err
		}
		Tx.Commit()
		return true, nil
	}
	oConf := &OAuth2GlobalConf{}
	oAuthList := make([]*oauth2Credential, 0)
	if oauth2CredentialList != "" {
		err = json.Unmarshal([]byte(oauth2CredentialList), &oAuthList)
		if err != nil {
			Tx.Rollback()
			return false, err
		}
	} else {

		sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
		_, err = Tx.Exec(sql, now, strategyID)
		if err != nil {
			Tx.Rollback()
			return false, err
		}
		Tx.Commit()
		return true, nil
	}
	err = json.Unmarshal([]byte(oAuthConf), &oConf)
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	oConf.Oauth2CredentialList = make([]*oauth2Credential, 0)
	for _, value := range oAuthList {
		if value.CredentialID == "" {
			value.CredentialID = utils.GetRandomString(16)
		}
		if value.ClientID == "" {
			value.ClientID = utils.GetRandomString(16)
		}
		oConf.Oauth2CredentialList = append(oConf.Oauth2CredentialList, value)
	}

	oAuthGlobalConf, err := json.Marshal(oConf)
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	_, err = Tx.Exec("UPDATE goku_conn_plugin_strategy SET pluginConfig = ?,updateTime = ? WHERE strategyID = ? AND pluginName = ? ", oAuthGlobalConf, now, strategyID, "goku-oauth2_auth")
	if err != nil {
		Tx.Rollback()
		return false, err
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	err = Tx.Commit()
	if err != nil {
		info := err.Error()
		log.Info(info)
	}
	return true, nil
}
