package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}

// [1,2,7,2,5,9]
// [5,9,9,8]
// [3,4,2,3]

func checkPossibility(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var count int
	for i := 1; i < len(nums) && count < 2; i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		count++
		if i >= 2 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return count <= 1
}
