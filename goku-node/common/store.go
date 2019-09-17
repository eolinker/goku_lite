package common

import goku_plugin "github.com/eolinker/goku-plugin"

type Store struct {
	value interface{}
}

func (s *Store) Set(value interface{}) {
	s.value = value
}

func (s *Store) Get() (value interface{}) {
	return s.value
}

type StoreHandler struct {
	Cache             map[string]interface{}
	Stores            map[string]goku_plugin.Store
	CurrentPluginName string
}

func (s *StoreHandler) SetPlugin(name string) {
	s.CurrentPluginName = name
}

func (s *StoreHandler) SetCache(name string, value interface{}) {
	if s.Cache == nil {
		s.Cache = make(map[string]interface{})
	}
	s.Cache[name] = value
}

func (s *StoreHandler) GetCache(name string) (value interface{}, has bool) {
	if s.Cache == nil {
		return nil, false
	}
	value, has = s.Cache[name]
	return
}

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

func NewStoreHandler() *StoreHandler {
	return new(StoreHandler)
}
