package entity

//Strategy strategy
type Strategy struct {
	StrategyID   string
	StrategyName string
	Auth         string
	EnableStatus int
	StrategyType int
}

//
//func (s *Strategy)Enable()bool  {
//	return s.EnableStatus ==1
//}

//StrategyPluginItem strategy plugin item
type StrategyPluginItem struct {
	StrategyID   string
	PluginName   string
	PluginConfig string
	UpdateTag    string
	//PluginInfo string
	//PluginStatus int
}

//StrategyAPIPlugin 策略接口插件
type StrategyAPIPlugin struct {
	APIId        string
	StrategyID   string
	PluginName   string
	PluginConfig string
	UpdateTag    string
}

//StrategyAPI 策略接口
type StrategyAPI struct {
	APIID      int
	StrategyID string
	Target     string
}
