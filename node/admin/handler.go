package admin

import (
	"fmt"
	"net/http"

	"github.com/eolinker/goku-api-gateway/module"
)

var (
	handler = &tHandler{
		serveMux: http.NewServeMux(),
		modules:  make(map[string]*HandlerItem),
	}
)

type tHandler struct {
	serveMux *http.ServeMux
	modules  map[string]*HandlerItem
}

//HandlerItem handlerItem
type HandlerItem struct {
	isOpen  bool
	handler http.Handler
	name    string
}

func (ht *HandlerItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !ht.isOpen {
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}
	ht.handler.ServeHTTP(w, r)
}

func (hd *tHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hd.serveMux == nil {
		w.WriteHeader(404)
		return
	}

	hd.serveMux.ServeHTTP(w, r)
}
func (hd *tHandler) Add(name string, pattern string, h http.Handler) {
	if _, has := hd.modules[pattern]; has {

		panic(fmt.Sprint("duplicate pattern:", pattern, " for module:", name))

		return
	}
	isOpen := module.IsOpen(name)

	ht := &HandlerItem{
		isOpen:  isOpen,
		handler: h,
		name:    name,
	}
	hd.modules[pattern] = ht
	hd.serveMux.Handle(pattern, ht)
}
func (hd *tHandler) Refresh() {
	for _, ht := range hd.modules {
		ht.isOpen = module.IsOpen(ht.name)
	}
}

//Handler handler
func Handler() http.Handler {
	return handler
}

//Add( add
func Add(name string, pattern string, h http.Handler) {
	handler.Add(name, pattern, h)
}

//Refresh refresh
func Refresh() {
	handler.Refresh()
}
