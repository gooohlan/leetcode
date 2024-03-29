package Merge

import (
	"fmt"
	"testing"
)

func Test912(t *testing.T) {
	nums := []int{4, 1, 7, 2, 1}
	// fmt.Println(sortArray(nums))
	fmt.Println(sortArrayFast(nums))
}

func Test315(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	fmt.Println(countSmaller(nums))
}

func Test493(t *testing.T) {
	nums := []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs(nums))
}

func Test327(t *testing.T) {
	nums := []int{-2, 5, -1}
	fmt.Println(countRangeSum(nums, -2, 2))
}

func Test215(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
