package static

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//DriverName 驱动名称
const DriverName = "static"

//Register 注册
func Register() {
	discovery.RegisteredDiscovery(DriverName, new(Driver))
}
