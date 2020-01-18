package listener

import "sync"

//InterceptFunc 拦截函数
type InterceptFunc func(event interface{}) error

//Intercept 拦截器
type Intercept struct {
	callbacks []InterceptFunc
	locker    sync.RWMutex
}

//NewIntercept 创建拦截器
func NewIntercept() *Intercept {
	return &Intercept{
		callbacks: nil,
		locker:    sync.RWMutex{},
	}
}

//Add add
func (i *Intercept) Add(f func(v interface{}) error) {
	i.locker.Lock()
	i.callbacks = append(i.callbacks, InterceptFunc(f))
	i.locker.Unlock()
}

//Call call
func (i *Intercept) Call(v interface{}) error {
	i.locker.RLock()
	fs := i.callbacks
	i.locker.RUnlock()

	for _, f := range fs {
		err := f(v)
		if err != nil {

			return err
		}
	}
	return nil
}
