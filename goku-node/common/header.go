package common

import (

	goku_plugin "github.com/eolinker/goku-plugin"
	"net/http"
	"net/url"

)

type Header struct {
	header http.Header
}

func (h *Header) Headers() http.Header {

	n := make(http.Header)
	for k, v := range h.header {
		n[k] = v
	}
	return n
}
func (h *Header) String()string  {

	return url.Values(h.header).Encode()
	//buf:=bytes.NewBuffer(nil)
	//for k,v:=range h.header{
	//	buf.WriteByte('&')
	//	buf.WriteString(k)
	//	buf.WriteString("=")
	//	buf.WriteString(strings.Join(v,","))
	//
	//}
	//data:=buf.Bytes()
	//if len(data)>1{
	//	data=data[1:]
	//}
	//return string(data)
}
func (h *Header) SetHeader(key, value string) {
	h.header.Set(key, value)
}

func (h *Header) AddHeader(key, value string) {
	h.header.Add(key, value)
}

func (h *Header) DelHeader(key string) {
	h.header.Del(key)
}

func (h *Header) GetHeader(name string) string {
	return h.header.Get(name)
}
func NewHeader(header http.Header) *Header {
	if header == nil {
		header = make(http.Header)
	}
	return &Header{
		header: header,
	}
}

type PriorityHeader struct {
	*Header
	setHeader    *Header
	appendHeader *Header
}

func (h *PriorityHeader) Set() goku_plugin.Header {
	if h.setHeader == nil {
		h.setHeader = NewHeader(nil)
	}
	return h.setHeader
}

func (h *PriorityHeader) Append() goku_plugin.Header {
	if h.appendHeader == nil {
		h.appendHeader = NewHeader(nil)
	}
	return h.setHeader
}

func NewPriorityHeader() *PriorityHeader {
	return &PriorityHeader{
		Header:       NewHeader(nil),
		setHeader:    nil,
		appendHeader: nil,
	}
}
