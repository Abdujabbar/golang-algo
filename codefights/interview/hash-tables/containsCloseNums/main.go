package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func containsCloseNums(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && abs(i-j) <= k {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{
		0, 1, 2, 3, 5, 2,
	}
	k := 2
	fmt.Println(containsCloseNums(nums, k))
}
