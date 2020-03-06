package database
const (
	//MysqlDriver mysql驱动器
	MysqlDriver = "mysql"
	//Sqlite3Driver Sqlite3Driver驱动
	Sqlite3Driver = "sqlite3"
)
//Config 数据库配置结构体
type Config interface {
	GetDriver() string
	GetSource() string
}
