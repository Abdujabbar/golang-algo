package main

import "fmt"

// ListNode Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

//
// func removeKFromList(l *ListNode, k int) *ListNode {
// 	if l.Value == k {
// 		l = l.Next
// 		return l
// 	} else {
// 		l.Next = removeKFromList(l.Next, k)
// 		return l
// 	}
// }
func removeKFromList(l *ListNode, k int) *ListNode {
	var res *ListNode
	hPrev := res
	for l != nil {
		if l.Value == k {
			continue
		}
		if res == nil {
			q := ListNode{l.Value, nil}
			res = &q
			hPrev = res
		} else {
			res.Next = &ListNode{l.Value, nil}
			res = res.Next
		}
		l = l.Next
	}
	return hPrev
}

func main() {
	var res ListNode
	fmt.Println(res)
	// q := ListNode{3, nil}
	// res = &q
	// fmt.Println(res)
	// n := ListNode{1, nil}
	// lst := *ListNode{2}
	// *lst.Next = l
	// l.Next = &n
	// q := ListNode{4, nil}
	// n.Next = &q
	// var qq *ListNode
	// qq = &l
	// fmt.Println(qq)
	// fmt.Println(removeKFromList(qq, 1))
	// for l.Value != nil {
	// 	fmt.Println(l)
	// 	if l.Next != nil {
	// 		l = *l.Next
	// 	}
	// }
}
