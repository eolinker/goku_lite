package config_log

import (
	"net/http"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

const operationLog = "logManagement"

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/console": NewLogHandler("console", factory),
		"/node":    NewLogHandler("node", factory),
		"/access":  NewAccessHandler(factory),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}
