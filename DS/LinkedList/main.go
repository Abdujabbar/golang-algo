package LinkedList

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

//Size method for returning size of linked list
func (a *LinkedList) Size() int {
	return a.Length
}

//Empty checking for linked list
func (a *LinkedList) Empty() bool {
	return a.Length == 0
}

//ValueAt get value of linked list by index
func (a *LinkedList) ValueAt(index int) interface{} {
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

//PushFront pushes value at front
func (a *LinkedList) PushFront(value interface{}) {
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

//PopFront return element from front
func (a *LinkedList) PopFront() interface{} {
	v := a.Head.Value
	a.Head = a.Head.Next
	a.decreaseSize()
	return v
}

//PushBack pushes elment at back of linked list
func (a *LinkedList) PushBack(value interface{}) {
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

//PopBack returns element from back of linked list
func (a *LinkedList) PopBack() interface{} {
	cnode := a.Head
	prev := cnode
	for cnode.Next != nil {
		prev = cnode
		cnode = cnode.Next
	}
	a.Tail = prev
	prev.Next = nil
	a.decreaseSize()
	return cnode.Value
}

//Front return value of front element
func (a *LinkedList) Front() interface{} {
	return a.Head.Value
}

//Back return value of back element
func (a *LinkedList) Back() interface{} {
	if !a.Empty() {
		return a.Tail.Value
	}
	return nil
}

//InsertAt inserts element at index
func (a *LinkedList) InsertAt(index int, value interface{}) {
	counter := 0
	chead := a.Head
	var prev *Node
	if chead == nil {
		q := Node{Value: value, Next: nil}
		a.Head = &q
	} else {
		q := Node{Value: value, Next: nil}
		for chead != nil {
			if index == counter {
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
			prev.Next = &q
			q.Next = chead
			if chead == nil {
				a.Tail = &q
			}

		}
	}
	a.increaseSize()
}

//EraseAt removes element by index
func (a *LinkedList) EraseAt(index int) {
	counter := 0
	chead := a.Head
	var prev *Node
	for chead != nil {
		if counter == index {
			break
		}
		prev = chead
		chead = chead.Next
		counter++
	}
	if prev != nil {
		if chead != nil {
			prev.Next = chead.Next
		} else {
			prev.Next = nil
		}
	} else {
		chead = chead.Next
		a.Head = chead
	}
	a.decreaseSize()
}

//Print printes elements of linkedlist
func (a *LinkedList) Print() {
	if a.Empty() {
		fmt.Println("No items in linked list")
	} else {
		c := a.Head
		for c != nil {
			fmt.Println(c.Value)
			c = c.Next
		}
	}
}

//Reverse reverses files in linkedlist
func (a *LinkedList) Reverse() {
	var r LinkedList
	chead := a.Head
	for chead != nil {
		r.PushFront(chead.Value)
		chead = chead.Next
	}
	*a = r
}

// Clear cleares all elements in linked list
func (a *LinkedList) Clear() {
	a.Head = nil
	a.Tail = nil
	a.Length = 0
}
