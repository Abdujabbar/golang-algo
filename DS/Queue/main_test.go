package Queue

import "testing"

var q Queue

func TestSize(t *testing.T) {
	if q.Size() != 0 {
		t.Error("Invalid size of queue")
	}
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Size() != 2 {
		t.Error("invalid size of queue after insert elements")
	}
}

func TestIsEmpty(t *testing.T) {
	if q.IsEmpty() != false {
		t.Error("Invalid error on is empty checking")
	}
	q.Clear()
	if q.IsEmpty() != true {
		t.Error("Invalid error after clearing")
	}
}

func TestEnqueue(t *testing.T) {
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Dequeue() != 10 {
		t.Error("Error on enqueu")
	}
	q.Clear()
}

func TestDequeue(t *testing.T) {
	q.Enqueue(20)
	if q.Dequeue() != 20 {
		t.Error("Error on queue")
	}
}
