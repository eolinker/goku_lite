package updater

import (
	"github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/updater"
)

//Updater 更新器
type Updater interface {
	GetVersion() string
	UpdateVersion()
	Exec() error
}

//Factory factory
type Factory interface {
	Add(version string, updater Updater)
}

//Manager manager
type Manager struct {
	updaters []manager
}

type manager struct {
	version string
	updater Updater
}

var updateManager = &Manager{updaters: make([]manager, 0, 10)}

//Add add
func Add(version string, updater Updater) {
	updateManager.Add(version, updater)
}

//InitUpdater 执行版本更新操作
func InitUpdater() error {
	for _, u := range updateManager.updaters {
		err := u.updater.Exec()
		if err != nil {
			updater.SetGokuVersion(u.version)
			return err
		}
	}
	return nil
}

//Add add
func (u *Manager) Add(version string, updater Updater) {
	u.updaters = append(u.updaters, manager{version: version, updater: updater})
}
