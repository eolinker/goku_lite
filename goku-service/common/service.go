package common

import (
	"math/rand"
	"sort"
	"sync"
)

//Service service
type Service struct {
	Name      string
	instances []*Instance
	//lastIndex int
	locker sync.RWMutex
}

//NewService 创建Service
func NewService(name string, Instances []*Instance) *Service {
	return &Service{
		Name:      name,
		instances: Instances,
		//lastIndex: 0,
		locker: sync.RWMutex{},
	}
}

//SetInstances setInstances
func (s *Service) SetInstances(instances []*Instance) {

	sort.Sort(sort.Reverse(PInstances(instances)))
	s.locker.Lock()
	//old:=s.instances
	s.instances = instances

	s.locker.Unlock()
	//
	//for _,instance:=range old{
	//	instance.ChangeStatus(InstanceChecking,InstanceDown)
	//}
}

//Weighting weighting
func (s *Service) Weighting() (*Instance, int, bool) {
	s.locker.RLock()
	instances := s.instances
	s.locker.RUnlock()

	if len(instances) == 0 {
		return nil, 0, false
	}
	weightSum := 0
	for _, ins := range instances {
		if ins.CheckStatus(InstanceRun) {
			weightSum += ins.Weight
		}
	}
	if weightSum == 0 {
		return nil, 0, false
	}
	weightValue := rand.Intn(weightSum) + 1
	for i, ins := range instances {
		if ins.CheckStatus(InstanceRun) {
			weightValue = weightValue - ins.Weight
			if weightValue <= 0 {
				return ins, i, true
			}
		}

	}
	return nil, 0, false
}

//Next next
func (s *Service) Next(lastIndex int) (*Instance, int, bool) {
	if lastIndex == -1 {
		return s.Weighting()
	}
	s.locker.RLock()
	instances := s.instances
	s.locker.RUnlock()

	size := len(instances)
	if size == 0 {
		return nil, 0, false
	}

	for i := 0; i < size; i++ {
		index := (lastIndex + i) % size
		instance := instances[index]
		if instance != nil {
			if instance.CheckStatus(InstanceRun) {
				return instance, index, true
			}
		}
	}
	return nil, 0, false
}
