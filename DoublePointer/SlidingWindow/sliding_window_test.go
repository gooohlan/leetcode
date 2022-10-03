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
