package internal

import (
	"sync"
	"time"
)

//IDCreate idCreate
type IDCreate struct {
	locker sync.Mutex
	lastID uint64
}

//NewIDCreate new idCreate
func NewIDCreate() *IDCreate {
	return &IDCreate{locker: sync.Mutex{}}
}

//Next next
func (c *IDCreate) Next() uint64 {
	var id uint64
	c.locker.Lock()

	id = uint64(time.Now().UnixNano())
	if id <= c.lastID {
		id = c.lastID + 1
	}

	c.lastID = id

	c.locker.Unlock()

	return id
}
