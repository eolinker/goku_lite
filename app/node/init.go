package main

import (
	"github.com/eolinker/goku-api-gateway/module/graphite"
	"github.com/eolinker/goku-api-gateway/module/prometheus"
)

func init() {

	prometheus.Register()
	graphite.Register()
}
