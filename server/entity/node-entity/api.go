package entity

type Api struct {
	ApiID         int
	ApiName       string
	RequestURL    string
	RequestMethod string
	Protocol      string
	BalanceName   string
	IsFollow      bool
	StripPrefix   bool
	Timeout       int
	RetryCount    int
	TargetMethod  string
	TargetURL     string
	AlertValve    int
	StripSlash    bool // 是否过滤斜杠
}

type ApiExtend struct {
	*Api
	Target       string

}
