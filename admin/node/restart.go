package node

import (
	goku_log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/admin/cmd"
)

func Restart(code cmd.Code, data []byte) error {
	goku_log.Info("restart")
	//endless.RestartServer()

	return nil
}

func Stop(code cmd.Code, data []byte) error {
	goku_log.Info("stop")
	//endless.StopServer()

	return nil
}
