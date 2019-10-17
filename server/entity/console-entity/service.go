package entity

//Service 服务发现
type Service struct {
	Name               string
	Driver             string
	Desc               string
	IsDefault          bool
	Config             string
	ClusterConfig      string
	HealthCheck        bool
	HealthCheckPath    string
	HealthCheckPeriod  int
	HealthCheckCode    string
	HealthCheckTimeOut int
	CreateTime         string
	UpdateTime         string
}
