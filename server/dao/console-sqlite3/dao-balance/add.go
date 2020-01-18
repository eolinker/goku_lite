package dao_balance

//AddStatic 新增静态负载
func (b *BalanceDao) AddStatic(name, serviceName, static, staticCluster, desc, now string) (string, error) {

	const sql = "INSERT INTO goku_balance (`balanceName`,`serviceName`,`static`,`staticCluster`,`balanceDesc`,`createTime`,`updateTime`,`appName`,`defaultConfig`,`clusterConfig`,`balanceConfig`) VALUES (?,?,?,?,?,?,?,'','','','');"

	db := b.db
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, serviceName, static, staticCluster, desc, now, now)
	if err != nil {
		return "[ERROR]Failed to add data!", err
	}
	return "", nil
}

//AddDiscovery 新增服务发现
func (b *BalanceDao) AddDiscovery(name, serviceName, appName, desc, now string) (string, error) {

	const sql = "INSERT INTO goku_balance (`balanceName`,`serviceName`,`appName`,`balanceDesc`,`createTime`,`updateTime`,`static`,`staticCluster`,`defaultConfig`,`clusterConfig`,`balanceConfig`) VALUES (?,?,?,?,?,?,'','','','','');"

	db := b.db
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, serviceName, appName, desc, now, now)
	if err != nil {
		return "[ERROR]Failed to add data!", err
	}
	return "", nil
}

//SaveStatic 保存静态负载信息
func (b *BalanceDao) SaveStatic(name, serviceName, static, staticCluster, desc string, now string) (string, error) {
	const sql = "UPDATE `goku_balance` SET `serviceName`=? ,`static` = ?,`staticCluster`=?,`balanceDesc` =?,`updateTime`=? WHERE `balanceName`=?;"
	db := b.db
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(serviceName, static, staticCluster, desc, now, name)
	if err != nil {
		return "[ERROR]Failed to add data!", err
	}
	return "", nil
}

//SaveDiscover 保存服务发现信息
func (b *BalanceDao) SaveDiscover(name, serviceName, appName, desc string, now string) (string, error) {
	const sql = "UPDATE `goku_balance` SET `serviceName`=? ,`appName` = ?,`balanceDesc` =?,`updateTime`=? WHERE `balanceName`=?;"
	db := b.db
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(serviceName, appName, desc, now, name)
	if err != nil {
		return "[ERROR]Failed to add data!", err
	}
	return "", nil
}

//Delete 删除负载
func (b *BalanceDao) Delete(name string) (string, error) {
	const sql = "DELETE FROM `goku_balance` WHERE  `balanceName`= ?;"
	db := b.db
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return "[ERROR]DELETE fail", err
	}
	return "", nil
}

//BatchDelete 批量删除负载
func (b *BalanceDao) BatchDelete(balanceNames []string) (string, error) {
	db := b.db
	sql := "DELETE FROM `goku_balance` WHERE  `balanceName` = ?;"
	sql2 := "UPDATE goku_conn_strategy_api SET target = '' WHERE target = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	stmt2, err := db.Prepare(sql2)
	if err != nil {
		return "[ERROR]Illegal SQL statement!", err
	}
	defer stmt2.Close()
	for _, balanceName := range balanceNames {
		stmt.Exec(balanceName)
		stmt2.Exec(balanceName)
	}
	return "", nil
}
