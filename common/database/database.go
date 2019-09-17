package database

type Config interface {
	GetDriver() string
	GetSource() string
}
