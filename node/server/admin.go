package server

import (
	"github.com/eolinker/goku-api-gateway/common/endless"
	"github.com/eolinker/goku-api-gateway/node/admin"
	"log"
	"sync"
)

var (adminOnce = sync.Once{})
func StartAdmin(address string)  {


	go adminOnce.Do(func() {
		err :=endless.ListenAndServe(address,admin.Handler())
		if err!=nil{
			log.Fatal(err)
		}
	})
}
