package main

import (
	"log"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	log.Println(arrayPairSum([]int{1,4,3,2}))
}
