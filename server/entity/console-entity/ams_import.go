package entity

//AmsProject ams项目
type AmsProject struct {
	ProjectInfo  AmsProjectInfo `json:"projectInfo"`
	APIGroupList []AmsGroupInfo `json:"apiGroupList"`
}

//AmsProjectInfo ams项目信息
type AmsProjectInfo struct {
	ProjectName string `json:"projectName"`
}

//AmsGroupInfo ams分组信息
type AmsGroupInfo struct {
	GroupName         string         `json:"groupName"`
	ChildGroupList    []AmsGroupInfo `json:"childGroupList"`
	APIList           []AmsAPIInfo   `json:"apiList"`
	APIGroupChildList []AmsGroupInfo `json:"apiGroupChildList"`
}

//AmsAPIInfo ams接口信息
type AmsAPIInfo struct {
	BaseInfo AmsAPI `json:"baseInfo"`
}

//AmsAPI ams接口
type AmsAPI struct {
	APIName        string `json:"apiName"`
	APIURI         string `json:"apiURI"`
	APIRequestType int    `json:"apiRequestType"`
	APIProtocol    int    `json:"apiProtocol"`
}
