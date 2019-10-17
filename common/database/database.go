package database

//Config 数据库配置结构体
type Config interface {
	GetDriver() string
	GetSource() string
}
