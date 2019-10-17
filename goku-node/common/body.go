package common

//BodyHandler 请求体处理器
type BodyHandler struct {
	body []byte
}

//GetBody 获取body内容
func (r *BodyHandler) GetBody() []byte {
	if r == nil {
		return nil
	}
	return r.body
}

//SetBody 设置body内容
func (r *BodyHandler) SetBody(body []byte) {
	r.body = body
}

//NewBodyHandler 创建BodyHandler
func NewBodyHandler(body []byte) *BodyHandler {
	return &BodyHandler{body: body}
}
