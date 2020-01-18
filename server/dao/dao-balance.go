package dao

import (
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity-service"
)

//BalanceDao balanceDao
type BalanceDao interface {
	//RegisterDao add
	//Add(name, serviceName, desc, appName, static, staticCluster, now string) (string, error)

	//AddStatic 新增静态负载
	AddStatic(name, serviceName, static, staticCluster, desc, now string) (string, error)
	//SaveStatic 保存静态负载信息
	SaveStatic(name, serviceName, static, staticCluster, desc string, now string) (string, error)

	//SaveDiscover 保存服务发现信息
	SaveDiscover(name, serviceName, appName, desc string, now string) (string, error)
	//AddDiscovery 新增服务发现
	AddDiscovery(name, serviceName, appName, desc, now string) (string, error)
	//Save save
	//Save(name, desc, static, staticCluster, now string) (string, error)

	//Delete 删除负载
	Delete(name string) (string, error)

	//BatchDelete 批量删除负载
	BatchDelete(balanceNames []string) (string, error)
	//GetBalanceNames 获取负载名称列表
	GetBalanceNames() (bool, []string, error)
	//Get 根据负载名获取负载配置
	Get(name string) (*entity.Balance, error)

	//GetAll 获取所有负载配置
	GetAll() ([]*entity.Balance, error)

	//Search 关键字获取负载列表
	Search(keyword string) ([]*entity.Balance, error)

	GetUseBalanceNames() (map[string]int, error)
}
