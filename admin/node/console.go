package node

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/admin/cmd"
	"github.com/eolinker/goku-api-gateway/common/listener"
	"github.com/eolinker/goku-api-gateway/common/manager"
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/node/console"
)

type TcpConsole struct {
	conn       *cmd.Connect
	addr       string
	lock       sync.Mutex
	instance   string
	register   *Register
	listener   *listener.Listener
	lastConfig *manager.Value

	ctx    context.Context
	cancel context.CancelFunc

	listenOnce sync.Once
}

func (c *TcpConsole) SendMonitor(data []byte) error {

	return c.conn.Send(cmd.Monitor, data)

}

func (c *TcpConsole) GetConfig() (*config.GokuConfig, error) {
	conf, b := c.lastConfig.Get()
	if b {

		return conf.(*config.GokuConfig), nil
	}
	return nil, errors.New("not register to console")
}

func (c *TcpConsole) Close() {
	c.cancel()
}

func (c *TcpConsole) AddListen(callback console.ConfigCallbackFunc) {
	c.listener.Listen(func(event interface{}) {
		conf := event.(*config.GokuConfig)
		callback(conf)
	})
}

func NewConsole(addr string, instance string) *TcpConsole {
	c := &TcpConsole{
		addr:     addr,
		instance: instance,
		conn:     nil,
		register: NewRegister(),
		listener: listener.New(),

		lastConfig: manager.NewValue(),
	}
	c.register.RegisterFunc(cmd.Config, c.OnConfigChange)
	c.register.RegisterFunc(cmd.Restart, Restart)
	c.register.RegisterFunc(cmd.Stop, Stop)

	return c
}

func connect(addr string) net.Conn {
	sleeps := []time.Duration{time.Second * 0, time.Second * 1, time.Second * 5, time.Second * 10}
	maxSleep := sleeps[len(sleeps)-1]
	retry := 0

	for {
		if retry > 0 {
			if retry > len(sleeps)-1 {
				time.Sleep(maxSleep)
			} else {
				time.Sleep(sleeps[retry])
			}
		}
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		return conn
	}
}

func (c *TcpConsole) RegisterToConsole() (*config.GokuConfig, error) {
	data, err := cmd.EncodeRegister(c.instance)
	if err != nil {
		return nil, err
	}
	for {

		conn := connect(c.addr)
		e := cmd.SendFrame(conn, cmd.NodeRegister, data)
		if e != nil {
			conn.Close()
			continue
		}

		frame, err := cmd.ReadFrame(conn)

		if err != nil {
			conn.Close()
			continue
		}

		code, data, err := cmd.GetCmd(frame)
		if err != nil {
			conn.Close()
			return nil, err
		}
		if code != cmd.NodeRegisterResult {
			conn.Close()
			return nil, ErrorNeedReadRegisterResult
		}

		result, err := cmd.DecodeRegisterResult(data)

		if err != nil {
			conn.Close()
			return nil, err
		}
		if result.Code != 0 {
			conn.Close()
			return nil, errors.New(result.Error)
		}

		c.conn = cmd.NewConnect(conn)

		return result.Config, nil
	}
}
func (c *TcpConsole) Listen() {

	c.listenOnce.Do(
		func() {
			go func() {
				for {
					c.listenRead()
					c.RegisterToConsole()
				}
			}()
		})

}

func (c *TcpConsole) listenRead() {
	defer c.conn.Close()
	for {
		select {
		case <-c.conn.Done():
			return
		case frame, ok := <-c.conn.ReadC():
			{
				if !ok {
					return
				}

				code, data, e := cmd.GetCmd(frame)
				if e != nil {
					return
				}

				err := c.register.Callback(code, data)
				if err != nil {
					return
				}
			}
		}
	}
}
