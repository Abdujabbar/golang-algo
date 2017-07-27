package main

import "fmt"

func areFollowingPatterns(strings []string, patterns []string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for _, v := range strings {
		map1[v]++
	}
	for _, v := range patterns {
		map2[v]++
	}
	if len(map1) != len(map2) {
		return false
	}
	for i := 1; i < len(strings); i++ {
		if strings[i] == strings[i-1] && patterns[i] != patterns[i-1] {
			return false
		}
		if strings[i] != strings[i-1] && patterns[i] == patterns[i-1] {
			return false
		}
	}
	return true
}

func main() {
	strings1 := []string{
		"cat",
		"dog",
		"doggy",
	}
	strings2 := []string{
		"a",
		"b",
		"b",
	}
	fmt.Println(areFollowingPatterns(strings1, strings2))
}
