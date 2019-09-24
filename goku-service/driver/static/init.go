package static

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

const DriverName = "static"

func init() {
	discovery.RegisteredDiscovery(DriverName, new(Driver))
}
