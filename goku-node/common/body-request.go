package common

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"

	goku_plugin "github.com/eolinker/goku-plugin"

	"io/ioutil"
	"net/http"

	"mime"
	"mime/multipart"
	"net/url"
)

const defaultMultipartMemory = 32 << 20 // 32 MB
var (
	errorNotForm      = errors.New("contentType is not Form")
	errorNotMultipart = errors.New("contentType is not Multipart")
	errorNotAllowRaw  = errors.New("contentType is not allow Raw")
)

//BodyRequestHandler body请求处理器
type BodyRequestHandler struct {
	form            url.Values
	rawBody         []byte
	orgContentParam map[string]string
	contentType     string
	files           map[string]*goku_plugin.FileHeader

	isInit     bool
	isWriteRaw bool

	object interface{}
}

//Files 获取文件参数
func (b *BodyRequestHandler) Files() (map[string]*goku_plugin.FileHeader, error) {

	err := b.Parse()

	if err != nil {
		return nil, err
	}
	return b.files, nil

}

//Parse 解析
func (b *BodyRequestHandler) Parse() error {
	if b.isInit {
		return nil
	}

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	switch contentType {
	case goku_plugin.JSON:
		{
			e := json.Unmarshal(b.rawBody, &b.object)
			if e != nil {
				return e
			}
		}
	case goku_plugin.AppLicationXML, goku_plugin.TextXML:
		{
			e := xml.Unmarshal(b.rawBody, &b.object)
			if e != nil {
				return e
			}

		}

	case goku_plugin.MultipartForm:
		{
			r, err := multipartReader(b.contentType, false, b.rawBody)
			if err != nil {
				return err
			}
			form, err := r.ReadForm(defaultMultipartMemory)
			if err != nil {
				return err
			}

			if b.form == nil {
				b.form = make(url.Values)
			}
			for k, v := range form.Value {
				b.form[k] = append(b.form[k], v...)
			}

			b.files = make(map[string]*goku_plugin.FileHeader)
			for k, fs := range form.File {

				if len(fs) > 0 {
					file, err := fs[0].Open()
					if err != nil {
						return err
					}
					fileData, err := ioutil.ReadAll(file)
					if err != nil {
						return err
					}

					b.files[k] = &goku_plugin.FileHeader{
						FileName: fs[0].Filename,
						Data:     fileData,
						Header:   fs[0].Header,
					}
				}
			}

			b.object = b.form
		}
	case goku_plugin.FormData:
		{
			form, err := url.ParseQuery(string(b.rawBody))
			if err != nil {
				return err
			}
			if b.form == nil {
				b.form = form
			} else {
				for k, v := range form {
					b.form[k] = append(b.form[k], v...)
				}

			}
			b.object = b.form

		}
	}
	b.isInit = true
	return nil
}

//GetForm 获取表单参数
func (b *BodyRequestHandler) GetForm(key string) string {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return ""
	}
	b.Parse()

	if !b.isInit || b.form == nil {
		return ""
	}
	return b.form.Get(key)
}

//GetFile 获取文件参数
func (b *BodyRequestHandler) GetFile(key string) (file *goku_plugin.FileHeader, has bool) {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return nil, false
	}

	err := b.Parse()
	if err != nil {
		return nil, false
	}

	if !b.isInit || b.files == nil {
		return nil, false
	}
	f, has := b.files[key]
	return f, has
}

//SetToForm 设置表单参数
func (b *BodyRequestHandler) SetToForm(key, value string) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errorNotForm
	}

	err := b.Parse()
	if err != nil {
		return err
	}
	b.isWriteRaw = false

	if b.form == nil {
		b.form = make(url.Values)
	}
	b.form.Set(key, value)
	b.isWriteRaw = false

	return nil
}

//AddForm 新增表单参数
func (b *BodyRequestHandler) AddForm(key, value string) error {
	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errorNotForm
	}
	err := b.Parse()
	if err != nil {
		return err
	}
	b.isWriteRaw = false

	if b.form == nil {
		b.form = make(url.Values)
	}
	b.form.Add(key, value)
	return nil
}

//AddFile 新增文件参数
func (b *BodyRequestHandler) AddFile(key string, file *goku_plugin.FileHeader) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errorNotMultipart
	}
	err := b.Parse()
	if err != nil {
		return err
	}
	b.isWriteRaw = false
	if file == nil && b.files != nil {
		delete(b.files, key)
		return nil
	}
	if b.files == nil {
		b.files = make(map[string]*goku_plugin.FileHeader)
	}
	b.files[key] = file

	return nil
}

//Clone 克隆body
func (b *BodyRequestHandler) Clone() *BodyRequestHandler {
	rawbody, _ := b.RawBody()
	return NewBodyRequestHandler(b.contentType, rawbody)

}

//ContentType 获取contentType
func (b *BodyRequestHandler) ContentType() string {
	return b.contentType
}

//BodyForm 获取表单参数
func (b *BodyRequestHandler) BodyForm() (url.Values, error) {

	err := b.Parse()
	if err != nil {
		return nil, err
	}
	return b.form, nil
}

//BodyInterface 获取请求体对象
func (b *BodyRequestHandler) BodyInterface() (interface{}, error) {
	err := b.Parse()
	if err != nil {
		return nil, err
	}

	return b.object, nil
}

//RawBody 获取raw数据
func (b *BodyRequestHandler) RawBody() ([]byte, error) {

	err := b.Encode()
	if err != nil {
		return nil, err
	}
	return b.rawBody, nil

}

//Encode encode
func (b *BodyRequestHandler) Encode() error {
	if b.isWriteRaw {
		return nil
	}

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		b.isWriteRaw = true
		return nil
	}

	if len(b.files) > 0 {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		for name, file := range b.files {
			part, err := writer.CreateFormFile(name, file.FileName)
			if err != nil {
				return err
			}
			_, err = part.Write(file.Data)
			if err != nil {
				return err
			}
		}

		for key, values := range b.form {
			temp := make(url.Values)
			temp[key] = values
			value := temp.Encode()
			err := writer.WriteField(key, value)
			if err != nil {
				return err
			}
		}
		err := writer.Close()
		if err != nil {
			return err
		}
		b.contentType = writer.FormDataContentType()
		b.rawBody = body.Bytes()
		b.isWriteRaw = true
	} else {
		if b.form != nil {
			b.rawBody = []byte(b.form.Encode())
		} else {
			b.rawBody = make([]byte, 0, 0)
		}
	}
	return nil
}

//SetForm 设置表单参数
func (b *BodyRequestHandler) SetForm(values url.Values) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errorNotForm
	}
	b.Parse()
	b.form = values
	b.isWriteRaw = false

	return nil
}

//SetFile 设置文件参数
func (b *BodyRequestHandler) SetFile(files map[string]*goku_plugin.FileHeader) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errorNotForm
	}
	b.Parse()
	b.files = files
	// b.form = values
	b.isWriteRaw = false

	return nil
}

//SetRaw 设置raw数据
func (b *BodyRequestHandler) SetRaw(contentType string, body []byte) {

	b.rawBody, b.contentType, b.isInit, b.isWriteRaw = body, contentType, false, true
	_, b.orgContentParam, _ = mime.ParseMediaType(contentType)
	return

}

//NewBodyRequestHandler 创建body请求处理器
func NewBodyRequestHandler(contentType string, body []byte) *BodyRequestHandler {
	b := new(BodyRequestHandler)
	b.SetRaw(contentType, body)
	return b
}

func multipartReader(contentType string, allowMixed bool, raw []byte) (*multipart.Reader, error) {

	if contentType == "" {
		return nil, http.ErrNotMultipart
	}
	d, params, err := mime.ParseMediaType(contentType)
	if err != nil || !(d == "multipart/form-data" || allowMixed && d == "multipart/mixed") {
		return nil, http.ErrNotMultipart
	}
	boundary, ok := params["boundary"]
	if !ok {
		return nil, http.ErrMissingBoundary
	}
	body := ioutil.NopCloser(bytes.NewBuffer(raw))
	return multipart.NewReader(body, boundary), nil
}
