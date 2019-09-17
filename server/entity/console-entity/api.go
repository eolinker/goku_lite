package entity

type Api struct {
	ApiID            int             `json:"apiID"`
	ApiName          string          `json:"apiName"`
	GroupID          int             `json:"groupID,omitempty"`
	ProjectID        int             `json:"projectID,omitempty"`
	RequestURL       string          `json:"requestURL"`
	ProxyURL         string          `json:"targetURL"`
	RequestMethod    string          `json:"requestMethod"`
	TargetServer     string          `json:"targetServer"`
	TargetMethod     string          `json:"targetMethod"`
	RequestParamList []*RequestParam `json:"requestParamList,omitempty"`
	ResultParamList  []string        `json:"resultParamList,omitempty"`
	IsFollow         bool            `json:"isFollow"`
	StripPrefix      bool            `json:"stripPrefix"`
	Timeout          int             `json:"timeout"`
	RetryConut       int             `json:"retryCount"`
	UpdateTime       string          `json:"updateTime"`
	CreateTime       string          `json:"createTime"`
	Valve            int             `json:"alertValve"`
	BalanceName      string          `json:"balanceName"`
	Protocol         string          `json:"protocol"`
	StripSlash       bool            `json:"stripSlash"`
	GroupPath        string          `json:"groupPath"`
	*ManagerInfo
}

type ManagerInfo struct {
	ManagerID      int    `json:"managerID"`
	UpdaterID      int    `json:"updaterID"`
	CreateUserID   int    `json:"createUserID"`
	ManagerName    string `json:"managerName"`
	UpdaterName    string `json:"updaterName"`
	CreateUserName string `json:"createUserName"`
}

type RequestParam struct {
	Key         string `json:"key"`
	KeyPosition string `json:"keyPosition"`
	NotEmpty    bool   `json:"notEmpty"`
}

type ApiPlugin struct {
	*Api
	StrategyID         string
	PluginList         []*PluginParams
	StrategyPluginList []*PluginParams
}
