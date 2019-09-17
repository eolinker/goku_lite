package static

import (
	"github.com/eolinker/goku/goku-service/discovery"
)

const DriverName = "static"

func init() {
	discovery.RegisteredDiscovery(DriverName, new(Driver))
}
