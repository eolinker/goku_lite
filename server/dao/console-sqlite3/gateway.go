package console_sqlite3

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//GatewayDao GatewayDao
type GatewayDao struct {
	db *sql.DB
}

//NewGatewayDao new GatewayDao
func NewGatewayDao() *GatewayDao {
	return &GatewayDao{}
}

//Create create
func (d *GatewayDao) Create(db *sql.DB) (interface{}, error) {

	d.db = db

	var i dao.GatewayDao = d

	return &i, nil
}

//GetGatewayConfig 获取网关配置
func (d *GatewayDao) GetGatewayConfig() (map[string]interface{}, error) {
	db := d.db
	var successCode string
	var nodeUpdatePeriod, monitorUpdatePeriod, monitorTimeout int
	sql := `SELECT successCode,nodeUpdatePeriod,monitorUpdatePeriod,monitorTimeout FROM goku_gateway WHERE id = 1;`
	err := db.QueryRow(sql).Scan(&successCode, &nodeUpdatePeriod, &monitorUpdatePeriod, &monitorTimeout)
	if err != nil {
		return nil, err
	}
	gatewayConfig := map[string]interface{}{
		"successCode":         successCode,
		"nodeUpdatePeriod":    nodeUpdatePeriod,
		"monitorUpdatePeriod": monitorUpdatePeriod,
		"monitorTimeout":      monitorTimeout,
	}
	return gatewayConfig, nil
}

//EditGatewayBaseConfig 编辑网关基本配置
func (d *GatewayDao) EditGatewayBaseConfig(config entity.GatewayBasicConfig) (bool, string, error) {
	db := d.db
	sql := "SELECT successCode FROM goku_gateway WHERE id = 1;"
	code := ""
	err := db.QueryRow(sql).Scan(&code)
	if err != nil {
		sql = "INSERT INTO goku_gateway (id,successCode,nodeUpdatePeriod,monitorUpdatePeriod,monitorTimeout) VALUES (1,?,?,?,?)"
	} else {
		sql = "UPDATE goku_gateway SET successCode = ?,nodeUpdatePeriod = ?,monitorUpdatePeriod = ?,monitorTimeout = ? WHERE id = 1;"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(config.SuccessCode, config.NodeUpdatePeriod, config.MonitorUpdatePeriod, config.MonitorTimeout)
	if err != nil {
		return false, "[ERROR]Fail to excute SQL Statement!", err
	}
	return true, "", nil
}

//GetGatewayInfo 获取网关信息
func (d *GatewayDao) GetGatewayInfo() (nodeStartCount, nodeStopCount, projectCount, apiCount, strategyCount int, err error) {
	db := d.db
	// 获取节点启动数量

	err = db.QueryRow("SELECT COUNT(0) FROM goku_node_info WHERE nodeStatus = 1;").Scan(&nodeStartCount)
	if err != nil {
		return
	}

	// 获取节点关闭数量

	err = db.QueryRow("SELECT COUNT(0) FROM goku_node_info WHERE nodeStatus = 0;").Scan(&nodeStopCount)
	if err != nil {
		return
	}
	// 获取项目数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_project;").Scan(&projectCount)
	if err != nil {
		return
	}

	// 获取api数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_api;").Scan(&apiCount)
	if err != nil {
		return
	}

	// 获取策略数量
	err = db.QueryRow("SELECT COUNT(0) FROM goku_gateway_strategy;").Scan(&strategyCount)
	if err != nil {
		return
	}
	return
}
