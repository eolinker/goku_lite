package console

import (
	"errors"

	"github.com/eolinker/goku-api-gateway/admin/cmd"
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

var (
	ErrorDuplicateInstance = errors.New("duplicate instance")
	ErrorNeedRegister      = errors.New("need register")
)

func NodeRegister(client *Client) error {

	err := clientManager.Add(client)
	if err != nil {
		return err
	}

	nodeInfo, err := node.GetNodeInfoByKey(client.instance)

	if err != nil {
		return ErrorDuplicateInstance
	}
	result, err := versionConfig.GetConfig(nodeInfo.Cluster)
	if err != nil {
		return err
	}
	nodeConf := toNodeConfig(result, nodeInfo)
	data, _ := cmd.EncodeRegisterResultConfig(nodeConf)

	return client.Send(cmd.NodeRegisterResult, data)
}

func NodeLeave(client *Client) {
	clientManager.Remove(client.instance)

}

func getNodeMapByCluster() (map[string][]*entity.Node, error) {
	nodes, e := node.GetAllNode()
	if e != nil {
		return nil, e
	}

	nodeMap := make(map[string][]*entity.Node)
	for _, node := range nodes {

		nodeMap[node.Cluster] = append(nodeMap[node.Cluster], node)
	}
	return nodeMap, nil
}
func OnConfigChange(conf map[string]*config.GokuConfig) {

	nodeMap, err := getNodeMapByCluster()
	if err != nil {
		log.Warn(err)
	}
	for cluster, c := range conf {
		for _, nodeInfo := range nodeMap[cluster] {

			client, has := clientManager.Get(nodeInfo.NodeKey)
			if has {

				_ = client.SendConfig(c, nodeInfo)

			}

		}
	}
}

func StopNode(nodeKey string) {

	client, has := clientManager.Get(nodeKey)
	if has {
		_ = client.SendRunCMD("stop")
		NodeLeave(client)

	}
}

func RestartNode(nodeKey string) {

	client, has := clientManager.Get(nodeKey)
	if has {
		NodeLeave(client)
		node.UnLock(nodeKey)
		_ = client.SendRunCMD("restart")
	}
}
