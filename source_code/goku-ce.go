package main

import (
	_ "goku-ce/utils"
	"goku-ce/goku"
	"goku-ce/middleware"
)

func main() {
	server := goku.New()
	server.RegisterRouter(server.ServiceConfig,middleware.Mapping)
	server.Listen()
	server.Run()
}


