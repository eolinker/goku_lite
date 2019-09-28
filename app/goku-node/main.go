package main

import (
	"encoding/json"
	"flag"
	"fmt"
	endless2 "github.com/eolinker/goku-api-gateway/common/endless"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/cmd"
	"runtime"

	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"

	"github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/common/general"
	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	goku_node "github.com/eolinker/goku-api-gateway/goku-node"
	"github.com/eolinker/goku-api-gateway/server/entity"
)

var (
	adminHost string
	//adminPort  int
	listenPort int
)

func initConfig(resultInfo map[string]interface{}) *entity.ClusterInfo {
	c := entity.ClusterInfo{}
	clusterConfig, err := json.Marshal(resultInfo["cluster"])
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(clusterConfig, &c)
	if err != nil {
		log.Panic(err)
	}
	return &c
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.StringVar(&adminHost, "admin", "127.0.0.1:7005", "Please provide a valid host!")
	//flag.IntVar(&adminPort, "P", 7005, "Please provide a valid port")
	flag.IntVar(&listenPort, "port", 6689, "Please provide a valid listen port!")
	isDebug := flag.Bool("debug", false, "")

	flag.Parse()
	if *isDebug {
		log.StartDebug()
	}
	//
	node_common.SetAdmin(adminHost)
	node_common.ListenPort = listenPort

	success, config := cmd.GetConfig(listenPort)
	if !success {
		log.Fatal(" Fail to get node config!")
		return
	}

	node_common.SetClusterName(config.Name)

	err := database.InitConnection(&config.DB)
	if err != nil {
		log.Fatal("Fail to Init db:", err)
		return
	}
	goku_node.InitLog()
	log.Debug("gokNode.InitLog")
	r := redis_manager.Create(&config.Redis)
	redis_manager.SetDefault(r)

	log.Debug("redis-manager.SetDefault")
	// 其他需要初始化的模块
	_ = general.General()

	log.Debug("general.General()")
	goku_node.InitServer()
	go cmd.Heartbeat(listenPort)
	server := goku_node.NewRouter()

	err = endless2.ListenAndServe(fmt.Sprintf(":%d", listenPort), server)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatalf("Server on :%d stoped \n", listenPort)

}
