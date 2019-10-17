package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

type Blacklist string

func (f Blacklist) Do(value *response.Response) {
	value.Delete(string(f))
}



