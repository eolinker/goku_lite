package discovery

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"sync"
)

var manager = &Manager{
	locker:  sync.RWMutex{},
	sources: make(map[string]ISource),
}

type Manager struct {
	locker sync.RWMutex

	sources map[string]ISource
}

func ResetAllServiceConfig(confs []*Config) {

	sources := make(map[string]ISource)
	manager.locker.RLock()
	oldSources := manager.sources
	manager.locker.RUnlock()
	for _, conf := range confs {

		name := conf.Name
		s, has := oldSources[name]
		if has && !s.CheckDriver(conf.Driver) {
			// 如果驱动不一样，关闭旧的
			has = false
			s.Close()
			s = nil
			delete(oldSources, name)
		}
		if !has {
			driverName := conf.Driver
			driver, has := drivers[driverName]
			if !has {
				log.Error("invalid driver:", driverName)
				continue
			}
			ns, err := driver.Open(name, conf.Config)
			if err != nil {
				continue
			}
			s = ns
		}

		sources[name] = s
		s.SetHealthConfig(&conf.HealthCheckConfig)

		err := s.SetDriverConfig(conf.Config)

		if err != nil {
			continue
		}
	}

	for name, s := range oldSources {
		if _, has := sources[name]; !has {
			s.Close()
		}
	}

	manager.locker.Lock()
	manager.sources = sources
	manager.locker.Unlock()
}

func GetDiscoverer(discoveryName string) (ISource, bool) {
	manager.locker.RLock()
	s, has := manager.sources[discoveryName]
	manager.locker.RUnlock()
	return s, has
}
