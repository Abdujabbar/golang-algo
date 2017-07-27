package arrays

func Swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x,y int) int {
	if x < y {
		return y
	}
	return x
}

func FindMaxThreeMultiple(a []int) int{
	highest := max(a[0], a[1])
	lowest := min(a[0], a[1])
	highest_of_three := a[0] * a[1] * a[2]
	lowest_of_two := a[0] * a[1]
	highest_of_two := a[0] * a[1]

	for i := 2; i < len(a); i++ {
		current := a[i]
		highest_of_three = max(highest_of_three, max(highest_of_two * current, lowest_of_two * current))

		highest_of_two = max(highest_of_two, max(highest * current, lowest * current))

		lowest_of_two = min(lowest_of_two, min(highest * current, lowest * current))

		highest = max(highest, current)

		lowest = min(lowest, current)
	}

	return highest_of_three
}