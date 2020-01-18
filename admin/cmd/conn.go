package cmd

import (
	"context"
	"errors"
	"net"
	"sync"
)

var (
	ErrorSendToClosedConnect = errors.New("send to closed connect")
)

type Connect struct {
	conn net.Conn

	//inputC chan _Frame
	outputC chan []byte

	doneC chan struct{}

	ctx        context.Context
	cancelFunc context.CancelFunc

	once sync.Once
}

func NewConnect(conn net.Conn) *Connect {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Connect{
		conn: conn,

		//inputC:     make(chan _Frame,10),
		outputC: make(chan []byte, 10),

		ctx:        ctx,
		cancelFunc: cancel,
	}

	go c.r()

	return c
}
func (c *Connect) r() {

	for {
		frame, err := ReadFrame(c.conn)
		if err != nil {
			break
		}
		c.outputC <- frame
	}
	close(c.outputC)
}

func (c *Connect) Close() error {

	c.once.Do(func() {
		c.cancelFunc()
		c.conn.Close()

	})

	return nil
}

func (c *Connect) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Connect) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connect) Send(code Code, data []byte) error {
	return SendFrame(c.conn, code, data)
}

func (c *Connect) ReadC() <-chan []byte {
	return c.outputC
}
func (c *Connect) Done() <-chan struct{} {
	return c.ctx.Done()
}
