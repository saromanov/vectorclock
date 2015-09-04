package vectorclock

type VectorClock struct {
	nodes []string
}

func New()*VectorClock {
	vc := new(VectorClock)
	vc.nodes = []string{}
	return vc
}

func (vc *VectorClock) Add(title string) {
	vc.nodes = append(vc.nodes, title)
}