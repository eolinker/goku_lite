package static

import (
	"github.com/eolinker/goku/goku-service/common"
	"github.com/eolinker/goku/goku-service/discovery"
)

type Driver struct {

}

func (d *Driver) Open(name string,config string)(discovery.ISource,error) {

	return NewStaticSources(name),nil
}

type StaticDiscovery struct {

}

func (d *StaticDiscovery) SetConfig(config string) error {
	return nil
}

func (d *StaticDiscovery) Driver() string {
	return DriverName
}

func (d *StaticDiscovery) SetCallback(callback func(services []*common.Service)) {
	return
}

func (d *StaticDiscovery) GetServers() ([]*common.Service, error) {
	return nil,nil
}

func (d *StaticDiscovery) Close() error {
return nil
}

func (d *StaticDiscovery) Open() error {
return nil
}
