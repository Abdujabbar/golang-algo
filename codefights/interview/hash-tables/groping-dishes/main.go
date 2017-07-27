package main

import (
	"fmt"
	"sort"
)

func groupingDishes(dishes [][]string) [][]string {
	ingridients := make(map[string][]string)
	ingridientCounter := make(map[string]bool)
	counter := 0
	keys := []string{}
	for i := 0; i < len(dishes); i++ {
		for j := 1; j < len(dishes[i]); j++ {
			ingridients[dishes[i][j]] = append(ingridients[dishes[i][j]], dishes[i][0])
			if !ingridientCounter[dishes[i][j]] && len(ingridients[dishes[i][j]]) > 1 {
				counter++
				ingridientCounter[dishes[i][j]] = true
				keys = append(keys, dishes[i][j])
			}
		}
	}
	resDishes := make([][]string, counter)
	sort.Strings(keys)
	c := 0
	for _, v := range keys {
		vr := []string{}
		vr = append(vr, v)
		sort.Strings(ingridients[v])
		vr = append(vr, ingridients[v]...)
		resDishes[c] = make([]string, len(vr))
		resDishes[c] = vr
		c++
	}
	return resDishes
}

func main() {
	dishes := [][]string{{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
		{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
		{"Quesadilla", "Chicken", "Cheese", "Sauce"},
		{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"}}
	fmt.Println(groupingDishes(dishes))
}
