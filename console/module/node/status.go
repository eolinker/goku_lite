package node

import (
	"sync"
	"time"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//EXPIRE 心跳检测过期时间
const EXPIRE = time.Second * 10

var (
	manager = _StatusManager{
		locker:        sync.RWMutex{},
		lastHeartBeat: make(map[string]time.Time),
	}
	instanceLocker = newInstanceLocker()
)

type _StatusManager struct {
	locker        sync.RWMutex
	lastHeartBeat map[string]time.Time
}

func (m *_StatusManager) refresh(id string) {
	t := time.Now()
	//heartBeat, err := nodeDao.GetHeartBeatTime(id)
	//if err == nil {
	//	t = heartBeat
	//}
	m.locker.Lock()

	m.lastHeartBeat[id] = t

	m.locker.Unlock()
	nodeDao.SetHeartBeatTime(id, time.Now())
}

func (m *_StatusManager) stop(id string) {

	m.locker.Lock()

	delete(m.lastHeartBeat, id)

	m.locker.Unlock()
}
func (m *_StatusManager) get(id string) (time.Time, bool) {
	m.locker.RLock()
	t, b := m.lastHeartBeat[id]
	m.locker.RUnlock()
	return t, b
}

type _InstanceLocker struct {
	locker    sync.RWMutex
	instances map[string]bool
}

func newInstanceLocker() *_InstanceLocker {
	return &_InstanceLocker{
		locker:    sync.RWMutex{},
		instances: make(map[string]bool),
	}
}
func (l *_InstanceLocker) IsLock(key string) bool {
	l.locker.RLock()
	locked := l.instances[key]
	l.locker.RUnlock()
	return locked
}
func (l *_InstanceLocker) Lock(key string) bool {

	locked := l.IsLock(key)
	if locked {
		return false
	}

	l.locker.Lock()
	locked = l.instances[key]
	if locked {
		l.locker.Unlock()
		return false
	}
	l.instances[key] = true
	l.locker.Unlock()
	return true
}
func (l *_InstanceLocker) UnLock(key string) {

	l.locker.Lock()
	l.instances[key] = false
	l.locker.Unlock()
}

//Refresh refresh

func Refresh(instance string) {

	manager.refresh(instance)
}

func GetLastHeartTime(instance string) (time.Time, bool) {
	return manager.get(instance)
}

//IsLive 通过ip和端口获取当前节点在线状态
func IsLive(instance string) bool {

	if instanceLocker.IsLock(instance) {
		return true
	}

	t, has := manager.get(instance)

	if !has {
		return false
	}
	now := time.Now()
	if now.Sub(t) > EXPIRE {
		return false
	}
	return true
}

//ResetNodeStatus 重置节点状态
func ResetNodeStatus(nodes ...*entity.Node) {
	for _, node := range nodes {
		if instanceLocker.IsLock(node.NodeKey) || IsLive(node.NodeKey) {
			node.NodeStatus = 1
		} else {
			if node.NodeStatus == 1 {
				node.NodeStatus = 2
			} else {
				node.NodeStatus = 0
			}
		}
	}
}

func Lock(key string) bool {
	return instanceLocker.Lock(key)
}
func UnLock(key string) {
	instanceLocker.UnLock(key)
	Refresh(key)

}
