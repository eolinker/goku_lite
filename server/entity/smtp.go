package entity

//SMTPInfo SMTP信息
type SMTPInfo struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Protocol int    `json:"protocol"`
	Sender   string `json:"sender"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
