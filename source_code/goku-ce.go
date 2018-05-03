package main

import (
	_ "goku-ce/utils"
	"goku-ce/goku"
	"goku-ce/middleware"
)

func main() {
	server := goku.New()
	server.Use(middleware.Mapping)
	server.Listen()
	server.Run()
}



