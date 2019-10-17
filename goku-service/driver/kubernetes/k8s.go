package kubernetes

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//Driver driver
type Driver struct {
}

//Open open
func (d *Driver) Open(name string, config string) (discovery.ISource, error) {
	panic("implement me")
}

//Discovery discovery
type Discovery struct {
}
