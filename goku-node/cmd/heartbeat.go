package cmd

import (
	"fmt"
	log "github.com/eolinker/goku/goku-log"
	node_common "github.com/eolinker/goku/goku-node/node-common"
	"net/http"

	"time"
)

const HeartBeatProid = time.Second * 5

var (
	serverPort = 0
	closeChan  chan bool
)

func Heartbeat(port int) {
	closeChan = make(chan bool)
	serverPort = port
	tick := time.NewTicker(HeartBeatProid)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			{
				sendHeartBeat(port)
			}

		case <-closeChan:
			{
				return
			}
		}
	}
}

func sendHeartBeat(port int) {

	addr := node_common.GetAdminUrl(fmt.Sprintf("node/heartbeat?port=%d", port))
	_, err := http.Get(addr)
	if err != nil {
		log.Warn("fail to send heartbeat:", err)
	}
}

func StopNode() {
	close(closeChan)

	log.Debug("stop node")
	addr := node_common.GetAdminUrl(fmt.Sprintf("node/stop?port=%d", serverPort))
	_, err := http.Get(addr)
	if err != nil {
		log.Warn("fail to send heartbeat:", err)
	}
}
