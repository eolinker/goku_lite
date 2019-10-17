package discovery

import (
	"errors"
	"reflect"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"github.com/eolinker/goku-api-gateway/goku-service/health"
)

//ISource iSource
type ISource interface {
	GetApp(app string) (*common.Service, health.CheckHandler, bool)
	SetHealthConfig(conf *config.HealthCheckConfig)
	SetDriverConfig(config string) error
	Close()
	CheckDriver(driverName string) bool
}

//SourceDiscovery sourceDiscovery
type SourceDiscovery struct {
	name               string
	discovery          Discovery
	healthCheckHandler health.CheckHandler

	locker   sync.RWMutex
	services map[string]*common.Service
}

//SetDriverConfig 设置服务发现配置
func (s *SourceDiscovery) SetDriverConfig(config string) error {
	return s.discovery.SetConfig(config)
}

//Close close
func (s *SourceDiscovery) Close() {
	instances := s.healthCheckHandler.Close()
	for _, instance := range instances {
		instance.ChangeStatus(common.InstanceChecking, common.InstanceRun)
	}

}

//CheckDriver checkDriver
func (s *SourceDiscovery) CheckDriver(driverName string) bool {
	if s.discovery == nil {
		return false
	}
	if s.discovery.Driver() == driverName {
		return true
	}
	return false
}

//SetHealthConfig 设置健康检查配置
func (s *SourceDiscovery) SetHealthConfig(conf *config.HealthCheckConfig) {
	if conf == nil || !conf.IsHealthCheck {
		s.Close()
		return
	}

	s.healthCheckHandler.Open(
		conf.URL,
		conf.StatusCode,
		conf.Second,
		time.Duration(conf.TimeOutMill)*time.Millisecond)
}

//GetApp getApp
func (s *SourceDiscovery) GetApp(name string) (*common.Service, health.CheckHandler, bool) {
	s.locker.RLock()
	service, has := s.services[name]
	s.locker.RUnlock()
	if has {
		return service, s.healthCheckHandler, true
	}
	return nil, nil, false
}

//SetServices setServices
func (s *SourceDiscovery) SetServices(services []*common.Service) {

	serviceMap := make(map[string]*common.Service)

	for _, se := range services {
		serviceMap[se.Name] = se
	}

	s.locker.Lock()
	s.services = serviceMap
	s.locker.Unlock()

}

//ErrorEmptyDiscovery errorEmptyDiscovery
var ErrorEmptyDiscovery = errors.New("discovery is nil")

//NewSource newSource
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

	d.SetCallback(s.SetServices)

	return s, d.Open()
}
