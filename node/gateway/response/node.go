package response

import (
	"strconv"
	"strings"
)

type _Node struct {
	data   interface{}
	parent *_Node
	index  int
	key    string
}

func (node *_Node) Set(v interface{}) {
	if node.parent == nil {
		return
	}
}

func (node *_Node) child(key string, callback func(*_Node)) bool {

	return false
}
func (node *_Node) get(key string) *_Node {

	if key == "" {
		return nil
	}

	if key == "*" {
		return node
	}
	switch node.data.(type) {
	case []interface{}:
		{

			sl := node.data.([]interface{})
			if i, e := strconv.Atoi(key); e == nil {
				if i < len(sl) {
					n := &_Node{
						data:   sl[i],
						parent: node,
						index:  i,
					}
					return n
				}

			}
		}
	case map[string]interface{}:
		{
			sm := node.data.(map[string]interface{})
			if v, has := sm[key]; has {
				n := &_Node{
					data:   v,
					parent: node,
					key:    key,
				}

				return n
			}
		}
	}

	return nil

}
func (node *_Node) Make(path []string) {

	if node.data == nil {
		node.data = make(map[string]interface{})
	}
	makePath(node.data, path)

}
func makePath(data interface{}, path []string) {
	if len(path) == 0 {
		return
	}
	k := path[0]
	next := path[1:]
	switch data.(type) {
	case []interface{}:
		dl := data.([]interface{})
		if k == "*" {
			for _, d := range dl {
				makePath(d, next)
			}
		} else if i, err := strconv.Atoi(k); err != nil {
			if i < len(dl) {
				makePath(dl[i], next)
			}
		}
	case map[string]interface{}:
		dm := data.(map[string]interface{})
		if k == "*" {
			for _, d := range dm {
				makePath(d, next)
			}
		} else {
			d, has := dm[k]
			if !has {
				d = make(map[string]interface{})
				dm[k] = d
			}
			makePath(d, next)
		}
	}

}

//Pattern 匹配节点并执行callback
//pattern 的格式为，使用 . 分割字段层级，对于数组，可以使用 * 或者 数字 指定项目，对于map也可以使用 * 表示匹配该层的所有字段
//优先匹配短路径，例如 a.b 将优先匹配  "a.b",
func (node *_Node) Pattern(pattern string, callback func(*_Node) bool) (match bool, isBreak bool) {

	if pattern == "" {
		return true, callback(node)
	}

	isMatch := false
	for key := pattern; len(key) > 0; key = spiltKey(key) {

		next := next(pattern, key)
		if key == "*" {
			data := node.data
			switch data.(type) {
			case []interface{}:
				{
					sl := data.([]interface{})

					for i, s := range sl {
						n := &_Node{
							data:   s,
							parent: node,
							index:  i,
						}
						match, isBreak := n.Pattern(next, callback)
						if isBreak {
							return true, true
						}
						isMatch = isMatch || match
					}
				}
			case map[string]interface{}:
				{
					sm := data.(map[string]interface{})

					if len(sm) > 0 {
						for k, s := range sm {
							n := &_Node{
								data:   s,
								parent: node,
								key:    k,
							}
							match, isBreak := n.Pattern(next, callback)
							if isBreak {
								return true, true
							}
							isMatch = isMatch || match
						}
					} else {
						n := &_Node{
							data:   nil,
							parent: node,
							key:    "",
						}
						match, isBreak := n.Pattern(next, callback)
						if isBreak {
							return true, true
						}
						isMatch = isMatch || match
					}

				}
			}

		} else {

			child := node.get(key)
			if child != nil {
				match, isBreak := child.Pattern(next, callback)
				if isBreak {
					return true, true
				}
				isMatch = match
			}
		}
		if isMatch {
			return true, false
		}
	}
	return false, false
}

func next(pattern, key string) string {

	l := len(key)

	if len(pattern) > l {
		return pattern[l+1:]
	}
	return ""
}
func spiltKey(key string) string {
	if index := strings.LastIndex(key, "."); index != -1 {
		return key[:index]
	}
	return ""
}
