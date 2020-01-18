package ksitigarbha

//ConfigHandler configHandler
type ConfigHandler interface {
	OnOpen(namespace, name, config string)
	OnClose(namespace, name string)
}

type _H interface {
	Handler(namespace string, handler ConfigHandler)
	Close(namespace string)
	Open(namespace string, config string)
}
