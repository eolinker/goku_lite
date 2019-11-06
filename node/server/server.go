package server

import (
	"errors"
	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/eolinker/goku-api-gateway/module"
	"github.com/eolinker/goku-api-gateway/node/admin"
	"github.com/eolinker/goku-api-gateway/node/monitor"
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
	//port    int
	//console *console.Console
	router  http.Handler
}

//NewServer newServer
func NewServer() *Server {
	return &Server{
		//port:    port,
		//console: nil,
		router:  nil,
	}
}

//SetRouter setRouter
func (s *Server) SetRouter(r http.Handler) error {
	s.router = r
	return nil
}


//Server server
func (s *Server) ServerWidthConsole(console *console.Console ) error {
	if  console == nil {
		return errors.New("can not start server widthout router and console")
	}

	if console != nil {

		conf, err := console.GetConfig()
		if err != nil {
			return err
		}

		console.AddListen(s.FlushRouter)
		console.AddListen(s.FlushModule)
		return 	s.ServerWidthConfig(conf)
	}
	return  errors.New("can not start server widthout router and console")
}

func (s *Server)ServerWidthConfig(conf *config.GokuConfig)error  {

	if conf == nil{
		return errors.New("can not start server width out config")
	}


	r, err := gateway.Parse(conf, httprouter.Factory())
	if err != nil {
		log.Panic("parse config error:", err)
	}
	e := s.SetRouter(r)
	if e != nil {
		return e
	}
	// 初始化监控模块
	monitor.Init(conf.Cluster,conf.Instance)

	s.FlushModule(conf)

	if conf.BindAddress == ""{
		log.Panic("invalid bind address")
	}
	//if conf.AdminAddress == ""{
	//	log.Panic("invalid admin address")
	//}
	// 启用管理接口
	if conf.AdminAddress != ""{
		StartAdmin(conf.AdminAddress)
	}

	return endless.ListenAndServe(conf.BindAddress, s)
}
//FlushRouter flushConfig
func (s *Server) FlushRouter(config *config.GokuConfig) {


		r, err := gateway.Parse(config, httprouter.Factory())
		if err != nil {
			log.Error("parse config error:", err)
			return
		}
		_=s.SetRouter(r)
}
//FlushRouter flushConfig
func (s *Server) FlushModule(conf *config.GokuConfig) {
	SetLog(conf.Log)
	SetAccessLog(conf.AccessLog)
	module.Refresh(nil)

	//demo:= map[string]string{
	//	"diting.prometheus":"",
	//}
	diting.Refresh(conf.MonitorModules)

	admin.Refresh()

}
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if s.router == nil {
		w.WriteHeader(404)
		return
	}

	s.router.ServeHTTP(w, req)

}
