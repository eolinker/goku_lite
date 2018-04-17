package utils

import (
	"goku-ce-1.0/conf"
	"flag"
	"log"
)

func init() {
	ParseArgs()
}

func ParseArgs() {
	var confFilepath string = "configure.json"
	flag.StringVar(&confFilepath, "c", "configure.json", "Please provide a valid configuration file path")
	flag.Parse()

	err := conf.ReadConfigure(confFilepath)
	if err != nil && confFilepath != "configure.json"{
		log.Fatalln("[error]: Not a valid configuration file, check if the file exists and the validation inside")
	}
}
