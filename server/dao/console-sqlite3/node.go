package console_sqlite3

import (
	SQL "database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	v "github.com/eolinker/goku-api-gateway/common/version"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//NodeDao NodeDao
type NodeDao struct {
	db *SQL.DB
}

//NewNodeDao new NodeDao
func NewNodeDao() *NodeDao {
	return &NodeDao{}
}

//Create create
func (d *NodeDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.NodeDao = d
	return &i, nil
}

//AddNode 新增节点信息
func (d *NodeDao) AddNode(clusterID int, nodeName, nodeKey, listenAddress, adminAddress, gatewayPath string, groupID int) (int64, string, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO goku_node_info (`clusterID`,`nodeName`,`groupID`,`nodeKey`,`listenAddress`,`adminAddress`,`updateTime`,`createTime`,`version`, `gatewayPath`,`nodeStatus`) VALUES (?,?,?,?,?,?,?,?,?,?,0);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, "", "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	res, err := stmt.Exec(clusterID, nodeName, groupID, nodeKey, listenAddress, adminAddress, now, now, v.Version, gatewayPath)
	if err != nil {
		return 0, "", "[ERROR]Failed to insert data!", err
	}
	nodeID, err := res.LastInsertId()
	if err != nil {
		return 0, "", "[ERROR]Failed to insert data!", err
	}
	return nodeID, v.Version, "", nil
}

//EditNode 修改节点信息
func (d *NodeDao) EditNode(nodeName, listenAddress, adminAddress, gatewayPath string, nodeID, groupID int) (string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE goku_node_info SET  nodeName = ?,listenAddress = ?,adminAddress = ?,updateTime = ?,groupID = ?,gatewayPath = ? WHERE nodeID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeName, listenAddress, adminAddress, now, groupID, gatewayPath, nodeID)
	if err != nil {
		return "[ERROR]Failed to update data!", err
	}
	return "", nil
}

//DeleteNode 删除节点信息
func (d *NodeDao) DeleteNode(nodeID int) (string, error) {
	db := d.db
	sql := "DELETE FROM goku_node_info WHERE nodeID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return "[ERROR]Failed to delete data!", err
	}
	return "", nil
}

// GetNodeList 获取节点列表
func (d *NodeDao) GetNodeList(clusterID, groupID int, keyword string) ([]*entity.Node, error) {

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
		ruleStr += " WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprint(nodeSQLAll, ruleStr, " ORDER BY updateTime DESC;")

	return d.getNodeInfo(sql)

	//rows, err := db.Query(sql)
	//if err != nil {
	//	return  nil, err
	//}
	////延时关闭Rows
	//defer rows.Close()
	////获取记录列
	//nodeList := make([]*entity.Node, 0)
	//for rows.Next() {
	//	node := entity.Node{}
	//	err = rows.Scan(&node.NodeID, &node.NodeName, &node.NodeKey, &node.ListenAddress, &node.AdminAddress, &node.UpdateTime, &node.CreateTime, &node.Version, &node.GatewayPath, &node.GroupID, &node.GroupName)
	//	if err != nil {
	//		return  nil, err
	//	}
	//	if node.Version == v.Version {
	//		// 判断节点版本号是否是最新
	//		node.IsUpdate = true
	//	}
	//	nodeList = append(nodeList, &node)
	//}
	//return nodeList, nil
}

const nodeSQLAll = "SELECT A.`nodeID` , A.`nodeName` , A.`listenAddress` , A.`adminAddress` , A.`nodeKey` , A.`updateTime` , A.`createTime` , A.`version` , A.`gatewayPath` , A.`groupID` , IFNULL(G.`groupName` , '未分类') , C.`name`As cluster , C.`title` As cluster_title FROM goku_node_info A LEFT JOIN goku_node_group G ON A.`groupID` = G.`groupID` LEFT JOIN `goku_cluster` C ON A.`clusterID`=C.`id`"
const nodeSQLID = nodeSQLAll + " WHERE A.`nodeID` = ? ;"
const nodeSQLInstance = nodeSQLAll + " WHERE A.`nodeKey` = ? ;"

func (d *NodeDao) getNodeInfo(sql string, args ...interface{}) ([]*entity.Node, error) {

	db := d.db

	rows, e := db.Query(sql, args...)
	if e != nil {
		return nil, e
	}
	nodes := make([]*entity.Node, 0, 10)
	for rows.Next() {
		node := &entity.Node{}
		err := rows.Scan(&node.NodeID,
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
		nodes = append(nodes, node)
	}
	return nodes, nil
}

//GetNodeInfoAll get all node
func (d *NodeDao) GetNodeInfoAll() ([]*entity.Node, error) {
	nodes, e := d.getNodeInfo(nodeSQLAll)
	if e != nil {
		return nil, e
	}

	return nodes, nil

}

//GetNodeInfo 获取节点信息
func (d *NodeDao) GetNodeInfo(nodeID int) (*entity.Node, error) {
	nodes, e := d.getNodeInfo(nodeSQLID, nodeID)
	if e != nil {
		return nil, e
	}
	if len(nodes) > 0 {
		return nodes[0], nil
	}
	return nil, fmt.Errorf("not exit node width noddID:%d", nodeID)
}

//GetNodeByKey 通过Key查询节点信息
func (d *NodeDao) GetNodeByKey(nodeKey string) (*entity.Node, error) {
	nodes, e := d.getNodeInfo(nodeSQLInstance, nodeKey)
	if e != nil {
		return nil, e
	}
	if len(nodes) > 0 {
		return nodes[0], nil
	}
	return nil, fmt.Errorf("not exit node width nodeKey:%s", nodeKey)
}

//GetAvaliableNodeListFromNodeList 从待操作节点中获取关闭节点列表
func (d *NodeDao) GetAvaliableNodeListFromNodeList(nodeIDList string, nodeStatus int) (string, error) {
	db := d.db
	sql := "SELECT nodeID FROM goku_node_info WHERE nodeID IN (" + nodeIDList + ") AND nodeStatus = ?"
	rows, err := db.Query(sql, nodeStatus)
	if err != nil {
		return "[ERROR]Fail to excute SQL statement!", err
	}
	defer rows.Close()
	idList := make([]string, 0)
	for rows.Next() {
		var nodeID int
		err = rows.Scan(&nodeID)
		if err != nil {
			return err.Error(), err
		}
		idList = append(idList, strconv.Itoa(nodeID))
	}
	return strings.Join(idList, ","), nil
}

//BatchEditNodeGroup 批量修改节点分组
func (d *NodeDao) BatchEditNodeGroup(nodeIDList string, groupID int) (string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "UPDATE goku_node_info SET groupID = ?,updateTime = ? WHERE nodeID IN (" + nodeIDList + ");"
	_, err := Tx.Exec(sql, groupID, now)
	if err != nil {
		Tx.Rollback()
		return "[ERROR]Fail to excute SQL statement!", err
	}
	Tx.Commit()
	return "", nil
}

//BatchDeleteNode 批量修改接口分组
func (d *NodeDao) BatchDeleteNode(nodeIDList string) (string, error) {
	db := d.db
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_node_info WHERE nodeID IN (" + nodeIDList + ");"
	_, err := Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return "[ERROR]Fail to excute SQL statement!", err
	}
	Tx.Commit()
	return "", nil
}

//UpdateAllNodeClusterID 更新节点集群ID
func (d *NodeDao) UpdateAllNodeClusterID(clusterID int) {
	db := d.db
	Tx, _ := db.Begin()
	sql := "UPDATE goku_node_info SET clusterID = ?;"
	_, err := Tx.Exec(sql, clusterID)
	if err != nil {
		Tx.Rollback()
		return
	}
	sql = "UPDATE goku_node_group SET clusterID = ?;"
	_, err = Tx.Exec(sql, clusterID)
	if err != nil {
		Tx.Rollback()
		return
	}
	Tx.Commit()
}

//GetHeartBeatTime 获取节点心跳时间
func (d *NodeDao) GetHeartBeatTime(nodeKey string) (time.Time, error) {
	db := d.db
	heartBeat := ""

	sql := "SELECT heartBeatTime FROM goku_node_info WHERE nodeKey = ?"
	err := db.QueryRow(sql, nodeKey).Scan(&heartBeat)
	if err != nil {
		return time.Time{}, err
	}
	heartBeatTime, err := time.ParseInLocation("2006-01-02 15:04:05", heartBeat, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return heartBeatTime, nil
}

//SetHeartBeatTime 设置节点心跳时间
func (d *NodeDao) SetHeartBeatTime(nodeKey string, heartBeatTime time.Time) error {
	db := d.db
	heartBeat := heartBeatTime.Format("2006-01-02 15:04:05")

	sql := "UPDATE goku_node_info SET heartBeatTime = ? WHERE nodeKey = ?"
	_, err := db.Exec(sql, heartBeat, nodeKey)
	if err != nil {
		return err
	}
	return nil
}
