package cmd

import (
	"encoding/json"
	"github.com/eolinker/goku-api-gateway/config"
)

type RegisterResult struct {
	Code   int
	Error  string
	Config *config.GokuConfig
}

func DecodeRegisterResult(data []byte) (*RegisterResult, error) {
	r := new(RegisterResult)
	err := json.Unmarshal(data, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func EncodeRegisterResultConfig(c *config.GokuConfig) ([]byte, error) {
	r := RegisterResult{
		Code:   0,
		Error:  "",
		Config: c,
	}
	return json.Marshal(r)
}
func EncodeRegisterResultError(err string) ([]byte, error) {
	r := RegisterResult{
		Code:   -1,
		Error:  err,
		Config: nil,
	}
	return json.Marshal(r)
}

func DecodeRegister(data []byte) (string, error) {
	if len(data) == 32 {
		return string(data), nil
	}
	return "", ErrorInvalidNodeInstance
}

func EncodeRegister(nodeKey string) ([]byte, error) {
	data := []byte(nodeKey)
	if len(data) == 32 {
		return data, nil
	}
	return nil, ErrorInvalidNodeInstance
}
