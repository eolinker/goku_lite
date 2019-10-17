package main

import "flag"

//ParseFlag 获取命令行参数
func ParseFlag() (port int, admin string, staticConfigFile string, isDebug bool) {
	adminP := flag.String("admin", "", "Please provide a valid host!")
	portP := flag.Int("port", 0, "Please provide a valid listen port!")
	staticConfigFileP := flag.String("config", "", "Please provide a config file")

	isDebugP := flag.Bool("debug", false, "")

	flag.Parse()

	return *portP, *adminP, *staticConfigFileP, *isDebugP

}
