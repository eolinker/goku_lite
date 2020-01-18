package dao

import (
	"time"

	"github.com/eolinker/goku-api-gateway/config"
	balanceentity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity"
	configLogEntry "github.com/eolinker/goku-api-gateway/server/entity/config-log"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//APIDao apiDao
type APIDao interface {
	// AddAPI 新增接口
	AddAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkAPIs, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType int) (bool, int, error)
	// EditAPI 修改接口
	EditAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkAPIs, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error)
	// GetAPIInfo 获取接口信息
	GetAPIInfo(apiID int) (bool, *entity.API, error)
	//GetAPIListByGroupList 通过分组列表获取接口列表
	GetAPIListByGroupList(projectID int, groupIDList string) (bool, []map[string]interface{}, error)

	// GetAPIIDList 获取接口ID列表
	GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error)
	// GetAPIList 获取所有接口列表
	GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error)
	//CheckURLIsExist 接口路径是否存在
	CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool
	//CheckAPIIsExist 检查接口是否存在
	CheckAPIIsExist(apiID int) (bool, error)
	//CheckAliasIsExist 检查别名是否存在
	CheckAliasIsExist(apiID int, alias string) bool
	//BatchEditAPIBalance 批量修改接口负载
	BatchEditAPIBalance(apiIDList []string, balance string) (string, error)
	//BatchEditAPIGroup 批量修改接口分组
	BatchEditAPIGroup(apiIDList []string, groupID int) (string, error)
	//BatchDeleteAPI 批量修改接口
	BatchDeleteAPI(apiIDList string) (bool, string, error)
}

//APIGroupDao apiGroupDao
type APIGroupDao interface {
	//AddAPIGroup 新建接口分组
	AddAPIGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error)
	//EditAPIGroup 修改接口分组
	EditAPIGroup(groupName string, groupID, projectID int) (bool, string, error)
	//DeleteAPIGroup 删除接口分组
	DeleteAPIGroup(projectID, groupID int) (bool, string, error)
	//GetAPIGroupList 获取接口分组列表
	GetAPIGroupList(projectID int) (bool, []map[string]interface{}, error)
}

//APIPluginDao apiPlugin.go
type APIPluginDao interface {
	//AddPluginToAPI 新增接口插件
	AddPluginToAPI(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error)
	//EditAPIPluginConfig 修改接口插件配置
	EditAPIPluginConfig(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error)
	//GetAPIPluginList 获取接口插件列表
	GetAPIPluginList(apiID int, strategyID string) (bool, []map[string]interface{}, error)
	//GetPluginIndex 获取插件优先级
	GetPluginIndex(pluginName string) (bool, int, error)
	//GetAPIPluginConfig 通过APIID获取配置信息
	GetAPIPluginConfig(apiID int, strategyID, pluginName string) (bool, map[string]string, error)
	//CheckPluginIsExistInAPI 检查策略组是否绑定插件
	CheckPluginIsExistInAPI(strategyID, pluginName string, apiID int) (bool, error)
	// GetAPIPluginInStrategyByAPIID 通过接口ID获取策略组中接口插件列表
	GetAPIPluginInStrategyByAPIID(strategyID string, apiID int, keyword string, condition int) (bool, []map[string]interface{}, map[string]interface{}, error)
	//GetAllAPIPluginInStrategy 获取策略组中所有接口插件列表
	GetAllAPIPluginInStrategy(strategyID string) (bool, []map[string]interface{}, error)
	//BatchEditAPIPluginStatus 批量修改策略组插件状态
	BatchEditAPIPluginStatus(connIDList, strategyID string, pluginStatus, userID int) (bool, string, error)
	//BatchDeleteAPIPlugin 批量删除策略组插件
	BatchDeleteAPIPlugin(connIDList, strategyID string) (bool, string, error)
	//GetAPIPluginName 通过connID获取插件名称
	GetAPIPluginName(connID int) (bool, string, error)
	//CheckAPIPluginIsExistByConnIDList 通过connIDList判断插件是否存在
	CheckAPIPluginIsExistByConnIDList(connIDList, pluginName string) (bool, []int, error)
	//GetAPIPluginListWithNotAssignAPIList 获取没有绑定嵌套插件列表
	GetAPIPluginListWithNotAssignAPIList(strategyID string) (bool, []map[string]interface{}, error)
}

//APIStrategyDao apiStrategy.go
type APIStrategyDao interface {
	//AddAPIToStrategy 将接口加入策略组
	AddAPIToStrategy(apiList []string, strategyID string) (bool, string, error)
	// SetAPITargetOfStrategy 重定向接口负载
	SetAPITargetOfStrategy(apiID int, strategyID string, target string) (bool, string, error)
	// BatchSetAPITargetOfStrategy 批量重定向接口负载
	BatchSetAPITargetOfStrategy(apiIds []int, strategyID string, target string) (bool, string, error)
	// GetAPIIDListFromStrategy 获取策略组接口列表
	GetAPIIDListFromStrategy(strategyID, keyword string, condition int, ids []int, balanceNames []string) (bool, []int, error)
	// GetAPIListFromStrategy 获取策略组接口列表
	GetAPIListFromStrategy(strategyID, keyword string, condition, page, pageSize int, ids []int, balanceNames []string) (bool, []map[string]interface{}, int, error)
	// CheckIsExistAPIInStrategy 检查插件是否添加进策略组
	CheckIsExistAPIInStrategy(apiID int, strategyID string) (bool, string, error)

	// GetAPIIDListNotInStrategy 获取未被该策略组绑定的接口ID列表(通过项目)
	GetAPIIDListNotInStrategy(strategyID string, projectID, groupID int, keyword string) (bool, []int, error)
	// GetAPIListNotInStrategy 获取未被该策略组绑定的接口列表(通过项目)
	GetAPIListNotInStrategy(strategyID string, projectID, groupID, page, pageSize int, keyword string) (bool, []map[string]interface{}, int, error)
	//BatchDeleteAPIInStrategy 批量删除策略组接口
	BatchDeleteAPIInStrategy(apiIDList, strategyID string) (bool, string, error)
}

//AuthDao auth.go
type AuthDao interface {
	//GetAuthStatus 获取认证状态
	GetAuthStatus(strategyID string) (bool, map[string]interface{}, error)
	//GetAuthInfo 获取认证信息
	GetAuthInfo(strategyID string) (bool, map[string]interface{}, error)
	//EditAuthInfo 编辑认证信息
	EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList string, delClientIDList []string) (bool, error)
}

//ClusterDao cluster.go
type ClusterDao interface {
	//AddCluster 新增集群
	AddCluster(name, title, note string) error
	//EditCluster 修改集群信息
	EditCluster(name, title, note string) error
	//DeleteCluster 删除集群
	DeleteCluster(name string) error
	//GetClusterCount 获取集群数量
	GetClusterCount() int
	//GetClusterNodeCount 获取集群节点数量
	GetClusterNodeCount(name string) int
	//GetClusterIDByName 通过集群名称获取集群ID
	GetClusterIDByName(name string) int
	//GetClusters 获取集群列表
	GetClusters() ([]*entity.Cluster, error)
	//GetCluster 获取集群信息
	GetCluster(name string) (*entity.Cluster, error)
	GetClusterByID(id int) (*entity.Cluster, error)
	//CheckClusterNameIsExist 判断集群名称是否存在
	CheckClusterNameIsExist(name string) bool
}

//ConfigLogDao config-log
type ConfigLogDao interface {
	//Get get
	Get(name string) (*configLogEntry.LogConfig, error)
	//Set set
	Set(ent *configLogEntry.LogConfig) error
}

//BalanceUpdateDao dao-balance-update
type BalanceUpdateDao interface {
	//GetAllOldVerSion 获取所有旧负载配置
	GetAllOldVerSion() ([]*balanceentity.BalanceInfoEntity, error)
	//GetDefaultServiceStatic 获取默认静态负载
	GetDefaultServiceStatic() string
}

//ServiceDao dao-service
type ServiceDao interface {
	//RegisterDao 新增服务
	Add(name, driver, desc, config, clusterConfig string, isDefault, healthCheck bool, healthCheckPath string, healthCheckCode string, healthCheckPeriod, healthCheckTimeOut int) error
	//SetDefault 设置默认服务
	SetDefault(name string) error
	//Delete 删除服务发现
	Delete(names []string) error
	//Get 获取服务发现信息
	Get(name string) (*entity.Service, error)
	//List 获取服务发现列表
	List(keyword string) ([]*entity.Service, error)
	//Save 存储服务发现信息
	Save(name, desc, config, clusterConfig string, healthCheck bool, healthCheckPath string, healthCheckCode string, healthCheckPeriod, healthCheckTimeOut int) error
}

//VersionConfigDao dao-version-config
type VersionConfigDao interface {
	//GetAPIContent 获取接口信息
	GetAPIContent() ([]*config.APIContent, error)
	//GetBalances 获取balance信息
	GetDiscoverConfig(clusters []*entity.Cluster) (map[string]map[string]*config.DiscoverConfig, error)
	//GetDiscoverConfig 获取服务发现信息
	GetBalances(clusters []*entity.Cluster) (map[string]map[string]*config.BalanceConfig, error)
	//GetGlobalPlugin 获取全局插件
	GetGlobalPlugin() (*config.GatewayPluginConfig, error)
	//GetAPIPlugins 获取接口插件
	GetAPIPlugins() (map[string][]*config.PluginConfig, error)
	//GetStrategyPlugins 获取策略插件
	GetStrategyPlugins() (map[string][]*config.PluginConfig, map[string]map[string]string, error)
	//GetAPIsOfStrategy 获取策略内接口数据
	GetAPIsOfStrategy() (map[string][]*config.APIOfStrategy, error)
	//GetStrategyConfig 获取策略配置
	GetStrategyConfig() (string, []*config.StrategyConfig, error)
	//GetLogInfo 获取日志信息
	GetLogInfo() (*config.LogConfig, *config.AccessLogConfig, error)
	//GetMonitorModules 获取监控模块信息
	GetMonitorModules(status int, isAll bool) (map[string]string, error)

	GetRouterRules(enable int) ([]*config.Router, error)

	GetGatewayBasicConfig() (*config.Gateway, error)
}

//GatewayDao gateway.go
type GatewayDao interface {
	//GetGatewayConfig 获取网关配置
	GetGatewayConfig() (map[string]interface{}, error)
	//EditGatewayBaseConfig 编辑网关基本配置
	EditGatewayBaseConfig(config entity.GatewayBasicConfig) (bool, string, error)
	//GetGatewayInfo 获取网关信息
	GetGatewayInfo() (nodeStartCount, nodeStopCount, projectCount, apiCount, strategyCount int, err error)
}

//GuestDao guest.go
type GuestDao interface {
	//Login 登录
	Login(loginCall, loginPassword string) (bool, int)
	//CheckLogin 检查用户是否登录
	CheckLogin(userToken string, userID int) bool
	//Register 用户注册
	Register(loginCall, loginPassword string) bool
}

//ImportDao import.go
type ImportDao interface {
	//ImportAPIGroupFromAms 导入分组
	ImportAPIGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error)
	//ImportProjectFromAms 导入项目
	ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error)
	//ImportAPIFromAms 从ams中导入接口
	ImportAPIFromAms(projectID, groupID, userID int, apiList []entity.AmsAPIInfo) (bool, string, error)
}

//MonitorModulesDao monitorModule.go
type MonitorModulesDao interface {
	//GetMonitorModules 获取监控模块列表
	GetMonitorModules() (map[string]*entity.MonitorModule, error)
	//SetMonitorModule 设置监控模块
	SetMonitorModule(moduleName string, config string, moduleStatus int) error

	CheckModuleStatus(moduleName string) int
}

//NodeDao node.go
type NodeDao interface {
	//AddNode 新增节点信息
	AddNode(clusterID int, nodeName, nodeKey, listenAddress, adminAddress, gatewayPath string, groupID int) (int64, string, string, error)
	//EditNode 修改节点信息
	EditNode(nodeName, listenAddress, adminAddress, gatewayPath string, nodeID, groupID int) (string, error)
	//DeleteNode 删除节点信息
	DeleteNode(nodeID int) (string, error)
	// GetNodeList 获取节点列表
	GetNodeList(clusterID, groupID int, keyword string) ([]*entity.Node, error)
	//GetNodeInfo 获取节点信息
	GetNodeInfo(nodeID int) (*entity.Node, error)
	//GetNodeByKey 通过Key查询节点信息
	GetNodeByKey(nodeKey string) (*entity.Node, error)
	//GetAvaliableNodeListFromNodeList 从待操作节点中获取关闭节点列表
	GetAvaliableNodeListFromNodeList(nodeIDList string, nodeStatus int) (string, error)
	//BatchEditNodeGroup 批量修改节点分组
	BatchEditNodeGroup(nodeIDList string, groupID int) (string, error)
	//BatchDeleteNode 批量修改接口分组
	BatchDeleteNode(nodeIDList string) (string, error)
	//UpdateAllNodeClusterID 更新节点集群ID
	UpdateAllNodeClusterID(clusterID int)
	//GetNodeInfoAll get all node
	GetNodeInfoAll() ([]*entity.Node, error)

	GetHeartBeatTime(nodeKey string) (time.Time, error)

	SetHeartBeatTime(nodeKey string, heartBeatTime time.Time) error
}

//NodeGroupDao nodeGroup.go
type NodeGroupDao interface {
	//AddNodeGroup 新建节点分组
	AddNodeGroup(groupName string, clusterID int) (bool, interface{}, error)
	//EditNodeGroup 修改节点分组信息
	EditNodeGroup(groupName string, groupID int) (bool, string, error)
	//DeleteNodeGroup 删除节点分组
	DeleteNodeGroup(groupID int) (bool, string, error)
	//GetNodeGroupInfo 获取节点分组信息
	GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error)
	//GetNodeGroupList 获取节点分组列表
	GetNodeGroupList(clusterID int) (bool, []map[string]interface{}, error)
	//CheckNodeGroupIsExist 检查节点分组是否存在
	CheckNodeGroupIsExist(groupID int) (bool, error)
	//GetRunningNodeCount 获取分组内启动节点数量
	GetRunningNodeCount(groupID int) (bool, interface{}, error)
}

//PluginDao plugin.go
type PluginDao interface {
	//GetPluginInfo 获取插件配置信息
	GetPluginInfo(pluginName string) (bool, *entity.Plugin, error)
	// GetPluginList 获取插件列表
	GetPluginList(keyword string, condition int) (bool, []*entity.Plugin, error)
	// GetPluginCount 获取插件数量
	GetPluginCount() int
	// AddPlugin 新增插件信息
	AddPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error)
	// EditPlugin 修改插件信息
	EditPlugin(pluginName, pluginConfig, pluginDesc, version string, pluginPriority, isStop, pluginType int) (bool, string, error)
	// DeletePlugin 删除插件信息
	DeletePlugin(pluginName string) (bool, string, error)
	//CheckIndexIsExist 判断插件ID是否存在
	CheckIndexIsExist(pluginName string, pluginPriority int) (bool, error)
	//GetPluginConfig 获取插件配置及插件信息
	GetPluginConfig(pluginName string) (bool, string, error)
	//CheckNameIsExist 检查插件名称是否存在
	CheckNameIsExist(pluginName string) (bool, error)
	//EditPluginStatus 修改插件开启状态
	EditPluginStatus(pluginName string, pluginStatus int) (bool, error)
	//GetPluginListByPluginType 获取不同类型的插件列表
	GetPluginListByPluginType(pluginType int) (bool, []map[string]interface{}, error)
	//BatchStopPlugin 批量关闭插件
	BatchStopPlugin(pluginNameList string) (bool, string, error)
	//BatchStartPlugin 批量关闭插件
	BatchStartPlugin(pluginNameList string) (bool, string, error)
	//EditPluginCheckStatus 更新插件检测状态
	EditPluginCheckStatus(pluginName string, isCheck int) (bool, string, error)
}

//ProjectDao project.go
type ProjectDao interface {
	//AddProject 新建项目
	AddProject(projectName string) (bool, interface{}, error)
	//EditProject 修改项目信息
	EditProject(projectName string, projectID int) (bool, string, error)
	//DeleteProject 修改项目信息
	DeleteProject(projectID int) (bool, string, error)
	//BatchDeleteProject 批量删除项目
	BatchDeleteProject(projectIDList string) (bool, string, error)
	//GetProjectInfo 获取项目信息
	GetProjectInfo(projectID int) (bool, entity.Project, error)
	//GetProjectList 获取项目列表
	GetProjectList(keyword string) (bool, []*entity.Project, error)
	//CheckProjectIsExist 检查项目是否存在
	CheckProjectIsExist(projectID int) (bool, error)
	//GetAPIListFromProjectNotInStrategy 获取项目列表中没有被策略组绑定的接口
	GetAPIListFromProjectNotInStrategy() (bool, []map[string]interface{}, error)
}

//StrategyDao strategy.go
type StrategyDao interface {
	//AddStrategy 新增策略组
	AddStrategy(strategyName string, groupID, userID int) (bool, string, error)
	//EditStrategy 修改策略组信息
	EditStrategy(strategyID, strategyName string, groupID, userID int) (bool, string, error)
	//DeleteStrategy 删除策略组
	DeleteStrategy(strategyID string) (bool, string, error)
	// GetStrategyList 获取策略组列表
	GetStrategyList(groupID int, keyword string, condition, page, pageSize int) (bool, []*entity.Strategy, int, error)
	// GetOpenStrategy 获取策略组列表
	GetOpenStrategy() (bool, *entity.Strategy, error)
	//GetStrategyInfo 获取策略组信息
	GetStrategyInfo(strategyID string) (bool, *entity.Strategy, error)
	//CheckStrategyIsExist 检查策略组ID是否存在
	CheckStrategyIsExist(strategyID string) (bool, error)
	//BatchEditStrategyGroup 批量修改策略组分组
	BatchEditStrategyGroup(strategyIDList string, groupID int) (bool, string, error)
	//BatchDeleteStrategy 批量修改策略组
	BatchDeleteStrategy(strategyIDList string) (bool, string, error)
	//CheckIsOpenStrategy 判断是否是开放策略
	CheckIsOpenStrategy(strategyID string) bool
	//BatchUpdateStrategyEnableStatus 更新策略启动状态
	BatchUpdateStrategyEnableStatus(strategyIDList string, enableStatus int) (bool, string, error)
	// GetBalanceListInStrategy 获取在策略中的负载列表
	GetBalanceListInStrategy(strategyID string, balanceType int) (bool, []string, error)
	// CopyStrategy 复制策略
	CopyStrategy(strategyID string, newStrategyID string, userID int) (string, error)
	//GetStrategyIDList 获取策略ID列表
	GetStrategyIDList(groupID int, keyword string, condition int) (bool, []string, error)
}

//StrategyGroupDao strategyGroup.go
type StrategyGroupDao interface {
	//AddStrategyGroup 新建策略组分组
	AddStrategyGroup(groupName string) (bool, interface{}, error)
	//EditStrategyGroup 修改策略组分组
	EditStrategyGroup(groupName string, groupID int) (bool, string, error)
	//DeleteStrategyGroup 删除策略组分组
	DeleteStrategyGroup(groupID int) (bool, string, error)
	//GetStrategyGroupList 获取策略组分组列表
	GetStrategyGroupList() (bool, []map[string]interface{}, error)
	//CheckIsOpenGroup 判断是否是开放分组
	CheckIsOpenGroup(groupID int) bool
}

//StrategyPluginDao strategyPlugin.go
type StrategyPluginDao interface {
	//AddPluginToStrategy 新增策略组插件
	AddPluginToStrategy(pluginName, config, strategyID string) (bool, interface{}, error)
	//EditStrategyPluginConfig 新增策略组插件配置
	EditStrategyPluginConfig(pluginName, config, strategyID string) (bool, string, error)
	// GetStrategyPluginList 获取策略插件列表
	GetStrategyPluginList(strategyID, keyword string, condition int) (bool, []map[string]interface{}, error)
	//GetStrategyPluginConfig 通过策略组ID获取配置信息
	GetStrategyPluginConfig(strategyID, pluginName string) (bool, string, error)
	//CheckPluginIsExistInStrategy 检查策略组是否绑定插件
	CheckPluginIsExistInStrategy(strategyID, pluginName string) (bool, error)
	//GetStrategyPluginStatus 检查策略组插件是否开启
	GetStrategyPluginStatus(strategyID, pluginName string) (bool, error)
	//GetConnIDFromStrategyPlugin 获取Connid
	GetConnIDFromStrategyPlugin(pluginName, strategyID string) (bool, int, error)
	//BatchEditStrategyPluginStatus 批量修改策略组插件状态
	BatchEditStrategyPluginStatus(connIDList, strategyID string, pluginStatus int) (bool, string, error)
	//BatchDeleteStrategyPlugin 批量删除策略组插件
	BatchDeleteStrategyPlugin(connIDList, strategyID string) (bool, string, error)
	//CheckStrategyPluginIsExistByConnIDList 通过connIDList判断插件是否存在
	CheckStrategyPluginIsExistByConnIDList(connIDList, pluginName string) (bool, error)
}

//UpdaterDao updater
type UpdaterDao interface {
	//IsTableExist 检查table是否存在
	IsTableExist(name string) bool
	//IsColumnExist 检查列是否存在
	IsColumnExist(name string, column string) bool
	//GetTableVersion 获取当前表版本号
	GetTableVersion(name string) string
	//UpdateTableVersion 更新表版本号
	UpdateTableVersion(name, version string) error
	//GetGokuVersion 获取goku当前版本号
	GetGokuVersion() string
	//SetGokuVersion 设置goku版本号
	SetGokuVersion(version string) error
}

//UserDao user.go
type UserDao interface {
	//EditPassword 修改账户信息
	EditPassword(oldPassword, newPassword string, userID int) (bool, string, error)
	//GetUserInfo 获取账户信息
	GetUserInfo(userID int) (bool, interface{}, error)
	//GetUserType 获取用户类型
	GetUserType(userID int) (bool, interface{}, error)
	//CheckUserIsAdmin 判断是否是管理员
	CheckUserIsAdmin(userID int) (bool, string, error)
	//CheckUserIsSuperAdmin 判断是否是超级管理员
	CheckUserIsSuperAdmin(userID int) (bool, string, error)
	//CheckSuperAdminCount 获取超级管理员数量
	CheckSuperAdminCount() (int, error)
}

//VersionDao version.go
type VersionDao interface {
	//GetVersionList 获取版本列表
	GetVersionList(keyword string) ([]config.VersionConfig, error)
	//AddVersionConfig 新增版本配置
	AddVersionConfig(name, version, remark, config, balanceConfig, discoverConfig, now string, userID int) (int, error)
	EditVersionBasicConfig(name, version, remark string, userID, versionID int) error
	//BatchDeleteVersionConfig 批量删除版本配置
	BatchDeleteVersionConfig(ids []int, publishID int) error
	//PublishVersion 发布版本
	PublishVersion(id, userID int, now string) error
	//GetVersionConfigCount 获取版本配置数量
	GetVersionConfigCount() int
	//GetPublishVersionID 获取发布版本ID
	GetPublishVersionID() int
	//GetVersionConfig 获取当前版本配置
	GetVersionConfig() (*config.GokuConfig, map[string]map[string]*config.BalanceConfig, map[string]map[string]*config.DiscoverConfig, error)
}
