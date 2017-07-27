package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func isListPalindrome(l *ListNode) bool {
	var r *ListNode
	cc := l
	for l != nil {
		if r == nil {
			q := ListNode{l.Value, nil}
			r = &q
		} else {
			q := ListNode{l.Value, r}
			r = &q
		}
		l = l.Next
	}
	for r != nil && cc != nil {
		if r.Value != cc.Value {
			return false
		}
		cc = cc.Next
		r = r.Next
	}
	return true
}
func main() {
	fmt.Println("HElloWorld")
}
