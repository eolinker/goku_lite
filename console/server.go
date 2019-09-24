package console

import (
	"net/http"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	_ "net/http/pprof"

	"github.com/eolinker/goku-api-gateway/console/controller/api"
	"github.com/eolinker/goku-api-gateway/console/controller/script"
	"github.com/eolinker/goku-api-gateway/console/controller/strategy"
	cluster2 "github.com/eolinker/goku-api-gateway/server/cluster"
	monitor_read "github.com/eolinker/goku-api-gateway/server/monitor/monitor-read"

	"github.com/eolinker/goku-api-gateway/common/conf"
	"github.com/eolinker/goku-api-gateway/console/admin"
	"github.com/eolinker/goku-api-gateway/console/module/account"
)

func Server() {
	//go monitor.MonitorNode()
	monitor_read.InitMonitorRead(cluster2.GetList())
	script.UpdateTables()
	api.UpdateAllApiPluginUpdateTag()
	strategy.UpdateAllStrategyPluginUpdateTag()
	bind, has := conf.Get("admin_bind")

	ec := make(chan error, 2)
	if has {
		go func() {
			err := admin.StartServer(bind)
			ec <- err
		}()
	} else {
		log.Panic("[ERROR] Illegal admin_bind!")
	}

	port, has := conf.Get("listen_port")
	if has {
		go func() {
			log.Print("Listen: ", port)
			log.Print("Start Successfully!")
			err := http.ListenAndServe(":"+port, nil)

			ec <- err
		}()
	} else {
		log.Panic("[ERROR] Illegal listen port!")
	}
	// go func() {
	// 	err := http.ListenAndServe(":6060", nil)
	// 	ec <- err
	// }()

	for {
		select {
		case e, ok := <-ec:
			if !ok {
				break
			}
			log.Fatal(e)
		}

	}

}

// 用户注册
func Register(loginCall, loginPassword string) bool {
	return account.Register(loginCall, loginPassword)
}
