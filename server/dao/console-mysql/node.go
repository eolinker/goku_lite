package consolemysql

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/database"
	v "github.com/eolinker/goku-api-gateway/common/version"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
	"strconv"
	"strings"
	"time"
)

// AddNode 新增节点信息
func AddNode(clusterID int, nodeName, nodeIP, nodePort, gatewayPath string, groupID int) (bool, map[string]interface{}, error) {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO goku_node_info (`clusterID`,`nodeName`,`groupID`,`nodeIP`,`nodePort`,`updateTime`,`createTime`,`version`, `gatewayPath`,`nodeStatus`) VALUES (?,?,?,?,?,?,?,?,?,0);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Illegal SQL statement!"}, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(clusterID, nodeName, groupID, nodeIP, nodePort, now, now, v.Version, gatewayPath)
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Failed to insert data!"}, err
	}
	nodeID, err := res.LastInsertId()
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Failed to insert data!"}, err
	}
	return true, map[string]interface{}{"nodeID": nodeID, "version": v.Version}, nil
}

// EditNode 修改节点信息
func EditNode(nodeName, nodeIP, nodePort, gatewayPath string, nodeID, groupID int) (bool, string, error) {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE goku_node_info SET  nodeName = ?,nodeIP = ?,nodePort = ?,updateTime = ?,groupID = ?,gatewayPath = ? WHERE nodeID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeName, nodeIP, nodePort, now, groupID, gatewayPath, nodeID)
	if err != nil {
		return false, "[ERROR]Failed to update data!", err
	}
	return true, "", nil
}

// DeleteNode 删除节点信息
func DeleteNode(nodeID int) (bool, string, error) {
	db := database.GetConnection()
	sql := "DELETE FROM goku_node_info WHERE nodeID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return false, "[ERROR]Failed to delete data!", err
	}
	return true, "", nil
}

// GetNodeList 获取节点列表
func GetNodeList(clusterID, groupID int, keyword string) (bool, []*entity.Node, error) {
	db := database.GetConnection()
	rule := make([]string, 0, 2)

	rule = append(rule, fmt.Sprintf("A.clusterID = %d", clusterID))
	if groupID > -1 {
		groupRule := fmt.Sprintf("A.groupID = %d", groupID)
		rule = append(rule, groupRule)
	}
	if keyword != "" {
		searchRule := "(A.nodeName LIKE '%" + keyword + "%' OR A.nodeIP LIKE '%" + keyword + "%' OR A.nodePort LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.nodeID,A.nodeName,A.nodeIP,A.nodePort,A.updateTime,A.createTime,A.version,A.gatewayPath,A.groupID,IFNULL(G.groupName,'未分类') FROM goku_node_info A LEFT JOIN goku_node_group G ON A.groupID = G.groupID %s ORDER BY updateTime DESC;", ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	nodeList := make([]*entity.Node, 0)
	for rows.Next() {
		node := entity.Node{}
		err = rows.Scan(&node.NodeID, &node.NodeName, &node.NodeIP, &node.NodePort, &node.UpdateTime, &node.CreateTime, &node.Version, &node.GatewayPath, &node.GroupID, &node.GroupName)
		if err != nil {
			panic(err)
			return false, nil, err
		}
		if node.Version == v.Version {
			// 判断节点版本号是否是最新
			node.IsUpdate = true
		}
		nodeList = append(nodeList, &node)
	}
	return true, nodeList, nil
}

const nodeSQLIPPort = "SELECT  A.`nodeID`, A.`nodeName`, A.`nodeIP`, A.`nodePort`, A.`updateTime`, A.`createTime`, A.`version`,   A.`gatewayPath`, A.`groupID`, IFNULL(G.`groupName`,'') , C.`name` As cluster, C.`title` As cluster_title  FROM goku_node_info A LEFT JOIN goku_node_group G ON A.`groupID` = G.`groupID` LEFT JOIN `goku_cluster` C ON A.`clusterID` = C.`id`WHERE A.`nodeIP` = ? and A.`nodePort`=?;"
const nodeSQLID = "SELECT  A.`nodeID`, A.`nodeName`, A.`nodeIP`, A.`nodePort`, A.`updateTime`, A.`createTime`, A.`version`,   A.`gatewayPath`, A.`groupID`, IFNULL(G.`groupName`,''),  C.`name` As cluster, C.`title` As cluster_title  FROM goku_node_info A LEFT JOIN goku_node_group G ON A.`groupID` = G.`groupID` LEFT JOIN `goku_cluster` C ON A.`clusterID` = C.`id`WHERE A.`nodeID` = ? ;"

func getNodeInfo(sql string, args ...interface{}) (bool, *entity.Node, error) {

	db := database.GetConnection()

	node := &entity.Node{}
	err := db.QueryRow(sql, args...).Scan(&node.NodeID,
		&node.NodeName,
		&node.NodeIP,
		&node.NodePort,
		&node.UpdateTime,
		&node.CreateTime,
		&node.Version,
		&node.GatewayPath,
		&node.GroupID,
		&node.GroupName,
		&node.Cluster,
		&node.ClusterTitle)
	if err != nil {
		return false, nil, err
	}
	return true, node, err
}

// GetNodeInfo 获取节点信息
func GetNodeInfo(nodeID int) (bool, *entity.Node, error) {

	return getNodeInfo(nodeSQLID, nodeID)
}

//GetNodeByIPPort 通过IP+端口获取节点信息
func GetNodeByIPPort(ip string, port int) (bool, *entity.Node, error) {

	return getNodeInfo(nodeSQLIPPort, ip, port)
}

// CheckIsExistRemoteAddr 节点IP查重
func CheckIsExistRemoteAddr(nodeID int, nodeIP, nodePort string) bool {
	db := database.GetConnection()
	sql := `SELECT nodeID FROM goku_node_info WHERE nodeIP = ? AND nodePort = ?;`
	var id int
	err := db.QueryRow(sql, nodeIP, nodePort).Scan(&id)
	if err != nil {
		return false
	}
	if id == nodeID {
		return false
	}
	return true
}

// GetNodeIPList 获取节点IP列表
func GetNodeIPList() (bool, []map[string]interface{}, error) {
	db := database.GetConnection()
	sql := `SELECT nodeID,nodeIP,nodePort FROM goku_node_info WHERE nodeStatus = 1;`
	rows, err := db.Query(sql)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	nodeList := make([]map[string]interface{}, 0)

	for rows.Next() {
		var nodeID int
		var nodeIP, nodePort string
		err = rows.Scan(&nodeID, &nodeIP, &nodePort)
		if err != nil {
			return false, make([]map[string]interface{}, 0), err
		}
		nodeList = append(nodeList, map[string]interface{}{
			"nodeID":   nodeID,
			"nodeIP":   nodeIP,
			"nodePort": nodePort,
		})
	}
	return true, nodeList, nil
}

// GetAvaliableNodeListFromNodeList 从待操作节点中获取关闭节点列表
func GetAvaliableNodeListFromNodeList(nodeIDList string, nodeStatus int) (bool, string, error) {
	db := database.GetConnection()
	sql := "SELECT nodeID FROM goku_node_info WHERE nodeID IN (" + nodeIDList + ") AND nodeStatus = ?"
	rows, err := db.Query(sql, nodeStatus)
	if err != nil {
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	defer rows.Close()
	idList := make([]string, 0)
	for rows.Next() {
		var nodeID int
		err = rows.Scan(&nodeID)
		if err != nil {
			return false, err.Error(), err
		}
		idList = append(idList, strconv.Itoa(nodeID))
	}
	fmt.Println(sql, nodeStatus)
	return true, strings.Join(idList, ","), nil
}

// BatchEditNodeGroup 批量修改节点分组
func BatchEditNodeGroup(nodeIDList string, groupID int) (bool, string, error) {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "UPDATE goku_node_info SET groupID = ?,updateTime = ? WHERE nodeID IN (" + nodeIDList + ");"
	_, err := Tx.Exec(sql, groupID, now)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	Tx.Commit()
	return true, "", nil
}

// BatchDeleteNode 批量修改接口分组
func BatchDeleteNode(nodeIDList string) (bool, string, error) {
	db := database.GetConnection()
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_node_info WHERE nodeID IN (" + nodeIDList + ");"
	_, err := Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	Tx.Commit()
	return true, "", nil
}

// UpdateAllNodeClusterID 更新节点集群ID
func UpdateAllNodeClusterID(clusterID int) {
	db := database.GetConnection()
	Tx, _ := db.Begin()
	sql := "UPDATE goku_node_info SET clusterID = ?;"
	_, err := Tx.Exec(sql, clusterID)
	if err != nil {
		Tx.Rollback()
	}
	sql = "UPDATE goku_node_group SET clusterID = ?;"
	_, err = Tx.Exec(sql, clusterID)
	if err != nil {
		Tx.Rollback()
	}
	Tx.Commit()
}
