package server

import (
	"log"
	"net/http"
	"sync"

	//"github.com/eolinker/goku-api-gateway/common/endless"

	"github.com/eolinker/goku-api-gateway/node/admin"
)

var (
	adminOnce = sync.Once{}
)

//StartAdmin 启动节点管理端
func StartAdmin(address string) {
	go adminOnce.Do(func() {
		//server := endless.NewServer(address, admin.Handler())
		//endless.SetAdminServer(server)
		//err := server.ListenAndServe()
		err := http.ListenAndServe(address, admin.Handler())
		if err != nil {
			log.Fatal(err)
		}
	})
}
