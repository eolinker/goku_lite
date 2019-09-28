package gokunode


import (
	// 驱动加载
	_ "github.com/eolinker/goku-api-gateway/goku-node/manager/service-manager"
	_ "github.com/eolinker/goku-api-gateway/goku-service/driver/consul"
	_ "github.com/eolinker/goku-api-gateway/goku-service/driver/eureka"
	_ "github.com/eolinker/goku-api-gateway/goku-service/driver/static"
)
