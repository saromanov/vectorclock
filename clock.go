package vectorclock

import (
	"sync"
)

type Clock struct {
	nodes map[string]int
	mutex *sync.Mutex
}

//Note: No need to check exist of elements

func (clock *Clock) Set(title string) {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	clock.nodes[title]++
}

func (clock *Clock) Get(title string) int {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	return clock.nodes[title]
}
