package entity

//UserInfo 用户信息
type UserInfo struct {
	UserID    int    `json:"userID"`
	LoginCall string `json:"loginCall"`
	Remark    string `json:"remark"`
	UserType  int    `json:"userType"`
	CanDelete bool   `json:"canDelete"`
}
