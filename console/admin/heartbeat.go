package admin

import (
	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/node"
	"net/http"
	"strconv"
)

func heartbead(w http.ResponseWriter, r *http.Request)  {

	ip, port, err := GetIpPort(r)
	if err != nil {

		controller.WriteError(w, "700000", "node", err.Error(), err)
		 return
	}
	node.Refresh(ip,strconv.Itoa(port))
	controller.WriteResultInfo(w, "node", "node", nil)
}


func stopNode(w http.ResponseWriter, r *http.Request)  {


	ip, port, err := GetIpPort(r)
	if err != nil {

		controller.WriteError(w, "700000", "node", err.Error(), err)
		return
	}
	node.NodeStop(ip,strconv.Itoa(port))
	controller.WriteResultInfo(w, "node", "node", nil)
}