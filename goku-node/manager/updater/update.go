package updater

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	gateway_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/gateway-manager"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var (
	isInit       = false
	handlers     = make(map[string]*updateHandler)
	ch           = make(chan bool, 1)
	periodCh     = make(chan int, 1)
	locker       sync.Mutex
	handlerExecs = make([]*updateHandlerExec, 0)
)

type UpdateHandleFunc func()

func init() {
	Add(func() {
		gateway_manager.LoadGatewayConfig()
		updatePeriod(gateway_manager.GetUpdatePeriod())
	}, 2, "goku_gateway")
}

type updateHandler struct {
	tables         []string
	UpdateHandlers []UpdateHandleFunc
	last           time.Time
}

func updatePeriod(period int) {
	periodCh <- period
}

func Add(handler UpdateHandleFunc, priority int, tables ...string) {
	sort.Strings(tables)
	key := strings.Join(tables, ":")
	locker.Lock()
	defer locker.Unlock()
	hfs, has := handlers[key]
	if !has {
		hfs = &updateHandler{
			tables: tables,
		}
	}
	hfs.UpdateHandlers = append(hfs.UpdateHandlers, handler)
	if !has {
		handlerExec := &updateHandlerExec{
			name:          key,
			priority:      priority,
			updateHandler: hfs,
		}
		handlerExecs = append(handlerExecs, handlerExec)
	}
	handlers[key] = hfs
}

func Update() {
	ch <- true
}

func InitUpdate() {
	locker.Lock()
	defer locker.Unlock()
	if isInit {
		return
	}
	isInit = true
	// 排序
	log.Debug("update sort handler")
	sort.Sort(handlerSlice(handlerExecs))
	log.Debug("update sort handler done")
	doFirst()
	go doUpdate(gateway_manager.GetUpdatePeriod())
}
func doFirst() {
	log.Debug("update doFirst")
	handlerExecsTmp := handlerExecs
	for _, handler := range handlerExecsTmp {

		if handler != nil {
			log.Debug("update ", handler.name, handler.tables)
			handler.last = time.Now()
			for _, fn := range handler.UpdateHandlers {
				fn()
			}
			log.Debug("update done")
		}
	}
}
func doUpdate(sec int) {
	fmt.Println(sec)
	period := time.Duration(sec) * time.Second
	t := time.NewTimer(period)
	for {
		select {
		case <-t.C:
		case <-ch:
		case p := <-periodCh:
			{
				period = time.Duration(p) * time.Second
				continue
			}
		}
		updates()
		t.Reset(period)
	}
}

func updates() {
	handlerExecsTmp := handlerExecs
	for _, handler := range handlerExecsTmp {
		if handler != nil {
			if u, nt := CheckUpdate(handler.last, handler.tables...); u {
				handler.last = nt
				for _, fn := range handler.UpdateHandlers {
					fn()
				}
			}
		}
	}
}

func CheckUpdate(last time.Time, tables ...string) (bool, time.Time) {
	t, err := dao.GetLastUpdateOfApi(tables...)
	if err != nil {
		return false, last
	}
	if t.After(last) {
		return true, t
	}
	return false, last
}
