package console_sqlite3

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/eolinker/goku-api-gateway/server/dao"

	log "github.com/eolinker/goku-api-gateway/goku-log"
)

type basicAuthConf struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

//APIKeyConf apiKey配置
type APIKeyConf struct {
	APIKey         string `json:"Apikey"`
	TokenPlace     string `json:"tokenPlace"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

//OAuth2GlobalConf oauth2配置
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
	AcceptHTTPIfAlreadyTerminated bool     `json:"acceptHttpIfAlreadyTerminated"` //accept_http_if_already_terminated = { required = false, type = "boolean", default = false },
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

//AuthDao AuthDao
type AuthDao struct {
	db *sql.DB
}

//NewAuthDao new AuthDao
func NewAuthDao() *AuthDao {
	return &AuthDao{}
}

//Create create
func (d *AuthDao) Create(db *sql.DB) (interface{}, error) {

	d.db = db

	var i dao.AuthDao = d

	return &i, nil
}

//GetAuthStatus 获取认证状态
func (d *AuthDao) GetAuthStatus(strategyID string) (bool, map[string]interface{}, error) {
	db := d.db
	var basicStatus, apikeyStatus int
	sql := `SELECT CASE WHEN goku_plugin.pluginStatus = 0 THEN 0 ELSE goku_conn_plugin_strategy.pluginStatus END AS pluginStatus FROM goku_conn_plugin_strategy INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_strategy.pluginName WHERE goku_conn_plugin_strategy.pluginName = ? AND goku_conn_plugin_strategy.strategyID = ?;`
	db.QueryRow(sql, "goku-basic_auth", strategyID).Scan(&basicStatus)

	db.QueryRow(sql, "goku-apikey_auth", strategyID).Scan(&apikeyStatus)

	authInfo := map[string]interface{}{
		"basicAuthStatus": basicStatus,
		"apiKeyStatus":    apikeyStatus,
		"jwtStatus":       0,
		"oAuthStatus":     0,
	}
	return true, authInfo, nil
}

//GetAuthInfo 获取认证信息
func (d *AuthDao) GetAuthInfo(strategyID string) (bool, map[string]interface{}, error) {
	db := d.db
	var strategyName, auth string
	sql := "SELECT IFNULL(auth,''),strategyName FROM goku_gateway_strategy WHERE strategyID = ?;"
	err := db.QueryRow(sql, strategyID).Scan(&auth, &strategyName)
	if err != nil {
		return false, make(map[string]interface{}), err
	}
	basicAuthList := make([]basicAuthConf, 0)
	apiKeyList := make([]APIKeyConf, 0)

	var basicConfig, apiKeyConfig string
	sql = `SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE pluginName = ? AND strategyID = ? AND pluginStatus = 1;`
	err = db.QueryRow(sql, "goku-basic_auth", strategyID).Scan(&basicConfig)
	if err == nil {
		if basicConfig != "" {
			json.Unmarshal([]byte(basicConfig), &basicAuthList)
			if err != nil {
				panic(err)
				return false, make(map[string]interface{}), err
			}
		}
	}
	err = db.QueryRow(sql, "goku-apikey_auth", strategyID).Scan(&apiKeyConfig)
	if err == nil {
		if apiKeyConfig != "" {
			err = json.Unmarshal([]byte(apiKeyConfig), &apiKeyList)
			if err != nil {
				panic(err)
				return false, make(map[string]interface{}), err
			}
		}
	}

	authInfo := map[string]interface{}{
		"strategyID":           strategyID,
		"strategyName":         strategyName,
		"auth":                 auth,
		"basicAuthList":        basicAuthList,
		"apiKeyList":           apiKeyList,
		"jwtCredentialList":    make([]interface{}, 0),
		"oauth2CredentialList": make([]interface{}, 0),
	}
	return true, authInfo, nil
}

//EditAuthInfo 编辑认证信息
func (d *AuthDao) EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error) {
	db := d.db
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
