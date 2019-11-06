package action

import "github.com/eolinker/goku-api-gateway/node/gateway/response"

//Blacklist 黑名单
type Blacklist string

//Do do
func (f Blacklist) Do(value *response.Response) {
	value.Delete(string(f))
}
