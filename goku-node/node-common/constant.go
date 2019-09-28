package nodecommon

import (
	"fmt"
	"strings"
)

var (
	//ListenPort 网关监听端口
	ListenPort  = 6689
	clusterName string
	adminURL    = ""
)

//SetClusterName 设置集群名称
func SetClusterName(name string) {
	clusterName = name
}

//ClusterName 获取集群名称
func ClusterName() string {
	return clusterName
}

//SetAdmin 设置admin地址
func SetAdmin(host string) {
	h := strings.TrimPrefix(host, "http://")
	h = strings.TrimSuffix(h, "/")
	adminURL = fmt.Sprintf("http://%s", h)
}

//GetAdminURL 获取adminURL
func GetAdminURL(path string) string {
	p := strings.TrimPrefix(path, "/")
	return fmt.Sprintf("%s/%s", adminURL, p)
}
