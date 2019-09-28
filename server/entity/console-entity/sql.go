package entity

//Table table
type Table struct {
	TableName   string   `json:"tableName"`
	TableColumn []string `json:"tableColumn"`
}

//TableData tableData
type TableData struct {
	Data []map[string]interface{}
}

//ColumnInfo column info
type ColumnInfo struct {
	FieldName string
	Type      interface{}
	Null      interface{}
	Key       interface{}
	Default   interface{}
	Extra     interface{}
}

//GokuAdmin 网关超级管理员信息
type GokuAdmin struct {
	UserID        int    `json:"userID"`
	LoginCall     string `json:"loginCall"`
	LoginPassword string `json:"loginPassword"`
	UserType      int    `json:"userType"`
}

//GokuBalance 网关负载
type GokuBalance struct {
	BalanceID     int    `json:"balanceID"`
	BalanceName   string `json:"balanceName"`
	BalanceConfig string `json:"balanceConfig"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

//GokuConnPluginAPI goku conn plgin api
type GokuConnPluginAPI struct {
	ConnID int `json:"connID"`
	APIID  int `json:"apiID"`
}

//ColumnValue column value
type ColumnValue struct {
	Value interface{}
}
