package routerRule

import "strings"

//Routers Routers
type Routers []*Router

func (p Routers) Len() int {
	return len(p)
}

func (p Routers) Less(i, j int) bool {
	if p[i].Host == "*" {
		return false
	}
	if p[j].Host == "*" {
		return true
	}
	if strings.Contains(p[i].Host, "*") {
		return false
	}
	if strings.Contains(p[j].Host, "*") {
		return true
	}
	return false
}

func (p Routers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
