package kubernetes

import (
	"github.com/eolinker/goku/goku-service/discovery"
)

type Driver struct {

}

func (d *Driver) Open(name string,config string)(discovery.ISource,error) {
	panic("implement me")
}


type KubernetesDiscovery struct {

}