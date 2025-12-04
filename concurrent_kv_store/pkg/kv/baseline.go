package kv

import "sync"

type BaselineMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewBaselineMap() *BaselineMap {
	return &BaselineMap{
		data: make(map[string]string),
	}
}

func (m *BaselineMap) Put(key, value string) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

func (m *BaselineMap) Get(key string) (string, bool) {
	m.mu.RLock()
	val, ok := m.data[key]
	m.mu.RUnlock()
	return val, ok
}

func (m *BaselineMap) Delete(key string) {
	m.mu.Lock()
	delete(m.data, key)
	m.mu.Unlock()
}
