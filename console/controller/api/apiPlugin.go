package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
	"github.com/eolinker/goku-api-gateway/console/module/plugin"
	plugin_config "github.com/eolinker/goku-api-gateway/console/module/plugin/plugin-config"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

const operationAPIPlugin = "strategyManagement"

//PluginHandlers 插件处理器
type PluginHandlers struct {
}

//Handlers handlers
func (p *PluginHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/addPluginToApi":    factory.NewAccountHandleFunction(operationAPIPlugin, true, AddPluginToAPI),
		"/edit":              factory.NewAccountHandleFunction(operationAPIPlugin, true, EditAPIPluginConfig),
		"/getInfo":           factory.NewAccountHandleFunction(operationAPIPlugin, false, GetAPIPluginConfig),
		"/getList":           factory.NewAccountHandleFunction(operationAPIPlugin, false, GetAPIPluginList),
		"/getListByStrategy": factory.NewAccountHandleFunction(operationAPIPlugin, false, GetAllAPIPluginInStrategy),
		"/batchStart":        factory.NewAccountHandleFunction(operationAPIPlugin, true, BatchStartAPIPlugin),
		"/batchStop":         factory.NewAccountHandleFunction(operationAPIPlugin, true, BatchStopAPIPlugin),
		"/batchDelete":       factory.NewAccountHandleFunction(operationAPIPlugin, true, BatchDeleteAPIPlugin),
		"/notAssign/getList": factory.NewAccountHandleFunction(operationAPIPlugin, false, GetAPIPluginListWithNotAssignAPIList),
	}
}

//NewPluginHandlers new插件处理器
func NewPluginHandlers() *PluginHandlers {
	return &PluginHandlers{}

}

//AddPluginToAPI 新增插件到接口
func AddPluginToAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")
	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse, "240002", "apiPlugin", "[ERROR]Illegal apiID!", err)
		return
	}
	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "apiPlugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}

	// 查询该插件是否存在
	flag, err = plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse, "240005", "apiPlugin", "[ERROR]The plugin does not exist!", err)
		return
	}
	flag, err = api.CheckAPIIsExist(aID)
	if !flag {

		controller.WriteError(httpResponse, "240012", "apiPlugin", "[ERROR]The api does not exist!", err)
		return
	}

	flag, err = strategy.CheckStrategyIsExist(strategyID)
	if !flag {

		controller.WriteError(httpResponse, "240013", "apiPlugin", "[ERROR]The strategy does not exist!", err)
		return
	}

	id := 0
	exist, err := api.CheckPluginIsExistInAPI(strategyID, pluginName, aID)

	if exist {
		flag, resultDesc, err := api.EditAPIPluginConfig(pluginName, pluginConfig, strategyID, aID, userID)
		if !flag {
			controller.WriteError(httpResponse, "240000", "apiPlugin", resultDesc.(string), err)
			return
		}
		id = resultDesc.(int)
	} else {
		flag, resultDesc, err := api.AddPluginToAPI(pluginName, pluginConfig, strategyID, aID, userID)
		if !flag {
			controller.WriteError(httpResponse, "240000", "apiPlugin", resultDesc.(string), err)
			return
		}
		id = resultDesc.(int)
	}

	controller.WriteResultInfo(httpResponse, "apiPlugin", "connId", id)

	return
}

//EditAPIPluginConfig 修改接口插件
func EditAPIPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")
	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")
	userID := goku_handler.UserIDFromRequest(httpRequest)

	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "apiPlugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240002",
			"apiPlugin",
			"[ERROR]Illegal apiID!",
			err)
		return

	}

	// 查询该插件是否存在
	flag, err = plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"200005",
			"apiPlugin",
			"[ERROR]The plugin does not exist!",
			err)
		return

	}
	flag, err = api.CheckAPIIsExist(aID)
	if !flag {
		controller.WriteError(httpResponse,
			"240012",
			"apiPlugin",
			"[ERROR]The api does not exist!",
			err)
		return

	}
	flag, err = strategy.CheckStrategyIsExist(strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240013",
			"apiPlugin",
			"[ERROR]The strategy does not exist!",
			err)
		return

	}
	flag, resultDesc, err := api.EditAPIPluginConfig(pluginName, pluginConfig, strategyID, aID, userID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			resultDesc.(string),
			err)
	}

	controller.WriteResultInfo(httpResponse, "apiPlugin", "", nil)
	return
}

//GetAPIPluginConfig 获取接口插件配置
func GetAPIPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	pluginName := httpRequest.PostFormValue("pluginName")
	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")

	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240002",
			"apiPlugin",
			"[ERROR]Illegal apiID!",
			err)
		return

	}
	flag, result, err := api.GetAPIPluginConfig(aID, strategyID, pluginName)
	if !flag {
		resultByte, _ := json.Marshal(map[string]interface{}{
			"type":            "apiPlugin",
			"statusCode":      "000000",
			"apiPluginConfig": "",
			// "apiName":         result["apiName"],
			// "requestURL":      result["requestURL"],
		})
		httpResponse.Write(resultByte)
		return
	}
	resultByte, _ := json.Marshal(map[string]interface{}{
		"type":            "apiPlugin",
		"statusCode":      "000000",
		"apiPluginConfig": result["pluginConfig"],
		"apiName":         result["apiName"],
		"requestURL":      result["requestURL"],
	})
	httpResponse.Write(resultByte)
	return

}

//GetAPIPluginList 获取接口插件配置
func GetAPIPluginList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240002",
			"apiPlugin",
			"[ERROR]Illegal apiID!",
			err)
		return

	}
	flag, result, err := api.GetAPIPluginList(aID, strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			"[ERROR]Empty plugin list!",
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "apiPlugin", "apiPluginList", result)
	return
}

//GetAllAPIPluginInStrategy 获取策略组中所有接口插件列表
func GetAllAPIPluginInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")

	flag, result, err := api.GetAllAPIPluginInStrategy(strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			"[ERROR]Empty api plugin list!",
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "apiPlugin", "apiPluginList", result)
	return
}

//BatchStartAPIPlugin 批量修改策略组插件状态
func BatchStartAPIPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")
	userID := goku_handler.UserIDFromRequest(httpRequest)

	if connIDList == "" {
		controller.WriteError(httpResponse,
			"240001",
			"apiPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return

	}
	flag, result, err := api.BatchEditAPIPluginStatus(connIDList, strategyID, 1, userID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "apiPlugin", "", nil)
	return
}

//BatchStopAPIPlugin 批量修改策略组插件状态
func BatchStopAPIPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")
	userID := goku_handler.UserIDFromRequest(httpRequest)

	if connIDList == "" {
		controller.WriteError(httpResponse,
			"240001",
			"apiPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return

	}
	flag, _, err := api.BatchEditAPIPluginStatus(connIDList, strategyID, 0, userID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			"[ERROR]The api plugin is stoped!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "apiPlugin", "", nil)
	return
}

//BatchDeleteAPIPlugin 批量删除策略组插件
func BatchDeleteAPIPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")

	if connIDList == "" {
		controller.WriteError(httpResponse,
			"240001",
			"apiPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return

	}
	flag, result, err := api.BatchDeleteAPIPlugin(connIDList, strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			result,
			err)
	}
	controller.WriteResultInfo(httpResponse, "apiPlugin", "", nil)
}

//GetAPIPluginListWithNotAssignAPIList 获取没有分配接口插件的接口列表
func GetAPIPluginListWithNotAssignAPIList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	if strategyID == "" {
		controller.WriteError(httpResponse,
			"240003",
			"apiPlugin",
			"[ERROR]Illegal strategyID",
			nil)
		return
	}
	flag, result, err := api.GetAPIPluginListWithNotAssignAPIList(strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			err.Error(),
			err)

	}
	controller.WriteResultInfo(httpResponse, "apiPlugin", "pluginList", result)

}
