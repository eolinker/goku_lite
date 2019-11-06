package console

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/updater"

	"github.com/eolinker/goku-api-gateway/common/conf"
	"github.com/eolinker/goku-api-gateway/console/admin"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	log "github.com/eolinker/goku-api-gateway/goku-log"
)

//Server 服务
func Server() {
	bind, has := conf.Get("admin_bind")
	moduleRegister()
	updater.InitUpdater()
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

//Register 用户注册
func Register(loginCall, loginPassword string) bool {
	return account.Register(loginCall, loginPassword)
}
