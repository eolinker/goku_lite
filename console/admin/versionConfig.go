package admin

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/module/node"

	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"

	"github.com/eolinker/goku-api-gateway/console/controller"
)

//GetVersionConfig 获取版本配置
func GetVersionConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	httpRequest.ParseForm()
	version := httpRequest.Form.Get("version")
	ip, port, err := GetIPPort(httpRequest)
	if err != nil {
		controller.WriteError(httpResponse, "700000", "cluster", err.Error(), err)
		return
	}
	node.Refresh(ip, strconv.Itoa(port))
	cluster, err := getCluster(ip, port)
	if err != nil {
		controller.WriteError(httpResponse, "700001", "cluster", err.Error()+ip, err)
		return
	}

	result := versionConfig.GetVersionConfig(cluster, version)

	httpResponse.Write(result)
}
