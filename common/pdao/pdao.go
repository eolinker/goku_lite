package pdao

import (
	"database/sql"
	"sync"
)

var (
	needsManager  = NewNeedsManager()
	factoryManage = NewFactoryManager()
)

//Need need
func Need(is ...interface{}) {
	for _, i := range is {
		needsManager.Need(i)
	}
}

//Check check
func Check() error {
	return needsManager.Check()
}

//RegisterDao 注册dao类
func RegisterDao(driver string, factories ...Factory) {
	for _, factory := range factories {
		factoryManage.RegisterDao(driver, factory)
	}

}

//RegisterDBBuilder 注册dbBuilder
func RegisterDBBuilder(driver string, builders ...DBBuilder) {
	for _, builder := range builders {
		factoryManage.RegisterDBBuilder(driver, builder)

	}
}

//Build build
func Build(driver string, db *sql.DB) error {
	return factoryManage.Build(driver, db, needsManager)
}

//FactoryManager 工厂管理者
type FactoryManager struct {
	factories map[string][]Factory
	builders  map[string][]DBBuilder
	locker    sync.Mutex
}

//NewFactoryManager new工厂管理者
func NewFactoryManager() *FactoryManager {
	return &FactoryManager{
		factories: make(map[string][]Factory),
		builders:  make(map[string][]DBBuilder),
		locker:    sync.Mutex{},
	}
}

//RegisterDBBuilder dbBuilder注册器
func (f *FactoryManager) RegisterDBBuilder(driver string, builder DBBuilder) {
	f.locker.Lock()
	f.builders[driver] = append(f.builders[driver], builder)
	f.locker.Unlock()
}

//RegisterDao dao类注册器
func (f *FactoryManager) RegisterDao(driver string, factory Factory) {
	f.locker.Lock()
	f.factories[driver] = append(f.factories[driver], factory)
	f.locker.Unlock()
}

func (f *FactoryManager) get(driver string) []Factory {
	f.locker.Lock()
	fs := f.factories[driver]
	delete(f.factories, driver)
	f.locker.Unlock()
	return fs
}

func (f *FactoryManager) callBuild(driver string, db *sql.DB) error {
	f.locker.Lock()

	bs := f.builders[driver]
	for _, b := range bs {
		err := b.Build(db)
		if err != nil {
			f.locker.Unlock()
			return err
		}
	}
	f.locker.Unlock()
	return nil
}

//Build build
func (f *FactoryManager) Build(driver string, db *sql.DB, m *NeedsManager) error {
	err := f.callBuild(driver, db)
	if err != nil {
		return err
	}
	fs := f.get(driver)
	for _, factory := range fs {

		i, err := factory.Create(db)
		if err != nil {
			return err
		}
		m.Set(i)
	}
	return nil

}
