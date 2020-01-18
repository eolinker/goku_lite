package strategy

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/api"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
)

const operationAPIStrategy = "strategyManagement"

//APIStrategyHandlers 接口策略handlers
type APIStrategyHandlers struct {
}

//Handlers handlers
func (h *APIStrategyHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":             factory.NewAccountHandleFunction(operationAPIStrategy, true, AddAPIToStrategy),
		"/target":          factory.NewAccountHandleFunction(operationAPIStrategy, true, ResetAPITargetOfStrategy),
		"/batchEditTarget": factory.NewAccountHandleFunction(operationAPIStrategy, true, BatchResetAPITargetOfStrategy),
		"/getList":         factory.NewAccountHandleFunction(operationAPIStrategy, false, GetAPIListFromStrategy),
		"/id/getList":      factory.NewAccountHandleFunction(operationAPIStrategy, false, GetAPIIDListFromStrategy),
		"/getNotInList":    factory.NewAccountHandleFunction(operationAPIStrategy, false, GetAPIListNotInStrategy),
		"/id/getNotInList": factory.NewAccountHandleFunction(operationAPIStrategy, false, GetAPIIDListNotInStrategyByProject),
		"/batchDelete":     factory.NewAccountHandleFunction(operationAPIStrategy, true, BatchDeleteAPIInStrategy),
		"/plugin/getList":  factory.NewAccountHandleFunction(operationAPIStrategy, false, GetAPIPluginInStrategyByAPIID),
	}
}

//NewAPIStrategyHandlers new接口策略处理器
func NewAPIStrategyHandlers() *APIStrategyHandlers {
	return &APIStrategyHandlers{}
}

//AddAPIToStrategy 将接口加入策略组
func AddAPIToStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")
	apiArray := strings.Split(apiID, ",")

	flag, err := strategy.CheckStrategyIsExist(strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240013",
			"apiStrategy",
			"[ERROR]The strategy does not exist!",
			err)
		return

	}
	flag, result, err := api.AddAPIToStrategy(apiArray, strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiStrategy",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "apiStrategy", "", nil)

}

// ResetAPITargetOfStrategy 将接口加入策略组
func ResetAPITargetOfStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	strategyID := httpRequest.PostFormValue("strategyID")
	target := httpRequest.PostFormValue("target")
	apiID := httpRequest.PostFormValue("apiID")
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240013",
			"apiStrategy",
			"[ERROR]The strategy does not exist!",
			err)
		return

	}
	flag, err := strategy.CheckStrategyIsExist(strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240013",
			"apiStrategy",
			"[ERROR]The strategy does not exist!",
			err)
		return

	}
	flag, result, err := api.SetTarget(aID, strategyID, target)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiStrategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "apiStrategy", "", nil)

}

// BatchResetAPITargetOfStrategy 将接口加入策略组
func BatchResetAPITargetOfStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	target := httpRequest.PostFormValue("target")
	apiIDs := httpRequest.PostFormValue("apiIDs")
	ids := make([]int, 0)
	err := json.Unmarshal([]byte(apiIDs), &ids)
	if err != nil || len(ids) < 1 {
		controller.WriteError(httpResponse,
			"240004",
			"apiStrategy",
			"[ERROR]Illegal apiIDs!",
			err)
		return
	}
	flag, err := strategy.CheckStrategyIsExist(strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240013",
			"apiStrategy",
			"[ERROR]The strategy does not exist!",
			err)
		return

	}
	flag, result, err := api.BatchSetTarget(ids, strategyID, target)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiStrategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "apiStrategy", "", nil)

}

// GetAPIIDListFromStrategy 获取策略组接口ID列表
func GetAPIIDListFromStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	idStr := httpRequest.Form.Get("ids")
	balanceNames := httpRequest.Form.Get("balanceNames")

	op, err := strconv.Atoi(condition)
	if err != nil {
	}
	var ids []int
	var names []string
	if op > 0 {
		switch op {
		case 1, 2:
			{
				err := json.Unmarshal([]byte(balanceNames), &names)
				if err != nil || len(names) < 1 {
					controller.WriteError(httpResponse, "240001", "apiStrategy", "[ERROR]Illegal balanceNames!", err)
					return
				}
				break

			}
		case 3, 4:
			{
				err := json.Unmarshal([]byte(idStr), &ids)
				if err != nil || len(ids) < 1 {
					controller.WriteError(httpResponse, "240002", "apiStrategy", "[ERROR]Illegal ids!", err)
					return
				}
				break
			}
		default:
			{
				controller.WriteError(httpResponse, "240003", "apiStrategy", "[ERROR]Illegal condition!", err)
				return
			}
		}

	}

	_, result, err := api.GetAPIIDListFromStrategy(strategyID, keyword, op, ids, names)
	controller.WriteResultInfoWithPage(httpResponse, "apiStrategy", "apiIDList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: len(result),
	})
	return
}

// GetAPIListFromStrategy 获取策略组接口列表
func GetAPIListFromStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	idStr := httpRequest.Form.Get("ids")
	balanceNames := httpRequest.Form.Get("balanceNames")
	page := httpRequest.Form.Get("page")
	pageSize := httpRequest.Form.Get("pageSize")

	p, e := strconv.Atoi(page)
	if e != nil {
		p = 1
	}
	pSize, e := strconv.Atoi(pageSize)
	if e != nil {
		pSize = 15
	}

	op, err := strconv.Atoi(condition)
	if err != nil {

	}
	var ids []int
	var names []string
	if op > 0 {
		switch op {
		case 1, 2:
			{
				err := json.Unmarshal([]byte(balanceNames), &names)
				if err != nil || len(names) < 1 {
					controller.WriteError(httpResponse, "240001", "apiStrategy", "[ERROR]Illegal balanceNames!", err)
					return
				}
				break

			}
		case 3, 4:
			{
				err := json.Unmarshal([]byte(idStr), &ids)
				if err != nil || len(ids) < 1 {
					controller.WriteError(httpResponse, "240002", "apiStrategy", "[ERROR]Illegal ids!", err)
					return
				}
				break
			}
		default:
			{
				controller.WriteError(httpResponse, "240003", "apiStrategy", "[ERROR]Illegal condition!", err)
				return
			}
		}

	}

	_, result, count, err := api.GetAPIListFromStrategy(strategyID, keyword, op, p, pSize, ids, names)

	controller.WriteResultInfoWithPage(httpResponse, "apiStrategy", "apiList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: count,
		Page:     p,
		PageSize: pSize,
	})
	return
}

//CheckIsExistAPIInStrategy 检查插件是否添加进策略组
func CheckIsExistAPIInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")
	apiID := httpRequest.PostFormValue("apiID")

	id, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse,
			"190001",
			"apiStrategy",
			"[ERROR]Illegal apiID",
			err)
		return

	}
	flag, _, err := api.CheckIsExistAPIInStrategy(id, strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiStrategy",
			"[ERROR]Can not find the api in strategy!",
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "apiStrategy", "", nil)

	return
}

// GetAPIIDListNotInStrategyByProject 获取未被该策略组绑定的接口ID列表(通过项目)
func GetAPIIDListNotInStrategyByProject(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	projectID := httpRequest.Form.Get("projectID")
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")

	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240008",
			"apiStrategy",
			"[ERROR]Illegal projectID!",
			err)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		if groupID != "" {
			controller.WriteError(httpResponse,
				"240009",
				"apiStrategy",
				"[ERROR]Illegal groupID!",
				err)
			return
		}
		gID = -1
	}
	_, result, _ := api.GetAPIIDListNotInStrategy(strategyID, pjID, gID, keyword)
	controller.WriteResultInfoWithPage(httpResponse, "apiStrategy", "apiIDList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: len(result),
	})
	return
}

//GetAPIListNotInStrategy 获取未被该策略组绑定的接口列表
func GetAPIListNotInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	projectID := httpRequest.Form.Get("projectID")
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")
	page := httpRequest.Form.Get("page")
	pageSize := httpRequest.Form.Get("pageSize")

	p, e := strconv.Atoi(page)
	if e != nil {
		p = 1
	}
	pSize, e := strconv.Atoi(pageSize)
	if e != nil {
		pSize = 15
	}

	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse,
			"240008",
			"apiStrategy",
			"[ERROR]Illegal projectID!",
			err)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		if groupID != "" {
			controller.WriteError(httpResponse,
				"240009",
				"apiStrategy",
				"[ERROR]Illegal groupID!",
				err)
			return
		}
		gID = -1
	}
	result := make([]map[string]interface{}, 0)
	_, result, count, err := api.GetAPIListNotInStrategy(strategyID, pjID, gID, p, pSize, keyword)
	controller.WriteResultInfoWithPage(httpResponse, "apiStrategy", "apiList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: count,
		Page:     p,
		PageSize: pSize,
	})
	return
}

//BatchDeleteAPIInStrategy 批量删除策略组接口
func BatchDeleteAPIInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	apiIDList := httpRequest.PostFormValue("apiIDList")
	strategyID := httpRequest.PostFormValue("strategyID")

	flag, result, err := api.BatchDeleteAPIInStrategy(apiIDList, strategyID)
	if !flag {

		controller.WriteError(httpResponse,
			"240000",
			"apiStrategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "apiStrategy", "", nil)
}

// GetAPIPluginInStrategyByAPIID 获取策略组中所有接口插件列表
func GetAPIPluginInStrategyByAPIID(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	apiID := httpRequest.Form.Get("apiID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")

	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse, "240002", "apiPlugin", "[ERROR]Illegal condition!", err)
		return
	}
	op, err := strconv.Atoi(condition)
	if err != nil && condition != "" {
		controller.WriteError(httpResponse, "270006", "apiPlugin", "[ERROR]Illegal condition!", err)
		return
	}

	flag, pluginList, apiInfo, err := api.GetAPIPluginInStrategyByAPIID(strategyID, aID, keyword, op)
	if !flag {
		controller.WriteError(httpResponse,
			"240000",
			"apiPlugin",
			"[ERROR]Empty api plugin list!",
			err)
		return

	}
	result := map[string]interface{}{
		"statusCode":    "000000",
		"type":          "apiPlugin",
		"resultDesc":    "",
		"apiPluginList": pluginList,
		"apiInfo":       apiInfo,
		"page": controller.PageInfo{
			ItemNum: len(pluginList),
		},
	}
	resultStr, _ := json.Marshal(result)
	httpResponse.Write(resultStr)
	// controller.WriteResultInfo(httpResponse, "apiPlugin", "apiPluginList", result)
}
