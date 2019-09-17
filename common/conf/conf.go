// 读取配置信息
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

func Get(name string) (string, bool) {
	v, has := _Configure[name]
	return v, has
}
func Set(name, value string) {
	_Configure[name] = value
}

func Value(name string) string {
	return _Configure[name]
}
func Reload() {
	ReadConfigure(lastFile)
}
func MastValue(name string, def string) string {
	v, h := _Configure[name]
	if h {
		return v
	}
	return def
}
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

// 更新配置文件
func Save() (bool, error) {
	//file, err := os.OpenFile(lastFile, os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	confStr, err := yaml.Marshal(_Configure)
	if err != nil {
		return false, err
	}

	ioutil.WriteFile(lastFile, confStr, 0666)
	return true, nil
}
