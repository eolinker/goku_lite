package updater

//GetTableVersion 获取当前表版本号
func (d *Dao)GetTableVersion(name string) string {
	db := d.db
	version := ""
	sql := "SELECT version FROM goku_table_version WHERE tableName = ?"
	err := db.QueryRow(sql, name).Scan(&version)
	if err != nil {
		return ""
	}
	return version
}

//UpdateTableVersion 更新表版本号
func (d *Dao)UpdateTableVersion(name, version string) error {
	db := d.db
	sql := "REPLACE INTO goku_table_version (tableName,version) VALUES (?,?);"
	_, err := db.Exec(sql, name, version)
	if err != nil {
		return err
	}
	return nil
}

//GetGokuVersion 获取goku当前版本号
func (d *Dao)GetGokuVersion() string {
	db := d.db
	version := ""
	sql := "SELECT version FROM goku_version;"
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		return ""
	}
	return version
}

//SetGokuVersion 设置goku版本号
func (d *Dao)SetGokuVersion(version string) error {
	db := d.db
	sql := "REPLACE INTO goku_version (sol,version) VALUES (?,?);"
	_, err := db.Exec(sql, 1, version)
	if err != nil {
		return err
	}
	return nil
}
