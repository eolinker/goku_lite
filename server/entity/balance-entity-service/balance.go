package entity

import "github.com/eolinker/goku/server/driver"

type Balance struct {
	Name          string
	ServiceName   string
	ServiceDriver string
	ServiceType   string
	AppName       string
	Static        string
	StaticCluster string
	Desc          string
	CreateTime    string
	UpdateTime    string
}

func (e *Balance) Type() *Balance {

	if e != nil {

		d, has := driver.Get(e.ServiceDriver)
		if has {
			e.ServiceType = d.Type
		}
	}

	return e
}
