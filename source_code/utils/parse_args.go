package utils

import (
	"goku-ce/conf"
	"flag"
	"log"
)

func init() {
	ParseArgs()
}

func ParseArgs() {
	var confFilepath string = "goku.conf"
	flag.StringVar(&confFilepath, "c", "configure.conf", "Please provide a valid configuration file path")
	flag.Parse()

	err := conf.ReadConfigure(confFilepath)
	if err != nil && confFilepath != "goku.conf"{
		log.Fatalln("[error]: Not a valid configuration file, check if the file exists and the validation inside")
	}
}
