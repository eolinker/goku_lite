package strategy

import (
	"net/http"
	"strconv"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/strategy"
)

// AddStrategyGroup 新建接口分组
func AddStrategyGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

	groupName := httpRequest.PostFormValue("groupName")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"300002",
			"strategyGroup",
			"[ERROR]Illegal groupName!",
			nil)
		return
	}
	flag, result, err := strategy.AddStrategyGroup(groupName)
	if !flag {
		controller.WriteError(httpResponse,
			"300000",
			"strategyGroup",
			result.(string),
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyGroup", "groupID", result)
}

// EditStrategyGroup 修改接口分组
func EditStrategyGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

	groupName := httpRequest.PostFormValue("groupName")
	groupID := httpRequest.PostFormValue("groupID")
	if groupName == "" {
		controller.WriteError(httpResponse,
			"300002",
			"strategyGroup",
			"[ERROR]Illegal groupName!",
			nil)
		return

	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"300001",
			"strategyGroup",
			"[ERROR]Illegal groupID",
			err)
		return
	}
	flag, result, err := strategy.EditStrategyGroup(groupName, gID)
	if !flag {
		controller.WriteError(httpResponse,
			"300000",
			"strategyGroup",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "strategyGroup", "", nil)

}

// DeleteStrategyGroup 删除接口分组
func DeleteStrategyGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

	groupID := httpRequest.PostFormValue("groupID")

	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse,
			"300001",
			"strategyGroup",
			"[ERROR]Illegal groupID",
			err)
		return
	}
	flag, result, err := strategy.DeleteStrategyGroup(gID)
	if !flag {

		controller.WriteError(httpResponse,
			"300000",
			"strategyGroup",
			result,
			err)
		return

	}
	controller.WriteResultInfo(httpResponse, "strategyGroup", "", nil)
}

// GetStrategyGroupList 获取接口分组列表
func GetStrategyGroupList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationStrategy, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := strategy.GetStrategyGroupList()
	if !flag {
		controller.WriteError(httpResponse,
			"300000",
			"strategyGroup",
			"[ERROR]Empty strategy group list!",
			err)
		return
	}

	controller.WriteResultInfo(httpResponse, "strategyGroup", "groupList", result)

	return

}
