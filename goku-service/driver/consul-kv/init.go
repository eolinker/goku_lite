package consul_kv

import (
	"github.com/eolinker/goku/goku-service/discovery"
)
const DriverName = "consulKv"
func init() {
	discovery.RegisteredDiscovery(DriverName,new(Driver))
}
