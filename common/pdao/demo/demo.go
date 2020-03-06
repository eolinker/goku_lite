package main

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/pdao"
)

type DemoGetDao interface {
	GetNameById(id int) string
	GetIdByName(name string) int
}
type DemoSetDao interface {
	Set(id int,name string) 
	 
}

type Demo struct {
	names map[string]int
	ids map[int]string
}

func NewDemo() *Demo {
	return &Demo{
		names: make(map[string]int),
		ids:   make(map[int]string),
	}}

func (d *Demo) Set(id int, name string) {
	 d.names[name]=id
	 d.ids[id]=name
}

func (d *Demo) GetNameById(id int) string {
	return d.ids[id]
}

func (d *Demo) GetIdByName(name string) int {
	  return d.names[name]
}
var (
	getDao DemoGetDao
	setDao DemoSetDao
)

func init() {
	pdao.Need(&getDao)
	pdao.Need(&setDao)
}
func main() {

	demo:=NewDemo()
	var seter DemoSetDao = demo
	pdao.Set(&seter)
	var getter DemoGetDao = demo
	pdao.Set(&getter)

	pdao.Check()

	setDao.Set(1,"test")
	fmt.Println(getDao.GetNameById(1))
	fmt.Println(getDao.GetIdByName("test"))
}