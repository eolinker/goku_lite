package common

import "sync"

//Instance instance
type Instance struct {
	InstanceID string
	IP         string
	Port       int
	Weight     int
	Status     InstanceStatus
	locker     sync.RWMutex
}

//PInstances PInstances
type PInstances []*Instance

func (p PInstances) Len() int {
	return len(p)
}

func (p PInstances) Less(i, j int) bool {
	return p[i].Weight < p[j].Weight
}

func (p PInstances) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//CheckStatus checkStatus
func (i *Instance) CheckStatus(status InstanceStatus) bool {
	i.locker.RLock()
	b := i.Status == status
	i.locker.RUnlock()
	return b
}

//ChangeStatus set status to desc  where status is org
func (i *Instance) ChangeStatus(org, dest InstanceStatus) bool {
	if org == dest {
		return i.CheckStatus(org)
	}

	i.locker.RLock()
	b := i.Status == org
	i.locker.RUnlock()

	if !b {
		return false
	}

	i.locker.Lock()
	b = i.Status == org
	if b {
		i.Status = dest
	}
	i.locker.Unlock()
	return b

}
