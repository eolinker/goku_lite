package entity

//AlertInfo 告警信息
type AlertInfo struct {
	ReceiverList string `json:"receiverList"`
	AlertAddr    string `json:"alertAddr"`
}
