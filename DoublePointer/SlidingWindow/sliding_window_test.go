package SlidingWindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	window := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println(window)
}
