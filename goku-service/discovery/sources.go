package discovery

import (
	"errors"
	"github.com/eolinker/goku/goku-service/common"
	"github.com/eolinker/goku/goku-service/health"
	"reflect"
	"sync"
	"time"
)

type ISource interface {
	GetApp(app string) (*common.Service, health.CheckHandler, bool)
	SetHealthConfig(conf *HealthCheckConfig)
	SetDriverConfig(config string) error
	Close()
	CheckDriver(driverName string) bool
}

type SourceDiscovery struct {
	name               string
	discovery          Discovery
	healthCheckHandler health.CheckHandler

	locker   sync.RWMutex
	services map[string]*common.Service
}

func (s *SourceDiscovery) SetDriverConfig(config string) error {
	return s.discovery.SetConfig(config)
}

func (s *SourceDiscovery) Close() {
	instances := s.healthCheckHandler.Close()
	for _, instance := range instances {
		instance.ChangeStatus(common.InstanceChecking, common.InstanceRun)
	}

}
func (s *SourceDiscovery) CheckDriver(driverName string) bool {
	if s.discovery == nil {
		return false
	}
	if s.discovery.Driver() == driverName {
		return true
	}
	return false
}

func (s *SourceDiscovery) SetHealthConfig(conf *HealthCheckConfig) {
	if conf == nil || !conf.IsHealthCheck {
		s.Close()
		return
	}

	s.healthCheckHandler.Open(
		conf.Url,
		conf.StatusCode,
		conf.Second,
		time.Duration(conf.TimeOutMill)*time.Millisecond)
}

func (s *SourceDiscovery) GetApp(name string) (*common.Service, health.CheckHandler, bool) {
	s.locker.RLock()
	service, has := s.services[name]
	s.locker.RUnlock()
	if has {
		return service, s.healthCheckHandler, true
	}
	return nil, nil, false
}

func (s *SourceDiscovery) SetServices(services []*common.Service) {

	serviceMap := make(map[string]*common.Service)

	for _, se := range services {
		serviceMap[se.Name] = se
	}

	s.locker.Lock()
	s.services = serviceMap
	s.locker.Unlock()

}

var ErrorEmptyDiscovery = errors.New("discovery is nil")

func NewSource(name string, d Discovery) (*SourceDiscovery, error) {

	if d == nil || reflect.ValueOf(d).IsNil() {
		return nil, ErrorEmptyDiscovery
	}

	s := &SourceDiscovery{
		name:               name,
		discovery:          d,
		healthCheckHandler: new(health.CheckBox),
		services:           make(map[string]*common.Service),
		locker:             sync.RWMutex{},
	}

	//services, e := d.GetServers()
	//if e!=nil{
	//	return nil,e
	//}

	//s.SetServices(services)

	d.SetCallback(s.SetServices)

	return s, d.Open()
}
