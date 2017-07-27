package main

import "fmt"

// ListNode Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

//
func reverseNodesInKGroups(l *ListNode, k int) *ListNode {
	var lstReversed *ListNode
	var lstCurrent *ListNode
	lstCurrent = l
	// var res *ListNode
	// var hRes *ListNode
	var clen int
	for l != nil {
		l = l.Next
		clen++
	}
	var res *ListNode
	var hRes *ListNode
	var i int
	var cclen int
	if clen%k == 0 {
		cclen = clen
	} else {
		cclen = clen - clen%k
	}
	for i = 0; i < cclen; i++ {

		if i > 0 && i%k == 0 {
			if res == nil {
				res = lstReversed
				hRes = res
				lstReversed = nil
			} else {
				for res.Next != nil {
					res = res.Next
				}
				res.Next = lstReversed
				lstReversed = nil
			}
		}
		if lstReversed == nil {
			q := ListNode{lstCurrent.Value, nil}
			lstReversed = &q
		} else {
			q := ListNode{lstCurrent.Value, lstReversed}
			lstReversed = &q
		}
		lstCurrent = lstCurrent.Next
	}
	if clen == k {
		return lstReversed
	}

	if clen%k == 0 {
		if lstReversed != nil {
			for res.Next != nil {
				res = res.Next
			}
			res.Next = lstReversed
			res = res.Next
		}
	} else {
		if lstReversed != nil {
			if res == nil {
				res = lstReversed
				hRes = res
			} else {
				for res.Next != nil {
					res = res.Next
				}
				res.Next = lstReversed
				res = res.Next
			}
		}

		if lstCurrent != nil {
			for res.Next != nil {
				res = res.Next
			}
			res.Next = lstCurrent
			res = res.Next
		}
	}

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
	var cc *ListNode

	cc = reverseNodesInKGroups(aPrev, 4)

	for cc != nil {
		fmt.Println(cc.Value)
		cc = cc.Next
	}
}
