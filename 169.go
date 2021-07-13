package main

import (
	"sort"
)

// map解法
func majorityElement(nums []int) int {
	mapArr := make(map[int]int)
	for _, v := range nums {
		if _, ok := mapArr[v]; !ok {
			mapArr[v] = 1
		} else {
			mapArr[v]++
		}
	}

	for k, v := range mapArr {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

func majorityElement0(nums []int) int {
	mapArr := make(map[int]int)
	for _, v := range nums {
		mapArr[v]++
		if mapArr[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// 题目中保证了众数出现次数多余一半,所以遇到众数+1,非众数-1的最终结果始终大于1
func majorityElement1(nums []int) int {
	count, repeat := 0, nums[0]
	for k, v := range nums {
		if v == repeat {
			count++
		} else {
			count--
			if count == 0 && k < len(nums) {
				repeat = nums[k+1]
			}
		}
	}
	return repeat
}

// 先排序直接找 len/2的数
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
