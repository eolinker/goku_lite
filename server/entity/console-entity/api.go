package entity

import "github.com/eolinker/goku-api-gateway/config"

//API 接口
type API struct {
	APIID            int                      `json:"apiID"`
	APIName          string                   `json:"apiName"`
	Alias            string                   `json:"alias"`
	GroupID          int                      `json:"groupID,omitempty"`
	ProjectID        int                      `json:"projectID,omitempty"`
	RequestURL       string                   `json:"requestURL"`
	ProxyURL         string                   `json:"targetURL"`
	RequestMethod    string                   `json:"requestMethod"`
	TargetServer     string                   `json:"targetServer"`
	TargetMethod     string                   `json:"targetMethod"`
	RequestParamList []*RequestParam          `json:"requestParamList,omitempty"`
	ResultParamList  []string                 `json:"resultParamList,omitempty"`
	IsFollow         bool                     `json:"isFollow"`
	StripPrefix      bool                     `json:"stripPrefix"`
	Timeout          int                      `json:"timeout"`
	RetryConut       int                      `json:"retryCount"`
	UpdateTime       string                   `json:"updateTime"`
	CreateTime       string                   `json:"createTime"`
	Valve            int                      `json:"alertValve"`
	BalanceName      string                   `json:"balanceName"`
	Protocol         string                   `json:"protocol"`
	StripSlash       bool                     `json:"stripSlash"`
	GroupPath        string                   `json:"groupPath"`
	APIType          int                      `json:"apiType"`
	LinkAPIs         []config.APIStepUIConfig `json:"linkApis"`
	StaticResponse   string                   `json:"staticResponse"`
	ResponseDataType string                   `json:"responseDataType"`
	*ManagerInfo
}

//ManagerInfo 用户管理者信息
type ManagerInfo struct {
	ManagerID      int    `json:"managerID"`
	UpdaterID      int    `json:"updaterID"`
	CreateUserID   int    `json:"createUserID"`
	ManagerName    string `json:"managerName"`
	UpdaterName    string `json:"updaterName"`
	CreateUserName string `json:"createUserName"`
}

//RequestParam 请求参数
type RequestParam struct {
	Key         string `json:"key"`
	KeyPosition string `json:"keyPosition"`
	NotEmpty    bool   `json:"notEmpty"`
}

//APIPlugin 接口插件
type APIPlugin struct {
	*API
	StrategyID         string
	PluginList         []*PluginParams
	StrategyPluginList []*PluginParams
}
