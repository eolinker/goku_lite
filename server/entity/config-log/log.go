package entity

//LogConfig 日志配置
type LogConfig struct {
	Name   string
	Enable int
	Dir    string
	File   string
	Level  string
	Period string
	Expire int
	Fields string
}
