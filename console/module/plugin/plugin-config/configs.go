package plugin_config

func init() {
	allConfigOfPlugin = map[string]interface{}{
		"goku-apikey_auth":             new(APIKeyConf),
		"goku-basic_auth":              new(basicAuthConf),
		"goku-circuit_breaker":         new(CircuitBreakerConf),
		"goku-cors":                    new(gokuCorsConfig),
		"goku-data_format_transformer": new(dataFormatTranformerConf),
		"goku-default_response":        new(defaultResponseConf),
		"goku-extra_params":            new(extraParamsConf),
		"goku-http_log":                new(Log),
		"goku-ip_restriction":          new(IPList),
		"goku-jwt_auth":                new(JwtConf),
		"goku-oauth2_auth":             new(Oauth2Conf),
		"goku-params_check":            new(paramsCheckConf),
		"goku-params_transformer":      new(paramsTransformerconf),
		"goku-proxy_caching":           new(ProxyCachingConf),
		"goku-rate_limiting":           new(_RateLimitingConf),
		"goku-replay_attack_defender":  new(ReplayAttackDefenderConf),
		"goku-request_size_limiting":   new(requestSizeLimit),
		"goku-response_headers":        new(responseHeader),
		"goku-service_downgrade":       new(serviceDowngradeConf),
	}
}

type APIKeyNode struct {
	APIKey         string `json:"Apikey"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type APIKeyConf []APIKeyNode

type basicAuthNode struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type basicAuthConf []basicAuthNode

type CircuitBreakerConf struct {
	ThresholdPercent float64           `json:"failurePercent"`
	SamplesSize      int               `json:"minimumRequests"`
	SamplesPeriod    int               `json:"monitorPeriod"`
	BreakPeriod      int64             `json:"breakPeriod"`
	SuccessCounts    int               `json:"successCounts"`
	MatchStatusCodes string            `json:"matchStatusCodes"`
	StatusCode       int               `json:"statusCode"`
	Headers          map[string]string `json:"headers"`
	Body             string            `json:"body"`
}

type gokuCorsConfig struct {
	AllowOrigin      string `json:"allowOrigin"`
	AllowMethod      string `json:"allowMethods"`
	AllowCredentials string `json:"allowCredentials"`
	AllowHeaders     string `json:"allowHeaders"`
	ExposeHeaders    string `json:"exposeHeaders"`
}
type dataFormatTranformerConf struct {
	EnableXMLToJSON           bool   `json:"enableXMLToJSON"`
	EnableJSONToXML           bool   `json:"enableJSONToXML"`
	ContinueIfTransformFailed bool   `json:"continueIfTransformFailed"`
	XMLRootTag                string `json:"XMLRootTag"`
}
type defaultResponseConf struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type extraParam struct {
	ParamName             string      `json:"paramName"`
	ParamPosition         string      `json:"paramPosition"`
	ParamValue            interface{} `json:"paramValue"`
	ParamConflictSolution string      `json:"paramConflictSolution"`
}

type extraParamsConf struct {
	Params []*extraParam `json:"params"`
}
type Log struct {
	LogName      string `json:"logName"`
	FileDir      string `json:"fileDir"`      // 文件夹路径
	RecordPeriod string `json:"recordPeriod"` // 记录周期
}
type IPList struct {
	IpListType  string   `json:"ipListType"`
	IpWhiteList []string `json:"ipWhiteList"`
	IpBlackList []string `json:"ipBlackList"`
}

type JwtCredential struct {
	ISS          string `json:"iss"`
	Secret       string `json:"secret"`
	RsaPublicKey string `json:"rsaPublicKey"`
	Algorithm    string `json:"algorithm"`
}

type JwtConf struct {
	KeyClaimName      string          `json:"keyClaimName"`
	SignatureIsBase64 bool            `json:"signatureIsBase64"`
	ClaimsToVerify    []string        `json:"claimsToVerify"`
	RunOnPreflight    bool            `json:"runOnPreflight"`
	JwtCredentials    []JwtCredential `json:"jwtCredentials"`
	HideCredentials   bool            `json:"hideCredentials"`
}

type Oauth2Conf struct {
	Scopes                  []string `json:"scopes"`                  //scopes = { required = false, type = "array" },
	MandatoryScope          bool     `json:"mandatoryScope"`          //mandatory_scope = { required = true, type = "boolean", default = false, func = check_mandatory_scope },
	TokenExpiration         int      `json:"tokenExpiration"`         //token_expiration = { required = true, type = "number", default = 7200 },
	EnableAuthorizationCode bool     `json:"enableAuthorizationCode"` //enable_authorization_code = { required = true, type = "boolean", default = false },
	EnableImplicitGrant     bool     `json:"enableImplicitGrant"`     //enable_implicit_grant = { required = true, type = "boolean", default = false },
	EnableClientCredentials bool     `json:"enableClientCredentials"` //enable_client_credentials = { required = true, type = "boolean", default = false },
	//EnablePasswordGrant           bool     `json:"enablePasswordGrant"`           //enable_password_grant = { required = true, type = "boolean", default = false },
	HideCredentials               bool `json:"hideCredentials"`               //hide_credentials = { type = "boolean", default = false },
	AcceptHttpIfAlreadyTerminated bool `json:"acceptHttpIfAlreadyTerminated"` //accept_http_if_already_terminated = { required = false, type = "boolean", default = false },
	RefreshTokenTTL               int  `json:"refreshTokenTTL"`               //refresh_token_ttl = {required = true, type = "number", default = 1209600} -- original hardcoded value - 14 days

	Oauth2Credentials []Oauth2Credential `json:"oauth2CredentialList"`
}

type Oauth2Credential struct {
	CredentialID string `json:"credentialID"`
	StrategyID   string `json:"strategyID"`
	Name         string `json:"name"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RedirectURI  string `json:"redirectURI"`
}
type paramsCheckInfo struct {
	ParamName     string `json:"paramName"`
	ParamPosition string `json:"paramPosition"`
	Regular       string `json:"regular"`
}

type paramsCheckConf struct {
	Params []paramsCheckInfo `json:"params"`
}

type paramsTransformerparam struct {
	ParamName             string `json:"paramName"`
	ParamPosition         string `json:"paramPosition"`
	ProxyParamName        string `json:"proxyParamName"`
	ProxyParamPosition    string `json:"proxyParamPosition"`
	Required              bool   `json:"required"`
	ParamConflictSolution string `json:"paramConflictSolution"`
}

type paramsTransformerconf struct {
	Params                 []paramsTransformerparam `json:"params"`
	RemoveAfterTransformed bool                     `json:"removeAfterTransformed"`
}

type ProxyCachingConf struct {
	ResponseCodes  string `json:"responseCodes"`  //缓存条件：返回的HTTP状态码在该状态码列表中
	RequestMethods string `json:"requestMethods"` //缓存条件：请求的Method在该列表中
	ContentTypes   string `json:"contentTypes"`   //缓存条件：返回的Content-Type在该列表中
	CacheTTL       int    `json:"cacheTTL"`
	RedisHost      string `json:"redisHost"`
	RedisTimeout   int    `json:"redisTimeout"`
	RedisPort      string `json:"redisPort"`
	RedisPassword  string `json:"redisPassword"`
	RedisDatabase  int    `json:"redisDatabase"`
}

type _RateLimitingConf struct {
	Second           int64 `json:"second,omitempty"`
	Minute           int64 `json:"minute,omitempty"`
	Hour             int64 `json:"hour,omitempty"`
	Day              int64 `json:"day,omitempty"`
	HideClientHeader bool  `json:"hideClientHeader"`
}
type ReplayAttackDefenderConf struct {
	TimeStampTTL int64 `json:"timestampTTL"`
	//EnableRequestVerify bool  `json:"enableRequestVerify"`
	Token string `json:"replayAttackToken"`
}
type requestSizeLimit struct {
	AllowedPayLoadSize int `json:"allowedPayLoadSize"`
}
type responseHeader struct {
	MatchStatusCode string            `json:"matchStatusCode"`
	Headers         map[string]string `json:"responseHeaders"`
}
type serviceDowngradeConf struct {
	MatchStatusCodes string            `json:"matchStatusCodes"`
	StatusCode       int               `json:"statusCode"`
	Headers          map[string]string `json:"headers"`
	Body             string            `json:"body"`
}
