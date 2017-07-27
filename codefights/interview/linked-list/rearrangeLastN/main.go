package main

import "fmt"

// ListNode Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func rearrangeLastN(l *ListNode, n int) *ListNode {
	// var startPartial *ListNode
	var c int
	startPartial := l
	var res *ListNode
	var hRes *ListNode
	cll := l
	for l != nil {
		l = l.Next
		c++
	}

	if c == n || n == 0 {
		return startPartial
	}
	cc := 0
	for cll != nil {
		if cc == c-n {
			res = cll
			hRes = res
			break
		}
		cll = cll.Next
		cc++
	}

	for res.Next != nil {
		res = res.Next
	}
	prevPartial := startPartial
	for i := 0; i < c-n-1; i++ {
		startPartial = startPartial.Next
	}
	startPartial.Next = nil
	res.Next = prevPartial
	res = res.Next
	return hRes
}

func main() {
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
	d := ListNode{4, nil}
	aPointer.Next = &d
	aPointer = aPointer.Next
	e := ListNode{5, nil}
	aPointer.Next = &e
	aPointer = aPointer.Next
	f := ListNode{6, nil}
	aPointer.Next = &f
	aPointer = aPointer.Next
	g := ListNode{7, nil}
	aPointer.Next = &g
	aPointer = aPointer.Next
	cc := rearrangeLastN(aPrev, 0)
	for cc != nil {
		fmt.Println(cc.Value)
		cc = cc.Next
	}
}
