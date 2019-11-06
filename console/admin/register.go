package admin

import (
	"github.com/eolinker/goku-api-gateway/console/module/node"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)
func GetNode(key string)(*entity.Node, error){
	  node, err := node.GetNodeInfoByKey(key)
	if err != nil {
		return nil, err
	}

	return node, nil
}
