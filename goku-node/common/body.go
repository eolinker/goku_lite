package common

type BodyHandler struct {
	body []byte
}

func (r *BodyHandler) GetBody() []byte {
	if r == nil {
		return nil
	}
	return r.body
}

func (r *BodyHandler) SetBody(body []byte) {
	r.body = body
}

func NewBodyHandler(body []byte) *BodyHandler {
	return &BodyHandler{body: body}
}
