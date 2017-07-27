package main

import (
	"fmt"
	"strconv"
)

func isCryptSolution(crypt []string, solution [][]string) bool {
	mapping := make(map[string]string)
	for _, v := range solution {
		mapping[string(v[0])] = v[1]
	}
	trans := [3]string{}
	for i := 0; i < len(crypt); i++ {
		v := ""
		for j := 0; j < len(crypt[i]); j++ {
			v += mapping[string(crypt[i][j])]
		}
		trans[i] = v
		if len(trans[i]) > 1 && string(trans[i][0]) == "0" {
			return false
		}
	}
	first, err := strconv.Atoi(trans[0])
	if err != nil {
		return false
	}
	second, err := strconv.Atoi(trans[1])
	if err != nil {
		return false
	}
	third, err := strconv.Atoi(trans[2])
	if err != nil {
		return false
	}
	return first+second == third
}
func main() {
	crypted := []string{"TEN", "TWO", "ONE"}
	sol := [][]string{{"O", "1"},
		{"T", "0"},
		{"W", "9"},
		{"E", "5"},
		{"N", "4"}}
	fmt.Println(isCryptSolution(crypted, sol))
}
