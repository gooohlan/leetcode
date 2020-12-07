package main

import (
	"fmt"
	"sort"
)

// 哈希
func missingNumber1(nums []int) int {
	mapArr := make(map[int]bool)
	for _, v := range nums {
		mapArr[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := mapArr[i]; !ok {
			return i
		}
	}
	return 0
}

// 排序

func missingNumber2(nums []int) int {
	sort.Ints(nums)
	for k, v := range nums {
		if k != v {
			return k
		}
	}
	return len(nums)
}

// 数学
func missingNumber3(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

// 异或位运算
func missingNumber4(nums []int) int {
	miss := len(nums)
	for k, v := range nums {
		miss ^= k ^ v
	}
	return miss
}

func main() {
	fmt.Println(missingNumber4([]int{0, 1, 2, 4, 5, 3}))
}
