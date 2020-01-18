package cmd

import (
	"encoding/json"
	"errors"

	"github.com/eolinker/goku-api-gateway/config"
)

var (
	ErrorInvalidNodeInstance = errors.New("invalid instance value")
	ErrorInvalidNodeConfig   = errors.New("invalid instance config")
)

type Code string

const (
	None               Code = "none"
	NodeRegister       Code = "register"
	NodeRegisterResult Code = "register-result"
	NodeLevel          Code = "level"
	Config             Code = "config"
	Restart            Code = "restart"
	Stop               Code = "stop"
	Monitor            Code = "monitor"
	EventClientLeave   Code = "leave"
	Error              Code = "error"
)

func EncodeConfig(c *config.GokuConfig) ([]byte, error) {
	if c == nil {
		return nil, ErrorInvalidNodeConfig
	}

	return json.Marshal(c)
}
func DecodeConfig(data []byte) (*config.GokuConfig, error) {
	if len(data) == 0 {
		return nil, ErrorInvalidNodeConfig
	}
	c := new(config.GokuConfig)
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
