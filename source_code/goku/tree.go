package goku

import (
	"strings"
	"unicode"
	"unicode/utf8"
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}


// 计算路径中参数数量
func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

type node struct {
	path      string
	wildChild bool
	nType     nodeType
	maxParams uint8
	indices   string
	children  []*node
	handle    Handle
	priority  uint32
	context   *Context
}


// 在必要时给定子项和排序的优先级
func (n *node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

// addRoute将具有给定句柄的节点添加到路径中
func (n *node) addRoute(path string, handle Handle,context Context) {
	fullPath := path
	n.priority++
	numParams := countParams(path)

	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// 更新当前节点的 maxParams
			if numParams > n.maxParams {
				n.maxParams = numParams
			}

			i := 0
			max := min(len(path), len(n.path))
			for i < max && path[i] == n.path[i] {
				i++
			}

			if i < len(n.path) {
				child := node{
					path:      n.path[i:],
					wildChild: n.wildChild,
					nType:     static,
					indices:   n.indices,
					children:  n.children,
					handle:    n.handle,
					context:   n.context,
					priority:  n.priority - 1,
				}

				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}

				n.children = []*node{&child}
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.handle = nil
				n.context = &Context{}
				n.wildChild = false
			}

			if i < len(path) {
				path = path[i:]

				if n.wildChild {
					n = n.children[0]
					n.priority++

					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--

					if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
						(len(n.path) >= len(path) || path[len(n.path)] == '/') {
						continue walk
					} else {
						var pathSeg string
						if n.nType == catchAll {
							pathSeg = path
						} else {
							pathSeg = strings.SplitN(path, "/", 2)[0]
						}
						prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
						panic("'" + pathSeg +
							"' in new path '" + fullPath +
							"' conflicts with existing wildcard '" + n.path +
							"' in existing prefix '" + prefix +
							"'")
					}
				}

				c := path[0]

				if n.nType == param && c == '/' && len(n.children) == 1 {
					n = n.children[0]
					n.priority++
					continue walk
				}

				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}

				if c != ':' && c != '*' {
					n.indices += string([]byte{c})
					child := &node{
						maxParams: numParams,
					}
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				n.insertChild(numParams, path, fullPath, handle,context)
				return

			} else if i == len(path) { 
				if n.handle != nil {
					panic("a handle is already registered for path '" + fullPath + "'")
				}
				n.handle = handle
				n.context = &context
			}
			return
		}
	} else { 
		n.insertChild(numParams, path, fullPath, handle,context)
		n.nType = root
	}
}

func (n *node) insertChild(numParams uint8, path, fullPath string, handle Handle,context Context) {
	var offset int // already handled bytes of the path

	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		if c != ':' && c != '*' {
			continue
		}

		end := i + 1
		for end < max && path[end] != '/' {
			switch path[end] {
			case ':', '*':
				panic("only one wildcard per path segment is allowed, has: '" +
					path[i:] + "' in path '" + fullPath + "'")
			default:
				end++
			}
		}

		if len(n.children) > 0 {
			panic("wildcard route '" + path[i:end] +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		if end-i < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		if c == ':' { // param
			if i > 0 {
				n.path = path[offset:i]
				offset = i
			}

			child := &node{
				nType:     param,
				maxParams: numParams,
			}
			n.children = []*node{child}
			n.wildChild = true
			n = child
			n.priority++
			numParams--

			if end < max {
				n.path = path[offset:end]
				offset = end

				child := &node{
					maxParams: numParams,
					priority:  1,
				}
				n.children = []*node{child}
				n = child
			}

		} else { // catchAll
			if end != max || numParams > 1 {
				panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
			}

			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
			}

			i--
			if path[i] != '/' {
				panic("no / before catch-all in path '" + fullPath + "'")
			}

			n.path = path[offset:i]

			child := &node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
			}
			n.children = []*node{child}
			n.indices = string(path[i])
			n = child
			n.priority++

			child = &node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				handle:    handle,
				context: &context,
				priority:  1,
			}
			n.children = []*node{child}

			return
		}
	}

	n.path = path[offset:]
	n.handle = handle
	n.context = &context
}

func (n *node) getValue(path string) (handle Handle, p Params, context *Context, tsr bool) {
walk: // outer loop for walking the tree
	for {
		if len(path) > len(n.path) {
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				if !n.wildChild {
					c := path[0]
					for i := 0; i < len(n.indices); i++ {
						if c == n.indices[i] {
							n = n.children[i]
							continue walk
						}
					}
					tsr = (path == "/" && n.handle != nil)
					fmt.Println(123)
					return
				}

				// handle wildcard child
				n = n.children[0]
				switch n.nType {
				case param:
					// find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// save param value
					if p == nil {
						// lazy allocation
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[1:]
					p[i].Value = path[:end]

					// we need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						tsr = (len(path) == end+1)
						fmt.Println(456)
						return
					}
					context = n.context
					if handle = n.handle; handle != nil {
						return
					} else if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						tsr = (n.path == "/" && n.handle != nil)
					}

					return

				case catchAll:
					// save param value
					if p == nil {
						// lazy allocation
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[2:]
					p[i].Value = path

					handle = n.handle
					context = n.context
					return

				default:
					panic("invalid node type")
				}
			}
		} else if path == n.path {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			context = n.context
			if handle = n.handle; handle != nil {
				return
			}
			
			if path == "/" && n.wildChild && n.nType != root {
				tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			for i := 0; i < len(n.indices); i++ {
				if n.indices[i] == '/' {
					n = n.children[i]
					tsr = (len(n.path) == 1 && n.handle != nil) ||
						(n.nType == catchAll && n.children[0].handle != nil)
					fmt.Println(541)
					return
				}
			}
			return
		}
		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		tsr = (path == "/") ||
			(len(n.path) == len(path)+1 && n.path[len(path)] == '/' &&
				path == n.path[:len(n.path)-1] && n.handle != nil)
		return
	}
}

// Makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (ciPath []byte, found bool) {
	return n.findCaseInsensitivePathRec(
		path,
		strings.ToLower(path),
		make([]byte, 0, len(path)+1), // preallocate enough memory for new path
		[4]byte{},                    // empty rune buffer
		fixTrailingSlash,
	)
}

// shift bytes in array by n bytes left
func shiftNRuneBytes(rb [4]byte, n int) [4]byte {
	switch n {
	case 0:
		return rb
	case 1:
		return [4]byte{rb[1], rb[2], rb[3], 0}
	case 2:
		return [4]byte{rb[2], rb[3]}
	case 3:
		return [4]byte{rb[3]}
	default:
		return [4]byte{}
	}
}

// recursive case-insensitive lookup function used by n.findCaseInsensitivePath
func (n *node) findCaseInsensitivePathRec(path, loPath string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) ([]byte, bool) {
	loNPath := strings.ToLower(n.path)

walk: // outer loop for walking the tree
	for len(loPath) >= len(loNPath) && (len(loNPath) == 0 || loPath[1:len(loNPath)] == loNPath[1:]) {
		// add common path to result
		ciPath = append(ciPath, n.path...)

		if path = path[len(n.path):]; len(path) > 0 {
			loOld := loPath
			loPath = loPath[len(loNPath):]

			// If this node does not have a wildcard (param or catchAll) child,
			// we can just look up the next child node and continue to walk down
			// the tree
			if !n.wildChild {
				// skip rune bytes already processed
				rb = shiftNRuneBytes(rb, len(loNPath))

				if rb[0] != 0 {
					// old rune not finished
					for i := 0; i < len(n.indices); i++ {
						if n.indices[i] == rb[0] {
							// continue with child node
							n = n.children[i]
							loNPath = strings.ToLower(n.path)
							continue walk
						}
					}
				} else {
					// process a new rune
					var rv rune

					// find rune start
					// runes are up to 4 byte long,
					// -4 would definitely be another rune
					var off int
					for max := min(len(loNPath), 3); off < max; off++ {
						if i := len(loNPath) - off; utf8.RuneStart(loOld[i]) {
							// read rune from cached lowercase path
							rv, _ = utf8.DecodeRuneInString(loOld[i:])
							break
						}
					}

					// calculate lowercase bytes of current rune
					utf8.EncodeRune(rb[:], rv)
					// skipp already processed bytes
					rb = shiftNRuneBytes(rb, off)

					for i := 0; i < len(n.indices); i++ {
						// lowercase matches
						if n.indices[i] == rb[0] {
							// must use a recursive approach since both the
							// uppercase byte and the lowercase byte might exist
							// as an index
							if out, found := n.children[i].findCaseInsensitivePathRec(
								path, loPath, ciPath, rb, fixTrailingSlash,
							); found {
								return out, true
							}
							break
						}
					}

					// same for uppercase rune, if it differs
					if up := unicode.ToUpper(rv); up != rv {
						utf8.EncodeRune(rb[:], up)
						rb = shiftNRuneBytes(rb, off)

						for i := 0; i < len(n.indices); i++ {
							// uppercase matches
							if n.indices[i] == rb[0] {
								// continue with child node
								n = n.children[i]
								loNPath = strings.ToLower(n.path)
								continue walk
							}
						}
					}
				}

				// Nothing found. We can recommend to redirect to the same URL
				// without a trailing slash if a leaf exists for that path
				return ciPath, (fixTrailingSlash && path == "/" && n.handle != nil)
			}

			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				k := 0
				for k < len(path) && path[k] != '/' {
					k++
				}

				// add param value to case insensitive path
				ciPath = append(ciPath, path[:k]...)

				// we need to go deeper!
				if k < len(path) {
					if len(n.children) > 0 {
						// continue with child node
						n = n.children[0]
						loNPath = strings.ToLower(n.path)
						loPath = loPath[k:]
						path = path[k:]
						continue
					}

					// ... but we can't
					if fixTrailingSlash && len(path) == k+1 {
						return ciPath, true
					}
					return ciPath, false
				}

				if n.handle != nil {
					return ciPath, true
				} else if fixTrailingSlash && len(n.children) == 1 {
					// No handle found. Check if a handle for this path + a
					// trailing slash exists
					n = n.children[0]
					if n.path == "/" && n.handle != nil {
						return append(ciPath, '/'), true
					}
				}
				return ciPath, false

			case catchAll:
				return append(ciPath, path...), true

			default:
				panic("invalid node type")
			}
		} else {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if n.handle != nil {
				return ciPath, true
			}

			// No handle found.
			// Try to fix the path by adding a trailing slash
			if fixTrailingSlash {
				for i := 0; i < len(n.indices); i++ {
					if n.indices[i] == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.handle != nil) ||
							(n.nType == catchAll && n.children[0].handle != nil) {
							return append(ciPath, '/'), true
						}
						return ciPath, false
					}
				}
			}
			return ciPath, false
		}
	}

	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
	if fixTrailingSlash {
		if path == "/" {
			return ciPath, true
		}
		if len(loPath)+1 == len(loNPath) && loNPath[len(loPath)] == '/' &&
			loPath[1:] == loNPath[1:len(loPath)] && n.handle != nil {
			return append(ciPath, n.path...), true
		}
	}
	return ciPath, false
}
