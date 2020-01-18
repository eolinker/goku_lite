package static

import (
	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//Driver driver
type Driver struct {
}

//Open open
func (d *Driver) Open(name string, config string) (discovery.ISource, error) {

	return NewStaticSources(name), nil
}

//StaticDiscovery staticDiscovery
type Discovery struct {
}

//SetConfig setConfig
func (d *Discovery) SetConfig(config string) error {
	return nil
}

//Driver driver
func (d *Discovery) Driver() string {
	return DriverName
}

//SetCallback setCallBack
func (d *Discovery) SetCallback(callback func(services []*common.Service)) {
	return
}

//GetServers getServers
func (d *Discovery) GetServers() ([]*common.Service, error) {
	return nil, nil
}

//Close close
func (d *Discovery) Close() error {
	return nil
}

//Open open
func (d *Discovery) Open() error {
	return nil
}
