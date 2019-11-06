package ksitigarbha

//Model 模板数据
type Model struct {
	Type     string                   `json:"type"`
	Label    string                   `json:"label"`
	Describe string                   `json:"descript"`
	Items    []map[string]interface{} `json:"items"`
}
