package monitor

import (
	"github.com/eolinker/goku-api-gateway/ksitigarbha"

	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	"github.com/pkg/errors"
)

type MonitorModule struct {
	Name         string      `json:"moduleName"`
	Config       interface{} `json:"config,omitempty"`
	ModuleStatus int         `json:"moduleStatus"`
	Desc         string      `json:"moduleDesc"`
	Models       interface{} `json:"layer"`
}

//GetMonitorModules 获取监控模块列表
func GetMonitorModules() ([]*MonitorModule, error) {
	m, err := console_sqlite3.GetMonitorModules()
	if err != nil {
		return make([]*MonitorModule, 0), nil
	}

	names := ksitigarbha.GetMonitorModuleNames()
	modules := make([]*MonitorModule, 0, len(names))

	for _, name := range names {
		model,_ := ksitigarbha.GetMonitorModuleModel(name)
		  mod :=&MonitorModule{
			  Name:         name,
			  Config:       model.GetDefaultConfig(),
			  ModuleStatus: 0,
			  Desc:          model.GetDesc(),
			  Models:        model.GetModel(),
		  }

		v, ok := m[name]
		if ok {
			mod.ModuleStatus = v.ModuleStatus
			c ,err := model.Decode(v.Config)
			if err == nil {
				mod.Config = c
			}
		}

		modules = append(modules, mod)
	}
	return modules, nil
}

func SetMonitorModule(moduleName string, config string, moduleStatus int) error {

	model,has := ksitigarbha.GetMonitorModuleModel(moduleName)
	if !has {
		return errors.New("[error]the module does not exist")
	}

	if moduleStatus == 1  {

		_ ,err:= model.Decode(config)
		if err != nil{
			//errInfo := "[error]invalid config"
			return err
		}

	}

	err := console_sqlite3.SetMonitorModule(moduleName, config, moduleStatus)
	if err != nil {
		return err
	}
	return nil
}
