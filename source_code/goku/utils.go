package goku

import (
	"reflect"
)

// 判定handler是否是函数类型
func ValidateHandler(handler Handler) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("goku handler must be a callable func")
	}
}
