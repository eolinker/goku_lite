package common

import (
	"bytes"
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
	errNotForm      = errors.New("contentType is not Form")
	errNotMultipart = errors.New("contentType is not Multipart")
	errNotAllowRaw  = errors.New("contentType is not allow Raw")
)

//BodyRequestHandler body request handler
type BodyRequestHandler struct {
	form            url.Values
	rawbody         []byte
	orgContentParam map[string]string
	contentType     string
	files           map[string]*goku_plugin.FileHeader

	isInit     bool
	isWriteRaw bool
}

//Files files
func (b *BodyRequestHandler) Files() (map[string]*goku_plugin.FileHeader, error) {

	err := b.Parse()

	if err != nil {
		return nil, err
	}
	return b.files, nil

}

//Parse parse
func (b *BodyRequestHandler) Parse() error {

	if b.isInit {
		return nil
	}

	contentType, _, _ := mime.ParseMediaType(b.contentType)

	switch contentType {
	case goku_plugin.MultipartForm:
		{
			r, err := multipartReader(b.contentType, false, b.rawbody)
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
		}
	case goku_plugin.FormData:
		{
			form, err := url.ParseQuery(string(b.rawbody))
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

		}
	}
	b.isInit = true
	return nil
}

//GetForm get form
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

//GetFile getFile
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

//SetToForm setToForm
func (b *BodyRequestHandler) SetToForm(key, value string) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errNotForm
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

//AddForm addForm
func (b *BodyRequestHandler) AddForm(key, value string) error {
	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errNotForm
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

//AddFile 新建文件参数
func (b *BodyRequestHandler) AddFile(key string, file *goku_plugin.FileHeader) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errNotMultipart
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

//Clone 请求克隆
func (b *BodyRequestHandler) Clone() *BodyRequestHandler {

	rawbody, _ := b.RawBody()

	return NewBodyRequestHandler(b.contentType, rawbody)

}

//ContentType contentType
func (b *BodyRequestHandler) ContentType() string {
	return b.contentType
}

//BodyForm 获取body参数
func (b *BodyRequestHandler) BodyForm() (url.Values, error) {

	err := b.Parse()
	if err != nil {
		return nil, err
	}
	return b.form, nil
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

		for fieldname, file := range b.files {
			part, err := writer.CreateFormFile(fieldname, file.FileName)
			if err != nil {
				return err
			}
			_, err = part.Write(file.Data)
			if err != nil {
				return err
			}
		}

		for fieldname, values := range b.form {
			temp := make(url.Values)
			temp[fieldname] = values
			value := temp.Encode()
			err := writer.WriteField(fieldname, value)
			if err != nil {
				return err
			}
		}
		err := writer.Close()
		if err != nil {
			return err
		}
		b.contentType = writer.FormDataContentType()
		b.rawbody = body.Bytes()
		b.isWriteRaw = true
	} else {
		if b.form != nil {
			b.rawbody = []byte(b.form.Encode())
		} else {
			b.rawbody = make([]byte, 0, 0)
		}
	}
	return nil
}

//RawBody rawBody
func (b *BodyRequestHandler) RawBody() ([]byte, error) {

	err := b.Encode()
	if err != nil {
		return nil, err
	}
	return b.rawbody, nil

}

//SetForm 设置表单参数
func (b *BodyRequestHandler) SetForm(values url.Values) error {

	contentType, _, _ := mime.ParseMediaType(b.contentType)
	if contentType != goku_plugin.FormData && contentType != goku_plugin.MultipartForm {
		return errNotForm
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
		return errNotForm
	}
	b.Parse()
	b.files = files
	// b.form = values
	b.isWriteRaw = false

	return nil
}

//SetRaw 设置Raw
func (b *BodyRequestHandler) SetRaw(contentType string, body []byte) {

	b.rawbody, b.contentType, b.isInit, b.isWriteRaw = body, contentType, false, true
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
