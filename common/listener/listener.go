package listener

import "sync"

//CallbackFunc 回调函数
type CallbackFunc func(event interface{})

//Listener
type Listener struct {
	callbacks     []CallbackFunc
	callbacksOnce []CallbackFunc
	locker        sync.Mutex
}

//New 创建监听
func New() *Listener {
	return &Listener{
		callbacks:     nil,
		callbacksOnce: nil,
		locker:        sync.Mutex{},
	}
}

//ListenOnce listenOnce
func (l *Listener) ListenOnce(callbackFunc CallbackFunc) {
	l.locker.Lock()
	l.callbacksOnce = append(l.callbacksOnce, callbackFunc)
	l.locker.Unlock()
}

//Listen listen
func (l *Listener) Listen(callbackFunc CallbackFunc) {
	l.locker.Lock()
	l.callbacks = append(l.callbacks, callbackFunc)
	l.locker.Unlock()
}

//Call call
func (l *Listener) Call(event interface{}) {
	l.locker.Lock()
	cbs := l.callbacks
	cbsO := l.callbacksOnce
	l.callbacksOnce = nil
	l.locker.Unlock()

	for _, cb := range cbs {
		cb(event)
	}
	for _, cb := range cbsO {
		cb(event)
	}
}
