package entity

type Project struct {
	ProjectID   int    `json:"projectID"`
	ProjectName string `json:"projectName"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
