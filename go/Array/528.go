package Array

import (
	"math/rand"
)

type Solution struct {
	pre []int
}

func Constructor528(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	x := rand.Intn(s.pre[len(s.pre)-1]) + 1
	return search(s.pre, x)
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
