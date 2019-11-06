package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

//ReadConfig 读取文件配置
func ReadConfig(file string) (*GokuConfig, error) {
	fp, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	c := &GokuConfig{}
	e := json.Unmarshal(data, c)

	return c, e
}
