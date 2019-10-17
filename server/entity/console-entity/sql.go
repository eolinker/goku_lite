package entity

//Table 表结构
type Table struct {
	TableName   string   `json:"tableName"`
	TableColumn []string `json:"tableColumn"`
}

//TableData 表数据
type TableData struct {
	Data []map[string]interface{}
}

//ColumnInfo 列信息
type ColumnInfo struct {
	FieldName string
	Type      interface{}
	Null      interface{}
	Key       interface{}
	Default   interface{}
	Extra     interface{}
}

//GokuAdmin admin信息
type GokuAdmin struct {
	UserID        int    `json:"userID"`
	LoginCall     string `json:"loginCall"`
	LoginPassword string `json:"loginPassword"`
	UserType      int    `json:"userType"`
}

//GokuBalance 负载信息
type GokuBalance struct {
	BalanceID     int    `json:"balanceID"`
	BalanceName   string `json:"balanceName"`
	BalanceConfig string `json:"balanceConfig"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

//GokuConnPluginAPI 接口和插件绑定信息
type GokuConnPluginAPI struct {
	ConnID int `json:"connID"`
	APIID  int `json:"apiID"`
}

//ColumnValue 列值
type ColumnValue struct {
	Value interface{}
}
