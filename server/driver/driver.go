package driver

const (
	//Static 静态服务
	Static = "static"
	//Discovery 服务发现
	Discovery = "discovery"
)

//Driver driver
type Driver struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var (
	all          []*Driver
	typeNames    map[string]string
	driverOfType map[string][]*Driver
	drivers      map[string]*Driver
)

func init() {
	typeNames = map[string]string{
		Static:    "静态服务",
		Discovery: "服务发现",
	}
	all = []*Driver{
		{
			Name:  "static",
			Title: "Static",
			Type:  Static,
			Desc:  "静态服务",
		},
		{
			Name:  "eureka",
			Title: "Eureka",
			Type:  Discovery,
			Desc:  "Eureka服务发现",
		},
		{
			Name:  "consul",
			Type:  Discovery,
			Title: "Consul",
			Desc:  "Consul catalog",
		},
	}

	drivers = make(map[string]*Driver)
	driverOfType = make(map[string][]*Driver)
	for t := range typeNames {
		driverOfType[t] = make([]*Driver, 0, len(all))
	}

	for _, d := range all {
		drivers[d.Name] = d

		driverOfType[d.Type] = append(driverOfType[d.Type], d)
	}
}

//All all
func All() []*Driver {
	return all
}

//GetByType 根据类型获取驱动列表
func GetByType(t string) []*Driver {
	return driverOfType[t]
}

//Get get
func Get(name string) (*Driver, bool) {
	d, has := drivers[name]
	return d, has
}

//TypeName typeName
func TypeName(t string) (string, bool) {
	n, has := typeNames[t]
	return n, has
}

//Types types
func Types() map[string]string {
	return typeNames
}
