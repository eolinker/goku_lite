package pdao

import "database/sql"

//FactoryFunc factoryFunc
type FactoryFunc func(db *sql.DB) (interface{}, error)

//Create create
func (f FactoryFunc) Create(db *sql.DB) (interface{}, error) {
	return f(db)
}

//Factory factory
type Factory interface {
	Create(db *sql.DB) (interface{}, error)
}

//DBBuilder dbBuilder
type DBBuilder interface {
	Build(db *sql.DB) error
}
