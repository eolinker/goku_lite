package goku

import (
	"net"
	"fmt"
	"reflect"
	"goku-ce/conf"
	"net/http"
	"log"
	"os"
)
type GokuServer interface{
	Run() error
	Address() string
	Listener() net.Listener
	Listen() error

}

type classicGoku struct {
	*Goku
}


type Goku struct{
 	*Router
	index int
	ServiceConfig conf.GlobalConfig
	logger   *log.Logger
	listener *net.Listener
	address string
	values map[reflect.Type]reflect.Value
	cClose chan bool
	isStopping bool
	activeCount int32
}


// 启动一个Goku实例
func New() *Goku{
	g := &Goku{
		Router:NewRouter(),
		values: make(map[reflect.Type]reflect.Value),
		logger:log.New(os.Stdout, "[Goku]", 0),
		}
		
	g.ServiceConfig = conf.ParseConfInfo()
	return g
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

