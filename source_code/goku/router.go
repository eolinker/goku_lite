package goku

import (
	"net/http"
	"goku-ce/conf"
	"strings"
)

// Handle是一个可以被注册到路由中去处理http请求，类似于HandlerFunc,但是有第三个参数值
type Handle func(http.ResponseWriter, *http.Request, Params,*Context)
// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params是一个参数切片，作为路由的返回结果，这个切片是有序的
// 第一个URL参数会作为第一个切片值，因此通过索引来读值是安全的
type Params []Param

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

// Router是一个可以被用来调度请求去不同处理函数的Handler 
type Router struct {
	trees map[string]*node

	context *Context

	handle Handle

	RedirectTrailingSlash bool

	RedirectFixedPath bool

	HandleMethodNotAllowed bool

	HandleOPTIONS bool

	NotFound http.Handler

	MethodNotAllowed http.Handler

	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

func NewRouter() *Router {
	return &Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}


func (r *Router) Use(handle Handle) {
	r.handle = handle
}


func (r *Router) Handle(method, path string, handle Handle,context Context) {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}
	root.addRoute(path, handle,context)
}

// // HandlerFunc 是一个适配器允许使用http.HandleFunc函数作为一个请求处理器
// func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
// 	r.Handler(method, path, handler)
// }

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}
// 查找允许手动查找方法 + 路径组合。
// 这对于构建围绕此路由器的框架非常有用。
// 如果找到路径, 它将返回句柄函数和路径参数值
// 否则, 第三个返回值指示是否应执行与附加/不带尾随斜线的同一路径的重定向
func (r *Router) Lookup(method, path string) (Handle, Params, *Context, bool) {
	if root := r.trees[method]; root != nil {
		return root.getValue(path)
	}
	return nil, nil,&Context{}, false
}

func (r *Router) allowed(path, reqMethod string) (allow string) {
	if path == "*" { // server-wide
		for method := range r.trees {
			if method == "OPTIONS" {
				continue
			}

			// 将请求方法添加到允许的方法列表中
			if len(allow) == 0 {
				allow = method
			} else {
				allow += ", " + method
			}

		}
	} else { // 特定路径
		for method := range r.trees {
			// 跳过请求的方法-我们已经尝试过这一项
			if method == reqMethod || method == "OPTIONS" {
				continue
			}

			handle, _, _,_ := r.trees[method].getValue(path)
			if handle != nil {
				// 将请求方法添加到允许的方法列表中
				if len(allow) == 0 {
					allow = method
				} else {
					allow += ", " + method
				}
			}
		}
	}
	if len(allow) > 0 {
		allow += ", OPTIONS"
	}
	return
}

// ServeHTTP使用路由实现http.Handler接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		
		defer r.recv(w, req)
	}
	// now := time.Now()
	path := req.URL.Path
	pathArray := strings.Split(path,"/")

	if len(pathArray) == 2 {
		w.WriteHeader(500)
		if pathArray[1] == "" {
			w.Write([]byte("Missing Gateway Alias"))
		} else {
			w.Write([]byte("Missing StrategyID"))
		}
		return
	} else if len(pathArray) == 3 {
		w.WriteHeader(500)
		if pathArray[2] == "" {
			w.Write([]byte("Missing StrategyID"))
		} else {
			w.Write([]byte("Invalid URI"))
		}
		return
	}
	
	if root := r.trees[req.Method]; root != nil {
		handle, ps, context,tsr := root.getValue(path); 
		if handle != nil {
			handle(w, req, ps,context)
			return
		} else if req.Method != "CONNECT" && path != "/" {
			code := 301 
			if req.Method != "GET" {
				code = 307
			}

			if tsr && r.RedirectTrailingSlash {
				if len(path) > 1 && path[len(path)-1] == '/' {
					req.URL.Path = path[:len(path)-1]
				} else {
					req.URL.Path = path + "/"
				}
				http.Redirect(w, req, req.URL.String(), code)
				return
			}
			// 尝试修复请求路径
			if r.RedirectFixedPath {
				fixedPath, found := root.findCaseInsensitivePath(
					CleanPath(path),
					r.RedirectTrailingSlash,
				)
				if found {
					req.URL.Path = string(fixedPath)
					http.Redirect(w, req, req.URL.String(), code)
					return
				}
			}
		}
	}

	if req.Method == "OPTIONS" && r.HandleOPTIONS {
		// Handle OPTIONS requests
		if allow := r.allowed(path, req.Method); len(allow) > 0 {
			w.Header().Set("Allow", allow)
			return
		}
	} else {
		// Handle 405
		if r.HandleMethodNotAllowed {
			if allow := r.allowed(path, req.Method); len(allow) > 0 {
				w.Header().Set("Allow", allow)
				if r.MethodNotAllowed != nil {
					r.MethodNotAllowed.ServeHTTP(w, req)
				} else {
					http.Error(w,
						http.StatusText(http.StatusMethodNotAllowed),
						http.StatusMethodNotAllowed,
					)
				}
				return
			}
		}
	}

	// Handle 404
	if r.NotFound != nil {
		r.NotFound.ServeHTTP(w, req)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Invalid URI"))
	}
}

// 注册路由
func (r *Router) RegisterRouter(c conf.GlobalConfig,handle Handle) {
	r.handle = handle
	for _, g := range c.GatewayList {
		if g.GatewayStatus != "on" {
			continue
		}
		gateway := Gateway{
			GatewayAlias: g.GatewayAlias,
			GatewayStatus: g.GatewayStatus, 
			IPLimitType: g.IPLimitType,
			IPWhiteList: g.IPWhiteList,
			IPBlackList: g.IPBlackList,
		}
		for _, s := range g.StrategyList.Strategy {
			strategy := Strategy{
				StrategyID: s.StrategyID,
				Auth: s.Auth,
				ApiKey: s.ApiKey,
				BasicUserName: s.BasicUserName,
				BasicUserPassword: s.BasicUserPassword,
				IPLimitType: s.IPLimitType,
				IPWhiteList: s.IPWhiteList,
				IPBlackList: s.IPBlackList,
				RateLimitList:s.RateLimitList,
			}
			for _, api := range g.ApiList.Apis {
				path := "/" + g.GatewayAlias + "/" + s.StrategyID + api.RequestURL
				backendPath := ""
				flag := false
				// 获取后端请求路径
				for _,b := range g.BackendList.Backend {
					if b.BackendID == api.BackendID{
						backendPath = b.BackendPath
						flag =true
						break
					}
				}
				if !flag && api.BackendID != -1{
					continue
				}
				apiInfo := Api{
					RequestURL: api.RequestURL,
					BackendPath: backendPath,
					ProxyURL: api.ProxyURL,
					IsRaw:api.IsRaw,
					ProxyMethod:api.ProxyMethod,
					ProxyParams:api.ProxyParams,
					ConstantParams:api.ConstantParams,
				}
				context := Context{
					GatewayInfo:gateway,
					StrategyInfo:strategy,
					ApiInfo:apiInfo,
					Rate:make(map[string]Rate),
				}
				for _,method := range api.RequestMethod {
					r.Handle(strings.ToUpper(method),path,r.handle,context)
				}
			}
		}
	}
}
