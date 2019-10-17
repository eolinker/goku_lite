package main

import (
	"flag"
	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	console2 "github.com/eolinker/goku-api-gateway/node/console"
	"github.com/eolinker/goku-api-gateway/node/gateway"
	"github.com/eolinker/goku-api-gateway/node/router/httprouter"
	"github.com/eolinker/goku-api-gateway/node/server"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	port, admin, staticConfigFile, isDebug := ParseFlag()

	if isDebug {
		log.StartDebug()
	}

	if port == 0{
		flag.Usage()

		return
	}

	if  admin != "" {
		// 从控制台启动，
		console := console2.NewConsole(port, admin)
		ser := server.NewServer(port)
		ser.SetConsole(console)
		log.Fatal(ser.Server())

	} else if staticConfigFile != "" {

		// 从静态文件启动
		c, err := config.ReadConfig(staticConfigFile)
		if err != nil {
			log.Panic("read config from :", staticConfigFile, "\t", err)
		}

		server.SetLog(c.Log)
		server.SetAccessLog(c.AccessLog)


		r, err := gateway.Parse(c, httprouter.Factory())
		if err != nil {
			log.Panic("parse config error:", err)
		}



		ser := server.NewServer(port)
		e := ser.SetRouter(r)
		if e != nil {
			log.Panic("init router error:", e)
		}
		log.Fatal(ser.Server())
	} else {
		//
		flag.Usage()
		return
	}
}
