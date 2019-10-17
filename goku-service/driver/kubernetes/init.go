package kubernetes

import (
	discoveryManager "github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//DriverName 驱动名称
const DriverName = "kubernetes"

func init() {

	discoveryManager.RegisteredDiscovery(DriverName, new(Driver))
}
