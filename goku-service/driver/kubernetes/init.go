package kubernetes

import (
	discoveryManager "github.com/eolinker/goku/goku-service/discovery"
)
const DriverName = "kubernetes"
func init() {

	discoveryManager.RegisteredDiscovery(DriverName,new(Driver))
}
