package main

import (
	"flag"
	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	console2 "github.com/eolinker/goku-api-gateway/node/console"
	"github.com/eolinker/goku-api-gateway/node/server"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	instance, admin, staticConfigFile, isDebug := ParseFlag()

	if isDebug {
		log.StartDebug()
	}

	if admin != "" && instance != ""{

		console := console2.NewConsole(instance, admin)
		ser := server.NewServer()
		log.Fatal(ser.ServerWidthConsole(console))
		return
	}

 	if staticConfigFile != "" {

		// 从静态文件启动
		c, err := config.ReadConfig(staticConfigFile)
		if err != nil {
			log.Panic("read config from :", staticConfigFile, "\t", err)
		}
		ser := server.NewServer()
		log.Fatal(ser.ServerWidthConfig(c))
	}

	flag.Usage()
}
