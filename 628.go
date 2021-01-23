package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4, 6, -11, -2}))
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return max(nums[0]*nums[1]*nums[l-1], nums[l-1]*nums[l-2]*nums[l-3])
}

func maximumProduct2(nums []int) int {
	min1, min2 := math.MaxInt64, math.MinInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 {
			min2 = v
		}

		if v > max1 {
			max1, max2, max3 = v, max1, max2
		} else if v > max2 {
			max2, max3 = v, max2
		} else if v > max3 {
			max3 = v
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

// func max(a,b int) int {
// 	if a>b {
// 		return a
// 	}
// 	return b
// }
