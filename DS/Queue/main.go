package Queue

import ll "golang-algo/DS/LinkedList"

// Queue DS
type Queue struct {
	Lst ll.LinkedList
}

//Enqueue adds element to Queue
func (a *Queue) Enqueue(element interface{}) {
	a.Lst.PushBack(element)
}

//Dequeue removes element from Queue
func (a *Queue) Dequeue() interface{} {
	if a.IsEmpty() {
		return "Queue is empty"
	}
	return a.Lst.PopFront()
}

//Front returns front element of Queue
func (a *Queue) Front() interface{} {
	if a.IsEmpty() {
		return "Queue is empty"
	}
	return a.Lst.Front()
}

//Size returns count elements in Queue
func (a *Queue) Size() int {
	return a.Lst.Size()
}

//IsEmpty check weither is empty Queue or not
func (a *Queue) IsEmpty() bool {
	return a.Size() == 0
}

//Clear clearing Queue
func (a *Queue) Clear() {
	a.Lst.Clear()
}
