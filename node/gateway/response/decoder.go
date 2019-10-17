package response

import "errors"

var (
	ErrorInvalidDecoder = errors.New("invalid decoder")
)
type DecodeHandle func(data []byte, v interface{}) error

type EncodeHandle func (v interface{},org []byte)([]byte,error)
type Encoder interface {
	Encode(v interface{},org []byte)([]byte,error)
	ContentType()string
}

func Decode(data []byte,handle DecodeHandle) (*Response,error) {

	if handle == nil{
		return nil,ErrorInvalidDecoder
	}

	var v interface{}
	err:=handle(data,&v)
	if err!=nil{
		return nil,err
	}
	return &Response{
		Data: v,
	},nil

}
