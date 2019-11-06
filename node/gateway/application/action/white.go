package action

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/eolinker/goku-api-gateway/node/gateway/response"
)

type _WhiteNode map[string]_WhiteNode

func (n _WhiteNode) String() string {

	data, e := json.Marshal(n)
	if e != nil {
		return e.Error()
	}
	return string(data)
}

type _WhiteRoot struct {
	root _WhiteNode
}

func (w *_WhiteRoot) Do(value *response.Response) {

	value.Data = w.root.Do(value.Data)
}

func newWhiteNode() _WhiteNode {
	return make(_WhiteNode)
}

//GenWhite genWhite
func GenWhite(paths []string) Filter {
	root := genWhite(paths)
	return &_WhiteRoot{
		root: root,
	}
}

func genWhite(paths []string) _WhiteNode {
	root := newWhiteNode()

	for _, path := range paths {
		root.add(strings.Split(path, "."))
	}
	return root
}
func (n _WhiteNode) add(path []string) {

	if len(path) == 0 {
		return
	}
	key := path[0]
	next := path[1:]

	node, has := n[key]
	if !has {
		node = newWhiteNode()
		n[key] = node
	}

	node.add(next)

}
func (n _WhiteNode) Do(value interface{}) interface{} {

	if value == nil {
		return nil
	}

	if len(n) == 0 {
		return value
	}

	switch value.(type) {
	case []interface{}:
		{
			list := value.([]interface{})
			l := len(list)
			vs := make([]interface{}, 0, len(n))
			for key, child := range n {
				if key == "*" {
					vsALl := make([]interface{}, 0, len(list))
					for _, item := range list {
						vsALl = append(vsALl, child.Do(item))
					}
					return vsALl
				}
				if i, err := strconv.Atoi(key); err == nil {
					if i < l {
						vs = append(vs, child.Do(list[i]))
					}
				}
			}
			return vs
		}

	case []map[string]interface{}:
		{
			list := value.([]map[string]interface{})
			l := len(list)
			vs := make([]interface{}, 0, len(n))
			for key, child := range n {
				if key == "*" {
					vsALl := make([]interface{}, 0, len(list))
					for _, item := range list {
						vsALl = append(vsALl, child.Do(item))
					}
					return vsALl
				}
				if i, err := strconv.Atoi(key); err == nil {
					if i < l {
						vs = append(vs, child.Do(list[i]))
					}
				}
			}
			return vs
		}
	case map[string]interface{}:

		m := value.(map[string]interface{})
		vm := make(map[string]interface{})
		for key, child := range n {
			if key == "*" {
				vmAll := make(map[string]interface{})
				for k, item := range m {
					vmAll[k] = child.Do(item)
				}
				return vmAll
			}

			if v, has := m[key]; has {
				vm[key] = child.Do(v)
			}
		}
		return vm

	}

	return value
}
