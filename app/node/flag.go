package main

import "flag"

//ParseFlag 获取命令行参数
func ParseFlag() (instance string, admin string, staticConfigFile string, isDebug bool) {
	adminP := flag.String("admin", "", "Please provide a valid host!")
	instanceP := flag.String("instance", "", "Please provide a valid instance!")
	staticConfigFileP := flag.String("config", "", "Please provide a config file")

	isDebugP := flag.Bool("debug", false, "")

	flag.Parse()

	return *instanceP, *adminP, *staticConfigFileP, *isDebugP

}
