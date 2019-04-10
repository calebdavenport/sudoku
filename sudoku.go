package main

import (
    "fmt"
    "reflect"
)

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
        grid[i] = tuple_remove(grid[i])
        grid[i] = trivial_reduce(grid[i])
        column := wrapColumn(grid, i)
        column = trivial_reduce(column)
        grid = unwrapColumn(grid, column, i)

		for j := 0; j < 9; j++ {
			if len(grid[i][j]) == 1 {
                // Clear Box
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

func trivial_reduce(grid [][]int) [][]int {
    for pivot := 0; pivot < 9; pivot++ {
        if len(grid[pivot]) > 1 {
            continue
        }
        for check := 0; check < 9; check++ {
            if pivot == check {
                continue
            }
            for i := 0; i < len(grid[check]); i++ {
                if grid[check][i] == grid[pivot][0] {
                    grid[check] = append(grid[check][:i], grid[check][i+1:]...)
                    break
                }
            }
        }
    }
    return grid
}


func tuple_remove(grid [][]int) [][]int {
    // Only works with rows for now
    for i := 0; i < 8; i++ {
        for j := i + 1; j < 9; j++ {
            if reflect.DeepEqual(grid[i], grid[j]) && len(grid[i]) == 2 {
                for x := 0; x < 9; x++ {
                    if x == i || x == j {
                        continue
                    }
                    for y := 0; y < len(grid[x]); y++ {
                        if grid[x][y] == grid[i][0] || grid[x][y] == grid[i][1] {
							grid[x] = append(grid[x][:y], grid[x][y+1:]...)
                        }
                    }
                }
            }
        }
    }
    return grid
}

func wrapColumn(grid [][][]int, column int) [][]int {
    result := [][]int{}
    for row := 0; row < 9; row++ {
        result = append(result, grid[row][column])
    }
    return result
}

func unwrapColumn(grid [][][]int, column_data [][]int, column_num int) [][][]int {
    for row := 0; row < 9; row++ {
        grid[row][column_num] = column_data[row]
    }
    return grid
}
