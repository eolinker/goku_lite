package console

import (
	"fmt"
	"sync"

	"github.com/eolinker/goku-api-gateway/common/listener"
)

type ClientManager struct {
	clients   map[string]*Client
	locker    sync.RWMutex
	intercept *listener.Intercept
}

var (
	clientManager = &ClientManager{
		clients:   make(map[string]*Client),
		locker:    sync.RWMutex{},
		intercept: listener.NewIntercept(),
	}
)

func (m *ClientManager) Add(client *Client) (err error) {

	if _, has := m.Get(client.instance); has {
		return ErrorDuplicateInstance
	}

	e := m.intercept.Call(client)
	if e != nil {
		return e
	}
	m.locker.Lock()

	_, has := m.clients[client.instance]
	if has {
		fmt.Println(client.instance)
		err = ErrorDuplicateInstance
	} else {
		m.clients[client.instance] = client
		err = nil
	}
	m.locker.Unlock()
	return
}
func (m *ClientManager) Get(instance string) (*Client, bool) {
	m.locker.RLock()
	c, has := m.clients[instance]
	m.locker.RUnlock()
	return c, has
}
func (m *ClientManager) Remove(instance string) {
	m.locker.Lock()
	delete(m.clients, instance)
	m.locker.Unlock()
}
func InterceptNodeRegister(f func(client *Client) error) {
	clientManager.intercept.Add(func(v interface{}) error {
		c := v.(*Client)
		return f(c)
	})
}
func IsLive(key string) bool {
	_, b := clientManager.Get(key)
	return b
}
