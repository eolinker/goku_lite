package server

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	//"github.com/eolinker/goku-api-gateway/common/endless"

	"github.com/eolinker/goku-api-gateway/goku-service/application"

	"github.com/eolinker/goku-api-gateway/node/routerRule"

	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	"github.com/eolinker/goku-api-gateway/node"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/eolinker/goku-api-gateway/module"
	"github.com/eolinker/goku-api-gateway/node/admin"
	"github.com/eolinker/goku-api-gateway/node/monitor"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"

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
	router http.Handler
}

//NewServer newServer
func NewServer() *Server {
	return &Server{
		//port:    port,
		//console: nil,
		router: nil,
	}
}

//SetRouter setRouter
func (s *Server) SetRouter(r http.Handler) error {
	s.router = r
	return nil
}

//ServerWidthConsole 开启节点监听服务
func (s *Server) ServerWidthConsole(console console.ConfigConsole) error {
	if console == nil {
		return errors.New("can not start server widthout router and console")
	}

	if console != nil {

		conf, err := console.RegisterToConsole()
		if err != nil {
			return err
		}

		console.AddListen(s.FlushRouter)
		console.AddListen(s.FlushModule)
		console.AddListen(s.FlushRedisConfig)
		console.AddListen(s.FlushRouterRule)
		console.AddListen(s.FlushGatewayBasicConfig)

		console.Listen()

		return s.ServerWidthConfig(conf)
	}
	return errors.New("can not start server widthout router and console")
}

//ServerWidthConfig 处理配置
func (s *Server) ServerWidthConfig(conf *config.GokuConfig) error {

	if conf == nil {
		return errors.New("can not start server width out config")
	}
	s.FlushRedisConfig(conf)

	r, err := gateway.Parse(conf, httprouter.Factory())
	if err != nil {
		log.Panic("parse config error:", err)
	}
	if conf.GatewayBasicInfo != nil {
		application.SetSkipCertificate(conf.GatewayBasicInfo.SkipCertificate)
	}
	e := s.SetRouter(r)
	if e != nil {
		return e
	}
	routerRule.Load(conf.Routers)

	// 初始化监控模块
	monitor.Init(conf.Cluster, conf.Instance)

	s.FlushModule(conf)

	if conf.BindAddress == "" {
		log.Panic("invalid bind address")
	}

	// 启用管理接口
	if conf.AdminAddress != "" {
		StartAdmin(conf.AdminAddress)
	}

	//return endless.ListenAndServe(conf.BindAddress, s)
	return http.ListenAndServe(conf.BindAddress, s)
}

//FlushRouter flushConfig
func (s *Server) FlushRouter(config *config.GokuConfig) {
	r, err := gateway.Parse(config, httprouter.Factory())
	if err != nil {
		log.Error("parse config error:", err)
		return
	}
	_ = s.SetRouter(r)
}

//FlushRouterRule flushConfig
func (s *Server) FlushRouterRule(config *config.GokuConfig) {
	routerRule.Load(config.Routers)
}

//FlushGatewayBasicConfig 刷新网关基础配置
func (s *Server) FlushGatewayBasicConfig(config *config.GokuConfig) {
	if config.GatewayBasicInfo != nil {
		application.SetSkipCertificate(config.GatewayBasicInfo.SkipCertificate)
	}
}

//FlushRedisConfig 刷新redis配置
func (s *Server) FlushRedisConfig(config *config.GokuConfig) {
	if r, ok := config.ExtendsConfig["redis"]; ok {
		if r == nil {
			return
		}
		rr := r.(map[string]interface{})
		rdsConfig := &entity.ClusterRedis{
			Addrs:    rr["addrs"].(string),
			DbIndex:  int(rr["dbIndex"].(float64)),
			Masters:  rr["masters"].(string),
			Mode:     rr["mode"].(string),
			Password: rr["password"].(string),
		}
		rds := redis_manager.Create(rdsConfig)
		log.Info(rdsConfig)
		redis_manager.SetDefault(rds)
		node.InitPluginUtils()
	}
}

//FlushModule 刷新模块配置
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			print(err)
			debug.PrintStack()
		}
	}()
	if s.router == nil {
		w.WriteHeader(404)
		return
	}

	s.router.ServeHTTP(w, req)

}
