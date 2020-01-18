package goku_handler

import (
	"fmt"
	"net/http"
	"strings"
)

//GokuHandler gokuHandler
type GokuHandler interface {
	Handlers(factory *AccountHandlerFactory) map[string]http.Handler
}

//GokuServer gokuServer
type GokuServer struct {
	factory *AccountHandlerFactory
	mux     *http.ServeMux
}

func (s *GokuServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

//NewGokuServer new gokuServer
func NewGokuServer(factory *AccountHandlerFactory) *GokuServer {
	return &GokuServer{
		factory: factory,
		mux:     http.NewServeMux(),
	}
}

//Add add
func (s *GokuServer) Add(pre string, handler GokuHandler) {
	handlers := handler.Handlers(s.factory)

	for sub, h := range handlers {

		path := join(pre, sub)
		s.mux.Handle(path, h)
	}
}

func join(pre, sub string) string {

	return fmt.Sprint(strings.TrimSuffix(pre, "/"), "/", strings.TrimPrefix(sub, "/"))

}
