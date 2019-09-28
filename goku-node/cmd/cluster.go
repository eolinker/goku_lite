package cmd

import "github.com/eolinker/goku-api-gateway/server/entity"

//CMD cmd
type CMD struct {
	StatusCode string `json:"statuscode"`
}

//ClusterConfig clusterConfig
type ClusterConfig struct {
	CMD
	Cluster *entity.ClusterInfo `json:"cluster"`
}
