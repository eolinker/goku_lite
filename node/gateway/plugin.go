package gateway

import (
	"reflect"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
	plugin_executor "github.com/eolinker/goku-api-gateway/node/gateway/plugin-executor"
	plugin "github.com/eolinker/goku-api-gateway/node/plugin-loader"
)

func genBeforPlugin(cfgs []*config.PluginConfig, cluster string) []plugin_executor.Executor {
	ps := make([]plugin_executor.Executor, 0, len(cfgs))
	for _, cfg := range cfgs {

		factory, e := plugin.LoadPlugin(cfg.Name)
		if e != nil {
			continue
		}
		obj, err := factory.Create(cfg.Config, cluster, cfg.UpdateTag, "", 0)
		if err != nil {
			continue
		}

		if obj.BeforeMatch != nil && !reflect.ValueOf(obj.BeforeMatch).IsNil() {
			ps = append(ps, plugin_executor.NewBeforeExecutor(cfg, obj.BeforeMatch))
		}
	}
	return ps
}

func genPlugins(cfgs []*config.PluginConfig, cluster string, strategyID string, apiID int) ([]plugin_executor.Executor, []plugin_executor.Executor, []plugin_executor.Executor) {
	psBefor := make([]plugin_executor.Executor, 0, len(cfgs))
	psAccess := make([]plugin_executor.Executor, 0, len(cfgs))
	psProxy := make([]plugin_executor.Executor, 0, len(cfgs))
	for _, cfg := range cfgs {

		factory, e := plugin.LoadPlugin(cfg.Name)
		if e != nil {
			continue
		}
		obj, err := factory.Create(cfg.Config, cluster, cfg.UpdateTag, strategyID, apiID)
		if err != nil {
			continue
		}

		if obj.BeforeMatch != nil && !reflect.ValueOf(obj.BeforeMatch).IsNil() {
			psBefor = append(psBefor, plugin_executor.NewBeforeExecutor(cfg, obj.BeforeMatch))
		}
		if obj.Access != nil && !reflect.ValueOf(obj.Access).IsNil() {
			if strings.Contains(cfg.Name, "_auth") {
				cfg.IsAuth = true
			}
			psAccess = append(psAccess, plugin_executor.NewAccessExecutor(cfg, obj.Access))
		}
		if obj.Proxy != nil && !reflect.ValueOf(obj.Proxy).IsNil() {
			psProxy = append(psProxy, plugin_executor.NewProxyExecutor(cfg, obj.Proxy))
		}
	}
	return psBefor, psAccess, psProxy

}
