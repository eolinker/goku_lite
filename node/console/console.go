package console

import (
	"context"
	"sync"
)

//Console console
type Console struct {
	adminHost   string
	port        int
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
func NewConsole(port int, adminHost string) *Console {
	ctx, cancel := context.WithCancel(context.Background())
	return &Console{
		port:      port,
		adminHost: adminHost,
		ctx:       ctx,
		cancel:    cancel,
	}
}
