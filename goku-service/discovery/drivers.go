package discovery

import "sync"

//Driver driver
type Driver interface {
	Open(name string, config string) (ISource, error)
}

//CreateHandler createHandler
type CreateHandler func(config string) Discovery

//DriverBase driverBase
type DriverBase struct {
	createFunc CreateHandler
	locker     sync.RWMutex
	sources    map[string]*SourceDiscovery
}

//NewDriver newDriver
func NewDriver(createFunc CreateHandler) *DriverBase {
	return &DriverBase{
		createFunc: createFunc,
		locker:     sync.RWMutex{},
		sources:    make(map[string]*SourceDiscovery),
	}
}

//Open open
func (d *DriverBase) Open(name string, config string) (ISource, error) {
	d.locker.RLock()

	s, h := d.sources[name]
	d.locker.RUnlock()

	if h {
		return s, s.SetDriverConfig(config)
	}
	d.locker.Lock()

	s, h = d.sources[name]
	if h {
		d.locker.Unlock()
		return s, s.SetDriverConfig(config)
	}

	ds := d.createFunc(config)

	s, _ = NewSource(name, ds)
	d.sources[name] = s
	d.locker.Unlock()
	return s, nil
}
