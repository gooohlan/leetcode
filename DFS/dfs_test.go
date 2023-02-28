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
