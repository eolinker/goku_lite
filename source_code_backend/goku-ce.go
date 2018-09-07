package main

import (
	"fmt"
	"goku-ce/goku"
	"goku-ce/middleware"
	"goku-ce/utils"
	"log"
	"os"
)

func main() {
	utils.ParseArgs()
	server := goku.New()
	server.RegisterRouter(server.ServiceConfig, middleware.Mapping, middleware.GetVisitCount)
	fmt.Println("Listen", server.ServiceConfig.Port)
	err := goku.ListenAndServe(":"+server.ServiceConfig.Port, server)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on " + server.ServiceConfig.Port + " stopped")
	os.Exit(0)
}
