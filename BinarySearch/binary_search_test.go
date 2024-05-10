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

func Test875(t *testing.T) {
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}

func Test1011(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(shipWithinDays(nums, 5))
}

func Test410(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(nums, 3))
}
