package goku_plugin

import (
	"net/http"
	"net/textproto"
	"net/url"
)

type HeaderReader interface {
	GetHeader(name string) string
	// 返回所有header，返回值为一个副本，对他的修改不会生效
	Headers() http.Header
}
type HeaderWriter interface {
	SetHeader(key, value string)
	AddHeader(key, value string)
	DelHeader(key string)
}
type Header interface {
	HeaderReader
	HeaderWriter
}

type CookieReader interface {
	Cookie(name string) (*http.Cookie, error)
	Cookies() []*http.Cookie
}
type CookieWriter interface {
	AddCookie(c *http.Cookie)
}
type BodyGet interface {
	GetBody() []byte
}

type BodySet interface {
	SetBody([]byte)
}
type FileHeader struct {
	FileName string
	Header   textproto.MIMEHeader
	Data     []byte
}
type BodyDataReader interface {
	//Parse() error

	//Protocol() RequestType
	ContentType() string
	//content-Type = application/x-www-form-urlencoded 或 multipart/form-data，与原生request.Form不同，这里不包括 query 参数
	BodyForm() (url.Values, error)
	//content-Type = multipart/form-data 时有效
	Files() (map[string]*FileHeader, error)
	GetForm(key string) string
	GetFile(key string) (file *FileHeader, has bool)
	RawBody() ([]byte, error)
	//Encode()[]byte // 最终数据
}
type BodyDatawriter interface {
	//设置form数据并将content-type设置 为 application/x-www-form-urlencoded 或 multipart/form-data
	SetForm(values url.Values) error
	SetToForm(key, value string) error
	AddForm(key, value string) error
	// 会替换掉对应掉file信息，并且将content-type 设置为 multipart/form-data
	AddFile(key string, file *FileHeader) error
	//设置 multipartForm 数据并将content-type设置 为 multipart/form-data

	// 重置body，会清除掉未处理掉 form和file
	SetRaw(contentType string, body []byte)
}
type Body interface {
	BodyDataReader
	BodyDatawriter
}

type StatusGet interface {
	StatusCode() int
	Status() string
}
type StatusSet interface {
	SetStatus(code int, status string)
}
type RequestData interface {
	BodyDataReader
	Method() string
	URL() *url.URL
	RequestURI() string
	Host() string
	RemoteAddr() string
	Proto() string
}

//type RequestGet interface {
//	CookieReader
//	HeaderReader
//	BodyGet
//}

// 原始请求数据的读
type RequestReader interface {
	CookieReader
	HeaderReader
	RequestData
}

// 用于组装转发的request
type Request interface {
	CookieReader
	HeaderReader
	HeaderWriter
	CookieWriter
	Body
	Querys() url.Values
	TargetServer() string
	TargetURL() string
}

// 读取转发结果的response
type ResponseReader interface {
	CookieReader
	HeaderReader
	BodyGet
	StatusGet
}

// 请求基本接口信息
type ContextApiInfo interface {
	StrategyId() string
	StrategyName() string
	ApiID() int
}

// 单存储
type Store interface {
	Set(value interface{})
	Get() (value interface{})
}

// 存储容器
type StoreContainer interface {
	Store() Store // 私有存储
	SetCache(name string, value interface{})
	GetCache(name string) (value interface{}, has bool)
}

// 带优先的header
type PriorityHeader interface {
	HeaderReader // 读已经设置的header
	HeaderWriter // 设置header
	// 非Priority的header会被 proxy 的同名项替换掉，
	Set() Header    // 这里设置的header会替换掉proxy的内容
	Append() Header // 这里设置的header会追加到proxy的内容
}

// 返回给client端的
type ResponseWriter interface {
	PriorityHeader
	CookieReader // 已经设置的cookie
	CookieWriter // 设置返回的cookie
	StatusGet
	StatusSet // 设置返回状态
	BodySet   // 设置返回内容
	BodyGet
}

type Context interface {
	ResponseWriter // 处理返回
	StoreContainer // cache
	RequestId() string
	FinalTargetServer()string
	RetryTargetServers()string
}

type ContextBeforeMatch interface {
	Context // 处理返回
	Request() RequestReader
	Proxy() Request // 请求信息，包含原始请求数据以及被更高优先级处理过的结果
}

// 转发前
type ContextAccess interface {
	Context        // 处理返回
	ContextApiInfo // api 信息
	Request() RequestReader
	Proxy() Request // 请求信息，包含原始请求数据以及被更高优先级处理过的结果
}

// 转发后
type ContextProxy interface {
	Context                        // 处理返回
	ContextApiInfo                 // api 信息
	ProxyResponse() ResponseReader // 转发后返回的结果
}
