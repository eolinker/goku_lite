package balance

import (
	"github.com/eolinker/goku/goku-service/application"
	"github.com/eolinker/goku/goku-service/discovery"
)

func ResetBalances(balances []*Balance)  {

	bmap:=make(map[string]*Balance)

	for _,b:=range balances{
		bmap[b.Name] = b

	}

	manager.set(bmap)

}

func GetByName(name string)(application.IHttpApplication,bool)  {
	b,has:=manager.get(name)
	if !has{
		return application.NewOrg(name),true
	}

	sources,has:=discovery.GetDiscoverer(b.Discovery)
	if has{

		service, handler, yes:= sources.GetApp(b.AppConfig)
		if yes{
			return application.NewApplication(service,handler),true
		}
	}

	return nil,false
}