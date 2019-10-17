package discovery

import (
	"strings"

	log "github.com/eolinker/goku-api-gateway/goku-log"
)

var (
	//isLock =false
	drivers     = make(map[string]Driver)
	driverNames = make([]string, 0)
)

//AllDrivers allDrivers
func AllDrivers() []string {
	return driverNames
}

// main里应该调用这个方法，以锁住driver, @为了线程安全并且避免锁操作@
//func LockDriver(){
//	if isLock{
//		return
//	}
//	isLock=true
//}

//RegisteredDiscovery registerdDiscovery
func RegisteredDiscovery(name string, driver Driver) {

	//if isLock{
	//	panic("can not Register now")
	//}

	name = strings.ToLower(name)

	_, has := drivers[name]
	if has {
		log.Panic("driver duplicate:" + name)
	}
	drivers[name] = driver

	driverNames = append(driverNames, name)
}
