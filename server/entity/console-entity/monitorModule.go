package entity

//MonitorModule 监控模板
type MonitorModule struct {
	Name         string `json:"moduleName"`
	Config       string `json:"config,omitempty"`
	ModuleStatus int    `json:"moduleStatus"`
}
