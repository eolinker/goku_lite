package monitor

import (
	"github.com/eolinker/goku-api-gateway/common/general"
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/ksitigarbha"
	"github.com/eolinker/goku-api-gateway/server/dao"
	"github.com/pkg/errors"
)
var(
	monitorModuleDao dao.MonitorModulesDao
)

func init() {
	pdao.Need(&monitorModuleDao)
	general.RegeditLater(InitModuleStatus)
}

type MonitorModule struct {
	Name         string      `json:"moduleName"`
	Config       interface{} `json:"config,omitempty"`
	ModuleStatus int         `json:"moduleStatus"`
	Desc         string      `json:"moduleDesc"`
	Models       interface{} `json:"layer"`
}
// 初始化监控模块的配置状态
func InitModuleStatus() error {
	modules, err := monitorModuleDao.GetMonitorModules()
	if err != nil {
		return  err
	}

	names := ksitigarbha.GetMonitorModuleNames()


	for _, name := range names {

		if m,has:= modules[name];has{
			if m.ModuleStatus == 1{
				ksitigarbha.Open(name,m.Config)
			}else{
				ksitigarbha.Close(name)
			}
		}else{
			ksitigarbha.Close(name)
		}
	}
	return nil
}
//GetMonitorModules 获取监控模块列表
func GetMonitorModules() ([]*MonitorModule, error) {
	m, err := monitorModuleDao.GetMonitorModules()
	if err != nil {
		return nil, err
	}

	names := ksitigarbha.GetMonitorModuleNames()
	modules := make([]*MonitorModule, 0, len(names))

	for _, name := range names {
		model, _ := ksitigarbha.GetMonitorModuleModel(name)
		mod := &MonitorModule{
			Name:         name,
			Config:       model.GetDefaultConfig(),
			ModuleStatus: 0,
			Desc:         model.GetDesc(),
			Models:       model.GetModel(),
		}

		v, ok := m[name]
		if ok {
			mod.ModuleStatus = v.ModuleStatus
			c, err := model.Decode(v.Config)
			if err == nil {
				mod.Config = c
			}
		}

		modules = append(modules, mod)
	}
	return modules, nil
}

func SetMonitorModule(moduleName string, config string, moduleStatus int) error {

	model, has := ksitigarbha.GetMonitorModuleModel(moduleName)
	if !has {
		return errors.New("[error]the module does not exist")
	}

	if moduleStatus == 1 {

		_, err := model.Decode(config)
		if err != nil {
			//errInfo := "[error]invalid config"
			return err
		}
	}

	err := monitorModuleDao.SetMonitorModule(moduleName, config, moduleStatus)
	if err != nil {
		return err
	}
	return nil
}
