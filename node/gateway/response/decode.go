package response

import (
	"encoding/json"
	"strings"

	"github.com/eolinker/goku-api-gateway/utils"
)

const (
	//JSON json
	JSON = "json"
	//XML xml
	XML = "xml"
	//String string
	String = "string"
	//JSONNoQuote 非标准json（key不带双引号）
	JSONNoQuote = "json-noquote"
)

var (
	jsonDecoder = func(data []byte, v interface{}) error {
		err := json.Unmarshal(data, v)
		return err
	}
	jsonNoQuoteDecoder = func(data []byte, v interface{}) error {
		d, err := utils.JSObjectToJSON(string(data))
		if err != nil {
			return err
		}
		err = json.Unmarshal(d, v)
		return err
	}
)

//GetDecoder getDecoder
func GetDecoder(decoder string) DecodeHandle {

	switch strings.ToLower(decoder) {
	case JSON:
		return jsonDecoder
	case JSONNoQuote:
		return jsonNoQuoteDecoder
	}

	return nil
}
