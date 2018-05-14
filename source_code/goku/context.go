package goku

import (
	"goku-ce/conf"
)
type Context struct {
	GatewayInfo Gateway
	StrategyInfo Strategy
	ApiInfo Api
	Rate map[string]Rate
}

type Gateway struct {
	GatewayAlias			string					`json:"gateway_alias" yaml:"gateway_alias"`
	GatewayStatus			string					`json:"gateway_status" yaml:"gateway_status"`
	IPLimitType				string					`json:"ip_limit_type" yaml:"ip_limit_type"`
	IPWhiteList				[]string				`json:"ip_white_list" yaml:"ip_white_list"`
	IPBlackList				[]string				`json:"ip_black_list" yaml:"ip_black_list"`
}

type Strategy struct {
	StrategyID				string					`json:"strategy_id" yaml:"strategy_id"`
	Auth					string					`json:"auth" yaml:"auth"`
	BasicUserName			string					`json:"basic_user_name" yaml:"basic_user_name"`
	BasicUserPassword		string					`json:"basic_user_password" yaml:"basic_user_password"`
	ApiKey					string					`json:"api_key" yaml:"api_key"`
	IPLimitType				string					`json:"ip_limit_type" yaml:"ip_limit_type"`
	IPWhiteList				[]string				`json:"ip_white_list" yaml:"ip_white_list"`
	IPBlackList				[]string				`json:"ip_black_list" yaml:"ip_black_list"`
	RateLimitList			[]conf.RateLimitInfo			`json:"rate_limit_list" yaml:"rate_limit_list"`
}

type Api struct {
	RequestURL				string					`json:"request_url" yaml:"request_url"`
	BackendPath				string					`json:"backend_path" yaml:"backend_path"`
	ProxyURL				string					`json:"proxy_url" yaml:"proxy_url"`
	ProxyMethod				string					`json:"proxy_method" yaml:"proxy_method"`
	IsRaw					bool					`json:"is_raw" yaml:"is_raw"`
	ProxyParams				[]conf.Param			`json:"proxy_params" yaml:"proxy_params"`
	ConstantParams			[]conf.ConstantParam	`json:"constant_params" yaml:"constant_params"`	
}
