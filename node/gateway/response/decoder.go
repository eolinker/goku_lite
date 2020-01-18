package response

import "errors"

var (
	//ErrorInvalidDecoder 非法Decoder
	ErrorInvalidDecoder = errors.New("invalid decoder")
)

//DecodeHandle 解码器
type DecodeHandle func(data []byte, v interface{}) error

//EncodeHandle 解码处理器
type EncodeHandle func(v interface{}, org []byte) ([]byte, error)

//Encoder 解码器
type Encoder interface {
	Encode(v interface{}, org []byte) ([]byte, error)
	ContentType() string
}

//Decode 解码
func Decode(data []byte, handle DecodeHandle) (*Response, error) {

	if handle == nil {
		return nil, ErrorInvalidDecoder
	}

	var v interface{}
	err := handle(data, &v)
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: v,
	}, nil

}
