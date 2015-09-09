package vectorclock

type Clock struct {
	nodes map[string]int
}

//Note: No need to check exist of elements

func (clock *Clock) Set(title string) {
	clock[title]++
}

func (clock *Clock) Get(title string) int {
	return clock[title]
}