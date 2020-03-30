package console_sqlite3

import (
	SQL "database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
)

//NodeGroupDao NodeGroupDao
type NodeGroupDao struct {
	db *SQL.DB
}

//NewNodeGroupDao new NodeGroupDao
func NewNodeGroupDao() *NodeGroupDao {
	return &NodeGroupDao{}
}

//Create create
func (d *NodeGroupDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.NodeGroupDao = d
	return &i, nil
}

//AddNodeGroup 新建节点分组
func (d *NodeGroupDao) AddNodeGroup(groupName string, clusterID int) (bool, interface{}, error) {
	db := d.db
	sql := "INSERT INTO goku_node_group (`groupName`,`clusterID`,`groupType`) VALUES (?,?,0);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	r, err := stmt.Exec(groupName, clusterID)
	if err != nil {
		return false, "[ERROR]Fail to insert data!", err
	}
	groupID, _ := r.LastInsertId()
	return true, groupID, nil
}

//EditNodeGroup 修改节点分组信息
func (d *NodeGroupDao) EditNodeGroup(groupName string, groupID int) (bool, string, error) {
	db := d.db
	sql := "UPDATE goku_node_group SET groupName = ? WHERE groupID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	_, err = stmt.Exec(groupName, groupID)
	if err != nil {
		return false, "[ERROR]Fail to update data!", err
	}
	return true, "", nil
}

//DeleteNodeGroup 删除节点分组
func (d *NodeGroupDao) DeleteNodeGroup(groupID int) (bool, string, error) {
	db := d.db
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_node_group WHERE groupID = ?;"
	_, err := Tx.Exec(sql, groupID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}
	sql = "DELETE FROM goku_node_info WHERE groupID = ?;"
	_, err = Tx.Exec(sql, groupID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}
	Tx.Commit()
	return true, "", nil
}

//GetNodeGroupInfo 获取节点分组信息
func (d *NodeGroupDao) GetNodeGroupInfo(groupID int) (bool, map[string]interface{}, error) {
	db := d.db

	sql := "SELECT G.`groupName`,C.`name` FROM goku_node_group G left join `goku_cluster` C ON C.`id` = G.`clusterID` WHERE G.`groupID` = ?;"
	var groupName string
	var clusterName string
	err := db.QueryRow(sql, groupID).Scan(&groupName, &clusterName)
	if err != nil {
		return false, nil, err
	}
	groupInfo := map[string]interface{}{
		"groupID":   groupID,
		"groupName": groupName,
		"cluster":   clusterName,
	}
	return true, groupInfo, nil
}

//GetNodeGroupList 获取节点分组列表
func (d *NodeGroupDao) GetNodeGroupList(clusterID int) (bool, []map[string]interface{}, error) {
	db := d.db
	sql := "SELECT G.`groupID`, G.groupName,C.`name` as cluster  FROM goku_node_group G left join `goku_cluster` C ON C.`id` = G.`clusterID` where G.`clusterID`=?;"
	rows, err := db.Query(sql, clusterID)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列

	nodeGroupList := make([]map[string]interface{}, 0)
	for rows.Next() {
		var groupID int
		var groupName string
		var clusterName string
		err = rows.Scan(&groupID, &groupName, &clusterName)
		if err != nil {
			return false, nil, err
		}
		groupInfo := map[string]interface{}{
			"groupID":   groupID,
			"groupName": groupName,
			"cluster":   clusterName,
		}
		nodeGroupList = append(nodeGroupList, groupInfo)
	}
	return true, nodeGroupList, nil
}

//CheckNodeGroupIsExist 检查节点分组是否存在
func (d *NodeGroupDao) CheckNodeGroupIsExist(groupID int) (bool, error) {
	db := d.db
	var id int
	sql := "SELECT groupID FROM goku_node_group WHERE groupID = ?;"
	err := db.QueryRow(sql, groupID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}

//GetRunningNodeCount 获取分组内启动节点数量
func (d *NodeGroupDao) GetRunningNodeCount(groupID int) (bool, interface{}, error) {
	db := d.db
	var count int
	sql := "SELECT COUNT(0) FROM goku_node_info WHERE groupID = ? AND nodeStatus = 1"
	err := db.QueryRow(sql, groupID).Scan(&count)
	if err != nil {
		return false, "[ERROR]Can not find the avaliable node", err
	}
	return true, count, nil
}
