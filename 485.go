package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var count, maxCount int
	for _, v := range nums {
		if v == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1}))
}
