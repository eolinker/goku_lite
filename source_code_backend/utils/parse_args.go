package utils

import (
	"fmt"
	"goku-ce/conf"
	"flag"
	"log"
)

func init() {
	ParseArgs()
}

var ConfFilepath string = "./config/goku.conf"
var command string = "start"

func ParseArgs() {
	flag.StringVar(&command,"s","start","send `signal` to a master process: stop, start, restart, reload")
	flag.StringVar(&ConfFilepath, "c", "./config/goku.conf", "Please provide a valid configuration file path")
	flag.Parse()

	err := conf.ReadConfigure(ConfFilepath)
	if err != nil && ConfFilepath != "./config/goku.conf"{
		log.Fatalln("[error]: Not a valid configuration file, check if the file exists and the validation inside")
	}
}

func ReloadConf() {
	fmt.Println(ConfFilepath)
	err := conf.ReadConfigure(ConfFilepath)
	if err != nil && ConfFilepath != "./config/goku.conf"{
		log.Fatalln("[error]: Not a valid configuration file, check if the file exists and the validation inside")
	}
}
