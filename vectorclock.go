package vectorclock

import
(
	"errors"
	"fmt"
)


type Clock struct {
	nodes map[string]int
}

type VectorClock struct {
	nodes []string
	clocks map[string] *Clock
}

//New provides creating a the new object of VectorClock
func New()*VectorClock {
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
func (vc *VectorClock) InitClocks() {
	for _, item := range vc.nodes {
		nodes := map[string]int{}
		for _, value := range vc.nodes {
			nodes[value] = 0
		}
		vc.clocks[item] = &Clock{nodes}
	}
}


//Send provides cause between node1 and node2
func (vc *VectorClock) Send(title1, title2 string)error {
	clock1, ok := vc.clocks[title1]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not registred", title1))
	}

	clock2, ok2 := vc.clocks[title2]
	if !ok2 {
		return errors.New(fmt.Sprintf("%s is not registred", title2))
	}
	return nil 
}
