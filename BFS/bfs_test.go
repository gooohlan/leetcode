package BFS

import (
    "fmt"
    "testing"
)

func TestMinDepth(t *testing.T) {
    node := &TreeNode{
        1,
        &TreeNode{2, &TreeNode{4, nil, nil}, nil},
        &TreeNode{3, nil, &TreeNode{5, nil, nil}},
    }
    depth := minDepth(node)
    fmt.Println(depth)
}

func TestOpenLock(t *testing.T) {
    deadends := []string{"0201", "0101", "0102", "1212", "2002"}
    // fmt.Println(openLock(deadends, "0202"))
    fmt.Println(openLock2(deadends, "0202"))
}

func TestSlidingPuzzle(t *testing.T) {
    board := [][]int{
        []int{4, 1, 2},
        []int{5, 0, 3},
    }
    fmt.Println(slidingPuzzle(board))
    fmt.Println(slidingPuzzle2(board))
}
