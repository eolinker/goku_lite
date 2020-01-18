package console

import (
	"net"

	"github.com/eolinker/goku-api-gateway/admin/cmd"
	"github.com/eolinker/goku-api-gateway/config"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

type Client struct {
	*cmd.Connect
	instance string
}

func NewClient(conn net.Conn, instance string) *Client {

	return &Client{
		Connect:  cmd.NewConnect(conn),
		instance: instance,
	}
}

func (c *Client) Instance() string {
	return c.instance
}

func (c *Client) SendConfig(conf *config.GokuConfig, nodeInfo *entity.Node) error {

	nodeConfig := toNodeConfig(conf, nodeInfo)

	data, err := cmd.EncodeConfig(nodeConfig)
	if err != nil {
		return err
	}

	return c.Send(cmd.Config, data)
}

func (c *Client) SendRunCMD(operate string) error {

	if operate == "stop" {
		return c.Send(cmd.Stop, []byte(""))
	} else if operate == "restart" {

		return c.Send(cmd.Restart, []byte(""))
	}
	return nil

}
