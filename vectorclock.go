package vectorclock

type VectorClock struct {
	nodes []string
}

//New provides creating a the new object of VectorClock
func New()*VectorClock {
	vc := new(VectorClock)
	vc.nodes = []string{}
	return vc
}

//Add provides append new item for clocks
func (vc *VectorClock) Add(title string) {
	vc.nodes = append(vc.nodes, title)
}