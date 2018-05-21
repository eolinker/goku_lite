package request

import (
	"io/ioutil"
	"net/http"
	"io"
	"compress/gzip"
)

type Response interface {
	StatusCode() int
	Headers() map[string][]string
	Protocol() string
	Body() []byte
	ContentLength() int64
}

type response struct {
	status         int
	protocol       string
	headers        map[string][]string
	body           []byte
	content_length int64
}

func newResponse(httpResponse *http.Response) (Response, error) {
	defer httpResponse.Body.Close()
	var headers map[string][]string = httpResponse.Header
	var reader io.ReadCloser
	switch httpResponse.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ = gzip.NewReader(httpResponse.Body)
			defer reader.Close()
		default:
			reader = httpResponse.Body
	}
	body, err := ioutil.ReadAll(reader)
	content_length := int64(len(body))
	if err != nil {
		return nil, err
	}
	res := &response{headers: headers,
		protocol:       httpResponse.Proto,
		status:         httpResponse.StatusCode,
		body:           body,
		content_length: int64(content_length)}
	return res, nil
}

func (this *response) StatusCode() int {
	return this.status
}

func (this *response) Headers() map[string][]string {
	return this.headers
}

func (this *response) Protocol() string {
	return this.protocol
}

func (this *response) Body() []byte {
	return this.body
}

func (this *response) ContentLength() int64 {
	return this.content_length
}
