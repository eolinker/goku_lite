package entity

type Message struct {
	MsgID      int    `json:"msgID"`
	Msg        string `json:"msg"`
	UpdateTime string `json:"updateTime"`
	MsgType    int    `json:"msgType"`
}
