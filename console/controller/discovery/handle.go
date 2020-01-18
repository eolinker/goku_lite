package discovery

import (
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":     factory.NewAccountHandleFunction(operationDiscovery, true, add),
		"/save":    factory.NewAccountHandleFunction(operationDiscovery, true, edit),
		"/delete":  factory.NewAccountHandleFunction(operationDiscovery, true, delete),
		"/info":    factory.NewAccountHandleFunction(operationDiscovery, false, getInfo),
		"/list":    factory.NewAccountHandleFunction(operationDiscovery, false, list),
		"/default": factory.NewAccountHandleFunction(operationDiscovery, true, setDefault),
		"/drivers": factory.NewAccountHandleFunction(operationDiscovery, false, getDrivices),
		"/simple":  factory.NewAccountHandleFunction(operationDiscovery, false, simple),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

const operationDiscovery = "balanceManagement"
