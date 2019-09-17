package entity

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

type StrategyPluginItem struct {
	StrategyID   string
	PluginName   string
	PluginConfig string
	UpdateTag    string
	//PluginInfo string
	//PluginStatus int
}
type StrategyApiPlugin struct {
	ApiId        string
	StrategyID   string
	PluginName   string
	PluginConfig string
	UpdateTag    string
}
type StrategyApi struct {
	ApiId      int
	StrategyID string
	Target     string
}
