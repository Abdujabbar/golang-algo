package main

import "fmt"

//ListNode Linked List Item
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var lst *ListNode
	var res *ListNode
	for l1 != nil && l2 != nil {
		var q ListNode
		if (l1.Value).(int) <= (l2.Value).(int) {
			q = ListNode{l1.Value, nil}
			l1 = l1.Next
		} else {
			q = ListNode{l2.Value, nil}
			l2 = l2.Next
		}
		if res == nil {
			lst = &q
			res = lst
		} else {
			lst.Next = &q
			lst = lst.Next
		}
	}

	for l1 != nil {
		q := ListNode{l1.Value, nil}
		if lst == nil {
			lst = &q
			res = lst
		} else {
			lst.Next = &q
			lst = lst.Next
		}
		l1 = l1.Next
	}
	for l2 != nil {
		q := ListNode{l2.Value, nil}
		if lst == nil {
			lst = &q
			res = lst
		} else {
			lst.Next = &q
			lst = lst.Next
		}
		l2 = l2.Next
	}
	return res
}

func main() {
	// var aPrev *ListNode
	// var bPrev *ListNode
	// r := mergeTwoLinkedLists(aPrev, bPrev)
	// fmt.Println(r)
	var aPointer *ListNode
	// var aPrev *ListNode
	a := ListNode{1, nil}
	aPointer = &a
	aPrev := aPointer
	b := ListNode{2, nil}
	aPointer.Next = &b
	aPointer = aPointer.Next
	c := ListNode{3, nil}
	aPointer.Next = &c
	aPointer = aPointer.Next
	var bPointer *ListNode
	//
	// aa := ListNode{0, nil}
	// bPointer = &aa
	// var bPrev *ListNode
	// bPrev = bPointer
	// bb := ListNode{4, nil}
	// bPointer.Next = &bb
	// bPointer = bPointer.Next
	// cc := ListNode{5, nil}
	// bPointer.Next = &cc
	// bPointer = bPointer.Next
	cRes := mergeTwoLinkedLists(aPrev, bPointer)
	for cRes != nil {
		fmt.Println(cRes.Value)
		cRes = cRes.Next
	}
}
