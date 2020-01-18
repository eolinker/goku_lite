package main

import (
	graphite "github.com/eolinker/goku-api-gateway/module/graphite/config"
	prometheus "github.com/eolinker/goku-api-gateway/module/prometheus/config"
)

func moduleRegister() {
	prometheus.Register()
	graphite.Register()

}
