package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	"github.com/eolinker/goku-api-gateway/server/entity"
	"io/ioutil"
	"net/http"
	"time"
)

// 获取节点配置
func GetConfig(listenPort int) (bool, *entity.ClusterInfo) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	addr := node_common.GetAdminUrl(fmt.Sprintf("/register?port=%d", listenPort))
	reader := bytes.NewReader([]byte(""))
	request, err := http.NewRequest("GET", addr, reader)
	if err != nil {
		return false, nil
	}
	// request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(request)
	if err != nil {
		log.Info(err)
		return false, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		errReason := "[Error] Fail to match code!"
		log.Info(errReason)
		return false, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info(err)
		return false, nil
	}
	content := new(ClusterConfig)
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Info(err)
		return false, nil
	}
	if content.StatusCode != "000000" {
		errReason := "[Error] Fail to match status code!"
		log.Info(errReason)
		return false, nil
	}
	return true, content.Cluster
}
