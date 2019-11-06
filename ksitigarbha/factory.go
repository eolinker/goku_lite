package ksitigarbha

type IModule interface {
	GetModel() []Model
	GetDesc() string
	GetName() string
	GetNameSpace()string
	GetDefaultConfig() interface{}
	//CheckConfig(interface{}) bool
	Decode(config string) (interface{},error)
	Encoder(v interface{}) (string,error)

}
