package main

import "fmt"

type grid struct {
	data   [][][]int
	solved int
	hints  int
}

func main() {
	g := grid{
		data: [][][]int{
			{{8}, {}, {}, {}, {1}, {}, {}, {5}, {}},
			{{}, {6}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {2}, {}, {}, {3}, {8}, {}, {6}, {4}},
			{{4}, {}, {}, {}, {}, {1}, {}, {9}, {}},
			{{9}, {}, {}, {}, {5}, {}, {}, {}, {2}},
			{{}, {1}, {}, {9}, {}, {}, {}, {}, {3}},
			{{1}, {9}, {}, {4}, {6}, {}, {}, {8}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {2}, {}},
			{{}, {7}, {}, {}, {8}, {}, {}, {}, {5}},
		},
		solved: 0,
		hints:  0,
	}

	initPencilMarks(g.data)
	reducePencilMarks(g.data)
	g.hints = countHints(g.data)
	prevHints := g.hints
	for {
		prevHints = g.hints
		reducePencilMarks(g.data)
		g.hints = countHints(g.data)
		if prevHints == g.hints {
			break
		}
	}
	g.solved = countSolved(g.data)
	for i := 0; i < 9; i++ {
		fmt.Println(g.data[i])
	}
	fmt.Println(g.solved)
	fmt.Println(g.hints)
}

func initPencilMarks(grid [][][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(grid[i][j]) == 0 {
				grid[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
	}
}

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

func countHints(grid [][][]int) int {
	hints := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			hints += len(grid[i][j])
		}
	}
	return hints
}

func reducePencilMarks(grid [][][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(grid[i][j]) == 1 {
				// Clear row
				for k := 0; k < 9; k++ {
					if k == j {
						continue
					}
					for y := 0; y < len(grid[i][k]); y++ {
						if grid[i][k][y] == grid[i][j][0] {
							grid[i][k] = append(grid[i][k][:y], grid[i][k][y+1:]...)
							break
						}
					}
				}
				// Clear column
				for k := 0; k < 9; k++ {
					if k == i {
						continue
					}
					for y := 0; y < len(grid[k][j]); y++ {
						if grid[k][j][y] == grid[i][j][0] {
							grid[k][j] = append(grid[k][j][:y], grid[k][j][y+1:]...)
							break
						}
					}
				}
				same := 3*(i%3) + (j % 3)
				for k := 0; k < 9; k++ {
					if k == same {
						continue
					}
					check_i := ((k) / 3) + 3*(i/3)
					check_j := ((k) % 3) + 3*(j/3)
					for y := 0; y < len(grid[check_i][check_j]); y++ {
						if grid[check_i][check_j][y] == grid[i][j][0] {
							grid[check_i][check_j] = append(grid[check_i][check_j][:y], grid[check_i][check_j][y+1:]...)
							break
						}

					}

				}

			}
		}
	}
}
