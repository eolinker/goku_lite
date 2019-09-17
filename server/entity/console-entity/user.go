package entity

type UserInfo struct {
	UserID    int    `json:"userID"`
	LoginCall string `json:"loginCall"`
	Remark    string `json:"remark"`
	UserType  int    `json:"userType"`
	CanDelete bool   `json:"canDelete"`
}
