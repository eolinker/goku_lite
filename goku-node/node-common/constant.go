package node_common

import (
	"fmt"
	"strings"
)

var (
	ListenPort  = 6689
	clusterName string
	adminUrl    = ""
)

func SetClusterName(name string) {
	clusterName = name
}
func ClusterName() string {
	return clusterName
}

func SetAdmin(host string) {
	h:= strings.TrimPrefix(host,"http://")
	h = strings.TrimSuffix(h,"/")
	adminUrl = fmt.Sprintf("http://%s", h)
}
func GetAdminUrl(path string) string {
	p:=strings.TrimPrefix(path,"/")
	return fmt.Sprintf("%s/%s", adminUrl, p)
}
