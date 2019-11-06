package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

//DeleteFilter deleteFilter
type DeleteFilter string

//Do do
func (f DeleteFilter) Do(value *response.Response) {
	value.Delete(string(f))
}
