package strategy

import (
	"net/http"
	"strconv"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

const operationStrategy = "strategyManagement"

//Handlers 策略处理器
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":            factory.NewAccountHandleFunction(operationStrategy, true, AddStrategy),
		"/edit":           factory.NewAccountHandleFunction(operationStrategy, true, EditStrategy),
		"/copy":           factory.NewAccountHandleFunction(operationStrategy, true, CopyStrategy),
		"/delete":         factory.NewAccountHandleFunction(operationStrategy, true, DeleteStrategy),
		"/getInfo":        factory.NewAccountHandleFunction(operationStrategy, false, GetStrategyInfo),
		"/getList":        factory.NewAccountHandleFunction(operationStrategy, false, GetStrategyList),
		"/batchEditGroup": factory.NewAccountHandleFunction(operationStrategy, true, BatchEditStrategyGroup),
		"/batchDelete":    factory.NewAccountHandleFunction(operationStrategy, true, BatchDeleteStrategy),
		"/batchStart":     factory.NewAccountHandleFunction(operationStrategy, true, BatchStartStrategy),
		"/batchStop":      factory.NewAccountHandleFunction(operationStrategy, true, BatchStopStrategy),
		"/id/getList":     factory.NewAccountHandleFunction(operationStrategy, false, GetStrategyIDList),
	}
}

//NewStrategyHandlers new策略处理器
func NewStrategyHandlers() *Handlers {
	return &Handlers{}
}

//AddStrategy 新增策略组
func AddStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyName := httpRequest.PostFormValue("strategyName")
	groupID := httpRequest.PostFormValue("groupID")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if strategyName == "" {
		controller.WriteError(httpResponse,
			"220006",
			"strategy",
			"[ERROR]Illegal strategyName!",
			nil)
		return

	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"220005",
			"strategy",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	flag, result, err := strategy.AddStrategy(strategyName, gID, userID)
	if !flag {

		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategy", "strategyID", result)
	return
}

//EditStrategy 修改策略组信息
func EditStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyName := httpRequest.PostFormValue("strategyName")
	strategyID := httpRequest.PostFormValue("strategyID")
	groupID := httpRequest.PostFormValue("groupID")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if strategyName == "" {
		controller.WriteError(httpResponse,
			"220006",
			"strategy",
			"[ERROR]Illegal strategyName!",
			nil)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"220005",
			"strategy",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	flag, result, err := strategy.EditStrategy(strategyID, strategyName, gID, userID)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategy", "strategyID", result)
}

//DeleteStrategy 删除策略组
func DeleteStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyID := httpRequest.PostFormValue("strategyID")

	flag, result, err := strategy.DeleteStrategy(strategyID)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategy", "", nil)
}

// GetOpenStrategy 获取策略组列表
func GetOpenStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	var flag bool
	var err error
	err = nil
	var result *entity.Strategy
	flag, result, err = strategy.GetOpenStrategy()
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			"[ERROR]The open strategy dosen't exist!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategy", "strategyInfo", result)
}

//GetStrategyList 获取策略组列表
func GetStrategyList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
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

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		if groupID != "" {
			controller.WriteError(httpResponse,
				"220005",
				"strategy",
				"[ERROR]Illegal groupID!",
				err)
			return
		}
		gID = -1
	}

	op, err := strconv.Atoi(condition)
	if err != nil {
		if condition != "" {
			controller.WriteError(httpResponse,
				"220007",
				"strategy",
				"[ERROR]Illegal condition!",
				err)
			return
		}
	}

	var flag bool
	err = nil
	result := make([]*entity.Strategy, 0)
	flag, result, count, err := strategy.GetStrategyList(gID, keyword, op, p, pSize)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			"[ERROR]Empty list!",
			err)
		return
	}
	controller.WriteResultInfoWithPage(httpResponse, "strategy", "strategyList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: count,
		Page:     p,
		PageSize: pSize,
	})
	return
}

//GetStrategyIDList 获取策略组列表
func GetStrategyIDList(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		if groupID != "" {
			controller.WriteError(httpResponse,
				"220005",
				"strategy",
				"[ERROR]Illegal groupID!",
				err)
			return
		}
		gID = -1
	}

	op, err := strconv.Atoi(condition)
	if err != nil {
		if condition != "" {
			controller.WriteError(httpResponse,
				"220007",
				"strategy",
				"[ERROR]Illegal condition!",
				err)
			return
		}
	}

	flag, result, err := strategy.GetStrategyIDList(gID, keyword, op)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			"[ERROR]Empty list!",
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategy", "strategyList", result)
}

// GetStrategyInfo 获取策略组信息
func GetStrategyInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()

	strategyID := httpRequest.Form.Get("strategyID")
	strategyType := httpRequest.Form.Get("strategyType")
	sType, err := strconv.Atoi(strategyType)
	if err != nil {

	}
	var result *entity.Strategy
	var flag bool
	if sType == 1 {
		flag, result, err = strategy.GetOpenStrategy()
		if !flag {
			controller.WriteError(httpResponse,
				"220000",
				"strategy",
				"[ERROR]The open strategy dosen't exist!",
				err)
			return
		}
	} else {
		flag, result, err = strategy.GetStrategyInfo(strategyID)
		if !flag {
			controller.WriteError(httpResponse,
				"220000",
				"strategy",
				"[ERROR]Can not find the strategy!",
				err)
			return

		}
	}

	controller.WriteResultInfo(httpResponse, "strategy", "strategyInfo", result)
}

//BatchEditStrategyGroup 批量修改策略组分组
func BatchEditStrategyGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyIDList := httpRequest.PostFormValue("strategyIDList")
	groupID := httpRequest.PostFormValue("groupID")

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"220001",
			"strategy",
			"[ERROR]Illegal groupID!",
			err)
		return
	}
	if strategyIDList == "" {
		controller.WriteError(httpResponse,
			"220002",
			"strategy",
			"[ERROR]Illegal strategyIDList!",
			err)
		return
	}
	flag, result, err := strategy.BatchEditStrategyGroup(strategyIDList, gID)
	if !flag {

		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategy", "", nil)
	return
}

//BatchDeleteStrategy 批量修改策略组
func BatchDeleteStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyIDList := httpRequest.PostFormValue("strategyIDList")
	if strategyIDList == "" {
		controller.WriteError(httpResponse,
			"220002",
			"strategy",
			"[ERROR]Illegal strategyIDList!",
			nil)
		return
	}
	flag, result, err := strategy.BatchDeleteStrategy(strategyIDList)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategy", "", nil)
}

//BatchStartStrategy 批量开启策略
func BatchStartStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyIDList := httpRequest.PostFormValue("strategyIDList")
	if strategyIDList == "" {
		controller.WriteError(httpResponse,
			"220002",
			"strategy",
			"[ERROR]Illegal strategyIDList!",
			nil)
		return
	}
	flag, result, err := strategy.BatchUpdateStrategyEnableStatus(strategyIDList, 1)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "strategy", "", nil)
}

//BatchStopStrategy 批量关闭策略
func BatchStopStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyIDList := httpRequest.PostFormValue("strategyIDList")
	if strategyIDList == "" {
		controller.WriteError(httpResponse,
			"220002",
			"strategy",
			"[ERROR]Illegal strategyIDList!",
			nil)
		return
	}

	flag, result, err := strategy.BatchUpdateStrategyEnableStatus(strategyIDList, 0)

	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategy", "", nil)
}

// GetBalanceListInStrategy 获取在策略中的负载列表
func GetBalanceListInStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	httpRequest.ParseForm()
	strategyID := httpRequest.Form.Get("strategyID")
	balanceType := httpRequest.Form.Get("balanceType")
	bt, err := strconv.Atoi(balanceType)
	if err != nil && balanceType != "" {

	}
	flag, result, err := strategy.GetBalanceListInStrategy(strategyID, bt)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			err.Error(),
			err)
		return
	}
	controller.WriteResultInfo(httpResponse, "strategy", "balanceNameList", result)
}

// CopyStrategy 复制策略
func CopyStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	strategyName := httpRequest.PostFormValue("strategyName")
	strategyID := httpRequest.PostFormValue("strategyID")
	groupID := httpRequest.PostFormValue("groupID")
	userID := goku_handler.UserIDFromRequest(httpRequest)
	if strategyID == "" {
		controller.WriteError(httpResponse,
			"220003",
			"strategy",
			"[ERROR]Illegal strategyID!",
			nil)
		return
	}
	if strategyName == "" {
		controller.WriteError(httpResponse,
			"220006",
			"strategy",
			"[ERROR]Illegal strategyName!",
			nil)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"220001",
			"strategy",
			"[ERROR]Illegal groupID!",
			err)
		return

	}
	flag, result, err := strategy.AddStrategy(strategyName, gID, userID)
	if !flag {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}
	_, err = strategy.CopyStrategy(strategyID, result, userID)
	if err != nil {
		controller.WriteError(httpResponse,
			"220000",
			"strategy",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "strategy", "strategyID", result)
	return
}
