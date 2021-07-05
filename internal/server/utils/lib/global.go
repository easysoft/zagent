package serverUitls

import (
	"sync"
)

var (
	Map = MyMap{}
)

type MyMap struct {
	v map[uint]int
	sync.RWMutex
}

func (m *MyMap) Put(key uint, value int) {
	m.Lock()
	defer m.Unlock()

	if m.v == nil {
		m.v = map[uint]int{}
	}

	m.v[key] = value
}
func (m *MyMap) Get(key uint) int {
	m.RLock()
	defer m.RUnlock()

	if m.v == nil {
		m.v = map[uint]int{}
	}

	return m.v[key]
}
