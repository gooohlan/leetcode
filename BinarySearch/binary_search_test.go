package BinarySearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
	fmt.Println(search2(nums, 2))
}

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(leftSearch(nums, 6))
	fmt.Println(rightSearch(nums, 10))
	fmt.Println(searchRange(nums, 4))
}
