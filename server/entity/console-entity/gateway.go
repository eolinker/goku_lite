package entity

//GatewayBasicConfig 网关基础配置
type GatewayBasicConfig struct {
	SuccessCode         string `json:"successCode"`
	NodeUpdatePeriod    int    `json:"nodeUpdatePeriod"`
	MonitorUpdatePeriod int    `json:"monitorUpdatePeriod"`
	MonitorTimeout      int    `json:"monitorTimeout"`
	SkipCertificate     int    `json:"skipCertificate"`
}
