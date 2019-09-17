package discovery

type HealthCheckConfig struct {
	IsHealthCheck bool
	Url string
	Second int
	TimeOutMill int
	StatusCode string

}
type Config struct {
	Name string
	Driver string
	Config string
	HealthCheckConfig
}