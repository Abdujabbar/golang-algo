package main

import "fmt"

// Node Item of linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList structure
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (a *LinkedList) size() int {
	return a.Length
}

func (a *LinkedList) empty() bool {
	return a.Length == 0
}

func (a *LinkedList) valueAt(index int) interface{} {
	counter := 0
	chead := a.Head
	for chead != nil {
		if index == counter {
			return chead.Value
		}
		counter++
	}
	return -1
}

func (a *LinkedList) increaseSize() {
	a.Length++
}

func (a *LinkedList) decreaseSize() {
	a.Length--
}

func (a *LinkedList) pushFront(value interface{}) {
	q := Node{Value: value, Next: a.Head}
	a.Head = &q
	if a.Head == nil {
		a.Tail = &q
	} else {
		chead := a.Head
		for chead.Next != nil {
			chead = chead.Next
		}
		a.Tail = chead
	}

	a.increaseSize()
}

func (a *LinkedList) popFront() interface{} {
	v := a.Head.Value
	a.Head = a.Head.Next
	a.decreaseSize()
	return v
}

func (a *LinkedList) pushBack(value interface{}) {
	q := Node{Value: value, Next: nil}
	if a.Head == nil {
		a.Head = &q
		a.Tail = &q
	} else {
		a.Tail.Next = &q
		a.Tail = a.Tail.Next
	}
	a.increaseSize()
}

func (a *LinkedList) popBack() interface{} {
	cnode := a.Head
	prev := cnode
	for cnode.Next != nil {
		prev = cnode
		cnode = cnode.Next
	}
	a.Tail = prev
	prev.Next = nil
	return cnode.Value
}

func (a *LinkedList) front() interface{} {
	return a.Head.Value
}

func (a *LinkedList) back() interface{} {
	if !a.empty() {
		return a.Tail.Value
	}
	return nil
}

func (a *LinkedList) insertAt(index int, value interface{}) {
	counter := 0
	chead := a.Head
	prev := chead
	if chead == nil {
		q := Node{Value: value, Next: nil}
		a.Head = &q
	} else {

		q := Node{Value: value, Next: nil}
		for chead != nil {
			if index == counter {
				// prev = &q
				// prev.Next = chead
				break
			}
			prev = chead
			chead = chead.Next
			counter++
		}
		if prev == nil {
			prev = &q
			prev.Next = chead
			a.Head = prev
		} else {

		}
		fmt.Println(prev)
		fmt.Println(chead)
		fmt.Println(q)
		// prev.Next = &q
		// q.Next = chead
		// prev.Next = chead.Next

	}
	a.increaseSize()
}

func (a *LinkedList) eraseAt(index int) {
	counter := 0
	chead := a.Head
	var prev *Node
	for chead != nil {
		if counter == index {
			break
		}
		prev = chead
		chead = chead.Next
	}
	if prev != nil {
		prev.Next = chead.Next
	} else {
		chead = chead.Next
		a.Head = chead
	}
	a.decreaseSize()
}

func (a *LinkedList) print() {
	if a.empty() {
		fmt.Println("No items in linked list")
	} else {
		c := a.Head
		for c != nil {
			fmt.Println(c.Value)
			c = c.Next
		}
	}
}

func main() {
	var q LinkedList
	q.pushFront(10)
	q.pushFront(20)
	q.pushBack(30)
	q.pushBack(40)
	// q.insert(0, 40)
	// q.insert(0, 20)
	q.print()
	// q.pushBack(10)
	// q.pushBack(20)
	// // q.pushFront(30)
	//
	// q.pushTop(40)
	// q.print()
	// fmt.Println("-----")
	// // fmt.Println(q.popFront())
	// fmt.Println("-----")
	// q.print()
	// fmt.Println(q.tail())
	// q.delete(10)
	// q.delete(20)
	// q.delete(30)
	// q.print()
}
