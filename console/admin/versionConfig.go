package admin

import (
	"encoding/json"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"
	"net/http"
)


//GetVersionConfig 获取版本配置
func GetVersionConfig(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	version := r.Form.Get("version")
	ip, instance, err := GetInstanceAndIP(r)
	if err != nil {
		controller.WriteError(w, "700000", "cluster", err.Error(), err)
		return
	}

	if !node.Lock(instance){
		controller.WriteError(w, "700007", "cluster", "invalid instance", nil)
		return
	}

	defer node.UnLock(instance)

	nodeInfo, err := GetNode(instance)

	if err != nil {
		controller.WriteError(w, "700001", "cluster", err.Error()+ip, err)
		return
	}
	ctx := r.Context()
	result ,err:= versionConfig.GetVersionConfig(ctx,nodeInfo.Cluster, version)
	if err!= nil{
		// client close, connect close
		return
	}

	result.Cluster = nodeInfo.Cluster
	result.BindAddress = nodeInfo.ListenAddress
	result.AdminAddress = nodeInfo.AdminAddress
	result.Instance = nodeInfo.NodeKey
	data,_:= json.Marshal(result)
	w.Write(data)


}
