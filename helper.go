package main

// Return the number of solved squares
// 81 denotes a solved puzzle
func countSolved(grid [][][]int) int {
	solved := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(grid[i][j]) == 1 {
				solved++
			}
		}
	}
	return solved
}

// Return the total number of possible values
// This includes already known values
// Thus, this will return a minimum of 81
func countHints(grid [][][]int) int {
	hints := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			hints += len(grid[i][j])
		}
	}
	return hints
}

// Wrap a row of the grid into an array of pointers
func wrapRow(grid [][][]int, row int) []*[]int {
	result := []*[]int{}
	for column := 0; column < 9; column++ {
		result = append(result, &grid[row][column])
	}
	return result
}

// Wrap a column of the grid into an array of pointers
func wrapColumn(grid [][][]int, column int) []*[]int {
	result := []*[]int{}
	for row := 0; row < 9; row++ {
		result = append(result, &grid[row][column])
	}
	return result
}

// Wrap a box of the grid into an array of pointers
func wrapBox(grid [][][]int, box int) []*[]int {
	result := []*[]int{}
	for i := 0; i < 9; i++ {
		result = append(result, &grid[3*(box/3)+(i/3)][3*(box%3)+(i%3)])
	}
	return result
}
