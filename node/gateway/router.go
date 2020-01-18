package gateway

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-service/balance"
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
	"github.com/eolinker/goku-api-gateway/node/gateway/application"
	plugin_executor "github.com/eolinker/goku-api-gateway/node/gateway/plugin-executor"
	plugin_loader "github.com/eolinker/goku-api-gateway/node/plugin-loader"
	"github.com/eolinker/goku-api-gateway/node/router"
)

var (
	errorConfig = errors.New("config is error")
)

//Parse 解析
func Parse(config *config.GokuConfig, factory router.Factory) (http.Handler, error) {

	if config == nil {
		return nil, errorConfig
	}

	f := genFactory(config, factory)

	return &HTTPHandler{router: f.create()}, nil
}

type _RootFactory struct {
	orgCfg        *config.GokuConfig
	beforePlugin  []plugin_executor.Executor
	gBefores      []plugin_executor.Executor
	gAccesses     []plugin_executor.Executor
	gProxies      []plugin_executor.Executor
	apis          map[int]*config.APIContent
	appFactory    *application.Factory
	routerFactory router.Factory
	cluster       string

	authPlugin map[string]string
}

func (f *_RootFactory) create() *Before {
	beforeRouter := &Before{
		pluginBefor:       f.beforePlugin,
		pluginGlobalBefor: f.gBefores,
		strategies:        f.createStrategy(),
		anonymousStrategy: f.orgCfg.AnonymousStrategyID,
	}

	return beforeRouter
}
func (f *_RootFactory) createStrategy() map[string]*Strategy {
	strategys := make(map[string]*Strategy)

	for _, cfg := range f.orgCfg.Strategy {

		strategyRouter := f.genStrategy(cfg)
		strategys[cfg.ID] = strategyRouter
	}

	return strategys
}

// 构造策略
func (f *_RootFactory) genStrategy(cfg *config.StrategyConfig) *Strategy {
	_, accesses, proxies := genPlugins(cfg.Plugins, f.cluster, cfg.ID, 0)

	s := &Strategy{
		ID:     cfg.ID,
		Name:   cfg.Name,
		Enable: cfg.Enable,

		accessPlugin:       accesses,
		globalAccessPlugin: f.gAccesses,
		authPlugin:         make(map[string]plugin_executor.Executor),
		isNeedAuth:         false,
	}
	if !s.Enable {
		return s
	}
	s.apiRouter = f.routerFactory.New()
	s.apiRouter.AddNotFound(s.HandlerAPINotFound)
	for authKey, authCfg := range cfg.AUTH {
		s.isNeedAuth = true
		pluginName, has := f.authPlugin[authKey]
		if has {
			pluginFactory, e := plugin_loader.LoadPlugin(pluginName)
			if e != nil {
				continue
			}
			pluginObj, err := pluginFactory.Create(authCfg, f.cluster, "", s.ID, 0)
			if err != nil {
				continue
			}

			if pluginObj.Access != nil && !reflect.ValueOf(pluginObj.Access).IsNil() {
				if cfg.ID == "68YAY7" {
					fmt.Println(authKey, pluginName)
				}
				s.authPlugin[authKey] = plugin_executor.NewAccessExecutor(&config.PluginConfig{
					Name:   pluginName,
					IsStop: true,
					Config: authCfg,
					IsAuth: true,
				}, pluginObj.Access)
			}
		}
	}

	factory := newAPIFactory(f, s.ID)
	for _, apiCfg := range cfg.APIS {

		iRouter, apiContent := factory.genAPIRouter(apiCfg, proxies)
		if iRouter == nil {
			continue
		}
		for _, method := range apiContent.Methods {

			s.apiRouter.AddRouter(strings.ToUpper(method), apiContent.RequestURL, iRouter)
		}
	}

	return s
}

type _ApiFactory struct {
	root       *_RootFactory
	strategyID string
}

func newAPIFactory(root *_RootFactory, strategyID string) *_ApiFactory {
	return &_ApiFactory{
		root:       root,
		strategyID: strategyID,
	}
}
func (f *_ApiFactory) genAPIRouter(cfg *config.APIOfStrategy, proxies []plugin_executor.Executor) (router.IRouter, *config.APIContent) {

	apiContend, has := f.root.apis[cfg.ID]
	if !has {
		return nil, nil
	}

	app, err := f.root.appFactory.GenApplication(cfg)
	if err != nil {
		return nil, nil
	}
	_, pluginAccesses, pluginProxies := genPlugins(cfg.Plugins, f.root.cluster, f.strategyID, cfg.ID)
	pro := make([]plugin_executor.Executor, 0, len(proxies)+len(pluginProxies))
	pro = append(pro, proxies...)
	pro = append(pro, pluginProxies...)

	return &API{
		strategyID:          f.strategyID,
		apiID:               cfg.ID,
		app:                 app,
		pluginAccess:        pluginAccesses,
		pluginProxies:       pro,
		pluginAccessGlobal:  f.root.gAccesses,
		pluginProxiesGlobal: f.root.gProxies,
	}, apiContend
}

func genFactory(cfg *config.GokuConfig, factory router.Factory) *_RootFactory {

	discovery.ResetAllServiceConfig(cfg.DiscoverConfig)
	balance.ResetBalances(cfg.Balance)

	beforePlugin := genBeforPlugin(cfg.Plugins.BeforePlugins, cfg.Cluster)

	gBefores, gAccesses, gProxies := genPlugins(cfg.Plugins.GlobalPlugins, cfg.Cluster, "", 0)
	apis := toMap(cfg.APIS)
	return &_RootFactory{
		beforePlugin:  beforePlugin,
		gBefores:      gBefores,
		gAccesses:     gAccesses,
		gProxies:      gProxies,
		appFactory:    application.NewFactory(apis),
		apis:          apis,
		routerFactory: factory,
		cluster:       cfg.Cluster,
		orgCfg:        cfg,
		authPlugin:    cfg.AuthPlugin,
	}
}

func toMap(cfgs []*config.APIContent) map[int]*config.APIContent {
	m := make(map[int]*config.APIContent)
	for _, cfg := range cfgs {
		//if cfg.ID == 1880 {
		//	fmt.Println(cfg)
		//}
		m[cfg.ID] = cfg
	}
	return m
}
