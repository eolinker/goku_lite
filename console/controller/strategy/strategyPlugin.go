package strategy

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/plugin"
	plugin_config "github.com/eolinker/goku-api-gateway/console/module/plugin/plugin-config"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
)

// 新增插件到接口
func AddPluginToStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

// 修改插件信息
func EditStrategyPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

// 获取策略组插件列表
func GetStrategyPluginList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}
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

// 获取策略组插件信息
func GetStrategyPluginConfig(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

	strategyID := httpRequest.PostFormValue("strategyID")
	pluginName := httpRequest.PostFormValue("pluginName")
	_, result, _ := strategy.GetStrategyPluginConfig(strategyID, pluginName)

	controller.WriteResultInfo(httpResponse, "strategyPlugin", "strategyPluginConfig", result)

}

// 检查策略组是否绑定插件
func CheckPluginIsExistInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

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

// 检查策略组插件是否开启
func GetStrategyPluginStatus(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

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

// 批量开启策略组插件
func BatchStartStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

// 批量修改策略组插件状态
func BatchStopStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

// 批量删除策略组插件
func BatchDeleteStrategyPlugin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

func UpdateAllStrategyPluginUpdateTag() error {
	return strategy.UpdateAllStrategyPluginUpdateTag()
}
