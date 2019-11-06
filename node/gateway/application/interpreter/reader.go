package interpreter

import (
	"fmt"
	"net/url"
	"reflect"
)

//Reader reder
type Reader interface {
	Read(variables *Variables) string
}

type _NotReader string

func (r _NotReader) Read(variables *Variables) string {
	return string(r)
}

type _OrgReader struct {
}

func (r *_OrgReader) Read(variables *Variables) string {
	return string(variables.Org)
}

type _BodyReader struct {
	Index int
	Path  []string
	Name  string
}

func (r *_BodyReader) Read(variables *Variables) string {

	if len(variables.Bodes) <= r.Index {
		return ""
	}
	if r.Index == 0 {
		body := variables.Bodes[r.Index]
		form, ok := body.(url.Values)
		if ok {
			return form.Get(r.Name)
		}
	}

	root := reflect.ValueOf(variables.Bodes[r.Index])
	return find(&root, r.Path)
}
func find(node *reflect.Value, path []string) string {

	if len(path) == 0 {
		return fmt.Sprint(node.Interface())
	}

	k := node.Kind()

	switch k {

	case reflect.Interface:

		next := node.Elem()
		return find(&next, path)
	case reflect.Map:
		{
			key := reflect.ValueOf(path[0])
			next := node.MapIndex(key)
			return find(&next, path[1:])
		}
	default:
		return ""
	}

	return ""
}

type _HeaderReader struct {
	Index int
	Key   string
}

func (r *_HeaderReader) Read(variables *Variables) string {
	if len(variables.Headers) <= r.Index {
		return ""
	}

	return variables.Headers[r.Index].Get(r.Key)
}

type _RestFulReader struct {
	Key string
}

func (r *_RestFulReader) Read(variables *Variables) string {
	return variables.Restful[r.Key]
}

type _QueryReader struct {
	Key string
}

func (r *_QueryReader) Read(variables *Variables) string {
	return variables.Query.Get(r.Key)
}

type _CookieReader struct {
	Index int
	Name  string
}

func (r *_CookieReader) Read(variables *Variables) string {
	if len(variables.Cookies) <= r.Index {
		return ""
	}
	for _, c := range variables.Cookies[r.Index] {
		if c.Name == r.Name {
			return c.Value
		}
	}
	return ""
}
