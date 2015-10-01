package vectorclock

import 
(
	"testing"
)

func constructClocks()*Model {
	clock := New()
	clock.Add("node2")
	clock.Add("node1")
	clock.Add("node3")
	clock.Add("node4")
	clock.Add("node5")
	return clock.InitClocks()
}

func TestFit(t *testing.T) {
	items := constructClocks()
	items.Fit("node1")
	node, err := items.GetClock("node1")
	if err != nil {
		t.Error(err)
	}
	result:= node.Get("node1")

	if result != 1 {
		t.Errorf("Expected %d, found %d", 1, result)
	}
}

func TestSend(t *testing.T) {
	items := constructClocks()
	items.Fit("node3")
	items.Send("node3", "node2")
	node, err := items.GetClock("node2")
	if err != nil {
		t.Error(err)
	}
	result:= node.Get("node2")

	if result != 1 {
		t.Errorf("Expected %d, found %d", 1, result)
	}
}