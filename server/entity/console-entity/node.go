package entity

type Node struct {
	NodeID   int    `json:"nodeID"`
	NodeName string `json:"nodeName"`
	NodeIP   string `json:"nodeIP"`
	NodePort string `json:"nodePort"`

	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	UpdatePeriod int    `json:"updatePeriod,omitempty"`
	Version      string `json:"version"`
	NodeStatus   int    `json:"nodeStatus"`
	GroupID      int    `json:"groupID,omitempty"`
	GroupName    string `json:"groupName,omitempty"`
	IsUpdate     bool   `json:"isUpdate"`
	Cluster      string `json:"cluster"`
	ClusterTitle string `json:"cluster_title"`
	GatewayPath  string `json:"gatewayPath"`
	//*SSHInfo
}

//type SSHInfo struct {
//	SSHPort     string `json:"sshPort"`
//	UserName    string `json:"userName"`
//	Password    string `json:"password"`
//
//	Key         string `json:"key"`
//	AuthMethod  int    `json:"authMethod"`
//}
