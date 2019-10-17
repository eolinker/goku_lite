package node

import (
	"fmt"
	"sync"
	"time"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//EXPIRE 心跳检测过期时间
const EXPIRE = time.Second * 20

var (
	manager = _StatusManager{
		locker:        sync.RWMutex{},
		lastHeartBeat: make(map[string]time.Time),
	}
)

type _StatusManager struct {
	locker        sync.RWMutex
	lastHeartBeat map[string]time.Time
}

func (m *_StatusManager) refresh(id string) {
	t := time.Now()
	m.locker.Lock()

	m.lastHeartBeat[id] = t

	m.locker.Unlock()
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

//Refresh refresh
func Refresh(ip string, port string) {
	id := fmt.Sprintf("%s:%d", ip, port)
	manager.refresh(id)
}

//IsLive 通过ip和端口获取当前节点在线状态
func IsLive(ip string, port string) bool {
	id := fmt.Sprintf("%s:%d", ip, port)
	t, has := manager.get(id)
	if !has {
		return false
	}

	if time.Now().Sub(t) > EXPIRE {
		return false
	}
	return true
}

//ResetNodeStatus 重置节点状态
func ResetNodeStatus(nodes ...*entity.Node) {
	for _, node := range nodes {

		if IsLive(node.NodeIP, node.NodePort) {
			node.NodeStatus = 1
		} else {
			node.NodeStatus = 0
		}
	}
}
