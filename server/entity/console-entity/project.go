package entity

//Project 项目
type Project struct {
	ProjectID   int    `json:"projectID"`
	ProjectName string `json:"projectName"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
