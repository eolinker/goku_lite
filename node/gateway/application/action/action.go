package action

import (
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/node/gateway/response"
)

const (
	//Delete delete
	Delete = "delete"
	//Rename rename
	Rename = "rename"
	//Move move
	Move = "move"
	//Black black
	Black = "black"
	//White white
	White = "white"
)

//Filter 过滤器
type Filter interface {
	Do(value *response.Response)
}

//Filters 过滤器列表
type Filters []Filter

//Do do
func (f Filters) Do(value *response.Response) {
	target := value
	for _, item := range f {
		item.Do(target)
	}

}

//GenByconfig 通过配置生成Filter
func GenByconfig(ac *config.ActionConfig) Filter {
	switch strings.ToLower(ac.ActionType) {
	case Delete:
		return DeleteFilter(ac.Original)
	case Rename:
		return &RenameFilter{
			pattern: ac.Original,
			name:    ac.Target,
		}
	case Move:
		return &MoveFilter{
			target: ac.Target,
			source: ac.Original,
		}
	}
	return nil
}
