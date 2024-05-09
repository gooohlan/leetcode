package main

import (
	"math"
	"sort"
)

// [1,2,3] => [2,3,3] => [3,4,3] => [4,4,4]  3
// [1,3,5] => [2,4,5] => [3,5,5] => [4,5,6] => [5,7,5] => [6,7,6] => [7,7,7]  6
// [1,2,5] => [2,3,5] => [3,4,5] => [4,5,5] => [5,6,5] => [6,6,6] => 5
// [1,2,3,4] => [2,3,4,4] => [3,4,5,4] => [4,5,6,4] => [5,6,6,5] => [6,6,7,6] => [7,7,7,7]  6

func minMoves(nums []int) int {
	sort.Ints(nums)
	var count int
	for i := len(nums) - 1; i > 0; i-- {
		count += nums[i] - nums[0]
	}
	return count
}

func minMoves2(nums []int) int {
	count, min := 0, math.MaxInt32
	for _, num := range nums {
		count += num
		if min > num {
			min = num
		}
	}
	return count - len(nums)*min
}



func minMoves3(nums []int) int {
	sort.Ints(nums)
	var moves int
	for i := 1; i < len(nums); i++ {
		diff := moves + nums[i]- nums[i-1]
		nums[i] += moves // 这里比较难理解,此时相加的实际上是完成上一次操作中的改变nums[2]
		moves += diff
	}
	return moves
}
