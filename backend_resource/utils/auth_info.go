package utils

type AuthInfo struct{
	AuthID			int				`json:"authID,omitempty"`
	AuthType		int				`json:"authType"`
	ApiKey			string			`json:"apiKey"`
	UserName		string			`json:"userName"`
	UserPassword	string			`json:"userPassword"`
	StrategyID		int				`json:"strategyID,omitempty"`
	Authorization	string			`json:"authorization,omitempty`
}