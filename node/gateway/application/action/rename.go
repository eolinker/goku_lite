package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

type RenameFilter struct {
	pattern string
	name string
}

func (f *RenameFilter) Do(value *response.Response) {

	value.ReName(f.pattern,f.name)
}

