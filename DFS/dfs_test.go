package DFS

import (
    "fmt"
    "testing"
)

func Test200(t *testing.T) {
    grid := [][]byte{
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    }
    fmt.Println(numIslands(grid))
}

func Test1254(t *testing.T) {
    grid := [][]int{
        {0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0},
    }
    fmt.Println(closedIsland(grid))
}
