package telegraph

import (
	"context"
	"errors"
	"sync"
)
var (
	ErrorContextDone = errors.New("context done")
)
//Telegraph telegraph
type Telegraph struct {
	value   interface{}
	version string
	locker  sync.RWMutex
	c       chan struct{}
}

//NewTelegraph 创建telegraph
func NewTelegraph(version string, value interface{}) *Telegraph {
	return &Telegraph{
		value:   value,
		version: version,
		locker:  sync.RWMutex{},
		c:       make(chan struct{}),
	}
}

//Set set
func (t *Telegraph) Set(version string, value interface{}) {

	t.locker.Lock()
	close(t.c)
	t.version = version
	t.value = value

	t.c = make(chan struct{})

	t.locker.Unlock()

}

func (t *Telegraph) get() (string, <-chan struct{}, interface{}) {

	t.locker.RLock()
	version, c, value := t.version, t.c, t.value
	t.locker.RUnlock()

	return version, c, value
}

//Get get
func (t *Telegraph) Get(version string) (interface{},error) {
	return t.GetWidthContext(context.Background(), version)
}

//Close close
func (t *Telegraph) Close() {
	t.locker.Lock()
	close(t.c)
	t.version = ""
	t.locker.Unlock()
}

//GetWidthContext 获取上下文
func (t *Telegraph) GetWidthContext(ctx context.Context, version string) (interface{} ,error){
	v, c, value := t.get()
	if v == "" {
		// closed
		return nil,nil
	}
	if version != v {
		return value,nil
	}

	select {
	case <-c:
		return t.GetWidthContext(ctx, version)
	case <-ctx.Done():
		return nil,ErrorContextDone
	}

}
