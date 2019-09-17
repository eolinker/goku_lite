package main

import (
	"flag"
	"github.com/eolinker/goku/console/module/account"
	log "github.com/eolinker/goku/goku-log"

	"github.com/eolinker/goku/common/conf"
	"github.com/eolinker/goku/common/general"
	"github.com/eolinker/goku/console"
	"github.com/eolinker/goku/utils"
)

var (
	UserPassword string
	UserName     string
	ConfFilePath = "./config/goku.conf"

)

func main() {
	flag.StringVar(&ConfFilePath, "c", "./config/goku.conf", "Please provide a valid configuration file path")
	flag.StringVar(&UserName, "u", "", "Please provide user name")
	flag.StringVar(&UserPassword, "p", "", "Please provide user password")
	isDebug := flag.Bool("debug",false,"")

	flag.Parse()
	if *isDebug{
		log.StartDebug()
	}
	// 初始化配置
	if err := conf.ReadConfigure(ConfFilePath); err != nil {
		log.Panic(err)
		return
	}
	// 初始化db
	console.InitDatabase()
	console.InitLog()

	console.InitClusters()
	// 其他需要初始化的模块
	_ = general.General()
	// 检测是否安装


	if s, err := account.CheckSuperAdminCount(); err!= nil {
		log.Panic(err)
		return
	}else if s == 0 {
		if UserName == "" {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			//fmt.Println("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			return
		}
		if UserPassword == "" {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			//fmt.Println("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			return
		}

		// 用户注册
		password := utils.Md5(utils.Md5(UserPassword))
		f := console.Register(UserName, password)
		if !f {
			log.Fatal("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			//fmt.Println("[ERROR] Fail to create administrator. Please try again or contact technical support of eoLinker GOKU API Gateway.")
			return
		}
	}


	console.Router()
	console.Server()
}
