package vectorclock

import
(
	"sync"
)

type Clock struct {
	nodes map[string]int
	mutex *sync.Mutex
}

//Note: No need to check exist of elements

func (clock *Clock) Set(title string) {
	clock[title]++
}

func (clock *Clock) Get(title string) int {
	return clock[title]
}