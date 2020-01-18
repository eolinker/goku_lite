package node

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var (
	nodeDao dao.NodeDao
	nodeGroupDao dao.NodeGroupDao
)

func init() {
	pdao.Need(&nodeDao,&nodeGroupDao)
}