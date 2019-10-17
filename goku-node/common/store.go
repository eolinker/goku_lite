package common

import goku_plugin "github.com/eolinker/goku-plugin"

//Store 存储
type Store struct {
	value interface{}
}

//Set set
func (s *Store) Set(value interface{}) {
	s.value = value
}

//Get get
func (s *Store) Get() (value interface{}) {
	return s.value
}

//StoreHandler 存储器
type StoreHandler struct {
	Cache             map[string]interface{}
	Stores            map[string]goku_plugin.Store
	CurrentPluginName string
}

//SetPlugin 设置plugin
func (s *StoreHandler) SetPlugin(name string) {
	s.CurrentPluginName = name
}

//SetCache 设置缓存
func (s *StoreHandler) SetCache(name string, value interface{}) {
	if s.Cache == nil {
		s.Cache = make(map[string]interface{})
	}
	s.Cache[name] = value
}

//GetCache 获取缓存
func (s *StoreHandler) GetCache(name string) (value interface{}, has bool) {
	if s.Cache == nil {
		return nil, false
	}
	value, has = s.Cache[name]
	return
}

//Store 存储器
func (s *StoreHandler) Store() goku_plugin.Store {
	if s.Stores == nil {
		s.Stores = make(map[string]goku_plugin.Store)
	}
	store, has := s.Stores[s.CurrentPluginName]
	if !has {
		store = &Store{}
		s.Stores[s.CurrentPluginName] = store
	}

	return store
}

//NewStoreHandler 储存器
func NewStoreHandler() *StoreHandler {
	return new(StoreHandler)
}
