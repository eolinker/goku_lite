package goku

import (
	"net"
	"fmt"
	"reflect"
	"goku-ce/conf"
	"net/http"
	"log"
	"os"
	"sync/atomic"
	"time"
)
type GokuServer interface{
	Run() error
	Use(handler ...Handler)
	Address() string
	Listener() net.Listener
	Listen() error

}
type Handler interface{}

type Injector interface {
	Get(reflect.Type) reflect.Value
	Map(interface{}) Injector
}


type classicGoku struct {
	*Goku
}

type Goku struct{
	handlers []Handler
	index int
	ServiceConfig conf.GlobalConfig
	logger   *log.Logger
	listener *net.Listener
	address string
	values map[reflect.Type]reflect.Value
	parent Injector
	cClose chan bool
	isStopping bool
	activeCount int32
	Rate map[string]Rate
}

func (i *Goku) Map(val interface{}) Injector {
	i.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
	return i
}

// 启动一个Goku实例
func New() GokuServer{
	g := &Goku{values: make(map[reflect.Type]reflect.Value),logger:log.New(os.Stdout, "[Goku]", 0),Rate:make(map[string]Rate)}
	g.ServiceConfig = conf.ParseConfInfo()
	g.Map(g)
	return &classicGoku{g}
}

func (g *Goku) Run() error{
	server := &http.Server{Addr: g.Address(),Handler: g}
	fmt.Println("Listen on: " + g.Address() )
	err := server.Serve(g.Listener())
    if err != nil {
		fmt.Println("ERROR: ", err)
		return err
	}
	<-g.cClose
	return nil
}

func (s *Goku) Stop() {
	s.isStopping = true
	s.Listener().Close()
	if s.activeCount == 0 {
		s.cClose <- true
	}
}

func (g *Goku) Address() string {
	if g.address == "" {
		port := g.ServiceConfig.Port
		if port == "" {
			port = "3000"
		}
		host := g.ServiceConfig.Host
		g.address = host + ":" + port
	}
	return g.address
}

func (g *Goku) Listen() error {
	listener, err := net.Listen("tcp",g.Address())
	g.SetListener(listener)
	return err
}

func (g *Goku) SetListener(listener net.Listener) {
	g.listener = &listener
}


func (g *Goku) Listener() net.Listener{
	return *g.listener
}

func (g *Goku) run() {
	for g.index < len(g.handlers) {
		handle := g.handlers[g.index]
		_, err := g.Invoke(handle)
		if err != nil {
			panic(err)
		}
		g.index += 1
		
	}
	g.index = 0
	return 
}



func (g *Goku) Use(handler ...Handler) {
	for _,h := range handler{
		ValidateHandler(h)
		g.handlers = append(g.handlers,h)
	}
}

func (i *Goku) Get(t reflect.Type) reflect.Value {
	val := i.values[t]

	if val.IsValid() {
		return val
	}

	if t.Kind() == reflect.Interface {
		for k, v := range i.values {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	// Still no type found, try to look it up on the parent
	if !val.IsValid() && i.parent != nil {
		val = i.parent.Get(t)
	}

	return val

}

// 调用函数
func (g *Goku) Invoke(handler Handler) ([]reflect.Value,error) {
	ValidateHandler(handler)
	t := reflect.TypeOf(handler)
	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := g.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}

		in[i] = val
	}
	return reflect.ValueOf(handler).Call(in),nil
}

func (c *Goku) IsStopped() bool {
	return !(c.index < len(c.handlers))
}


func (g *Goku) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if(req.RequestURI == "/favicon.ico"){
		return
	}
	g.Map(res)
	g.Map(req)
	activeCount := atomic.AddInt32(&g.activeCount, -1)
	if g.isStopping && activeCount == 0 {
		time.Sleep(1)
		g.cClose <- true
	}
	g.run()
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}

	return t
}
