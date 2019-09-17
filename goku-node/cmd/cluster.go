package cmd

import "github.com/eolinker/goku/server/entity"

type CMD struct {
	StatusCode string `json:"statuscode"`
}

type ClusterConfig struct {
	CMD
	Cluster *entity.ClusterInfo `json:"cluster"`
}
