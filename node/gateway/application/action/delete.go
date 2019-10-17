package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

type DeleteFilter string

func (f DeleteFilter) Do(value *response.Response) {
	value.Delete(string(f))
}



