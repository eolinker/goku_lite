package main

import (
	"flag"
	log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/common/conf"
	"github.com/eolinker/goku-api-gateway/common/general"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/utils"
)

var (
	userPassword string
	userName     string
	confFilePath = "./config/goku.conf"
)

func main() {
	flag.StringVar(&confFilePath, "c", "./config/goku.conf", "Please provide a valid configuration file path")
	flag.StringVar(&userName, "u", "", "Please provide user name")
	flag.StringVar(&userPassword, "p", "", "Please provide user password")
	isDebug := flag.Bool("debug", false, "")

	flag.Parse()
	if *isDebug {
		log.StartDebug()
	}
	// 初始化配置
	if err := conf.ReadConfigure(confFilePath); err != nil {
		log.Panic(err)
		return
	}


	// 初始化db
	InitDatabase()
	InitLog()

	// 其他需要初始化的模块
	_ = general.General()
	// 检测是否安装
	s, err := account.CheckSuperAdminCount()
	if err != nil {
			log.Panic(err)
			return

	}
	if s == 0 {
		if userName == "" {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			return
		}
		if userPassword == "" {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")

			return
		}

		// 用户注册
		password := utils.Md5(utils.Md5(userPassword))
		f := account.Register(userName, password)
		if !f {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			return
		}
	}
	Server()
	//console.Router()
	//console.Server()
}
