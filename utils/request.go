package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//CheckPluginIsAvailiable 检查插件是否可用
func CheckPluginIsAvailiable(pluginName string, nodeList []map[string]interface{}) (bool, []map[string]interface{}) {
	errNodeList := make([]map[string]interface{}, 0)
	for _, v := range nodeList {
		if v["nodePort"] == "" {
			v["nodePort"] = "6689"
		}
		client := &http.Client{
			Timeout: time.Second * 15,
		}
		var data = url.Values{}
		data.Add("pluginName", pluginName)

		request, err := http.NewRequest("POST", "http://"+v["nodeIP"].(string)+":"+v["nodePort"].(string)+"/goku-check_plugin", strings.NewReader(data.Encode()))
		if err != nil {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           "[ERROR] Fail to create request",
				"errorStatusCode": "210014",
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(request)
		if err != nil {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           "[ERROR] Connect timeout",
				"errorStatusCode": "210011",
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           "[ERROR] Fail to read body",
				"errorStatusCode": "210013",
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
		resp.Body.Close()
		var bodyJSON map[string]string
		err = json.Unmarshal(body, &bodyJSON)
		if err != nil {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           "[ERROR] Fail to parse json",
				"errorStatusCode": "210010",
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
		if _, ok := bodyJSON["statusCode"]; !ok {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           "[ERROR] Fail to get statusCode",
				"errorStatusCode": "210012",
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
		if bodyJSON["statusCode"] != "000000" {
			errNode := map[string]interface{}{
				"nodeAddress":     v["nodeIP"].(string) + ":" + v["nodePort"].(string),
				"error":           bodyJSON["resultDesc"],
				"errorStatusCode": bodyJSON["statusCode"],
			}
			errNodeList = append(errNodeList, errNode)
			continue
		}
	}
	if len(errNodeList) > 0 {
		return false, errNodeList
	}
	return true, errNodeList
}
