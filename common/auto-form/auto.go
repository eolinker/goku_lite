package auto

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

//Values values
type Values interface {
	Get(key string) (string, bool)
}

//ValuesForm valuesForm
type ValuesForm map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v ValuesForm) Get(key string) (string, bool) {
	if v == nil {
		return "", false
	}
	vs := v[key]
	if len(vs) == 0 {
		return "", false
	}
	return vs[0], true
}

type (
	// DefaultsHandle  自定义部分默认值的初始化
	DefaultsHandle interface {
		SetDefaults() error
	}
	// ContextHandle  自定义部分从context中初始化字段
	ContextHandle interface {
		SetContext(ctx Values) error
	}
	//CheckOptHandler 检查配置处理器
	CheckOptHandler interface {
		IsOpt(key string) bool
	}
	// ConfigValidate 自定义校验
	ConfigValidate interface {
		Validate() error
	}
)

//IsOpt 判断是否是opt标签
func IsOpt(key string, c interface{}) bool {
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		opt := field.Tag.Get("opt")
		if opt == key {
			return true
		}
	}
	hh, ok := c.(CheckOptHandler)
	if ok {
		return hh.IsOpt(key)
	}
	return false
}

//ErrorNil 当target为nil时，报错
var ErrorNil = errors.New("target is nil")

// SetValues 从url.Values中完成配置字段，context中不存在时，使用 default
// 字段格式  opt:"name,require" default:"default value" min:"min value" max:"max value"
// require 为可选，表示该字段是否为必填
func SetValues(values url.Values, c interface{}) error {
	if c == nil {
		return ErrorNil
	}
	return setValues(ValuesForm(values), c)
}

func setValues(ctx Values, c interface{}) error {

	val := reflect.ValueOf(c).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {

		field := typ.Field(i)
		tag, has := field.Tag.Lookup("opt")
		if !has {
			continue
		}
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if !isValidTag(name) {
			continue
		}
		value, has := ctx.Get(name)
		if has {
			fieldVal := val.FieldByName(field.Name)
			if err := set(&field, &fieldVal, value); err != nil {
				return err
			}
		} else {
			if opts.Contains("require") {
				return fmt.Errorf("require value of [%s] but has no", name)
			}

			defaultVal := field.Tag.Get("default")
			if defaultVal == "" || name == "" {
				continue
			}
			fieldVal := val.FieldByName(field.Name)
			if err := set(&field, &fieldVal, defaultVal); err != nil {
				return err
			}
		}
	}
	hh, ok := c.(ContextHandle)
	if ok {
		if err := hh.SetContext(ctx); err != nil {
			return err
		}
	}

	return nil
}

// SetDefaults 对目标设置default
func SetDefaults(c interface{}) error {
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		opt := field.Tag.Get("opt")
		defaultVal := field.Tag.Get("default")
		if defaultVal == "" || opt == "" {
			continue
		}
		fieldVal := val.FieldByName(field.Name)

		if err := set(&field, &fieldVal, defaultVal); err != nil {
			return err
		}
	}
	hh, ok := c.(DefaultsHandle)
	if ok {
		if err := hh.SetDefaults(); err != nil {
			return err
		}
	}
	return nil
}

// Set values based on parameters in StructTags
func set(field *reflect.StructField, fieldVal *reflect.Value, value interface{}) error {

	//fieldVal := val.FieldByName(field.Name)
	dest := unsafeValueOf(fieldVal)
	coercedVal, err := coerce(value, field.Type)
	if err != nil {
		return fmt.Errorf("failed to coerce option %s (%v) - %s",
			field.Tag.Get("opt"), value, err)
	}
	if min, has := field.Tag.Lookup("min"); has && min != "" {
		coercedMinVal, _ := coerce(min, field.Type)
		if valueCompare(coercedVal, coercedMinVal) == -1 {
			return fmt.Errorf("invalid %s ! %v < %v",
				field.Tag.Get("opt"), coercedVal.Interface(), coercedMinVal.Interface())
		}
	}
	if max, has := field.Tag.Lookup("max"); has && max != "" {
		coercedMaxVal, _ := coerce(max, field.Type)
		if valueCompare(coercedVal, coercedMaxVal) == 1 {
			return fmt.Errorf("invalid %s ! %v > %v",
				field.Tag.Get("opt"), coercedVal.Interface(), coercedMaxVal.Interface())
		}
	}
	dest.Set(coercedVal)
	return nil

}

// because Config contains private structs we can't use reflect.Value
// directly, instead we need to "unsafely" address the variable
func unsafeValueOf(val *reflect.Value) reflect.Value {
	uptr := unsafe.Pointer(val.UnsafeAddr())
	return reflect.NewAt(val.Type(), uptr).Elem()
}

// Validate 校验目标的字段值
func Validate(c interface{}) error {

	val := reflect.ValueOf(c).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		min := field.Tag.Get("min")
		max := field.Tag.Get("max")

		if min == "" && max == "" {
			continue
		}

		value := val.FieldByName(field.Name)

		if min != "" {
			coercedMinVal, _ := coerce(min, field.Type)
			if valueCompare(value, coercedMinVal) == -1 {
				return fmt.Errorf("invalid %s ! %v < %v",
					field.Name, value.Interface(), coercedMinVal.Interface())
			}
		}
		if max != "" {
			coercedMaxVal, _ := coerce(max, field.Type)
			if valueCompare(value, coercedMaxVal) == 1 {
				return fmt.Errorf("invalid %s ! %v > %v",
					field.Name, value.Interface(), coercedMaxVal.Interface())
			}
		}
	}

	hh, ok := c.(ConfigValidate)
	if ok {
		if err := hh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func valueCompare(v1 reflect.Value, v2 reflect.Value) int {
	switch v1.Type().String() {
	case "int", "int16", "int32", "int64":
		if v1.Int() > v2.Int() {
			return 1
		} else if v1.Int() < v2.Int() {
			return -1
		}
		return 0
	case "uint", "uint16", "uint32", "uint64":
		if v1.Uint() > v2.Uint() {
			return 1
		} else if v1.Uint() < v2.Uint() {
			return -1
		}
		return 0
	case "float32", "float64":
		if v1.Float() > v2.Float() {
			return 1
		} else if v1.Float() < v2.Float() {
			return -1
		}
		return 0
	case "time.Duration":
		if v1.Interface().(time.Duration) > v2.Interface().(time.Duration) {
			return 1
		} else if v1.Interface().(time.Duration) < v2.Interface().(time.Duration) {
			return -1
		}
		return 0
	}
	panic("impossible")
}

func coerce(v interface{}, typ reflect.Type) (reflect.Value, error) {
	var err error
	if typ.Kind() == reflect.Ptr {
		return reflect.ValueOf(v), nil
	}
	switch typ.String() {
	case "string":
		v, err = coerceString(v)
	case "int", "int16", "int32", "int64":
		v, err = coerceInt64(v)
	case "uint", "uint16", "uint32", "uint64":
		v, err = coerceUint64(v)
	case "float32", "float64":
		v, err = coerceFloat64(v)
	case "bool":
		v, err = coerceBool(v)
	case "time.Duration":
		v, err = coerceDuration(v)
	case "net.Addr":
		v, err = coerceAddr(v)

	default:
		v = nil
		err = fmt.Errorf("invalid type %s", typ.String())
	}
	return valueTypeCoerce(v, typ), err
}

func valueTypeCoerce(v interface{}, typ reflect.Type) reflect.Value {
	val := reflect.ValueOf(v)
	if reflect.TypeOf(v) == typ {
		return val
	}
	tval := reflect.New(typ).Elem()
	switch typ.String() {
	case "int", "int16", "int32", "int64":
		tval.SetInt(val.Int())
	case "uint", "uint16", "uint32", "uint64":
		tval.SetUint(val.Uint())
	case "float32", "float64":
		tval.SetFloat(val.Float())
	default:
		tval.Set(val)
	}
	return tval
}

func coerceString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int, int16, int32, int64, uint, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	}
	return fmt.Sprintf("%s", v), nil
}

func coerceDuration(v interface{}) (time.Duration, error) {
	switch v := v.(type) {
	case string:
		return time.ParseDuration(v)
	case int, int16, int32, int64:
		// treat like ms
		return time.Duration(reflect.ValueOf(v).Int()) * time.Millisecond, nil
	case uint, uint16, uint32, uint64:
		// treat like ms
		return time.Duration(reflect.ValueOf(v).Uint()) * time.Millisecond, nil
	case time.Duration:
		return v, nil
	}
	return 0, errors.New("invalid value type")
}

func coerceAddr(v interface{}) (net.Addr, error) {
	switch v := v.(type) {
	case string:
		return net.ResolveTCPAddr("tcp", v)
	case net.Addr:
		return v, nil
	}
	return nil, errors.New("invalid value type")
}

func coerceBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	case int, int16, int32, int64:
		return reflect.ValueOf(v).Int() != 0, nil
	case uint, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint() != 0, nil
	}
	return false, errors.New("invalid value type")
}

func coerceFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case int, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	}
	return 0, errors.New("invalid value type")
}

func coerceInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int, int16, int32, int64:
		return reflect.ValueOf(v).Int(), nil
	case uint, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func coerceUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case int, int16, int32, int64:
		return uint64(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint(), nil
	}
	return 0, errors.New("invalid value type")
}
