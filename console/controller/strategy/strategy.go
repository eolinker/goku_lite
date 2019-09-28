package strategy

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddStrategy 新增策略组
func AddStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

	strategyName := httpRequest.PostFormValue("strategyName")
	groupID := httpRequest.PostFormValue("groupID")
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
	flag, result, err := strategy.AddStrategy(strategyName, gID)
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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

	strategyName := httpRequest.PostFormValue("strategyName")
	strategyID := httpRequest.PostFormValue("strategyID")
	groupID := httpRequest.PostFormValue("groupID")
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
	flag, result, err := strategy.EditStrategy(strategyID, strategyName, gID)
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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}
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

	var flag bool
	err = nil
	result := make([]*entity.Strategy, 0)
	flag, result, err = strategy.GetStrategyList(gID, keyword, op)
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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}
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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

//BatchStartStrategy 更新策略启用状态
func BatchStartStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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

//BatchStopStrategy 更新策略启用状态
func BatchStopStrategy(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

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
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}
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
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationEDIT)
	if e != nil {
		return
	}

	strategyName := httpRequest.PostFormValue("strategyName")
	strategyID := httpRequest.PostFormValue("strategyID")
	groupID := httpRequest.PostFormValue("groupID")
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
	flag, result, err := strategy.AddStrategy(strategyName, gID)
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
