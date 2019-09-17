package balance

import (
	"errors"
	"fmt"
	"github.com/eolinker/goku/console/module/service"
	driver2 "github.com/eolinker/goku/server/driver"
	entity "github.com/eolinker/goku/server/entity/balance-entity-service"
	"time"

	"github.com/eolinker/goku/server/dao"
	dao_balance "github.com/eolinker/goku/server/dao/dao-balance"
)

func Add(info *Param) (string, error) {
	const TableName = "goku_balance"
	serviceInfo,err:= service.Get(info.ServiceName)
	if err!=nil{
		return fmt.Sprintf("serviceName:%s",err.Error()),err
	}
	switch serviceInfo.Type {
	case driver2.Static:{
		if info.Static == ""&& info.StaticCluster ==""{
			return "param:static 和 staticCluster 不能同时为空",errors.New( "param:static 和 staticCluster 不能同时为空")
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		result, err :=dao_balance.AddStatic(info.Name,info.ServiceName,info.Static,info.StaticCluster,info.Desc,now)
		if err == nil {
			dao.UpdateTable(TableName)
		}
		return result, err
	}
	case driver2.Discovery:{
		if info.AppName == ""{
			return "param:appName 不能为空",errors.New( "param:appName 不能为空")
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		result, err :=dao_balance.AddDiscovery(info.Name,info.ServiceName,info.AppName,info.Desc,now)
		if err == nil {
			dao.UpdateTable(TableName)
		}
		return result, err
	}

	}

	return "无效serviceName", errors.New("invalid serviceName")
}


func Save(info *Param) (string, error) {
	const TableName = "goku_balance"
	serviceInfo,err:= service.Get(info.ServiceName)
	if err!=nil{
		return fmt.Sprintf("serviceName:%s",err.Error()),err
	}
	switch serviceInfo.Type {
	case driver2.Static:{
		if info.Static == ""&& info.StaticCluster ==""{
			return "param:static 和 staticCluster 不能同时为空",errors.New( "param:static 和 staticCluster 不能同时为空")
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		result, err :=dao_balance.SaveStatic(info.Name,info.ServiceName,info.Static,info.StaticCluster,info.Desc,now)
		if err == nil {
			dao.UpdateTable(TableName)
		}
		return result, err
	}
	case driver2.Discovery:{
		if info.AppName == ""{
			return "param:appName 不能为空",errors.New( "param:appName 不能为空")
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		result, err :=dao_balance.SaveDiscover(info.Name,info.ServiceName,info.AppName,info.Desc,now)
		if err == nil {
			dao.UpdateTable(TableName)
		}
		return result, err
	}

	}

	return "无效serviceName", errors.New("invalid serviceName")
}
func Get(name string) (*Info, error) {
	b, e := dao_balance.Get(name)
	if e != nil {
		return nil, e
	}

	return ReadInfo(b),nil
}
func Search(keyworkd string)([]*Info, error){
	var entities []*entity.Balance
	if keyworkd == ""{
		es, e:= dao_balance.GetAll()
		if e != nil {
			return nil, e
		}
		entities = es
	}else{
		es, e:= dao_balance.Search(keyworkd)
		if e != nil {
			return nil, e
		}
		entities = es
	}

	infos := make([]*Info, 0, len(entities))

	for _, ent := range entities {
		infos = append(infos, ReadInfo(ent))
	}
	return infos, nil
}
func GetAll() ([]*Info, error) {

	entities, e := dao_balance.GetAll()
	if e != nil {
		return nil, e
	}
	infos := make([]*Info, 0, len(entities))

	for _, ent := range entities {
		infos = append(infos, ReadInfo(ent))
	}
	return infos, nil
}

func Delete(name string) (string, error) {
	tableName := "goku_balance"
	result, err := dao_balance.Delete(name)
	if err == nil {
		dao.UpdateTable(tableName)
	}
	return result, err
}

func GetBalancNames() (bool, []string, error) {
	return dao_balance.GetBalanceNames()
}

// 批量删除负载
func BatchDeleteBalance(balanceNames []string) (string, error) {
	tableName := "goku_balance"
	result, err := dao_balance.BatchDelete(balanceNames)
	if err == nil {
		dao.UpdateTable(tableName)
	}
	return result, err
}
