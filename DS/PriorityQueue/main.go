package PriorityQueue

import (
	"fmt"
	hhp "golang-algo/DS/Heap"
	ll "golang-algo/DS/Queue"
)

//PqItem PriorityQueue item
type PqItem struct {
	Value    interface{}
	Priority int
}

//PriorityQueue ds
type PriorityQueue struct {
	Collection   []*ll.Queue
	Length       int
	CurrentSize  int
	PriorityHeap *hhp.Heap
}

//InitPriorityQueue priority queue
func InitPriorityQueue(sz int) *PriorityQueue {
	var pq PriorityQueue
	pq.Length = sz
	pq.Collection = make([]*ll.Queue, sz)
	pq.PriorityHeap = hhp.InitHeap(sz)
	return &pq
}

//Size return size of queue
func (a *PriorityQueue) Size() int {
	return a.Length
}

//IsEmpty return true when pq is empty
func (a *PriorityQueue) IsEmpty() bool {
	return a.CurrentSize == 0
}

//Enqueue enqueues element to pq
func (a *PriorityQueue) Enqueue(value PqItem) {
	if a.CurrentSize >= a.Length {
		fmt.Println("PQ is full")
		return
	}
	if a.Collection[value.Priority] == nil {
		var q ll.Queue
		a.Collection[value.Priority] = &q
		a.PriorityHeap.Insert(value.Priority)
		a.CurrentSize++
	}
	a.Collection[value.Priority].Enqueue(value)
}

//Dequeue dequeues element from pq
func (a *PriorityQueue) Dequeue() interface{} {
	if !a.IsEmpty() {
		priority := a.PriorityHeap.GetMin()
		v := a.Collection[priority].Dequeue()
		if a.Collection[priority].IsEmpty() {
			a.PriorityHeap.ExtractMin()
			a.CurrentSize--
		}
		return v
	}
	return "PQ is empty"
}
