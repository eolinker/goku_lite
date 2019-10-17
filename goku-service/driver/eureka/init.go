package eureka

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//DriverName 驱动名称
const DriverName = "eureka"

//EurekaStatusUp eureka状态
const EurekaStatusUp = "UP"

//Register 注册
func Register() {
	discovery.RegisteredDiscovery(DriverName, discovery.NewDriver(Create))
}

//Create 创建
func Create(config string) discovery.Discovery {
	return NewEurekaDiscovery(config)
}
