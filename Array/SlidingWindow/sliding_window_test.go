package SlidingWindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	window := minWindow2("ADOBECODEBANC", "ABC")
	fmt.Println(window)
}

func TestCheckInclusion(t *testing.T) {
	s1 := "abcabd"
	s2 := "abdcba"
	fmt.Println(checkInclusion(s1, s2))
}

func Test438(t *testing.T) {
	s1 := "cbaebabacd"
	s2 := "abc"
	fmt.Println(findAnagrams(s1, s2))
	fmt.Println(findAnagrams2(s1, s2))
}

func Test3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func Test1004(t *testing.T) {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(nums, 3))
}

func Test187(t *testing.T) {
	str := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println()
}
