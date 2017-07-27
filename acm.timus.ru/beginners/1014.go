package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbs := []int{}
	switch n {
		case 0 :
			fmt.Println(10)
			break
		case 1 :
			fmt.Println(1)
			break
		default :
			for devider := 9; devider >= 2; {
				for n%devider == 0 {
					numbs = append(numbs, devider)
					n /= devider
				}
				devider--
			}

			if n == 1 {
				for j := len(numbs) - 1; j >= 0; j-- {
					fmt.Print(numbs[j])
				}
 			} else {
				fmt.Println(-1)
			}
			break

	}
}
