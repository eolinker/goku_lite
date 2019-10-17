package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/eolinker/goku-api-gateway/common/endless"
	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/node/console"
	"github.com/eolinker/goku-api-gateway/node/gateway"
	"github.com/eolinker/goku-api-gateway/node/router/httprouter"
)

//Server server
type Server struct {
	port    int
	console *console.Console
	router  http.Handler
}

//NewServer newServer
func NewServer(port int) *Server {
	return &Server{
		port:    port,
		console: nil,
		router:  nil,
	}
}

//SetRouter setRouter
func (s *Server) SetRouter(r http.Handler) error {
	s.router = r
	return nil
}

//SetConsole setConsole
func (s *Server) SetConsole(c *console.Console) {

	s.console = c
	return
}

//Server server
func (s *Server) Server() error {
	if s.router == nil && s.console == nil {
		return errors.New("can not start server widthout router and console")
	}

	if s.console != nil {

		conf, err := s.console.GetConfig()
		if err != nil {
			return err
		}
		SetLog(conf.Log)
		SetAccessLog(conf.AccessLog)

		r, err := gateway.Parse(conf, httprouter.Factory())
		if err != nil {
			log.Panic("parse config error:", err)
		}
		e := s.SetRouter(r)
		if e != nil {
			return e
		}

		s.console.AddListen(s.FlushConfig)
	}

	return endless.ListenAndServe(fmt.Sprintf(":%d", s.port), s)
}

//FlushConfig flushConfig
func (s *Server) FlushConfig(config *config.GokuConfig) {

	go func() {
		r, err := gateway.Parse(config, httprouter.Factory())
		if err != nil {
			log.Error("parse config error:", err)
			return
		}
		s.SetRouter(r)

	}()

}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if s.router == nil {
		w.WriteHeader(404)
		return
	}

	s.router.ServeHTTP(w, req)

}
