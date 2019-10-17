package common

import (
	"net/http"
	"net/url"

	goku_plugin "github.com/eolinker/goku-plugin"
)

//Header header
type Header struct {
	header http.Header
}

//Headers 获取头部
func (h *Header) Headers() http.Header {

	n := make(http.Header)
	for k, v := range h.header {
		n[k] = v
	}
	return n
}
func (h *Header) String() string {

	return url.Values(h.header).Encode()

}

//SetHeader 设置请求头部
func (h *Header) SetHeader(key, value string) {
	h.header.Set(key, value)
}

//AddHeader 新增头部
func (h *Header) AddHeader(key, value string) {
	h.header.Add(key, value)
}

//DelHeader 删除头部
func (h *Header) DelHeader(key string) {
	h.header.Del(key)
}

//GetHeader 根据名称获取头部
func (h *Header) GetHeader(name string) string {
	return h.header.Get(name)
}

//NewHeader 创建Header
func NewHeader(header http.Header) *Header {
	if header == nil {
		header = make(http.Header)
	}
	return &Header{
		header: header,
	}
}

//PriorityHeader priorityHeader
type PriorityHeader struct {
	*Header
	setHeader    *Header
	appendHeader *Header
}

//Set set
func (h *PriorityHeader) Set() goku_plugin.Header {
	if h.setHeader == nil {
		h.setHeader = NewHeader(nil)
	}
	return h.setHeader
}

//Append append
func (h *PriorityHeader) Append() goku_plugin.Header {
	if h.appendHeader == nil {
		h.appendHeader = NewHeader(nil)
	}
	return h.setHeader
}

//NewPriorityHeader 创建PriorityHeader
func NewPriorityHeader() *PriorityHeader {
	return &PriorityHeader{
		Header:       NewHeader(nil),
		setHeader:    nil,
		appendHeader: nil,
	}
}
