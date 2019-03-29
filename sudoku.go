package main

import "fmt"

type grid struct {
    data    [][][]int
    empty   int
}

func main() {
    g := grid{
        data: [][][]int{
    //var grid = [][][]int{
            {{7}, {5}, {}, {2}, {6}, {}, {}, {3}, {}},
            {{}, {6}, {8}, {}, {}, {}, {}, {}, {}},
            {{3}, {}, {}, {}, {}, {8}, {6}, {1}, {}},
            {{}, {}, {}, {9}, {}, {}, {4}, {5}, {}},
            {{}, {}, {4}, {8}, {}, {7}, {1}, {}, {}},
            {{}, {7}, {5}, {}, {}, {6}, {}, {}, {}},
            {{}, {2}, {9}, {6}, {}, {}, {}, {}, {1}},
            {{}, {}, {}, {}, {}, {}, {5}, {4}, {}},
            {{}, {1}, {}, {}, {7}, {2}, {}, {8}, {6}},
        },
        empty: 81,
    }

    g.empty = initPencilMarks(g.data)
    reducePencilMarks(g.data)
    for i := 0; i < 9; i++ {
        fmt.Println(g.data[i])
        fmt.Println(g.empty)
    }
}

func initPencilMarks(grid [][][]int) int {
    empty := 0
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if len(grid[i][j]) == 0 {
                grid[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
                empty++
            }
        }
    }
    return empty
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
                // TODO: clear box
            }
        }
    }
}
