package console

import (
	"github.com/eolinker/goku-api-gateway/config"
)

//ConfigConsole configConsole
type ConfigConsole interface {
	Close()
	AddListen(callback ConfigCallbackFunc)
	GetConfig() (*config.GokuConfig, error)
	RegisterToConsole() (*config.GokuConfig, error)
	Listen()
}

//ConfigCallbackFunc configCallbackFunc
type ConfigCallbackFunc func(conf *config.GokuConfig)
