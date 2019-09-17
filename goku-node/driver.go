package goku_node

import (
	_ "github.com/eolinker/goku/goku-node/manager/service-manager"
	_ "github.com/eolinker/goku/goku-service/driver/consul"
	_ "github.com/eolinker/goku/goku-service/driver/eureka"
	_ "github.com/eolinker/goku/goku-service/driver/static"
)
