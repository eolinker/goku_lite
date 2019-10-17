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
type StaticDiscovery struct {
}

//SetConfig setConfig
func (d *StaticDiscovery) SetConfig(config string) error {
	return nil
}

//Driver driver
func (d *StaticDiscovery) Driver() string {
	return DriverName
}

//SetCallback setCallBack
func (d *StaticDiscovery) SetCallback(callback func(services []*common.Service)) {
	return
}

//GetServers getServers
func (d *StaticDiscovery) GetServers() ([]*common.Service, error) {
	return nil, nil
}

//Close close
func (d *StaticDiscovery) Close() error {
	return nil
}

//Open open
func (d *StaticDiscovery) Open() error {
	return nil
}
