package conf

import (
	"io/ioutil"
	"os"

)

var (
	Configure string
)

type GlobalConfig struct {
	Host					string					`json:"host" yaml:"host"`
	Port					string					`json:"port" yaml:"port"`
	GatewayConfPath			string					`json:"gateway_conf_path" yaml:"gateway_conf_path"`
	GatewayList				[]GatewayInfo
}

type GatewayInfo struct {
	GatewayName				string					`json:"gateway_name" yaml:"gateway_name"`
	GatewayAlias			string					`json:"gateway_alias" yaml:"gateway_alias"`
	GatewayStatus			string					`json:"gateway_status" yaml:"gateway_status"`
	PluginConfPath			string					`json:"plugin_conf_path" yaml:"plugin_conf_path"`
	ApiConfPath				string					`json:"api_conf_path" yaml:"api_conf_path"`
	ApiGroupConfPath		string					`json:"api_group_conf_path" yaml:"api_group_conf_path"`
	StrategyConfPath		string					`json:"strategy_conf_path" yaml:"strategy_conf_path"`
	BackendConfPath			string					`json:"backend_conf_path" yaml:"backend_conf_path"`
	Timeout					int						`json:"timeout" yaml:"timeout"`
	IPLimitType				string					`json:"ip_limit_type" yaml:"ip_limit_type"`
	IPWhiteList				[]string				`json:"ip_white_list" yaml:"ip_white_list"`
	IPBlackList				[]string				`json:"ip_black_list" yaml:"ip_black_list"`
	StrategyList			Strategy				
	ApiList					Api
	BackendList				Backend
}

type Strategy struct {
	Strategy				[]StrategyInfo			`json:"strategy" yaml:"strategy"`
}

type StrategyInfo struct {
	StrategyName			string					`json:"strategy_name" yaml:"strategy_name"`
	StrategyID				string					`json:"strategy_id" yaml:"strategy_id"`
	Auth					string					`json:"auth" yaml:"auth"`
	BasicUserName			string					`json:"basic_user_name" yaml:"basic_user_name"`
	BasicUserPassword		string					`json:"basic_user_password" yaml:"basic_user_password"`
	ApiKey					string					`json:"api_key" yaml:"api_key"`
	IPLimitType				string					`json:"ip_limit_type" yaml:"ip_limit_type"`
	IPWhiteList				[]string				`json:"ip_white_list" yaml:"ip_white_list"`
	IPBlackList				[]string				`json:"ip_black_list" yaml:"ip_black_list"`
	RateLimitList			[]RateLimitInfo			`json:"rate_limit_list" yaml:"rate_limit_list"`
}

type RateLimitInfo struct {
	Allow					bool					`json:"allow" yaml:"allow"`
	Period					string					`json:"period" yaml:"period"`
	Limit					int						`json:"limit" yaml:"limit"`
	Priority				int						`json:"priority" yaml:"priority"`
	StartTime				int						`json:"start_time" yaml:"start_time"`
	EndTime					int						`json:"end_time" yaml:"end_time"`
}

type ApiGroupInfo struct {
	GroupID					int						`json:"group_id" yaml:"group_id"`
	GroupName				string					`json:"group_name" yaml:"group_name"`
}

type ApiGroup struct {
	Group					ApiGroupInfo			`json:"group" yaml:"group"`
}

type Api struct {
	Apis					[]ApiInfo					`json:"apis" yaml:"apis"`
}

type ApiInfo struct {
	ApiName					string					`json:"api_name" yaml:"api_name"`
	GroupID					int						`json:"group_id" yaml:"group_id"`
	RequestURL				string					`json:"request_url" yaml:"request_url"`
	RequestMethod			[]string				`json:"request_method" yaml:"request_method"`
	BackendID				int						`json:"backend_id" yaml:"backend_id"`
	ProxyURL				string					`json:"proxy_url" yaml:"proxy_url"`
	ProxyMethod				string					`json:"proxy_method" yaml:"proxy_method"`
	ProxyBodyType			string					`json:"proxy_body_type" yaml:"proxy_body_type"`
	ProxyBody				string					`json:"proxy_body" yaml:"proxy_body"`
	ProxyParams				[]Param					`json:"proxy_params" yaml:"proxy_params"`
	ConstantParams			[]ConstantParam			`json:"constant_params" yaml:"constant_params"`						
}

type Backend struct {
	Backend					[]BackendInfo			`json:"backend" yaml:"backend"`
}

type BackendInfo struct {
	BackendID				int						`json:"backend_id" yaml:"backend_id"`
	BackendName				string					`json:"backend_name" yaml:"backend_name"`
	BackendPath				string					`json:"backend_path" yaml:"backend_path"`
}

type Param struct {
	Key						string					`json:"key" yaml:"key"`
	KeyPosition				string					`json:"key_position" yaml:"key_position"`
	NotEmpty				bool					`json:"not_empty" yaml:"not_empty"`
	ProxyKey				string					`json:"proxy_key" yaml:"proxy_key"`
	ProxyKeyPosition		string					`json:"proxy_key_position" yaml:"proxy_key_position"`
}

type ConstantParam struct {
	Position				string					`json:"position" yaml:"position"`
	Key						string					`json:"key" yaml:"key"`
	Value					string					`json:"value" yaml:"value"`
}


func init() {
	Configure = ""
}

func ReadConfigure(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	Configure = string(content)
	return
}

// 获取目录内文件夹
func GetDir(dirPth string) (files []string, err error) {
	files = make([]string, 0)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
	 	return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			files = append(files, dirPth+PthSep+fi.Name())
	 	}
	}
	return files, nil
}

