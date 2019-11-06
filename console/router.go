package console

import (
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller/monitor"

	"github.com/eolinker/goku-api-gateway/console/controller/updater"

	config_log "github.com/eolinker/goku-api-gateway/console/controller/config-log"
	"github.com/eolinker/goku-api-gateway/console/controller/gateway"

	"github.com/eolinker/goku-api-gateway/console/controller/account"
	"github.com/eolinker/goku-api-gateway/console/controller/api"
	"github.com/eolinker/goku-api-gateway/console/controller/auth"
	"github.com/eolinker/goku-api-gateway/console/controller/balance"
	"github.com/eolinker/goku-api-gateway/console/controller/cluster"
	"github.com/eolinker/goku-api-gateway/console/controller/discovery"

	"github.com/eolinker/goku-api-gateway/console/controller/node"
	"github.com/eolinker/goku-api-gateway/console/controller/plugin"
	"github.com/eolinker/goku-api-gateway/console/controller/project"
	"github.com/eolinker/goku-api-gateway/console/controller/strategy"
)

//Router 路由
func Router() {

	// 游客
	http.HandleFunc("/guest/login", account.Login)

	// 用户
	http.HandleFunc("/user/logout", account.Logout)
	http.HandleFunc("/user/password/edit", account.EditPassword)
	http.HandleFunc("/user/getInfo", account.GetUserInfo)
	http.HandleFunc("/user/getUserType", account.GetUserType)
	http.HandleFunc("/user/checkIsAdmin", account.CheckUserIsAdmin)
	http.HandleFunc("/user/checkIsSuperAdmin", account.CheckUserIsSuperAdmin)
	http.HandleFunc("/user/checkPermission", account.CheckUserPermission)

	// 网关
	//http.HandleFunc("/gateway/config/base/getInfo", gateway.GetGatewayConfig)
	//http.HandleFunc("/gateway/config/base/edit", gateway.EditGatewayBaseConfig)
	//http.HandleFunc("/gateway/config/alert/edit", gateway.EditGatewayAlarmConfig)
	//http.HandleFunc("/gateway/config/alert/getInfo", alert.GetAlertConfig)

	// 项目
	http.HandleFunc("/project/add", project.AddProject)
	http.HandleFunc("/project/edit", project.EditProject)
	http.HandleFunc("/project/delete", project.DeleteProject)
	http.HandleFunc("/project/getInfo", project.GetProjectInfo)
	http.HandleFunc("/project/getList", project.GetProjectList)
	http.HandleFunc("/project/strategy/getList", project.GetAPIListFromProjectNotInStrategy)
	http.HandleFunc("/project/batchDelete", project.BatchDeleteProject)

	// 接口分组
	http.HandleFunc("/apis/group/add", api.AddAPIGroup)
	http.HandleFunc("/apis/group/edit", api.EditAPIGroup)
	http.HandleFunc("/apis/group/delete", api.DeleteAPIGroup)
	http.HandleFunc("/apis/group/getList", api.GetAPIGroupList)

	// API
	http.HandleFunc("/apis/add", api.AddAPI)
	http.HandleFunc("/apis/edit", api.EditAPI)
	http.HandleFunc("/apis/copy", api.CopyAPI)
	http.HandleFunc("/apis/getInfo", api.GetAPIInfo)
	http.HandleFunc("/apis/getList", api.GetAPIList)
	http.HandleFunc("/apis/id/getList", api.GetAPIIDList)
	http.HandleFunc("/apis/batchEditGroup", api.BatchEditAPIGroup)
	http.HandleFunc("/apis/batchDelete", api.BatchDeleteAPI)
	http.HandleFunc("/apis/batchEditBalance", api.BatchSetBalanceAPI)

	http.HandleFunc("/apis/manager/getList", api.GetAPIManagerList)

	// API绑定插件
	http.HandleFunc("/plugin/api/addPluginToApi", api.AddPluginToAPI)
	http.HandleFunc("/plugin/api/edit", api.EditAPIPluginConfig)
	http.HandleFunc("/plugin/api/getInfo", api.GetAPIPluginConfig)
	http.HandleFunc("/plugin/api/getList", api.GetAPIPluginList)

	http.HandleFunc("/plugin/api/getListByStrategy", api.GetAllAPIPluginInStrategy)
	http.HandleFunc("/plugin/api/batchStart", api.BatchStartAPIPlugin)
	http.HandleFunc("/plugin/api/batchStop", api.BatchStopAPIPlugin)
	http.HandleFunc("/plugin/api/batchDelete", api.BatchDeleteAPIPlugin)
	http.HandleFunc("/plugin/api/notAssign/getList", api.GetAPIPluginListWithNotAssignAPIList)

	// 策略绑定插件
	http.HandleFunc("/plugin/strategy/addPluginToStrategy", strategy.AddPluginToStrategy)
	http.HandleFunc("/plugin/strategy/edit", strategy.EditStrategyPluginConfig)
	http.HandleFunc("/plugin/strategy/getInfo", strategy.GetStrategyPluginConfig)
	http.HandleFunc("/plugin/strategy/getList", strategy.GetStrategyPluginList)
	http.HandleFunc("/plugin/strategy/checkPluginIsExist", strategy.CheckPluginIsExistInStrategy)
	http.HandleFunc("/plugin/strategy/getStatus", strategy.GetStrategyPluginStatus)
	http.HandleFunc("/plugin/strategy/batchStart", strategy.BatchStartStrategyPlugin)
	http.HandleFunc("/plugin/strategy/batchStop", strategy.BatchStopStrategyPlugin)
	http.HandleFunc("/plugin/strategy/batchDelete", strategy.BatchDeleteStrategyPlugin)

	// 插件
	http.HandleFunc("/plugin/add", plugin.AddPlugin)
	http.HandleFunc("/plugin/edit", plugin.EditPlugin)
	http.HandleFunc("/plugin/delete", plugin.DeletePlugin)
	http.HandleFunc("/plugin/checkNameIsExist", plugin.CheckIndexIsExist)
	http.HandleFunc("/plugin/checkIndexIsExist", plugin.CheckNameIsExist)
	http.HandleFunc("/plugin/getList", plugin.GetPluginList)
	http.HandleFunc("/plugin/getInfo", plugin.GetPluginInfo)
	http.HandleFunc("/plugin/getConfig", plugin.GetPluginConfig)
	http.HandleFunc("/plugin/start", plugin.StartPlugin)
	http.HandleFunc("/plugin/stop", plugin.StopPlugin)
	http.HandleFunc("/plugin/getListByType", plugin.GetPluginListByPluginType)
	http.HandleFunc("/plugin/batchStop", plugin.BatchStopPlugin)
	http.HandleFunc("/plugin/batchStart", plugin.BatchStartPlugin)
	http.HandleFunc("/plugin/availiable/check", plugin.CheckPluginIsAvailable)

	// 策略组
	http.HandleFunc("/strategy/add", strategy.AddStrategy)
	http.HandleFunc("/strategy/edit", strategy.EditStrategy)
	http.HandleFunc("/strategy/copy", strategy.CopyStrategy)
	http.HandleFunc("/strategy/delete", strategy.DeleteStrategy)
	http.HandleFunc("/strategy/getList", strategy.GetStrategyList)
	http.HandleFunc("/strategy/getInfo", strategy.GetStrategyInfo)
	http.HandleFunc("/strategy/batchEditGroup", strategy.BatchEditStrategyGroup)
	http.HandleFunc("/strategy/batchDelete", strategy.BatchDeleteStrategy)
	http.HandleFunc("/strategy/batchStart", strategy.BatchStartStrategy)
	http.HandleFunc("/strategy/batchStop", strategy.BatchStopStrategy)
	http.HandleFunc("/strategy/id/getList", strategy.GetStrategyIDList)

	http.HandleFunc("/monitor/gateway/getSummaryInfo", gateway.GetGatewayBasicInfo)
	// http.HandleFunc("/strategy/openStrategy/getInfo", strategy.GetOpenStrategy)

	// 策略组分组
	http.HandleFunc("/strategy/group/add", strategy.AddStrategyGroup)
	http.HandleFunc("/strategy/group/edit", strategy.EditStrategyGroup)
	http.HandleFunc("/strategy/group/delete", strategy.DeleteStrategyGroup)
	http.HandleFunc("/strategy/group/getList", strategy.GetStrategyGroupList)

	// 接口策略组
	http.HandleFunc("/strategy/api/add", strategy.AddAPIToStrategy)
	http.HandleFunc("/strategy/api/target", strategy.ResetAPITargetOfStrategy)
	http.HandleFunc("/strategy/api/batchEditTarget", strategy.BatchResetAPITargetOfStrategy)
	http.HandleFunc("/strategy/api/getList", strategy.GetAPIListFromStrategy)
	http.HandleFunc("/strategy/api/id/getList", strategy.GetAPIIDListFromStrategy)
	http.HandleFunc("/strategy/api/getNotInList", strategy.GetAPIListNotInStrategy)
	http.HandleFunc("/strategy/api/id/getNotInList", strategy.GetAPIIDListNotInStrategyByProject)
	http.HandleFunc("/strategy/api/batchDelete", strategy.BatchDeleteAPIInStrategy)
	http.HandleFunc("/strategy/api/plugin/getList", api.GetAPIPluginInStrategyByAPIID)

	http.HandleFunc("/strategy/balance/getList", strategy.GetBalanceListInStrategy)

	// 节点
	http.HandleFunc("/node/add", node.AddNode)
	http.HandleFunc("/node/edit", node.EditNode)
	http.HandleFunc("/node/delete", node.DeleteNode)
	http.HandleFunc("/node/getInfo", node.GetNodeInfo)
	http.HandleFunc("/node/getList", node.GetNodeList)

	http.HandleFunc("/node/batchEditGroup", node.BatchEditNodeGroup)
	http.HandleFunc("/node/batchDelete", node.BatchDeleteNode)

	// 节点分组
	http.HandleFunc("/node/group/add", node.AddNodeGroup)
	http.HandleFunc("/node/group/edit", node.EditNodeGroup)
	http.HandleFunc("/node/group/delete", node.DeleteNodeGroup)
	http.HandleFunc("/node/group/getInfo", node.GetNodeGroupInfo)
	http.HandleFunc("/node/group/getList", node.GetNodeGroupList)

	// 负载均衡
	http.HandleFunc("/balance/add", balance.AddBalance)
	http.HandleFunc("/balance/edit", balance.SaveBalance)
	http.HandleFunc("/balance/delete", balance.DeleteBalance)
	http.HandleFunc("/balance/getInfo", balance.GetBalanceInfo)
	http.HandleFunc("/balance/getList", balance.GetBalanceList)
	http.HandleFunc("/balance/batchDelete", balance.BatchDeleteBalance)
	http.HandleFunc("/balance/simple", balance.GetSimpleList)

	// 服务发现
	http.Handle("/balance/service/", discovery.Handle("/balance/service"))

	// 鉴权
	http.HandleFunc("/auth/getStatus", auth.GetAuthStatus)
	http.HandleFunc("/auth/getInfo", auth.GetAuthInfo)
	http.HandleFunc("/auth/editInfo", auth.EditAuthInfo)

	// 导入
	http.HandleFunc("/import/ams/api", api.ImportAPIFromAms)
	http.HandleFunc("/import/ams/group", api.ImportAPIGroupFromAms)
	http.HandleFunc("/import/ams/project", api.ImportProjectFromAms)

	// 	集群
	http.HandleFunc("/cluster/add", cluster.AddCluster)
	http.HandleFunc("/cluster/edit", cluster.EditCluster)
	http.HandleFunc("/cluster/delete", cluster.DeleteCluster)
	http.HandleFunc("/cluster/list", cluster.GetClusterInfoList)
	http.HandleFunc("/cluster/simpleList", cluster.GetClusterList)

	// 配置发布
	http.HandleFunc("/version/config/add", cluster.AddVersionConfig)
	http.HandleFunc("/version/config/getList", cluster.GetVersionList)
	http.HandleFunc("/version/config/delete", cluster.BatchDeleteVersionConfig)
	http.HandleFunc("/version/config/publish", cluster.PublishVersion)

	// 表更新
	http.HandleFunc("/updater/check/table/isExist", updater.IsTableExist)
	http.HandleFunc("/updater/check/column/isExist", updater.IsColumnExist)

	// 监控模块
	http.HandleFunc("/monitor/module/config/get", monitor.GetMonitorModules)
	http.HandleFunc("/monitor/module/config/set", monitor.SetMonitorModule)

	// 配置
	http.Handle("/config/log/", config_log.Handle("/config/log/"))
	http.HandleFunc("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))).ServeHTTP)

}
