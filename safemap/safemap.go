package safemap

import "sync"

type SafeMap struct {
	sync.RWMutex
	data map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	m := new(SafeMap)
	m.data = make(map[interface{}]interface{})
	return m
}

func (m *SafeMap) Set(k, v interface{}) {
	m.Lock()
	m.data[k] = v
	m.Unlock()
}

func (m *SafeMap) Get(k interface{}) (interface{}, bool) {
	m.RLock()
	v, ok := m.data[k]
	m.RUnlock()
	return v, ok
}

func (m *SafeMap) Delete(k interface{}) {
	m.Lock()
	delete(m.data, k)
	m.Unlock()
}
