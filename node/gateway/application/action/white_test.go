package action

import (
	"encoding/json"
	"testing"

	"github.com/eolinker/goku-api-gateway/node/gateway/response"
)

type M = map[string]interface{}

func TestGenWhildRoot(t *testing.T) {

	paths := []string{
		"a.a.*.d",
	}
	value := M{
		"a": M{
			"a": []M{
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
			},
			"b": []M{
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
				{
					"a": 1,
					"b": 2,
					"c": 3,
					"d": "4",
				},
			},
		},
	}
	r := &response.Response{Data: value}
	filter := GenWhite(paths)

	filter.Do(r)

	data, _ := json.Marshal(r.Data)
	t.Log(string(data))

}
