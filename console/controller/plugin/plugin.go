package plugin

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/node"
	"github.com/eolinker/goku-api-gateway/console/module/plugin"
	plugin_config "github.com/eolinker/goku-api-gateway/console/module/plugin/plugin-config"
	"github.com/eolinker/goku-api-gateway/utils"
)

// GetPluginList 获取插件列表
func GetPluginList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationREAD)
	if e != nil {
		return
	}
	httpRequest.ParseForm()
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	op, err := strconv.Atoi(condition)
	if err != nil && condition != "" {
		controller.WriteError(httpResponse, "210006", "plugin", "[ERROR]Illegal condition!", err)
		return
	}
	flag, result, err := plugin.GetPluginList(keyword, op)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]Empty plugin list!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "pluginList", result)
	return
}

//AddPlugin 新增插件信息
func AddPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginPriority := httpRequest.PostFormValue("pluginPriority")
	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")
	pluginDesc := httpRequest.PostFormValue("pluginDesc")
	isStop := httpRequest.PostFormValue("isStop")
	pluginType := httpRequest.PostFormValue("pluginType")
	version := httpRequest.PostFormValue("version")
	index, err := strconv.Atoi(pluginPriority)

	if err != nil {
		controller.WriteError(httpResponse, "210001", "plugin", "[ERROR]Illegal pluginPriority!", err)
		return
	}

	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "plugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}

	pStop, err := strconv.Atoi(isStop)
	if err != nil {

		controller.WriteError(httpResponse, "210005", "plugin", "[ERROR]Illegal isStop!", err)
		return
	}

	pType, err := strconv.Atoi(pluginType)
	if err != nil {

		controller.WriteError(httpResponse, "210009", "plugin", "[ERROR]Illegal pluginType!", err)
		return
	}

	exits, err := plugin.CheckIndexIsExist("", index)
	if exits {

		controller.WriteError(httpResponse, "210003", "plugin", "[ERROR]Plugin pluginPriority is existed!", err)
		return
	}

	exits, err = plugin.CheckNameIsExist(pluginName)
	if exits {

		controller.WriteError(httpResponse, "210004", "plugin", "[ERROR]Plugin name is existed!", err)
		return
	}

	flag, result, err := plugin.AddPlugin(pluginName, pluginConfig, pluginDesc, version, index, pStop, pType)
	if !flag {
		controller.WriteError(httpResponse, "210000", "plugin", result, err)
		return
	}

	controller.WriteResultInfo(httpResponse, "plugin", "", nil)

	return
}

//EditPlugin 修改插件信息
func EditPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginPriority := httpRequest.PostFormValue("pluginPriority")
	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")

	pluginDesc := httpRequest.PostFormValue("pluginDesc")
	isStop := httpRequest.PostFormValue("isStop")
	pluginType := httpRequest.PostFormValue("pluginType")
	version := httpRequest.PostFormValue("version")

	index, err := strconv.Atoi(pluginPriority)
	if err != nil {
		controller.WriteError(httpResponse, "210001", "plugin", "[ERROR]Illegal pluginPriority!", err)
		return
	}
	pStop, err := strconv.Atoi(isStop)
	if err != nil {
		controller.WriteError(httpResponse, "210005", "plugin", "[ERROR]Illegal isStop!", err)
		return
	}
	pType, err := strconv.Atoi(pluginType)
	if err != nil {
		controller.WriteError(httpResponse, "210009", "plugin", "[ERROR]Illegal pluginType!", err)
		return
	}

	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "plugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}
	exits, err := plugin.CheckIndexIsExist(pluginName, index)
	if exits {
		controller.WriteError(httpResponse, "210003", "plugin", "[ERROR]Plugin priority is existed!", err)
		return
	}
	flag, err = plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse, "210004", "plugin", "[ERROR]Plugin name does not exist!", err)
		return
	}
	flag, result, err := plugin.EditPlugin(pluginName, pluginConfig, pluginDesc, version, index, pStop, pType)
	if !flag {
		controller.WriteError(httpResponse, "210000", "plugin", result, err)
		return
	}

	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
}

//DeletePlugin 删除插件信息
func DeletePlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")
	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"210004",
			"plugin",
			"[ERROR]Plugin name does not exist!",
			err)
		return

	}
	flag, result, err := plugin.DeletePlugin(pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
	return
}

//GetPluginInfo 获取插件信息
func GetPluginInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationREAD)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")

	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"210004",
			"plugin",
			"[ERROR]Plugin name does not exist!",
			err)
		return

	}
	flag, result, err := plugin.GetPluginInfo(pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]The plugin does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "pluginInfo", result)
}

//GetPluginConfig 获取插件配置
func GetPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationREAD)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")
	flag, result, err := plugin.GetPluginConfig(pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]The plugin does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "pluginConfig", result)

	return
}

//CheckIndexIsExist 判断插件优先级是否存在
func CheckIndexIsExist(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginPriority := httpRequest.PostFormValue("pluginPriority")

	index, err := strconv.Atoi(pluginPriority)
	if err != nil {
		controller.WriteError(httpResponse,
			"210001",
			"plugin",
			"[ERROR]Illegal pluginPriority!",
			err)
		return

	}
	flag, err := plugin.CheckIndexIsExist("", index)
	if !flag {

		controller.WriteError(httpResponse,
			"210011",
			"plugin",
			"[ERROR]Plugin pluginPriority does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
	return
}

//CheckNameIsExist 检查插件名称是否存在
func CheckNameIsExist(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")

	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]Plugin name does not exist!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)

}

//StartPlugin 开启插件
func StartPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")

	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"210010",
			"plugin",
			"[ERROR]Plugin name does not exist!",
			err)
		return

	}
	_, _ = plugin.EditPluginStatus(pluginName, 1)
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
}

//StopPlugin 关闭插件
func StopPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")

	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"210010",
			"plugin",
			"[ERROR]Plugin name does not exist!",
			err)
		return
	}
	_, _ = plugin.EditPluginStatus(pluginName, 0)
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
}

//GetPluginListByPluginType 获取不同类型的插件列表
func GetPluginListByPluginType(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationREAD)
	if e != nil {
		return
	}

	pluginType := httpRequest.PostFormValue("pluginType")

	pt, err := strconv.Atoi(pluginType)
	if err != nil {
		controller.WriteError(httpResponse,
			"210009",
			"plugin",
			"[ERROR]Illegal pluginType",
			err)
		return

	}
	flag, result, err := plugin.GetPluginListByPluginType(pt)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]Empty plugin list",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "pluginList", result)
}

//BatchStopPlugin 批量关闭插件
func BatchStopPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginNameList := httpRequest.PostFormValue("pluginNameList")

	flag, result, err := plugin.BatchStopPlugin(pluginNameList)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)
	return
}

//BatchStartPlugin 批量关闭插件
func BatchStartPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginNameList := httpRequest.PostFormValue("pluginNameList")

	flag, result, err := plugin.BatchStartPlugin(pluginNameList)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "plugin", "", nil)

}

//CheckPluginIsAvailable 检测插件
func CheckPluginIsAvailable(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationPlugin, controller.OperationEDIT)
	if e != nil {
		return
	}

	pluginName := httpRequest.PostFormValue("pluginName")
	if len(pluginName) < 2 {
		controller.WriteError(httpResponse,
			"210002",
			"plugin",
			"[ERROR]Illegal pluginName",
			nil)
		return
	}
	flag, nodeList, err := node.GetNodeIPList()
	if !flag {
		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			"[ERROR]The Open node list is empty",
			err)
		return
	}
	flag, errPluginList := utils.CheckPluginIsAvailiable(pluginName, nodeList)
	if !flag {

		controller.WriteResultInfoWithCode(httpResponse,
			"210000",
			"plugin",
			"errNodeList",
			errPluginList)
		return
	}
	flag, res, _ := plugin.EditPluginCheckStatus(pluginName, 1)
	if !flag {

		controller.WriteError(httpResponse,
			"210000",
			"plugin",
			res,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "plugin", "", nil)

	return
}
