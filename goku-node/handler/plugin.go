package handler

import (
	"encoding/json"
	log "github.com/eolinker/goku/goku-log"

	plugin_manager "github.com/eolinker/goku/goku-node/manager/plugin-manager"
	"net/http"
)

func gokuCheckPlugin(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	pluginName := req.PostFormValue("pluginName")

	code, e := plugin_manager.Check(pluginName)

	if code != plugin_manager.LoadOk {
		switch code {
		case plugin_manager.LoadFileError:
			log.Info(e)
			result := map[string]string{
				"statusCode": "210015",
				"type":       "plugin",
				"resultDesc": "[GOKU] " + pluginName + ".so can not be found in plugin catalog",
			}
			resultByte, _ := json.Marshal(result)
			w.Write(resultByte)
			return
		case plugin_manager.LoadLookupError:
			log.Info(e)
			result := map[string]string{
				"statusCode": "210017",
				"type":       "plugin",
				"resultDesc": "[GOKU] Object named '" + pluginName + "' can not be found in " + pluginName + ".so",
			}
			resultByte, _ := json.Marshal(result)
			w.Write(resultByte)
		case plugin_manager.LoadInterFaceError:

			result := map[string]string{
				"statusCode": "210016",
				"type":       "plugin",
				"resultDesc": "[GOKU] Object named '" + pluginName + "' did not inherit necessary methods from utils.PluginHandle",
			}
			log.Info(result["resultDesc"])
			resultByte, _ := json.Marshal(result)
			w.Write(resultByte)

		}
		return
	}
	result := map[string]string{
		"statusCode": "000000",
		"type":       "plugin",
		"resultDesc": "[GOKU]Success",
	}
	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)

	return
}
