package conf

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	_Configure map[string]string
	lastFile   = ""
)

//Get get
func Get(name string) (string, bool) {
	v, has := _Configure[name]
	return v, has
}

//Set set
func Set(name, value string) {
	_Configure[name] = value
}

//Value value
func Value(name string) string {
	return _Configure[name]
}

//Reload reload
func Reload() {
	ReadConfigure(lastFile)
}

//MastValue mastValue
func MastValue(name string, def string) string {
	v, h := _Configure[name]
	if h {
		return v
	}
	return def
}

//ReadConfigure 读取配置
func ReadConfigure(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	configure := make(map[string]string)
	err = yaml.Unmarshal(content, &configure)
	if err != nil {
		return err
	}
	_Configure = configure
	lastFile = filepath
	return nil
}

//Save 更新配置文件
func Save() (bool, error) {
	confStr, err := yaml.Marshal(_Configure)
	if err != nil {
		return false, err
	}

	ioutil.WriteFile(lastFile, confStr, 0666)
	return true, nil
}
