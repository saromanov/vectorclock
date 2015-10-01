package vectorclock

import (
	"sync"
)

type Clock struct {
	Nodes map[string]int
	mutex *sync.Mutex
}

//Note: No need to check exist of elements

func (clock *Clock) Inc(title string) {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	clock.Nodes[title]++
}

func (clock *Clock) Set(title string, value int) {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	newvalue := clock.Nodes[title] + value
	clock.Nodes[title] = newvalue
}

func (clock *Clock) Get(title string) int {
	clock.mutex.Lock()
	defer clock.mutex.Unlock()
	return clock.Nodes[title]
}
