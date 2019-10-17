package plugin_config

func init() {
	allConfigOfPlugin = map[string]interface{}{
		"goku-apikey_auth":        new(APIKeyConf),
		"goku-basic_auth":         new(basicAuthConf),
		"goku-extra_params":       new(extraParamsConf),
		"goku-ip_restriction":     new(IPList),
		"goku-params_transformer": new(paramsTransformerconf),
	}
}

//APIKeyNode apiKey节点
type APIKeyNode struct {
	APIKey         string `json:"Apikey"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

//APIKeyConf apiKey配置
type APIKeyConf []APIKeyNode

type basicAuthNode struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type basicAuthConf []basicAuthNode

type extraParam struct {
	ParamName             string      `json:"paramName"`
	ParamPosition         string      `json:"paramPosition"`
	ParamValue            interface{} `json:"paramValue"`
	ParamConflictSolution string      `json:"paramConflictSolution"`
}

type extraParamsConf struct {
	Params []*extraParam `json:"params"`
}

//IPList IP黑白名单
type IPList struct {
	IPListType  string   `json:"ipListType"`
	IPWhiteList []string `json:"ipWhiteList"`
	IPBlackList []string `json:"ipBlackList"`
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
