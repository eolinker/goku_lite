package cmd

import (
	"fmt"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"net/http"

	"time"
)

//HeartBeatProid heartBeatProid
const HeartBeatProid = time.Second * 5

var (
	serverPort = 0
	closeChan  chan bool
)

//Heartbeat heartBeat
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

	addr := node_common.GetAdminURL(fmt.Sprintf("node/heartbeat?port=%d", port))
	_, err := http.Get(addr)
	if err != nil {
		log.Warn("fail to send heartbeat:", err)
	}
}

//StopNode stopNode
func StopNode() {
	close(closeChan)

	log.Debug("stop node")
	addr := node_common.GetAdminURL(fmt.Sprintf("node/stop?port=%d", serverPort))
	_, err := http.Get(addr)
	if err != nil {
		log.Warn("fail to send heartbeat:", err)
	}
}
