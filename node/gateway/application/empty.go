package application

import "github.com/eolinker/goku-api-gateway/goku-node/common"

//EmptyApplication empty application
type EmptyApplication struct {
	response string
}

//Execute execute
func (app *EmptyApplication) Execute(ctx *common.Context) {
	panic("implement me")
}

//NewEmptyApplication 创建空应用
func NewEmptyApplication(response string) *EmptyApplication {
	return &EmptyApplication{
		response: response,
	}
}
