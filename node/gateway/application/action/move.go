package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

//MoveFilter moveFilter
type MoveFilter struct {
	source string
	target string
}

//Do do
func (f *MoveFilter) Do(value *response.Response) {
	value.Move(f.source, f.target)
}
