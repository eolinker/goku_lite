package discovery

import (
	"github.com/eolinker/goku-api-gateway/goku-service/common"
)

//Discovery discovery
type Discovery interface {
	SetConfig(config string) error
	Driver() string
	SetCallback(callback func(services []*common.Service))
	GetServers() ([]*common.Service, error)
	Close() error
	Open() error
}
