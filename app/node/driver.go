package main

import (
	"github.com/eolinker/goku-api-gateway/goku-service/driver/consul"
	"github.com/eolinker/goku-api-gateway/goku-service/driver/eureka"
	"github.com/eolinker/goku-api-gateway/goku-service/driver/static"
)

func init() {
	consul.Register()
	eureka.Register()
	static.Register()
}
