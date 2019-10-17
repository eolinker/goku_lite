package consul

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//DriverName 驱动名称
const DriverName = "consul"

//Register register
func Register() {
	discovery.RegisteredDiscovery(DriverName, discovery.NewDriver(Create))
}

//Create 创建
func Create(config string) discovery.Discovery {
	return NewConsulDiscovery(config)
}
