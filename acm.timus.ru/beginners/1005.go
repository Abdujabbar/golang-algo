package main

import "fmt"

var Res = 987654321
var sum = 0

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(v int) int {
	if v < 0 {
		return (-1) * v
	}
	return v
}
func p(a [21]int, _len int, pos, curSum int) {
	if pos == _len {
		Res = min(Res, abs(sum-2*curSum))
		return
	}
	p(a, _len, pos+1, curSum)
	p(a, _len, pos+1, curSum+a[pos])
}

func main() {
	var n int
	fmt.Scan(&n)
	items := [21]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i])
		sum += items[i]
	}
	p(items, n, 0, 0)
	fmt.Println(Res)
}
