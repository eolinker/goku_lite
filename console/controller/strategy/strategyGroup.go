package strategy

import (
	"net/http"
	"strconv"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/strategy"
)

const operationStrategyGroup = "strategyManagement"

//GroupHandlers 策略分组处理器
type GroupHandlers struct {
}

//Handlers handlers
func (h *GroupHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/add":     factory.NewAccountHandleFunction(operationStrategyGroup, true, AddStrategyGroup),
		"/edit":    factory.NewAccountHandleFunction(operationStrategyGroup, true, EditStrategyGroup),
		"/delete":  factory.NewAccountHandleFunction(operationStrategyGroup, true, DeleteStrategyGroup),
		"/getList": factory.NewAccountHandleFunction(operationStrategyGroup, false, GetStrategyGroupList),
	}
}

//NewGroupHandlers new groupHandlers
func NewGroupHandlers() *GroupHandlers {
	return &GroupHandlers{}
}

// AddStrategyGroup 新建接口分组
func AddStrategyGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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
