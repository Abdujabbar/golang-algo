package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNodes Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func fillLeadingZeros(s string) string {
	if len(s) < 4 {
		clen := len(s)
		for i := 0; i < 4-clen; i++ {
			s = "0" + s
		}
	}
	return s
}
func trimZeros(s string) string {

	return strings.TrimLeft(s, "0")
}
func addTwoNumbers(a string, b string) string {
	r := 0
	c := ""
	var res int
	var i int
	for i = 0; i < len(a) && i < len(b); i++ {
		aVal, err := strconv.Atoi(string(a[len(a)-i-1]))
		if err != nil {
			fmt.Println(err)
			return ""
		}

		bVal, err := strconv.Atoi(string(b[len(b)-i-1]))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		res = aVal + bVal + r
		if res >= 10 {
			r = 1
			res -= 10
		} else {
			r = 0
		}
		c += strconv.Itoa(res)
	}
	for i < len(a) {
		aVal, err := strconv.Atoi(string(a[len(a)-i-1]))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		aVal += r
		if aVal >= 10 {
			aVal -= 10
			r = 1
		} else {
			r = 0
		}
		c += strconv.Itoa(aVal)
		i++
	}

	for i < len(b) {
		bVal, err := strconv.Atoi(string(b[len(b)-i-1]))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		bVal += r
		if bVal >= 10 {
			bVal -= 10
			r = 1
		} else {
			r = 0
		}
		c += strconv.Itoa(bVal)
		i++
	}

	if r != 0 {
		c += "1"
	}

	return c
}
func reverseStr(str string) (out string) {
	for _, s := range str {
		out = string(s) + out
	}

	return
}
func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	var firstNumber string
	var secondNumber string
	for a != nil {
		if firstNumber == "" {
			firstNumber = strconv.Itoa((a.Value).(int))
		} else {
			firstNumber += fillLeadingZeros(strconv.Itoa((a.Value).(int)))
		}
		a = a.Next
	}
	for b != nil {
		if secondNumber == "" {
			secondNumber = strconv.Itoa((b.Value).(int))
		} else {
			secondNumber += fillLeadingZeros(strconv.Itoa((b.Value).(int)))
		}
		b = b.Next
	}

	thirdNumber := addTwoNumbers(firstNumber, secondNumber)
	var res *ListNode
	var index int
	for index < len(thirdNumber) {

		// v := thirdNumber[index : index+4]
		// value := trimZeros(v)
		var value int
		if index+4 < len(thirdNumber) {
			value, _ = strconv.Atoi(reverseStr(thirdNumber[index : index+4]))
		} else {
			value, _ = strconv.Atoi(reverseStr(thirdNumber[index:len(thirdNumber)]))
		}
		if res == nil {
			q := ListNode{value, nil}
			res = &q
		} else {
			q := ListNode{value, res}
			res = &q
		}
		index += 4
	}

	return res
}

func main() {
	// var aPointer *ListNode
	// // var aPrev *ListNode
	// a := ListNode{123, nil}
	// aPointer = &a
	// aPrev := aPointer
	// b := ListNode{4, nil}
	// aPointer.Next = &b
	// aPointer = aPointer.Next
	// c := ListNode{5, nil}
	// aPointer.Next = &c
	// aPointer = aPointer.Next
	// var bPointer *ListNode
	//
	// aa := ListNode{100, nil}
	// bPointer = &aa
	// bPrev := bPointer
	// bb := ListNode{100, nil}
	// bPointer.Next = &bb
	// bPointer = bPointer.Next
	// cc := ListNode{100, nil}
	// bPointer.Next = &cc
	// bPointer = bPointer.Next
	// lst := addTwoHugeNumbers(aPrev, bPrev)
	// for lst != nil {
	// 	fmt.Println(lst.Value)
	// 	lst = lst.Next
	// }
	// fmt.Println(fillLeadingZeros("123"))
	// a := "12300040005"
	// b := ""
	// v := 10
	// var vv interface{}
	// fmt.Printf("%T", vv)
	// vv = v
	// fmt.Printf("%T", vv)

	// v, err := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", "03"))
	// if err != nil {
	// fmt.Println(err)
	// } else {
	// fmt.Println(v)
	// }
}
