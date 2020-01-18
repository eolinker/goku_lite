package response

import (
	"encoding/json"
	"encoding/xml"

	"strings"
)

var (
	jsonEncoder = &EncoderH{
		contentType: "application/json",
		handleFunc: func(v interface{}, org []byte) ([]byte, error) {
			return json.Marshal(v)
		},
	}
	xmlEncoder = &EncoderH{
		contentType: "text/xml; charset=utf-8",
		handleFunc: func(v interface{}, org []byte) ([]byte, error) {
			return xml.Marshal(v)
		},
	}
	stringEncoder = &EncoderH{
		contentType: "text/plain",
		handleFunc: func(v interface{}, org []byte) ([]byte, error) {

			return org, nil
		},
	}
	notEncoder = &EncoderH{
		contentType: "",
		handleFunc: func(v interface{}, org []byte) (bytes []byte, e error) {
			return org, nil
		},
	}
)

//EncoderH encodeH
type EncoderH struct {
	contentType string
	handleFunc  EncodeHandle
}

//Encode encode
func (e *EncoderH) Encode(v interface{}, org []byte) ([]byte, error) {
	return e.handleFunc(v, org)
}

//ContentType contentType
func (e *EncoderH) ContentType() string {
	return e.contentType
}

//GetEncoder 获取编码器
func GetEncoder(encoder string) Encoder {

	switch strings.ToLower(encoder) {
	case JSON:
		return jsonEncoder
	case XML:
		return xmlEncoder
	case String:
		return stringEncoder
	}
	return notEncoder
}
