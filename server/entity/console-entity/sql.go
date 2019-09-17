package entity

type Table struct {
	TableName   string   `json:"tableName"`
	TableColumn []string `json:"tableColumn"`
}

type TableData struct {
	Data []map[string]interface{}
}

type ColumnInfo struct {
	FieldName string
	Type      interface{}
	Null      interface{}
	Key       interface{}
	Default   interface{}
	Extra     interface{}
}

type GokuAdmin struct {
	UserID        int    `json:"userID"`
	LoginCall     string `json:"loginCall"`
	LoginPassword string `json:"loginPassword"`
	UserType      int    `json:"userType"`
}

type GokuBalance struct {
	BalanceID     int    `json:"balanceID"`
	BalanceName   string `json:"balanceName"`
	BalanceConfig string `json:"balanceConfig"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

type GokuConnPluginApi struct {
	ConnID int `json:"connID"`
	ApiID  int `json:"apiID"`
}

type ColumnValue struct {
	Value interface{}
}
