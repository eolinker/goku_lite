package application

import "github.com/eolinker/goku-api-gateway/goku-node/common"

type EmptyApplication struct {
	response string
}

func (app *EmptyApplication) Execute(ctx *common.Context) {
	panic("implement me")
}

func NewEmptyApplication(response string) *EmptyApplication {
	return &EmptyApplication{
		response: response,
	}
}
