package vectorclock

import (
	"sync"
)

type Clock struct {
	nodes map[string]int
	mutex *sync.Mutex
}

//Note: No need to check exist of elements

func (clock *Clock) Inc(title string) {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	clock.nodes[title]++
}

func (clock *Clock) Set(title string, value int) {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	newvalue := clock.nodes[title] + value
	clock.nodes[title] = newvalue
}

func (clock *Clock) Get(title string) int {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	return clock.nodes[title]
}
