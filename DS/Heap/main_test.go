package Heap

import "testing"

var hp *Heap

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

	if hp.CurrentSize != 4 {
		t.Error("Problem calc size")
	}
}

func TestExtractMin(t *testing.T) {
	v := hp.GetMin()
	cl := hp.CurrentSize
	if v != hp.ExtractMin() {
		t.Error("Error on extract min")
	}
	if cl-1 != hp.CurrentSize {
		t.Error("Error on changing size")
	}
}

func TestParent(t *testing.T) {
	index := 10
	parent := (index - 1) / 2
	if parent != hp.Parent(index) {
		t.Error("Calculation Parent is invalid")
	}
}

func TestLeft(t *testing.T) {
	index := 9
	left := (index * 2) + 1
	if left != hp.Left(index) {
		t.Error("Error on calculation left child")
	}
}

func TestRight(t *testing.T) {
	index := 0
	right := (index * 2) + 2
	if right != hp.Right(index) {
		t.Error("Error on calculation right")
	}
}
