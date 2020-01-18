package balance

import (
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

const operationBalance = "balanceManagement"

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":         factory.NewAccountHandleFunction(operationBalance, true, AddBalance),
		"/edit":        factory.NewAccountHandleFunction(operationBalance, true, SaveBalance),
		"/delete":      factory.NewAccountHandleFunction(operationBalance, true, DeleteBalance),
		"/getInfo":     factory.NewAccountHandleFunction(operationBalance, false, GetBalanceInfo),
		"/getList":     factory.NewAccountHandleFunction(operationBalance, false, GetBalanceList),
		"/batchDelete": factory.NewAccountHandleFunction(operationBalance, true, BatchDeleteBalance),
		"/simple":      factory.NewAccountHandleFunction(operationBalance, true, GetSimpleList),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}
