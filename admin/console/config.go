package console

import (
	"github.com/eolinker/goku-api-gateway/config"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

func toNodeConfig(c *config.GokuConfig, nodeInfo *entity.Node) *config.GokuConfig {

	conf := *c
	conf.Cluster = nodeInfo.Cluster
	conf.BindAddress = nodeInfo.ListenAddress
	conf.AdminAddress = nodeInfo.AdminAddress
	conf.Instance = nodeInfo.NodeKey

	return &conf
}
