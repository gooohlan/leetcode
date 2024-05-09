package Array

import (
	"fmt"
	"sort"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func TestRemoveElement(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	val := 2
	fmt.Println(removeElement(arr, val))
}

func TestTwoSum(t *testing.T) {
	arr := []int{1, 3, 6, 7}
	fmt.Println(twoSum(arr, 9))
}

func TestMoveZeroes(t *testing.T) {
	arr := []int{1, 0, 2, 3, 0, 0, 5}
	moveZeroes(arr)
	fmt.Println(arr)
}

func TestReverseString(t *testing.T) {
	str := "abcdABCD"

	fmt.Println(string(reverseString([]byte(str))))
}

func TestLongestPalindrome(t *testing.T) {
	str := "cbbd"
	fmt.Println(longestPalindrome(str))
}

func Test15(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func Test18(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(nums, 8))
}

func Test1(t *testing.T) {
	nums := []int{3, 2, 4}
	sort.Ints(nums)
	fmt.Println(nums)
}

func Test557(t *testing.T) {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
