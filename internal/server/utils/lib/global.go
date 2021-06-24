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

func (this *MyMap) Put(key uint, value int) {
	this.Lock()
	defer this.Unlock()

	if this.v == nil {
		this.v = map[uint]int{}
	}

	this.v[key] = value
}
func (this *MyMap) Get(key uint) int {
	this.RLock()
	defer this.RUnlock()

	if this.v == nil {
		this.v = map[uint]int{}
	}

	return this.v[key]
}
