package utils

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

//JSObjectToJSON 将js对象转为json
func JSObjectToJSON(s string) ([]byte, error) {
	vm := otto.New()
	v, err := vm.Run(fmt.Sprintf(`
		cs = %s
		JSON.stringify(cs)
`, s))
	if err != nil {
		return nil, err
	}
	return []byte(v.String()), nil
}
