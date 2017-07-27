package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var Res = 0

func sp(prefix, remain string, n int) {
	if n == 0 {
		fmt.Println(prefix)
		return
	}
	if len(remain) == 0 {
		return
	}
	pr := prefix + remain[:1]
	sp(pr, remain[1:], n-1)
	sp(prefix, remain[1:], n)
}

func p(prefix, remain []int, n int) {
	if n == 0 {
		_ans := prefix[0] & prefix[1]
		for i := 2; i < len(prefix); i++ {
			_ans = _ans & prefix[i]
		}
		Res = max(Res, _ans)
		return
	}
	if len(remain) == 0 {
		return
	}
	pr := prefix
	pr = append(prefix, remain[:1]...)
	p(pr, remain[1:], n-1)
	p(prefix, remain[1:], n)
}

func swap(inp []int, x, y int) {
	t := inp[x]
	inp[x] = inp[y]
	inp[y] = t
}

func pp(inp []int) {
	left := 0
	right := len(inp) - 1

	for left < right {
		for inp[right] == 0 && right > left {
			right--
		}
		for inp[left] != 0 && left < right {
			left++
		}

		if left != right && left < right {
			swap(inp, left, right)
		}
	}
}

func main() {
	nums := []int{1, 2, 0, 3, 4, 5}
	pp(nums)
	fmt.Println(nums)
	// prefix := []int{}
	// remain := []int{3, 5, 6}
	// n := 2
	//
	// p(prefix, remain, n)
	// fmt.Println(Res)
}
