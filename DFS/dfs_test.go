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

func Test1020(t *testing.T) {
    grid := [][]int{
        {0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0},
    }
    
    fmt.Println(numEnclaves(grid))
}

func Test695(t *testing.T) {
    grid := [][]int{
        {1, 1, 0, 0, 0},
        {1, 1, 0, 0, 0},
        {0, 0, 1, 0, 0},
        {0, 0, 0, 1, 1},
    }
    fmt.Println(maxAreaOfIsland(grid))
}

func Test694(t *testing.T) {
    grid := [][]int{
        {1, 1, 0, 1, 1},
        {1, 0, 0, 0, 0},
        {0, 0, 0, 0, 1},
        {1, 1, 0, 1, 1},
    }
    islands := numDistinctIslands(grid)
    fmt.Println(islands)
}
