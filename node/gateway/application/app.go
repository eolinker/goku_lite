package application

import (
	"github.com/eolinker/goku-api-gateway/goku-node/common"
)

type Application interface {
	Execute(ctx *common.Context)
}
