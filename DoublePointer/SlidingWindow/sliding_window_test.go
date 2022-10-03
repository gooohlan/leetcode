package SlidingWindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	window := minWindow("ADOBECODEBANC", "ABC")
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
