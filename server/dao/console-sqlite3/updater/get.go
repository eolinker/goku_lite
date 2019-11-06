package updater

import "github.com/eolinker/goku-api-gateway/common/database"

//GetTableVersion 获取当前表版本号
func GetTableVersion(name string) string {
	db := database.GetConnection()
	version := ""
	sql := "SELECT version FROM goku_table_version WHERE tableName = ?"
	err := db.QueryRow(sql, name).Scan(&version)
	if err != nil {
		return ""
	}
	return version
}

//UpdateTableVersion 更新表版本号
func UpdateTableVersion(name, version string) error {
	db := database.GetConnection()
	sql := "REPLACE INTO goku_table_version (tableName,version) VALUES (?,?);"
	_, err := db.Exec(sql, name, version)
	if err != nil {
		return err
	}
	return nil
}

//GetGokuVersion 获取goku当前版本号
func GetGokuVersion() string {
	db := database.GetConnection()
	version := ""
	sql := "SELECT version FROM goku_version;"
	err := db.QueryRow(sql).Scan(&version)
	if err != nil {
		return ""
	}
	return version
}

//SetGokuVersion 设置goku版本号
func SetGokuVersion(version string) error {
	db := database.GetConnection()
	sql := "REPLACE INTO goku_version (sol,version) VALUES (?,?);"
	_, err := db.Exec(sql, 1, version)
	if err != nil {
		return err
	}
	return nil
}
