package Heap

import "testing"

var hp *Heap

// hp = InitHeap(10)

func TestInitHeap(t *testing.T) {
	hp = InitHeap(10)
	if hp.Capacity != 10 {
		t.Error("Length initilizing error")
	}
}

func TestInsert(t *testing.T) {
	hp.Insert(10)
	hp.Insert(9)
	hp.Insert(8)
	hp.Insert(1)
	if hp.GetMin() != 1 {
		t.Error("Problem on GetMin")
	}

}
