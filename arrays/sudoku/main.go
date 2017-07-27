package main

import "strconv"

func findEmptyCell(grid [][]string) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == "." {
				return i, j
			}
		}
	}
	return -1, -1
}
func isSafeRow(grid [][]string, row, col int, condidate string) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == condidate {
			return false
		}
	}
	return true
}
func isSafeCol(grid [][]string, row, col int, condidate string) bool {
	for i := 0; i < 9; i++ {
		if grid[i][col] == condidate {
			return false
		}
	}
	return true
}

func isSafeBox(grid [][]string, boxRowStar, boxColStar int, condidate string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxRowStar+i][boxColStar+j] == condidate {
				return false
			}
		}
	}
	return true
}

func isSafe(grid [][]string, row, col int, condidate string) bool {

	return isSafeRow(grid, row, col, condidate) &&
		isSafeCol(grid, row, col, condidate) &&
		isSafeBox(grid, row-row%3, col-col%3, condidate)
}

func solve(grid [][]string) bool {
	row, col := findEmptyCell(grid)
	if row == col && row == -1 {
		return true
	}
	for i := 1; i <= 9; i++ {
		if isSafe(grid, row, col, strconv.Itoa(i)) {
			grid[row][col] = strconv.Itoa(i)
			if solve(grid) {
				return true
			}
			grid[row][col] = "."
		}
	}
	return false
}

func sudoku2(grid [][]string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				continue
			}
			v := grid[i][j]
			//checking row
			for k := 0; k < 9; k++ {
				if v == grid[i][k] && k != j {
					return false
				}

				if v == grid[k][j] && k != i {
					return false
				}
			}
			boxStarRow := i - i%3
			boxStarCol := j - j%3
			for x := boxStarRow; x < boxStarRow+3; x++ {
				for y := boxStarCol; y < boxStarCol+3; y++ {
					if x != i && y != j && v == grid[x][y] {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {

	// matrix := [][]string{{".", ".", ".", ".", "2", ".", ".", "9", "."},
	// 	{".", ".", ".", ".", "6", ".", ".", ".", "."},
	// 	{"7", "1", ".", ".", "7", "5", ".", ".", "."},
	// 	{".", "7", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", "8", "3", ".", ".", "."},
	// 	{".", ".", "8", ".", ".", "7", ".", "6", "."},
	// 	{".", ".", ".", ".", ".", "2", ".", ".", "."},
	// 	{".", "1", ".", "2", ".", ".", ".", ".", "."},
	// 	{".", "2", ".", ".", "3", ".", ".", ".", "."}}
	// fmt.Println(solve(matrix))
	// for _, v := range matrix {
	// 	fmt.Println(v)
	// }
}
