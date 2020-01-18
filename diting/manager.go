package diting

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/ksitigarbha"
)

var (
	constructorMap = make(map[string]Constructor)
	refresher      = NewRefreshers()
)

//Register register
func Register(namespace string, constructor Constructor) {

	_, has := constructorMap[namespace]
	if has {
		panic(fmt.Sprint("duplicate namespace of constructor by", namespace))
	}
	constructorMap[namespace] = constructor

}

func get(namespace string) (Constructor, bool) {

	constructor, has := constructorMap[namespace]
	return constructor, has

}

func construct(confs map[string]string) Factories {

	lives:= make(map[string]int)
	factories := make(Factories,0,len(confs))
	defer func() {
		// close 关闭不用的模块

		for name,constructor:= range constructorMap{

			if _,has:=lives[name];!has{
				constructor.Close()
			}
		}
	}()
	if confs == nil{
		return factories
	}

	for name, conf := range confs {
		namespace,_:=ksitigarbha.GetNameSpaceByName(name)
		lives[namespace] = 1
		constructor, has := get(namespace)
		if !has {
			continue
		}

		factory, err := constructor.Create(conf)
		if err != nil {
			continue
		}

		factories = append(factories, factory)
	}
	return factories
}

//Refresh refresh
func Refresh(confs map[string]string) {
	factories := construct(confs)
	refresher.Refresh(factories)
}
