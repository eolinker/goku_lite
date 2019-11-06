package console

import (
	"context"
	"sync"
)

//Console console
type Console struct {
	adminHost   string
	instance        string
	ctx         context.Context
	cancel      context.CancelFunc
	lastVersion int
	once        sync.Once
}

//Close close
func (c *Console) Close() {
	c.once.Do(c.cancel)
}

//NewConsole newConsole
func NewConsole(instance string, adminHost string) *Console {
	ctx, cancel := context.WithCancel(context.Background())
	return &Console{
		instance:      instance,
		adminHost: adminHost,
		ctx:       ctx,
		cancel:    cancel,
	}
}
