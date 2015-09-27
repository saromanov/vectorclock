package vectorclock

import (
	"errors"
	"fmt"
	"sync"
	"sort"
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
func (vc *VectorClock) InitClocks() {
	for _, item := range vc.nodes {
		nodes := map[string]int{}
		for _, value := range vc.nodes {
			nodes[value] = 0
		}
		vc.clocks[item] = &Clock{nodes:nodes, mutex:&sync.Mutex{}}
	}
}

//Fit provides add new event
func (vc *VectorClock) Fit(title string) error {
	_, ok := vc.clocks[title]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not registred", title))
	}
	vc.clocks[title].Inc(title)
	return nil
}


//Send provides cause between node1 and node2
func (vc *VectorClock) Send(title1, title2 string) error {
	item1, ok := vc.clocks[title1]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not registred", title1))
	}

	item2, ok2 := vc.clocks[title2]
	if !ok2 {
		return errors.New(fmt.Sprintf("%s is not registred", title2))
	}
	vc.clocks[title2].Inc(title2)

	for key, _ := range vc.clocks {
		newvalue := item1.Get(key)
		if newvalue > item2.Get(key) {
			vc.clocks[title2].Set(key, newvalue)
		}
	}
	return nil
}

//GetState returns current state of clocks for registed nodes
func (vc *VectorClock) ShowState()map[string]*Clock {
	sort.Strings(vc.nodes)
	for _, key := range vc.nodes {
		fmt.Printf("Node %s\n", key)
		for _, value := range vc.nodes {
			fmt.Println(value, vc.clocks[key].Get(value))
		}
	}
	return vc.clocks
}


func (vc *VectorClock) GetClock(title string)(*Clock, error) {
	item, ok := vc.clocks[title]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s is not registred", title))
	}
	return item, nil
}
