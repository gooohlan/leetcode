package Backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	list := permute(nums)
	fmt.Println(list)
}

func TestSolveNQueens(t *testing.T) {
	queens := solveNQueens(4)
	fmt.Println(queens)
}

func TestSubsets(t *testing.T) {
	list := subsets([]int{1, 2, 3})
	fmt.Println(list)
}
