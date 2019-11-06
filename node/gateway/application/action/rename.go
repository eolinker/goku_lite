package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

//RenameFilter renameFilter
type RenameFilter struct {
	pattern string
	name    string
}

//Do do
func (f *RenameFilter) Do(value *response.Response) {

	value.ReName(f.pattern, f.name)
}
