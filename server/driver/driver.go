package driver

const (
	Static    = "static"
	Discovery = "discovery"
)

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
func All() []*Driver {
	return all
}

func GetByType(t string) []*Driver {
	return driverOfType[t]
}

func Get(name string) (*Driver, bool) {
	d, has := drivers[name]
	return d, has
}

func TypeName(t string) (string, bool) {
	n, has := typeNames[t]
	return n, has
}

func Types() map[string]string {
	return typeNames
}
