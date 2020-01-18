package application

import (
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
)

type staticeResponse struct {
	body     []byte
	strategy config.StaticResponseStrategy
}

func newStaticeResponse(body string, strategy config.StaticResponseStrategy) *staticeResponse {
	return &staticeResponse{body: []byte(body), strategy: strategy}
}

func (sp *staticeResponse) Do(ctx *common.Context) {
	ctx.SetBody(sp.body)
	ctx.SetStatus(200, "200")
}
