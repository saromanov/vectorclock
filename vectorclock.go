package vectorclock

import (
	"sync"
)

type VectorClock struct {
	nodes  []string
	clocks map[string]*Clock
}

//New provides creating a the new object of VectorClock
func New() *VectorClock {
	vc := new(VectorClock)
	vc.nodes = []string{}
	vc.clocks = map[string]*Clock{}
	return vc
}

//Add provides append new item for clocks
func (vc *VectorClock) Add(title string) {
	vc.nodes = append(vc.nodes, title)
}

//InitClocks provides basic initialization for clocks
func (vc *VectorClock) InitClocks()*Model {
	for _, item := range vc.nodes {
		nodes := map[string]int{}
		for _, value := range vc.nodes {
			nodes[value] = 0
		}
		vc.clocks[item] = &Clock{Nodes:nodes, mutex:&sync.Mutex{}}
	}

	return newmodel(vc)
}

