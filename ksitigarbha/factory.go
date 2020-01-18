package ksitigarbha

//IModule IModule
type IModule interface {
	GetModel() []Model
	GetDesc() string
	GetName() string
	GetNameSpace() string
	GetDefaultConfig() interface{}
	//CheckConfig(interface{}) bool
	Decode(config string) (interface{}, error)
	Encode(v interface{}) (string, error)
}
