package response

import "strings"

//Response response
type Response struct {
	// root 数据
	Data interface{}
}

//Delete delete
func (r *Response) Delete(pattern string) *Response {
	if pattern == "" {
		return r
	}
	root := _Node{
		data: r.Data,
	}

	root.Pattern(pattern, func(node *_Node) bool {
		if node.parent == nil {
			return false
		}
		parent := node.parent
		switch parent.data.(type) {
		case []interface{}:
			index := node.index
			sl := parent.data.([]interface{})

			nl := sl[:index]
			sl = append(nl, sl[index+1])
			parent.data = sl

		case map[string]interface{}:
			mp := parent.data.(map[string]interface{})
			delete(mp, node.key)
		}
		return false
	})
	return r
}

//SetValue 设置目标值，如果目标不存在，会对路径进行创建
func (r *Response) SetValue(pattern string, value interface{}) {
	if pattern == "" {
		r.Data = value
		return
	}
	root := _Node{
		data: r.Data,
	}
	root.Make(strings.Split(pattern, "."))

	root.Pattern(pattern, func(node *_Node) bool {

		if node.parent == nil {
			return false
		}
		parent := node.parent
		switch parent.data.(type) {
		case []interface{}:
			sl := parent.data.([]interface{})
			index := node.index
			sl[index] = value
			parent.data = sl

		case map[string]interface{}:
			mp := parent.data.(map[string]interface{})
			mp[node.key] = value
		}
		return false
	})

}

//ReTarget 选择目标重新设置为root
func (r *Response) ReTarget(pattern string) {
	if pattern == "" {
		return
	}
	root := _Node{
		data: r.Data,
	}

	match, _ := root.Pattern(pattern, func(node *_Node) bool {
		r.Data = node.data

		return true
	})
	if !match {
		r.Data = make(map[string]interface{})
	}
	return
}

//Group group
func (r *Response) Group(path []string) {
	l := len(path)
	if l == 0 {
		return
	}
	root := make(map[string]interface{})
	node := root

	lastKey := path[l-1]
	if l > 1 {
		for _, key := range path[:l-1] {
			v := make(map[string]interface{})
			node[key] = v
			node = v
		}
	}

	node[lastKey] = r.Data
	r.Data = root
}

//ReName 重命名
func (r *Response) ReName(pattern string, newName string) {
	if pattern == "" {
		return
	}
	root := _Node{
		data: r.Data,
	}

	root.Pattern(pattern, func(node *_Node) bool {
		if node.parent == nil {
			return false
		}
		parent := node.parent
		switch parent.data.(type) {
		case []interface{}:
			return false

		case map[string]interface{}:
			mp := parent.data.(map[string]interface{})
			delete(mp, node.key)
			mp[newName] = node.data
			return false
		}
		return false
	})
}

//Move move
func (r *Response) Move(source, target string) {

	if strings.Index(source, "*") != -1 {
		return
	}
	if strings.Index(target, "*") != -1 {
		return
	}
	root := _Node{
		data: r.Data,
	}
	var oldValues *_Node
	match, _ := root.Pattern(source, func(node *_Node) bool {
		oldValues = node

		if node.parent == nil {
			return false
		}
		parent := node.parent
		switch parent.data.(type) {
		case []interface{}:
			index := node.index
			sl := parent.data.([]interface{})

			nl := sl[:index]
			sl = append(nl, sl[index+1])
			parent.data = sl

		case map[string]interface{}:
			mp := parent.data.(map[string]interface{})
			delete(mp, node.key)
		}
		return false
	})
	if match {
		r.SetValue(target, oldValues.data)
	} else {
		r.SetValue(target, nil)
	}

}
