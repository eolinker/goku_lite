package action

import (
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/node/gateway/response"
	"strings"
)

const (
	Delete ="delete"
	Rename = "rename"
	Move = "move"
	Black  = "black"
	White  = "white"


)
type Filter interface {
	Do(value *response.Response)
}

type Filters []Filter

func (f Filters) Do(value *response.Response) {
	target := value
	for _,item:=range f{
		  item.Do(target)
	}

}

func GenByconfig( ac * config.ActionConfig) Filter {
	switch strings.ToLower(ac.ActionType) {
	case Delete:
		return DeleteFilter(ac.Original)
	case Rename:
		return &RenameFilter{
			pattern:ac.Original,
			name:ac.Target,
		}
	case Move:
		return &MoveFilter{
			target:ac.Target,
			source:ac.Original,
		}
	}
	return nil
}