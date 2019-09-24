package eureka

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

const DriverName = "eureka"
const EurekaStatusUp = "UP"

func init() {

	discovery.RegisteredDiscovery(DriverName, discovery.NewDriver(Create))

}
func Create(config string) discovery.Discovery {
	return NewEurekaDiscovery(config)
}
