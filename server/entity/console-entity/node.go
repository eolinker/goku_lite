package entity

//Node 节点信息
type Node struct {
	NodeID        int    `json:"nodeID"`
	NodeName      string `json:"nodeName"`
	NodeKey       string `json:"nodeKey"`
	ListenAddress string `json:"listenAddress"`
	AdminAddress  string `json:"adminAddress"`
	Cluster       string `json:"cluster,omitempty"`
	ClusterTitle  string `json:"cluster_title,omitempty"`
	Version       string `json:"version"`
	NodeStatus    int    `json:"nodeStatus"`
	GroupID       int    `json:"groupID,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	IsUpdate      bool   `json:"isUpdate"`
	GatewayPath   string `json:"gatewayPath"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
	UpdatePeriod  int    `json:"updatePeriod,omitempty"`
	*SSHInfo
}

//SSHInfo sshInfo
type SSHInfo struct {
	SSHAddress string `json:"sshAddress"`
	UserName   string `json:"sshUserName"`
	Password   string `json:"sshPassword"`
	Key        string `json:"sshKey"`
	AuthMethod int    `json:"authMethod"`
	IsSave     int    `json:"isSave"`
}
