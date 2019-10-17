package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

type MoveFilter struct {
	source string
	target string
}

func (f *MoveFilter) Do(value *response.Response) {
	value.Move(f.source,f.target)
}

