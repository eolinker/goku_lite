package admin

import (
	"github.com/eolinker/goku-api-gateway/console/module/node"
)

func getCluster(ip string, port int) (string, error) {

	has, node, err := node.GetNodeInfoByIPPort(ip, port)
	if err != nil {
		return "", err
	}
	if !has {
		return "", err
	}

	return node.Cluster, nil
}
