package Stack

import ll "golang-algo/DS/LinkedList"

//Stack structure
type Stack struct {
	Lst ll.LinkedList
}

//Push append element to Stack
func (a *Stack) Push(value interface{}) {
	a.Lst.PushFront(value)
}

//Pop element from Stack
func (a *Stack) Pop() interface{} {
	return a.Lst.PopFront()
}

//Size gets the size of stack
func (a *Stack) Size() int {
	return a.Lst.Size()
}

//IsEmpty checks is empty stack or not
func (a *Stack) IsEmpty() bool {
	return a.Size() == 0
}

//Clear cleares all elements in stack
func (a *Stack) Clear() {
	a.Lst.Clear()
}
