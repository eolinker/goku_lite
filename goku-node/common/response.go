package common

import (
	"io/ioutil"
	"net/http"
)

//ResponseReader 响应结构体
type ResponseReader struct {
	*CookiesHandler
	*Header
	*BodyHandler
	*StatusHandler
}

func newResponseReader(response *http.Response) *ResponseReader {
	if response == nil {
		return nil
	}
	r := new(ResponseReader)
	r.Header = NewHeader(response.Header)
	r.CookiesHandler = newCookieHandle(response.Header)
	r.StatusHandler = NewStatusHandler()
	r.SetStatus(response.StatusCode, response.Status)
	// if response.ContentLength > 0 {
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	r.BodyHandler = NewBodyHandler(body)
	// } else {
	// 	r.BodyHandler = NewBodyHandler(nil)
	// }
	body, _ := ioutil.ReadAll(response.Body)
	r.BodyHandler = NewBodyHandler(body)

	return r
}

//NewResponseReader 新增ResponseReader
func NewResponseReader(header http.Header, statusCode int, status string, body []byte) *ResponseReader {
	r := new(ResponseReader)
	r.Header = NewHeader(header)
	r.CookiesHandler = newCookieHandle(header)
	r.StatusHandler = NewStatusHandler()
	r.SetStatus(statusCode, status)
	// if response.ContentLength > 0 {
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	r.BodyHandler = NewBodyHandler(body)
	// } else {
	// 	r.BodyHandler = NewBodyHandler(nil)
	// }

	r.BodyHandler = NewBodyHandler(body)
	return r
}
