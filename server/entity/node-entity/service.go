package entity

type Service struct {
	Name string
	Driver string
	Desc string
	IsDefault bool
	Config string
	ClusterConfig string
	HealthCheck bool
	HealthCheckPath string
	HealthCheckPeriod int
	HealthCheckCode string
	HealthCheckTimeOut int

}
