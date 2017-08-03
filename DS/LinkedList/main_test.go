package LinkedList

import "testing"

var q LinkedList

func TestEmpty(t *testing.T) {
	if q.Empty() != true {
		t.Error("Not Valid Empty Function")
	}

	q.PushBack(10)

	if q.Empty() == true {
		t.Error("Not Valid Empty function after add item")
	}
	q.PopBack()
}

func TestSize(t *testing.T) {
	q.PushBack(10)
	q.PushBack(20)
	if q.Size() != 2 {
		t.Error("Count elements Error")
	}
	q.PopBack()
	if q.Size() != 1 {
		t.Error("Count Elements Error")
	}
	q.PopBack()
	if q.Size() != 0 {
		t.Error("Count Elements Error")
	}
}

func TestClear(t *testing.T) {
	q.PushBack(10)
	q.PushBack(20)
	q.Clear()
	if q.Size() != 0 {
		t.Error("Clearing error")
	}
}
