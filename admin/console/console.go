package console

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/eolinker/goku-api-gateway/admin/cmd"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"
)

var (
	register   *Register
	ctx        context.Context
	cancelFunc context.CancelFunc

	once sync.Once
)

func Stop() {
	cancelFunc()
}

func Start(addr string) error {
	once.Do(func() {
		versionConfig.InitVersionConfig()
		register = doRegister()
	})

	var lc net.ListenConfig
	ctx, cancelFunc = context.WithCancel(context.Background())

	listener, err := lc.Listen(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	go doAccept(listener)
	return nil

}

func doAccept(listener net.Listener) {
	for {
		conn, e := listener.Accept()
		if e != nil {
			listener.Close()
			return
		}

		go startClient(conn)

	}
}

func readClient(conn net.Conn) (string, error) {
	frame, e := cmd.ReadFrame(conn)
	if e != nil {
		return "", e
	}

	code, data, e := cmd.GetCmd(frame)
	if e != nil {
		return "", e
	}
	if code != cmd.NodeRegister {
		return "", ErrorNeedRegister
	}

	instance, err := cmd.DecodeRegister(data)
	if err != nil {
		return "", err
	}

	return instance, nil

}
func startClient(conn net.Conn) {

	instance, err := readClient(conn)
	if err != nil {
		return
	}
	fmt.Println(instance)
	if !node.Lock(instance) {
		data, err := cmd.EncodeRegisterResultError(ErrorDuplicateInstance.Error())
		if err == nil {
			cmd.SendFrame(conn, cmd.NodeRegisterResult, data)
		}
		//conn.Close()
		return
	}

	client := NewClient(conn, instance)
	defer func() {
		node.UnLock(instance)
		NodeLeave(client)
		_ = client.Close()
	}()

	e := NodeRegister(client)
	if e != nil {
		data, err := cmd.EncodeRegisterResultError(e.Error())
		if err == nil {
			cmd.SendFrame(conn, cmd.NodeRegisterResult, data)
		}
		return
	}

	for {
		select {
		case frame, ok := <-client.ReadC():
			{
				if !ok {
					return
				}
				code, data, e := cmd.GetCmd(frame)
				if e != nil {
					return
				}
				err := register.Callback(code, data, client)
				if err != nil {
					return
				}
			}
		case <-client.Done():
			return
		}
	}
}
