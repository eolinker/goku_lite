package cmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/eolinker/goku-api-gateway/common/ioutils"
)

var (
	ErrorEmptyFrame  = errors.New("empty frame")
	ErrorInvalidCode = errors.New("invalid code")
)

func ReadFrame(reader io.Reader) ([]byte, error) {

	sizeBuf := make([]byte, 4, 4)
	// 获取报文头部信息
	_, err := io.ReadFull(reader, sizeBuf)
	if err != nil {
		return nil, err
	}
	// 获取报文数据大小
	size := binary.BigEndian.Uint32(sizeBuf)

	data := make([]byte, size, size)

	_, e := io.ReadFull(reader, data)
	if e != nil {
		return nil, err
	}
	return data, nil
}
func GetCmd(frame []byte) (Code, []byte, error) {
	frameLen := len(frame)
	if frameLen < 5 {
		// 长度小于5时，报文没有数据
		return "", nil, ErrorEmptyFrame
	}
	buf := bytes.NewBuffer(frame)
	codeData, n, err := ioutils.ReadLField(buf, nil)
	if err != nil {
		return "", nil, err
	}

	return Code(codeData), frame[n:], nil
}

func SendError(w io.Writer, err error) {
	if err == nil {
		return
	}
	data := []byte(err.Error())
	SendFrame(w, Error, data)
}
func SendFrame(w io.Writer, code Code, data []byte) error {
	codeData := []byte(code)

	size := uint32(len(data) + len(codeData) + 1)
	sizeAll := size + 4
	buf := bytes.NewBuffer(make([]byte, sizeAll, sizeAll))
	buf.Reset()

	err := binary.Write(buf, binary.BigEndian, size)
	if err != nil {
		return err
	}

	_, err = ioutils.WriteLField(buf, codeData)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		_, err := buf.Write(data)
		if err != nil {
			return err
		}
	}
	//b:= buf.Bytes()
	//_,err =w.Write(b)
	_, err = buf.WriteTo(w)
	return err
}
