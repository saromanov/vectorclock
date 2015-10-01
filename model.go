package vectorclock

import
(
	"errors"
	"fmt"
	"sort"
)

type Model struct {
	Nodes  []string
	Clocks map[string]*Clock
}

func newmodel(clock *VectorClock)*Model {
	model := new(Model)
	model.Clocks = clock.clocks
	model.Nodes = clock.nodes
	return model
}

//Fit provides add new event
func (vc *Model) Fit(title string) error {
	_, ok := vc.Clocks[title]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not registred", title))
	}
	vc.Clocks[title].Inc(title)
	return nil
}


//Send provides cause between node1 and node2
func (vc *Model) Send(title1, title2 string) error {
	item1, ok := vc.Clocks[title1]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not registred", title1))
	}

	item2, ok2 := vc.Clocks[title2]
	if !ok2 {
		return errors.New(fmt.Sprintf("%s is not registred", title2))
	}

	//Increase event on the current clock
	vc.Clocks[title2].Inc(title2)

	//Increase counter on all clocks if the new value > old value on clock
	for key, _ := range vc.Clocks {
		newvalue := item1.Get(key)
		if newvalue > item2.Get(key) {
			vc.Clocks[title2].Set(key, newvalue)
		}
	}
	return nil
}

//GetState returns current state of Clocks for registed Nodes
func (vc *Model) ShowState()map[string]*Clock {
	sort.Strings(vc.Nodes)
	for _, key := range vc.Nodes {
		fmt.Printf("Node %s\n", key)
		for _, value := range vc.Nodes {
			fmt.Println(value, vc.Clocks[key].Get(value))
		}
	}
	return vc.Clocks
}


func (vc *Model) GetClock(title string)(*Clock, error) {
	item, ok := vc.Clocks[title]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s is not registred", title))
	}
	return item, nil
}