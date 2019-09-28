package admin

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	cluster2 "github.com/eolinker/goku-api-gateway/server/cluster"
	"github.com/eolinker/goku-api-gateway/server/entity"
	"net/http"
	"strconv"
)

//Register 注册
func Register(w http.ResponseWriter, r *http.Request) {

	ip, port, err := GetIpPort(r)
	if err != nil {
		controller.WriteError(w, "700000", "cluster", err.Error(), err)
		return
	}
	cluster, err := regedister(ip, port)
	if err != nil {
		controller.WriteError(w, "700001", "cluster", err.Error()+ip, err)
		return
	}
	node.Refresh(ip, strconv.Itoa(port))
	controller.WriteResultInfo(w, "cluster", "cluster", cluster)
}

func regedister(ip string, port int) (*entity.ClusterInfo, error) {

	has, node, err := node.GetNodeInfoByIPPort(ip, port)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, err
	}
	cName := node.Cluster
	info, has := cluster2.Get(cName)
	if has {
		return info, nil
	}
	return nil, fmt.Errorf("not has that node[%s:%d]", ip, port)

}
