package service

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"

	driver2 "github.com/eolinker/goku-api-gateway/server/driver"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

const _TableName = "goku_service_config"
var (
	serviceDao dao.ServiceDao
)

func init() {
	pdao.Need(&serviceDao)
}
//RegisterDao 新增服务发现
func Add(param *AddParam) error {
	err := serviceDao.Add(param.Name, param.Driver, param.Desc, param.Config, param.ClusterConfig, false, param.HealthCheck, param.HealthCheckPath, param.HealthCheckCode, param.HealthCheckPeriod, param.HealthCheckTimeOut)

	return err
}

//Save 保存服务发现
func Save(param *AddParam) error {

	v, e := serviceDao.Get(param.Name)
	if e != nil {
		return e
	}

	if v.Driver != param.Driver {
		return fmt.Errorf("not allowed change dirver from %s to %s for service", v.Driver, param.Driver)
	}

	err := serviceDao.Save(param.Name, param.Desc, param.Config, param.ClusterConfig, param.HealthCheck, param.HealthCheckPath, param.HealthCheckCode, param.HealthCheckPeriod, param.HealthCheckTimeOut)

	return err
}

//Get 通过名称获取服务发现信息
func Get(name string) (*Info, error) {
	v, err := serviceDao.Get(name)
	if err != nil {
		return nil, err
	}

	return &Info{
		Service:            tran(v),
		Config:             v.Config,
		ClusterConfig:      v.ClusterConfig,
		HealthCheckPath:    v.HealthCheckPath,
		HealthCheckPeriod:  v.HealthCheckPeriod,
		HealthCheckCode:    v.HealthCheckCode,
		HealthCheckTimeOut: v.HealthCheckTimeOut,
	}, nil
}

//Delete 批量删除服务发现
func Delete(names []string) error {

	for _, n := range names {
		if !ValidateName(n) {
			return fmt.Errorf("invalid name:%s", n)
		}
	}

	return serviceDao.Delete(names)
}

//SetDefaut 设置默认服务发现
func SetDefaut(name string) error {
	return serviceDao.SetDefault(name)
}
func tran(v *entity.Service) *Service {
	s := &Service{
		Simple: Simple{
			Name:   v.Name,
			Driver: v.Driver,
		},
		Desc:        v.Desc,
		IsDefault:   v.IsDefault,
		HealthCheck: v.HealthCheck,
		UpdateTime:  v.UpdateTime,
		CreateTime:  v.CreateTime,
	}

	d, has := driver2.Get(v.Driver)
	if has {
		s.DriverTitle = d.Title
		s.Type = d.Type
	} else {
		s.DriverTitle = "unknown"
		s.Type = "unknown"
	}
	return s
}

//List 获取服务发现列表
func List(keyword string) ([]*Service, error) {
	vs, e := serviceDao.List(keyword)
	if e != nil {
		return nil, e
	}
	list := make([]*Service, 0, len(vs))

	for _, v := range vs {

		list = append(list, tran(v))

	}
	return list, nil
}

//SimpleList 获取简易服务发现列表
func SimpleList() ([]*Simple, string, error) {
	vs, e := serviceDao.List("")
	if e != nil {
		return nil, "", e
	}
	list := make([]*Simple, 0, len(vs))
	defaultName := ""
	for _, v := range vs {

		if v.IsDefault {
			defaultName = v.Name
		}
		s := &Simple{
			Name:   v.Name,
			Driver: v.Driver,
		}

		d, has := driver2.Get(v.Driver)
		if has {
			s.DriverTitle = d.Title
			s.Type = d.Type
		} else {
			s.DriverTitle = "unknown"
			s.Type = "unknown"
		}

		list = append(list, s)
	}
	return list, defaultName, nil
}
