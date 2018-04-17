package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	Configure map[string]string
)

func init() {
	Configure = make(map[string]string)
}

func ReadConfigure(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &Configure)
	return
}
