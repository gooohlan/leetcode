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

func TestCombine(t *testing.T) {
    fmt.Println(combine(3, 2))
}

func TestSubsetsWithDup(t *testing.T) {
    nums := []int{1, 2, 2}
    fmt.Println(subsetsWithDup(nums))
}

func TestCombinationSum2(t *testing.T) {
    nums := []int{10, 1, 2, 7, 6, 1, 5}
    fmt.Println(combinationSum2(nums, 8))
}

func TestPermuteUnique(t *testing.T) {
    nums := []int{1, 2, 3}
    fmt.Println(permuteUnique(nums))
    nums = []int{1, 2, 1}
    fmt.Println(permuteUnique(nums))
}

func TestCombinationSum(t *testing.T) {
    nums := []int{2, 3, 5}
    fmt.Println(combinationSum(nums, 8))
}

func Test52(t *testing.T) {
    fmt.Println(totalNQueens(4))
}

func Test698(t *testing.T) {
    fmt.Println(canPartitionKSubsets2([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}

func Test37(t *testing.T) {
    board := [][]byte{
        []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
    solveSudoku(board)
    for _, bytes := range board {
        fmt.Println(string(bytes))
    }
}
