package main

import (
	"log"
	"net/http"

	"github.com/eolinker/goku-api-gateway/admin/console"
	"github.com/eolinker/goku-api-gateway/common/conf"
)

//Server 控制台服务
func Server() {
	moduleRegister()
	bind, has := conf.Get("admin_bind")
	if !has {
		log.Panic("[ERROR] Illegal admin_bind!")
		return
	}
	err := console.Start(bind)
	if err != nil {
		log.Fatal(err)
		return
	}

	ec := make(chan error, 1)

	port, has := conf.Get("listen_port")
	if has {
		go func() {
			log.Print("Listen: ", port)
			log.Print("Start Successfully!")
			err := http.ListenAndServe(":7000", router())

			ec <- err
		}()
	} else {
		log.Panic("[ERROR] Illegal listen port!")
	}

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
