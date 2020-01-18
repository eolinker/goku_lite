package ioutils

import (
	"encoding/binary"
	"io"
)

//WriteLField 写域内容
func WriteLField(writer io.Writer, s []byte) (int, error) {

	l := uint8(len(s))

	e := binary.Write(writer, binary.BigEndian, l)
	if e != nil {
		return 0, e
	}
	n, err := writer.Write(s)
	if err != nil {
		return n + 1, err
	}
	return n + 1, nil
}

//ReadLField 读取域内容
func ReadLField(reader io.Reader, buf []byte) ([]byte, int, error) {
	l := uint8(0)

	e := binary.Read(reader, binary.BigEndian, &l)
	if e != nil {
		return nil, 0, e
	}
	tmpbuf := buf
	if int(l) > len(buf) {
		tmpbuf = make([]byte, l, l)
	} else {
		tmpbuf = buf[:l]
	}

	_, err := reader.Read(tmpbuf)
	if err != nil {
		return nil, 0, err
	}
	return tmpbuf, int(l) + 1, nil
}
