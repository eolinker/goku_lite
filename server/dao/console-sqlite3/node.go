package console_sqlite3

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/common/database"
	v "github.com/eolinker/goku-api-gateway/common/version"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddNode 新增节点信息
func AddNode(clusterID int, nodeName, nodeKey, listenAddress, adminAddress, gatewayPath string, groupID int) (bool, map[string]interface{}, error) {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO goku_node_info (`clusterID`,`nodeName`,`groupID`,`nodeKey`,`listenAddress`,`adminAddress`,`updateTime`,`createTime`,`version`, `gatewayPath`,`nodeStatus`) VALUES (?,?,?,?,?,?,?,?,?,?,0);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Illegal SQL statement!"}, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(clusterID, nodeName, groupID, nodeKey, listenAddress, adminAddress, now, now, v.Version, gatewayPath)
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Failed to insert data!"}, err
	}
	nodeID, err := res.LastInsertId()
	if err != nil {
		return false, map[string]interface{}{"error": "[ERROR]Failed to insert data!"}, err
	}
	return true, map[string]interface{}{"nodeID": nodeID, "version": v.Version}, nil
}

//EditNode 修改节点信息
func EditNode(nodeName, listenAddress, adminAddress, gatewayPath string, nodeID, groupID int) (bool, string, error) {
	db := database.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE goku_node_info SET  nodeName = ?,listenAddress = ?,adminAddress = ?,updateTime = ?,groupID = ?,gatewayPath = ? WHERE nodeID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeName, listenAddress, adminAddress, now, groupID, gatewayPath, nodeID)
	if err != nil {
		return false, "[ERROR]Failed to update data!", err
	}
	return true, "", nil
}

//DeleteNode 删除节点信息
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
		searchRule := fmt.Sprint("(A.nodeName LIKE '%", keyword, "%' OR A.`listenAddress` LIKE '%", keyword, "%'  OR A.`nodeKey` LIKE '%", keyword, "%')")
		rule = append(rule, searchRule)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.nodeID,A.nodeName,A.nodeKey,A.listenAddress,A.adminAddress,A.updateTime,A.createTime,A.version,A.gatewayPath,A.groupID,IFNULL(G.groupName,'未分类') FROM goku_node_info A LEFT JOIN goku_node_group G ON A.groupID = G.groupID %s ORDER BY updateTime DESC;", ruleStr)
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
		err = rows.Scan(&node.NodeID, &node.NodeName, &node.NodeKey, &node.ListenAddress, &node.AdminAddress, &node.UpdateTime, &node.CreateTime, &node.Version, &node.GatewayPath, &node.GroupID, &node.GroupName)
		if err != nil {
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

const nodeSQLID = "SELECT  A.`nodeID`, A.`nodeName`,A.`listenAddress`,A.`adminAddress`, A.`nodeKey`,  A.`updateTime`, A.`createTime`, A.`version`,   A.`gatewayPath`, A.`groupID`, IFNULL(G.`groupName`,''),  C.`name` As cluster, C.`title` As cluster_title  FROM goku_node_info A LEFT JOIN goku_node_group G ON A.`groupID` = G.`groupID` LEFT JOIN `goku_cluster` C ON A.`clusterID` = C.`id` WHERE A.`nodeID` = ? ;"
const nodeSQLInstance = "SELECT  A.`nodeID`, A.`nodeName`,A.`listenAddress`,A.`adminAddress`, A.`nodeKey`,  A.`updateTime`, A.`createTime`, A.`version`,   A.`gatewayPath`, A.`groupID`, IFNULL(G.`groupName`,'') , C.`name` As cluster, C.`title` As cluster_title  FROM goku_node_info A LEFT JOIN goku_node_group G ON A.`groupID` = G.`groupID` LEFT JOIN `goku_cluster` C ON A.`clusterID` = C.`id` WHERE A.`nodeKey` = ? ;"

func getNodeInfo(sql string, args ...interface{}) (*entity.Node, error) {

	db := database.GetConnection()

	node := &entity.Node{}
	err := db.QueryRow(sql, args...).Scan(&node.NodeID,
		&node.NodeName,
		&node.ListenAddress,
		&node.AdminAddress,
		&node.NodeKey,
		&node.UpdateTime,
		&node.CreateTime,
		&node.Version,
		&node.GatewayPath,
		&node.GroupID,
		&node.GroupName,
		&node.Cluster,
		&node.ClusterTitle)
	if err != nil {
		return nil, err
	}
	return node, err
}

//GetNodeInfo 获取节点信息
func GetNodeInfo(nodeID int) (*entity.Node, error) {

	return getNodeInfo(nodeSQLID, nodeID)
}

//GetNodeByKey 通过Key查询节点信息
func GetNodeByKey(nodeKey string) (*entity.Node, error) {
	return getNodeInfo(nodeSQLInstance, nodeKey)
}

//GetAvaliableNodeListFromNodeList 从待操作节点中获取关闭节点列表
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
	return true, strings.Join(idList, ","), nil
}

//BatchEditNodeGroup 批量修改节点分组
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

//BatchDeleteNode 批量修改接口分组
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

//UpdateAllNodeClusterID 更新节点集群ID
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
