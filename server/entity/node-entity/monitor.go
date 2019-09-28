package entity

//MonitorConfig 监控配置
type MonitorConfig struct {
	APIMonitorStatus           int
	StrategyMonitorStatus      int
	StrategyMonitorStatusInAPI int
	APIMonitorStatusInStrategy int
}
