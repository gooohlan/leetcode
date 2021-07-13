package main

func thirdMax(nums []int) int {
	a := nums[0]
	if len(nums) <= 2 {
		for _, v := range nums {
			if v > a {
				a = v
			}
		}
		return a
	}

	minNum := ^int(^uint(0) >> 1)
	b, c := minNum, minNum
	for i := 1; i < len(nums); i++ {
		if nums[i] > a {
			a, b, c = nums[i], a, b
		} else if nums[i] > b && nums[i] < a {
			b, c = nums[i], b
		} else if nums[i] > c && nums[i] < b {
			c = nums[i]
		}
	}
	if c == minNum {
		return a
	}
	return c
}
