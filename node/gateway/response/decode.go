package response

import (
	"encoding/json"
	"strings"
)
const(
	JSON ="json"
	XML = "xml"
	String = "string"
)
var (
	jsonDecoder =  func(data []byte, v interface{}) error {
		err:=json.Unmarshal(data,v)
		return err
	}

)
func GetDecoder(decoder string) DecodeHandle {

	switch strings.ToLower(decoder) {
	case JSON:
		return jsonDecoder
	}

	return nil
}

