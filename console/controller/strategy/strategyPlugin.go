package strategy

import (
	"net/http"
	"strconv"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/plugin"
	plugin_config "github.com/eolinker/goku-api-gateway/console/module/plugin/plugin-config"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
)

const operationStrategyPlugin = "strategyManagement"

//PluginHandlers 策略插件处理器
type PluginHandlers struct {
}

//Handlers handlers
func (h *PluginHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/addPluginToStrategy": factory.NewAccountHandleFunction(operationAPIStrategy, true, AddPluginToStrategy),
		"/edit":                factory.NewAccountHandleFunction(operationAPIStrategy, true, EditStrategyPluginConfig),
		"/getInfo":             factory.NewAccountHandleFunction(operationAPIStrategy, false, GetStrategyPluginConfig),
		"/getList":             factory.NewAccountHandleFunction(operationAPIStrategy, false, GetStrategyPluginList),
		"/checkPluginIsExist":  factory.NewAccountHandleFunction(operationAPIStrategy, false, CheckPluginIsExistInStrategy),
		"/getStatus":           factory.NewAccountHandleFunction(operationAPIStrategy, false, GetStrategyPluginStatus),
		"/batchStart":          factory.NewAccountHandleFunction(operationAPIStrategy, false, BatchStartStrategyPlugin),
		"/batchStop":           factory.NewAccountHandleFunction(operationAPIStrategy, true, BatchStopStrategyPlugin),
		"/batchDelete":         factory.NewAccountHandleFunction(operationAPIStrategy, true, BatchDeleteStrategyPlugin),
	}

}

//NewPluginHandlers new pluginHandlers
func NewPluginHandlers() *PluginHandlers {
	return &PluginHandlers{}
}

//AddPluginToStrategy 新增插件到接口
func AddPluginToStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")
	strategyID := httpRequest.PostFormValue("strategyID")
	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "strategyPlugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}

	// 查询该插件是否存在,如果存在，则修改信息,不存在，则新增
	flag, err = plugin.CheckNameIsExist(pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"200005",
			"strategyPlugin",
			"[ERROR]The plugin does not exist!",
			err)
		return
	}
	// 检查插件是否已经绑定该插件
	exits, _ := strategy.CheckPluginIsExistInStrategy(strategyID, pluginName)
	if exits {
		flag, resultDesc, err := strategy.EditStrategyPluginConfig(pluginName, pluginConfig, strategyID)
		if !flag {
			controller.WriteError(httpResponse,
				"270000",
				"strategyPlugin",
				resultDesc,
				err)
			return
		}
		_, connID, _ := strategy.GetConnIDFromStrategyPlugin(pluginName, strategyID)
		controller.WriteResultInfo(httpResponse, "strategyPlugin", "connID", connID)
		return
	}
	flag, resultDesc, err := strategy.AddPluginToStrategy(pluginName, pluginConfig, strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			resultDesc.(string),
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "connID", resultDesc)

}

//EditStrategyPluginConfig 修改插件信息
func EditStrategyPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	pluginName := httpRequest.PostFormValue("pluginName")
	pluginConfig := httpRequest.PostFormValue("pluginConfig")
	strategyID := httpRequest.PostFormValue("strategyID")
	flag, err := plugin_config.CheckConfig(pluginName, []byte(pluginConfig))
	if !flag {
		controller.WriteError(httpResponse, "500000", "strategyPlugin", "[ERROR]插件配置无效:"+err.Error(), err)
		return
	}

	// 查询该插件是否存在
	flag, err = plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"200005",
			"strategyPlugin",
			"[ERROR]The plugin does not exist!",
			err)
		return

	}
	flag, resultDesc, err := strategy.EditStrategyPluginConfig(pluginName, pluginConfig, strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			resultDesc,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)

}

//GetStrategyPluginList 获取策略组插件列表
func GetStrategyPluginList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	op, err := strconv.Atoi(condition)
	if err != nil && condition != "" {
		controller.WriteError(httpResponse, "270002", "plugin", "[ERROR]Illegal condition!", err)
		return
	}

	// 查询该插件是否存在
	flag, result, err := strategy.GetStrategyPluginList(strategyID, keyword, op)
	if !flag {
		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			"[ERROR]Empty strategy plugin list!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "strategyPluginList", result)
}

//GetStrategyPluginConfig 获取策略组插件信息
func GetStrategyPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	pluginName := httpRequest.PostFormValue("pluginName")
	_, result, _ := strategy.GetStrategyPluginConfig(strategyID, pluginName)

	controller.WriteResultInfo(httpResponse, "strategyPlugin", "strategyPluginConfig", result)

}

//CheckPluginIsExistInStrategy 检查策略组是否绑定插件
func CheckPluginIsExistInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	pluginName := httpRequest.PostFormValue("pluginName")

	// 查询该插件是否存在
	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"200005",
			"strategyPlugin",
			"[ERROR]The plugin does not exist!",
			err)
		return
	}
	flag, err = strategy.CheckPluginIsExistInStrategy(strategyID, pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			"[ERROR]The api plugin is existed!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)
}

//GetStrategyPluginStatus 检查策略组插件是否开启
func GetStrategyPluginStatus(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	pluginName := httpRequest.PostFormValue("pluginName")

	// 查询该插件是否存在
	flag, err := plugin.CheckNameIsExist(pluginName)
	if !flag {
		controller.WriteError(httpResponse,
			"200005",
			"strategyPlugin",
			"[ERROR]The plugin does not exist!",
			err)
		return

	}
	flag, err = strategy.GetStrategyPluginStatus(strategyID, pluginName)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			"[ERROR]The api plugin does not exist!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)

}

//BatchStartStrategyPlugin 批量开启策略组插件
func BatchStartStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")
	if connIDList == "" {
		controller.WriteError(httpResponse,
			"270001",
			"strategyPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return
	}
	flag, result, err := strategy.BatchEditStrategyPluginStatus(connIDList, strategyID, 1)
	if !flag {
		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)
}

//BatchStopStrategyPlugin 批量修改策略组插件状态
func BatchStopStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")
	if connIDList == "" {
		controller.WriteError(httpResponse,
			"270001",
			"strategyPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return
	}
	flag, result, err := strategy.BatchEditStrategyPluginStatus(connIDList, strategyID, 0)
	if !flag {
		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			result,
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)
}

//BatchDeleteStrategyPlugin 批量删除策略组插件
func BatchDeleteStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	connIDList := httpRequest.PostFormValue("connIDList")
	if connIDList == "" {
		controller.WriteError(httpResponse,
			"270001",
			"strategyPlugin",
			"[ERROR]Illegal connIDList",
			nil)
		return
	}
	flag, result, err := strategy.BatchDeleteStrategyPlugin(connIDList, strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"270000",
			"strategyPlugin",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyPlugin", "", nil)
}
