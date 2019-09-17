package entity

type AmsProject struct {
	ProjectInfo  AmsProjectInfo `json:"projectInfo"`
	ApiGroupList []AmsGroupInfo `json:"apiGroupList"`
}

type AmsProjectInfo struct {
	ProjectName string `json:"projectName"`
}

type AmsGroupInfo struct {
	GroupName         string         `json:"groupName"`
	ChildGroupList    []AmsGroupInfo `json:"childGroupList"`
	ApiList           []AmsApiInfo   `json:"apiList"`
	ApiGroupChildList []AmsGroupInfo `json:"apiGroupChildList"`
}

type AmsApiInfo struct {
	BaseInfo AmsApi `json:"baseInfo"`
}

type AmsApi struct {
	ApiName        string `json:"apiName"`
	ApiURI         string `json:"apiURI"`
	ApiRequestType int    `json:"apiRequestType"`
}
